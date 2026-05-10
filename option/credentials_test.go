// Tests for the dual-header api-key auth path (v1.0).

package option

import (
	"fmt"
	"strings"
	"testing"

	core "github.com/AIEngineering26/promptvm-go-sdk/core"
	assert "github.com/stretchr/testify/assert"
)

func TestWithCredentialsSetsBothHeaders(t *testing.T) {
	opts := core.NewRequestOptions(WithCredentials("pk_live_smoke", "sk_live_smoke"))
	header := opts.ToHeader()

	assert.Equal(t, "pk_live_smoke", header.Get(core.HeaderPublicKey))
	assert.Equal(t, "sk_live_smoke", header.Get(core.HeaderSecretKey))
}

func TestWithCredentialsBothEmptyOmitsHeaders(t *testing.T) {
	// Defensive: passing empty strings produces no auth headers (the
	// backend would 401 anyway). This guards against accidentally
	// setting an empty bearer that some proxies may interpret as
	// "authenticated".
	opts := core.NewRequestOptions(WithCredentials("", ""))
	header := opts.ToHeader()

	assert.Empty(t, header.Get(core.HeaderPublicKey))
	assert.Empty(t, header.Get(core.HeaderSecretKey))
}

func TestWithAPIKeyPanicsWithMigrationMessage(t *testing.T) {
	// Legacy v0.x callers must hit a clear panic, not a silent 401.
	defer func() {
		r := recover()
		assert.NotNil(t, r, "WithAPIKey must panic in v1.0")
		msg, ok := r.(string)
		if !ok {
			// recover() returned non-string panic value — still
			// surface it so the test fails with useful context.
			t.Fatalf("WithAPIKey panic value was not a string: %v", r)
		}
		assert.Contains(t, msg, "WithCredentials", "panic message must point at the new API")
	}()
	WithAPIKey("legacy-key")
}

func TestRequestOptionsStringRedactsSecretKey(t *testing.T) {
	// FR-3: the secret key MUST NOT appear in any default debug/trace
	// output. Both %s and %#v default to a stringification path; both
	// must redact.
	opts := core.NewRequestOptions(WithCredentials("pk_redact_test", "sk_redact_VERY_SECRET_VALUE"))

	pretty := fmt.Sprintf("%v", opts)
	debug := fmt.Sprintf("%#v", opts)

	assert.False(t, strings.Contains(pretty, "sk_redact_VERY_SECRET_VALUE"),
		"secret key leaked in %%v output: %s", pretty)
	assert.False(t, strings.Contains(debug, "sk_redact_VERY_SECRET_VALUE"),
		"secret key leaked in %%#v output: %s", debug)

	// Sanity: the public key still appears so debug output stays useful.
	assert.Contains(t, pretty, "pk_redact_test")
}

func TestCredentialsOptionStringRedactsSecretKey(t *testing.T) {
	// Defence in depth: a caller logging the option directly (rather
	// than the RequestOptions container) must also see [REDACTED].
	opt := WithCredentials("pk_test", "sk_super_secret")
	pretty := fmt.Sprintf("%v", opt)
	debug := fmt.Sprintf("%#v", opt)

	assert.False(t, strings.Contains(pretty, "sk_super_secret"),
		"secret key leaked in %%v output: %s", pretty)
	assert.False(t, strings.Contains(debug, "sk_super_secret"),
		"secret key leaked in %%#v output: %s", debug)
}
