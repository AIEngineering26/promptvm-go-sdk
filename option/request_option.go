// PromptVM SDK options. WithCredentials supersedes the v0.x WithAPIKey;
// see MIGRATION.md for the one-line change.
//
// Hand-maintained (see ../.fernignore) — Fern would otherwise reset the
// auth option back to a single-header WithAPIKey on each regeneration.

package option

import (
	core "github.com/AIEngineering26/promptvm-go-sdk/core"
	http "net/http"
	url "net/url"
)

// RequestOption adapts the behavior of an individual request.
type RequestOption = core.RequestOption

// WithBaseURL sets the base URL, overriding the default
// environment, if any.
func WithBaseURL(baseURL string) *core.BaseURLOption {
	return &core.BaseURLOption{
		BaseURL: baseURL,
	}
}

// WithHTTPClient uses the given HTTPClient to issue the request.
func WithHTTPClient(httpClient core.HTTPClient) *core.HTTPClientOption {
	return &core.HTTPClientOption{
		HTTPClient: httpClient,
	}
}

// WithHTTPHeader adds the given http.Header to the request.
func WithHTTPHeader(httpHeader http.Header) *core.HTTPHeaderOption {
	return &core.HTTPHeaderOption{
		// Clone the headers so they can't be modified after the option call.
		HTTPHeader: httpHeader.Clone(),
	}
}

// WithBodyProperties adds the given body properties to the request.
func WithBodyProperties(bodyProperties map[string]interface{}) *core.BodyPropertiesOption {
	copiedBodyProperties := make(map[string]interface{}, len(bodyProperties))
	for key, value := range bodyProperties {
		copiedBodyProperties[key] = value
	}
	return &core.BodyPropertiesOption{
		BodyProperties: copiedBodyProperties,
	}
}

// WithQueryParameters adds the given query parameters to the request.
func WithQueryParameters(queryParameters url.Values) *core.QueryParametersOption {
	copiedQueryParameters := make(url.Values, len(queryParameters))
	for key, values := range queryParameters {
		copiedQueryParameters[key] = values
	}
	return &core.QueryParametersOption{
		QueryParameters: copiedQueryParameters,
	}
}

// WithMaxAttempts configures the maximum number of retry attempts.
func WithMaxAttempts(attempts uint) *core.MaxAttemptsOption {
	return &core.MaxAttemptsOption{
		MaxAttempts: attempts,
	}
}

// WithCredentials sets the dual-header api-key credentials. The PromptVM
// backend's api-key middleware requires both X-PromptVM-Public-Key and
// X-PromptVM-Secret-Key on every authenticated request; passing only
// one is a no-op for auth (the backend would 401 the request anyway).
//
// Find both values on the API keys page in the dashboard:
// https://app.promptvm.ai/settings/api-keys
func WithCredentials(publicKey, secretKey string) *core.CredentialsOption {
	return &core.CredentialsOption{
		PublicKey: publicKey,
		SecretKey: secretKey,
	}
}

// WithAPIKey is a v0.x stub that panics with a clear migration message.
// It is kept so callers compiling against v0.x land on a loud, actionable
// error rather than a silent 401 at request time. Remove in a future
// major.
//
// Deprecated: use WithCredentials(publicKey, secretKey).
func WithAPIKey(_ string) RequestOption {
	panic(
		"PromptVM SDK v1.0 removed WithAPIKey — use option.WithCredentials(publicKey, secretKey). " +
			"Find both values on the API keys page in the dashboard.",
	)
}
