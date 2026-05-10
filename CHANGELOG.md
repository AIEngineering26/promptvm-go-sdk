# Changelog

## v1.0.0 — 2026-05-10

### Breaking
- **Authentication**: the SDK now sends both `X-PromptVM-Public-Key`
  and `X-PromptVM-Secret-Key` on every authenticated request. The old
  single-header behaviour was broken — the PromptVM backend has always
  required both headers, so v0.x clients always received 401 on authed
  routes.
- **Option**: `option.WithAPIKey(string)` removed. Use
  `option.WithCredentials(publicKey, secretKey string)` instead. The
  old `WithAPIKey` is kept as a panic stub that points at the new
  option so callers compiling against v0.x land on a clear migration
  message.
- **Core**: `core.RequestOptions.APIKey` field removed; `PublicKey`
  and `SecretKey` fields added.
- **Redaction**: `core.RequestOptions.String()` and `GoString()` (and
  the same on `core.CredentialsOption`) now redact the secret-key
  value, so default `%v` / `%#v` debug printing never leaks the
  secret half (FR-3).

### Migration
See [MIGRATION.md](./MIGRATION.md).

### Why no compatibility shim
The dual-header requirement is a backend-side invariant
(`services/backend/src/modules/api-keys/auth-middleware.ts`). The
single-header path could not have been working in production; a
single major version bump with a panic stub on the old function is
the cleanest way to communicate that.
