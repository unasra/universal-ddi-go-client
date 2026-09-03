package internal

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// RetryConfig controls automatic retry behavior for transient HTTP errors.
type RetryConfig struct {
	MaxAttempts int
	MinWait     time.Duration
	MaxWait     time.Duration
}

func isRetryableStatusCode(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return true
	}
	return false
}

func parseRetryAfter(resp *http.Response) time.Duration {
	header := resp.Header.Get("Retry-After")
	if header == "" {
		return 0
	}

	if seconds, err := strconv.Atoi(header); err == nil && seconds > 0 {
		return time.Duration(seconds) * time.Second
	}

	if t, err := http.ParseTime(header); err == nil {
		delay := time.Until(t)
		if delay > 0 {
			return delay
		}
	}

	return 0
}

func retryBackoff(attempt int, minWait, maxWait time.Duration) time.Duration {
	wait := float64(minWait) * math.Pow(2, float64(attempt))
	if wait > float64(maxWait) {
		wait = float64(maxWait)
	}

	// Add jitter: ±25% of the computed wait
	jitter := wait * 0.25 * (rand.Float64()*2 - 1)
	wait += jitter

	if wait < float64(minWait) {
		wait = float64(minWait)
	}

	return time.Duration(wait)
}
