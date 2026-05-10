# Migrating to v1.0

v1.0 fixes a longstanding auth bug: the SDK now sends both
`X-PromptVM-Public-Key` **and** `X-PromptVM-Secret-Key` on every
authenticated request, matching what the PromptVM backend has always
required. Pre-1.0 releases sent only the public key, so authenticated
calls returned `401 unauthorized` regardless of whether the credentials
were valid.

This is a **breaking change**. There is no compatibility shim —
`option.WithAPIKey` is kept only as a panic stub that points users at
the new option, so v0.x callers land on a loud, actionable error rather
than a silent 401.

## What changed

- `option.WithAPIKey(string)` → `option.WithCredentials(publicKey, secretKey string)`
- Calling `option.WithAPIKey(...)` panics with a message that names the
  new option and links the dashboard.
- Every authed request now carries both header values.
- `core.RequestOptions.APIKey` field removed; `PublicKey` + `SecretKey`
  fields added.
- `core.RequestOptions.String()` and `GoString()` redact `SecretKey` so
  default `%v` / `%#v` printing of options never leaks the secret half.

## Find your keys

Open the API keys page in the dashboard:
<https://app.promptvm.ai/settings/api-keys>

The `publicKey` is the `pk_live_…` value. The `secretKey` is the
`sk_live_…` value, shown once when the key is created. Rotate the key
if you lost the secret half.

## Before (v0.x)

```go
c := client.NewClient(
    option.WithAPIKey("pk_live_only_was_never_enough"),
)
```

## After (v1.0)

```go
c := client.NewClient(
    option.WithCredentials(
        os.Getenv("PROMPTVM_PUBLIC_KEY"),
        os.Getenv("PROMPTVM_SECRET_KEY"),
    ),
)
```

## What's not changing

- The module path stays at `github.com/AIEngineering26/promptvm-go-sdk`
  (no `/v2` suffix). The first stable major release tags as `v1.0.0`,
  which Go's import compatibility rule allows on the existing path.
- All resource clients (`c.Prompts`, `c.Workspaces`, …) keep their
  existing method shapes and request/response types.
- Environments, retries, custom HTTP headers, and HTTP client config
  are unchanged.

## Why no compatibility shim

The dual-header requirement is a backend-side invariant (see
`services/backend/src/modules/api-keys/auth-middleware.ts`). The
single-header path could not have been working in production; a
single major version bump with a panic stub on the old function is
the cleanest way to communicate that.
