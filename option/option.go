package option

import (
	"net/http"

	"github.com/infobloxopen/universal-ddi-go-client/internal"
	"golang.org/x/time/rate"
)

// ClientOption is a function that applies configuration options to the API Client.
type ClientOption func(configuration *internal.Configuration)

// WithCSPUrl returns a ClientOption that sets the URL for Universal DDI Cloud Services Portal.
// Can also be configured using the `INFOBLOX_PORTAL_URL` environment variable.
// Optional. Default is https://csp.infoblox.com
func WithCSPUrl(cspURL string) ClientOption {
	return func(configuration *internal.Configuration) {
		if cspURL != "" {
			configuration.CSPURL = cspURL
		}
	}
}

// WithAPIKey returns a ClientOption that sets the API Key for accessing the Universal DDI API.
// Can also be configured by using the `INFOBLOX_PORTAL_KEY` environment variable.
// You can configure an API key for your user account in the Universal DDI Cloud Services Portal.
// Please refer to the following link for more information: https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys
// Required
func WithAPIKey(apiKey string) ClientOption {
	return func(configuration *internal.Configuration) {
		if apiKey != "" {
			configuration.APIKey = apiKey
		}
	}
}

// WithHTTPClient returns a ClientOption that sets the HTTPClient to use for the SDK.
// Optional. The default HTTPClient will be used if not provided.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(configuration *internal.Configuration) {
		if httpClient != nil {
			configuration.HTTPClient = httpClient
		}
	}
}

// WithDefaultTags returns a ClientOption that sets the tags the client can set by default for objects that has tags support.
// Optional.
func WithDefaultTags(defaultTags map[string]string) ClientOption {
	return func(configuration *internal.Configuration) {
		configuration.DefaultTags = defaultTags
	}
}

// WithClientName returns a ClientOption that sets the name of the client using the SDK.
// This can be used to identify the client in the audit logs.
// Optional. If not provided, the client name will be set to "universal-ddi-go-client".
func WithClientName(clientName string) ClientOption {
	return func(configuration *internal.Configuration) {
		if clientName != "" {
			configuration.ClientName = clientName
		}
	}
}

// WithDebug returns a ClientOption that sets the debug mode.
// Enabling the debug flag will write the request and response to the log.
func WithDebug(debug bool) ClientOption {
	return func(configuration *internal.Configuration) {
		configuration.Debug = debug
	}
}

// WithRateLimit returns a ClientOption that enables or disables the built-in rate limiter.
// When enable is true, the default rate of 25 requests/second with a burst of 25 is applied.
// When enable is false, rate limiting is disabled entirely.
// Use WithRateLimiter to supply a custom rate limiter with non-default values.
// Can also be configured using the INFOBLOX_RATE_LIMIT environment variable.
func WithRateLimit(enable bool) ClientOption {
	return func(configuration *internal.Configuration) {
		if enable {
			configuration.RateLimiter = rate.NewLimiter(rate.Limit(internal.DefaultRateLimit), internal.DefaultRateLimitBurst)
		} else {
			configuration.RateLimiter = nil
		}
	}
}

// WithRateLimiter returns a ClientOption that sets a custom rate limiter implementation.
// Use this for advanced use cases like distributed rate limiting or
// per-resource-type limiting at the provider level.
func WithRateLimiter(limiter internal.RateLimiter) ClientOption {
	return func(configuration *internal.Configuration) {
		if limiter != nil {
			configuration.RateLimiter = limiter
		}
	}
}
