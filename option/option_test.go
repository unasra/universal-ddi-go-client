package option

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infobloxopen/universal-ddi-go-client/internal"
)

func TestWithCSPUrl(t *testing.T) {
	config := &internal.Configuration{}
	url := "http://test.com"
	opt := WithCSPUrl(url)
	opt(config)
	assert.Equal(t, url, config.CSPURL)
}

func TestWithAPIKey(t *testing.T) {
	config := &internal.Configuration{}
	apiKey := "testKey"
	opt := WithAPIKey(apiKey)
	opt(config)
	assert.Equal(t, apiKey, config.APIKey)
}

func TestWithHTTPClient(t *testing.T) {
	config := &internal.Configuration{}
	client := &http.Client{}
	opt := WithHTTPClient(client)
	opt(config)
	assert.Equal(t, client, config.HTTPClient)
}

func TestWithDefaultTags(t *testing.T) {
	config := &internal.Configuration{}
	tags := map[string]string{"tag1": "value1"}
	opt := WithDefaultTags(tags)
	opt(config)
	assert.Equal(t, tags, config.DefaultTags)
}

func TestWithClientName(t *testing.T) {
	config := &internal.Configuration{}
	name := "testClient"
	opt := WithClientName(name)
	opt(config)
	assert.Equal(t, name, config.ClientName)
}

func TestWithDebug(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithDebug(true)
	opt(config)
	assert.Equal(t, true, config.Debug)
}

func TestWithRateLimitEnable(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithRateLimit(true)
	opt(config)
	assert.NotNil(t, config.RateLimiter)
}

func TestWithRateLimitDisable(t *testing.T) {
	config := &internal.Configuration{}
	config.RateLimiter = &mockRateLimiter{}
	opt := WithRateLimit(false)
	opt(config)
	assert.Nil(t, config.RateLimiter)
}

type mockRateLimiter struct{}

func (m *mockRateLimiter) Wait(ctx context.Context) error { return nil }

func TestWithRateLimiter(t *testing.T) {
	config := &internal.Configuration{}
	limiter := &mockRateLimiter{}
	opt := WithRateLimiter(limiter)
	opt(config)
	assert.Equal(t, limiter, config.RateLimiter)
}

func TestWithRateLimiterNil(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithRateLimiter(nil)
	opt(config)
	assert.Nil(t, config.RateLimiter)
}
