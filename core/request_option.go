// PromptVM SDK — dual-header api-key authentication.
//
// The PromptVM backend's api-key middleware requires BOTH
// X-PromptVM-Public-Key and X-PromptVM-Secret-Key on every authenticated
// request (services/backend/src/modules/api-keys/auth-middleware.ts).
// Earlier SDK releases shipped only the public key, which meant authed
// requests always returned 401. This file is the single source of truth
// for the auth header set.
//
// This file is intentionally hand-maintained (see .fernignore). The
// upstream Fern overlay declares the dual-header scheme so a clean
// regeneration produces the same shape; pinning here keeps these manual
// edits load-bearing in case generator behaviour changes.

package core

import (
	fmt "fmt"
	http "net/http"
	url "net/url"
)

const (
	HeaderPublicKey = "X-PromptVM-Public-Key"
	HeaderSecretKey = "X-PromptVM-Secret-Key"
)

// RequestOption adapts the behavior of the client or an individual request.
type RequestOption interface {
	applyRequestOptions(*RequestOptions)
}

// RequestOptions defines all of the possible request options.
//
// This type is primarily used by the generated code and is not meant
// to be used directly; use the option package instead.
type RequestOptions struct {
	BaseURL         string
	HTTPClient      HTTPClient
	HTTPHeader      http.Header
	BodyProperties  map[string]interface{}
	QueryParameters url.Values
	MaxAttempts     uint
	PublicKey       string
	SecretKey       string
}

// NewRequestOptions returns a new *RequestOptions value.
//
// This function is primarily used by the generated code and is not meant
// to be used directly; use RequestOption instead.
func NewRequestOptions(opts ...RequestOption) *RequestOptions {
	options := &RequestOptions{
		HTTPHeader:      make(http.Header),
		BodyProperties:  make(map[string]interface{}),
		QueryParameters: make(url.Values),
	}
	for _, opt := range opts {
		opt.applyRequestOptions(options)
	}
	return options
}

// ToHeader maps the configured request options into a http.Header used
// for the request(s).
//
// Both PublicKey and SecretKey must be set together; setting only one
// is a no-op for the auth headers (the backend would reject the
// request anyway). String() on RequestOptions never echoes SecretKey
// — see the custom GoString/String methods below.
func (r *RequestOptions) ToHeader() http.Header {
	header := r.cloneHeader()
	if r.PublicKey != "" && r.SecretKey != "" {
		header.Set(HeaderPublicKey, fmt.Sprintf("%v", r.PublicKey))
		header.Set(HeaderSecretKey, fmt.Sprintf("%v", r.SecretKey))
	}
	return header
}

// String redacts SecretKey so default %s/%v printing of a
// *RequestOptions never leaks the secret half. The non-secret fields
// are included so debug output stays useful.
func (r *RequestOptions) String() string {
	return fmt.Sprintf(
		"core.RequestOptions{BaseURL:%q MaxAttempts:%d PublicKey:%q SecretKey:[REDACTED]}",
		r.BaseURL,
		r.MaxAttempts,
		r.PublicKey,
	)
}

// GoString mirrors String for the %#v verb so spew/pretty printers
// don't otherwise dump every field including SecretKey.
func (r *RequestOptions) GoString() string {
	return r.String()
}

func (r *RequestOptions) cloneHeader() http.Header {
	headers := r.HTTPHeader.Clone()
	headers.Set("X-Fern-Language", "Go")
	headers.Set("X-Fern-SDK-Name", "github.com/AIEngineering26/promptvm-go-sdk")
	headers.Set("X-Fern-SDK-Version", "v1.0.0")
	return headers
}

// BaseURLOption implements the RequestOption interface.
type BaseURLOption struct {
	BaseURL string
}

func (b *BaseURLOption) applyRequestOptions(opts *RequestOptions) {
	opts.BaseURL = b.BaseURL
}

// HTTPClientOption implements the RequestOption interface.
type HTTPClientOption struct {
	HTTPClient HTTPClient
}

func (h *HTTPClientOption) applyRequestOptions(opts *RequestOptions) {
	opts.HTTPClient = h.HTTPClient
}

// HTTPHeaderOption implements the RequestOption interface.
type HTTPHeaderOption struct {
	HTTPHeader http.Header
}

func (h *HTTPHeaderOption) applyRequestOptions(opts *RequestOptions) {
	opts.HTTPHeader = h.HTTPHeader
}

// BodyPropertiesOption implements the RequestOption interface.
type BodyPropertiesOption struct {
	BodyProperties map[string]interface{}
}

func (b *BodyPropertiesOption) applyRequestOptions(opts *RequestOptions) {
	opts.BodyProperties = b.BodyProperties
}

// QueryParametersOption implements the RequestOption interface.
type QueryParametersOption struct {
	QueryParameters url.Values
}

func (q *QueryParametersOption) applyRequestOptions(opts *RequestOptions) {
	opts.QueryParameters = q.QueryParameters
}

// MaxAttemptsOption implements the RequestOption interface.
type MaxAttemptsOption struct {
	MaxAttempts uint
}

func (m *MaxAttemptsOption) applyRequestOptions(opts *RequestOptions) {
	opts.MaxAttempts = m.MaxAttempts
}

// CredentialsOption sets the dual-header api-key credentials.
type CredentialsOption struct {
	PublicKey string
	SecretKey string
}

func (c *CredentialsOption) applyRequestOptions(opts *RequestOptions) {
	opts.PublicKey = c.PublicKey
	opts.SecretKey = c.SecretKey
}

// String redacts SecretKey on this option type too — defence in depth
// in case a caller logs the option directly rather than the
// RequestOptions container.
func (c *CredentialsOption) String() string {
	return fmt.Sprintf("core.CredentialsOption{PublicKey:%q SecretKey:[REDACTED]}", c.PublicKey)
}

func (c *CredentialsOption) GoString() string {
	return c.String()
}
