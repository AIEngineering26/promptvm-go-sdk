# Reference
## Authentication
<details><summary><code>client.Authentication.McpAuthorize() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Browser-facing endpoint used by the MCP server to authenticate users via the backend session. If the user has a valid session, issues a one-time auth code and redirects back to the MCP callback URL.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.McpAuthorizeRequest{
        RedirectURI: "redirect_uri",
        State: "state",
    }
client.Authentication.McpAuthorize(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**redirectURI:** `string` — MCP callback URL to redirect to after auth
    
</dd>
</dl>

<dl>
<dd>

**state:** `string` — Opaque state from the MCP OAuth flow
    
</dd>
</dl>

<dl>
<dd>

**scope:** `*string` — Requested OAuth scopes (space-separated)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Authentication.UpdateUIPreferences(request) -> *promptvmgosdk.UpdateUIPreferencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Merges the provided UI preferences (panel collapse state, widths, etc.) with the user's existing preferences. Cached in Redis with a 1-hour TTL using a write-through strategy so subsequent reads hit cache.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateUIPreferencesRequest{}
client.Authentication.UpdateUIPreferences(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sidebarOpen:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**leftPanelCollapsed:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**leftPanelWidth:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**rightPanelCollapsed:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**rightPanelWidth:** `*float64` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CliAuth
<details><summary><code>client.CliAuth.CliAuthorize(request) -> *promptvmgosdk.CliAuthorizeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Called by the frontend consent page after the user authorizes a CLI login request. Returns a one-time authorization code tied to the PKCE challenge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CliAuthorizeRequest{
        CodeChallenge: "code_challenge",
        CodeChallengeMethod: promptvmgosdk.CliAuthorizeRequestCodeChallengeMethodS256,
        RedirectURI: "redirect_uri",
        ClientID: promptvmgosdk.CliAuthorizeRequestClientIDPromptvmCli,
    }
client.CliAuth.CliAuthorize(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**codeChallenge:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**codeChallengeMethod:** `*promptvmgosdk.CliAuthorizeRequestCodeChallengeMethod` 
    
</dd>
</dl>

<dl>
<dd>

**redirectURI:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**deviceName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*promptvmgosdk.CliAuthorizeRequestClientID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CliAuth.CliDeviceAuthorize(request) -> *promptvmgosdk.CliDeviceAuthorizeResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CliDeviceAuthorizeRequest{
        UserCode: "user_code",
        Action: promptvmgosdk.CliDeviceAuthorizeRequestActionApprove,
    }
client.CliAuth.CliDeviceAuthorize(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userCode:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**action:** `*promptvmgosdk.CliDeviceAuthorizeRequestAction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CliAuth.ListCliSessions() -> []*promptvmgosdk.ListCliSessionsResponseItem</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.CliAuth.ListCliSessions(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CliAuth.RevokeCliSession(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RevokeCliSessionRequest{
        ID: "id",
    }
client.CliAuth.RevokeCliSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CliAuth.CliTokenExchange(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

OAuth 2.0 token endpoint supporting grant_type=authorization_code (with PKCE) and grant_type=refresh_token. Public — no client secret.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CliTokenExchangeRequest{
        AuthorizationCode: &promptvmgosdk.CliTokenExchangeRequestAuthorizationCode{
            Code: "code",
            CodeVerifier: "code_verifier",
            RedirectURI: "redirect_uri",
            ClientID: promptvmgosdk.CliTokenExchangeRequestAuthorizationCodeClientIDPromptvmCli,
        },
    }
client.CliAuth.CliTokenExchange(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*promptvmgosdk.CliTokenExchangeRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CliAuth.CliDeviceCode(request) -> *promptvmgosdk.CliDeviceCodeResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CliDeviceCodeRequest{
        ClientID: promptvmgosdk.CliDeviceCodeRequestClientIDPromptvmCli,
    }
client.CliAuth.CliDeviceCode(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `*promptvmgosdk.CliDeviceCodeRequestClientID` 
    
</dd>
</dl>

<dl>
<dd>

**deviceName:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CliAuth.CliDeviceToken(request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CliDeviceTokenRequest{
        GrantType: promptvmgosdk.CliDeviceTokenRequestGrantTypeUrnIetfParamsOauthGrantTypeDeviceCode,
        DeviceCode: "device_code",
        ClientID: promptvmgosdk.CliDeviceTokenRequestClientIDPromptvmCli,
    }
client.CliAuth.CliDeviceToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**grantType:** `*promptvmgosdk.CliDeviceTokenRequestGrantType` 
    
</dd>
</dl>

<dl>
<dd>

**deviceCode:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*promptvmgosdk.CliDeviceTokenRequestClientID` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Billing
<details><summary><code>client.Billing.GetPromoOffer(Token) -> *promptvmgosdk.GetPromoOfferResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Public endpoint that returns the offer metadata and status for a single-use redemption token. No authentication is required. Returns 200 when the token is valid and the offer is available. Returns 410 Gone for any non-available status (already used, disabled, expired, or exhausted) so the frontend can render a user-friendly "link no longer valid" view without leaking token existence. Rate-limited to 30 requests/minute per IP.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetPromoOfferRequest{
        Token: "token",
    }
client.Billing.GetPromoOffer(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**token:** `string` — URL-safe redemption token (base64url, 32 chars). e.g. ABCDEFGHIJKLMNOPQRSTUVWXYZabcdef
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ListBillingPlans() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all active, public subscription plans sorted by display order. No authentication required. Values are DB-driven: tier names, prices, feature bullets, and seat limits are all set via SQL — no frontend code changes are needed when the catalog changes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Billing.ListBillingPlans(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.CreateBillingCheckoutSession(request) -> *promptvmgosdk.CreateBillingCheckoutSessionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a Stripe-hosted Checkout URL for the caller to complete payment. Owner/admin only. Rate-limited 5/min/org. The frontend redirects the browser to the returned URL; the actual subscription state lands in our DB via the `customer.subscription.created` webhook (US-01-4), not via the success URL. Returns 503 `billing_not_live` when the `FEATURE_BILLING_LIVE` kill-switch is disengaged (US-05-7 / FR-05-9) — clients should render an Upgrade-disabled state and route users to sales.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateBillingCheckoutSessionRequest{
        OrgID: "018e4a3b-0000-0000-0000-000000000001",
        PlanSlug: "pro",
        Interval: promptvmgosdk.CreateBillingCheckoutSessionRequestIntervalMonth,
    }
client.Billing.CreateBillingCheckoutSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>

<dl>
<dd>

**planSlug:** `string` — Subscription plan slug from the catalog (e.g. "pro", "teams"). Free is rejected — Free orgs use the backfill row, not Checkout.
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*promptvmgosdk.CreateBillingCheckoutSessionRequestInterval` — Billing interval. Must match a `subscription_plans` row paired with `planSlug`.
    
</dd>
</dl>

<dl>
<dd>

**seats:** `*int` — Seat quantity for per-seat plans. Required when the resolved plan has `is_per_seat=true`; rejected otherwise.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.CreateBillingPortalSession() -> *promptvmgosdk.CreateBillingPortalSessionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a Stripe-hosted Customer Portal URL the caller can use to update their payment method, view invoices, edit billing details, or cancel their subscription. Owner/admin only. Rate-limited 5/min/org. Plan switching is intentionally disabled in the Portal — use `/billing/change-plan` for that. Returns 503 `billing_not_live` when the `FEATURE_BILLING_LIVE` kill-switch is disengaged (US-05-7 / FR-05-9).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateBillingPortalSessionRequest{
        OrgID: "018e4a3b-0000-0000-0000-000000000001",
    }
client.Billing.CreateBillingPortalSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ChangeBillingPlan(request) -> *promptvmgosdk.ChangeBillingPlanResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Routes through the Phase 02 change-plan algorithm: detects upgrade/downgrade/seat-adjust direction, releases any pending `subscription_schedule` first (FR-02-5a), then writes the change to Stripe (immediate proration on upgrade, scheduled phase on downgrade). Owner/admin only. Rate-limited per org via Redis SET NX with TTL `BILLING_CHANGE_PLAN_COOLDOWN_SECONDS` (default 60, 0 in tests). The authoritative state lands in our DB via the `customer.subscription.updated` webhook (US-01-4) — clients should re-read `/billing/status` after a 200. Returns 503 `billing_not_live` when the `FEATURE_BILLING_LIVE` kill-switch is disengaged (US-05-7 / FR-05-9).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ChangeBillingPlanRequest{
        OrgID: "018e4a3b-0000-0000-0000-000000000001",
        PlanSlug: "teams",
        Interval: promptvmgosdk.ChangeBillingPlanRequestIntervalMonth,
    }
client.Billing.ChangeBillingPlan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>

<dl>
<dd>

**planSlug:** `string` — Target subscription plan slug from the catalog (e.g. "pro", "teams"). Free is rejected — use `/billing/cancel` to schedule a downgrade to Free.
    
</dd>
</dl>

<dl>
<dd>

**interval:** `*promptvmgosdk.ChangeBillingPlanRequestInterval` — Target billing interval. month→year=upgrade, year→month=downgrade (FR-02-4).
    
</dd>
</dl>

<dl>
<dd>

**seats:** `*int` — Seat quantity for per-seat plans. Required when the target plan has `is_per_seat=true`; rejected otherwise. Must satisfy `min_seats ≤ seats ≤ max_seats` and `seats ≥ used_seats` (FR-02-5b).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.AdjustBillingSeats(request) -> *promptvmgosdk.AdjustBillingSeatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Owner/admin only. Pushes a `quantity` change to the Stripe subscription item with `proration_behavior=create_prorations`. Increases generate a prorated charge on the next invoice; decreases (still ≥ used_seats) generate a prorated credit on the next invoice (no card refund). Per-seat plans only. Rate-limited 5/min/org (US-05-8 F-1). Returns 503 `billing_not_live` when the `FEATURE_BILLING_LIVE` kill-switch is disengaged (US-05-7 / FR-05-9).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdjustBillingSeatsRequest{
        OrgID: "018e4a3b-0000-0000-0000-000000000001",
        Seats: 5,
    }
client.Billing.AdjustBillingSeats(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>

<dl>
<dd>

**seats:** `int` — New total seat count for the active subscription. Must satisfy `min_seats ≤ seats ≤ max_seats` from the plan, and `seats ≥ used_seats`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.CancelBillingSubscription() -> *promptvmgosdk.CancelBillingSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Owner/admin only. Sets `cancel_at_period_end = true` on the Stripe subscription. Releases any pending downgrade schedule first (FR-02-15) so the update can apply cleanly. The cancellation is scheduled, not immediate — billing continues through the current period and access ends at `cancelAt`. Re-invoking with the same period is a no-op (FR-02-14 idempotency). Returns 503 `billing_not_live` when the `FEATURE_BILLING_LIVE` kill-switch is disengaged (US-05-7 / FR-05-9).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CancelBillingSubscriptionRequest{
        OrgID: "018e4a3b-0000-0000-0000-000000000001",
    }
client.Billing.CancelBillingSubscription(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ResumeBillingSubscription() -> *promptvmgosdk.ResumeBillingSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Owner/admin only. Sets `cancel_at_period_end = false` on the Stripe subscription. Does NOT recreate any previously scheduled downgrade — the user must re-request that via `/change-plan` (FR-02-15). Re-invoking with the same period is a no-op (FR-02-14 idempotency). Returns 503 `billing_not_live` when the `FEATURE_BILLING_LIVE` kill-switch is disengaged (US-05-7 / FR-05-9).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ResumeBillingSubscriptionRequest{
        OrgID: "018e4a3b-0000-0000-0000-000000000001",
    }
client.Billing.ResumeBillingSubscription(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ListBillingInvoices() -> *promptvmgosdk.ListBillingInvoicesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated, owner/admin-gated history of Stripe invoices mirrored locally by the webhook worker (US-02-4). Rows are ordered `created_at desc, id desc` and paginated via an opaque cursor (`?cursor`). Default page size is 20, maximum 50 (FR-02-12). The response always includes Stripe-hosted `hostedInvoiceUrl` + `invoicePdfUrl` so the frontend can deep-link to receipts without a second round-trip.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListBillingInvoicesRequest{
        OrgID: "018e4a3b-0000-0000-0000-000000000001",
    }
client.Billing.ListBillingInvoices(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Maximum number of invoices to return on this page. Default 20, max 50.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque base64 cursor from a prior response's `nextCursor`. Omit to start at the most recent invoice.
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.RedeemPromotionalOffer(Token) -> *promptvmgosdk.RedeemPromotionalOfferResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Owner/admin only. Atomically locks the token, increments the offer counter, and returns a Stripe Checkout URL configured with `trial_period_days` and `payment_method_collection: 'always'`. On Stripe failure runs a compensating reversal so the token is reusable (FR-06-8). One promotional trial per org per lifetime — enforced at the DB layer via the `promo_one_trial_per_org` partial UNIQUE index (FR-06-9). Rate-limited 10/min/IP AND 5/min/authenticated user (FR-06-5). Returns 503 `billing_not_live` when the `FEATURE_BILLING_LIVE` kill-switch is disengaged (US-05-7 / FR-05-9).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RedeemPromotionalOfferRequest{
        Token: "abcdef0123456789ABCDEF0123456789",
        OrgID: "018e4a3b-0000-0000-0000-000000000001",
    }
client.Billing.RedeemPromotionalOffer(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**token:** `string` — Single-use promotional-redemption token from a /invite/:token link.
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.GetBillingStatus() -> *promptvmgosdk.GetBillingStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the active org's subscription projection, plan-derived entitlements, and a usage snapshot. Any role inside the org may read. Cached server-side in Redis for 30s; invalidated by the Stripe webhook worker on every event apply.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetBillingStatusRequest{
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.Billing.GetBillingStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Billing.ListBillingWebhookErrors() -> *promptvmgosdk.ListBillingWebhookErrorsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the FR-05-11 dead-letter view: `stripe_webhook_events` rows where `error IS NOT NULL` AND `org_id = X-Org-Id` (US-05-7a-FIX). Ordered `received_at DESC, id DESC` with opaque cursor pagination. Default `limit=50`, max `200`. Supports `?since=<ISO8601>` and `?kind=<error-keyword>` filters. The `payload` jsonb column is INTENTIONALLY EXCLUDED — Stripe payloads can carry PII (customer email, billing address); ops reads the full payload from the worker logs or the Stripe Dashboard, not from this endpoint. Events whose `org_id` could not be resolved at ingress time (`org_id IS NULL`) are intentionally NOT returned — a future platform-admin route will surface them. Currently gated by `requireOwnerOrAdmin` (org-scoped) as a placeholder.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListBillingWebhookErrorsRequest{
        OrgID: "018e4a3b-0000-0000-0000-000000000001",
    }
client.Billing.ListBillingWebhookErrors(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Maximum number of events to return on this page. Default 50, max 200.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque base64 cursor from a prior response's `nextCursor`. Omit to start at the most recent error.
    
</dd>
</dl>

<dl>
<dd>

**since:** `*time.Time` — ISO-8601 timestamp lower bound — returns events with `received_at > since`. Use to scope to the last 24h / 1h.
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*string` — Case-insensitive substring match against the error text. Useful for parked classes like `unknown_price`, `unmapped_subscription`.
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AgentAPI
<details><summary><code>client.AgentAPI.AgentResolvePrompt(Slug) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns resolved prompt content as plain text or JSON. Pass ?version=latest (default) or ?version=<number>. Pass ?format=json for structured output. Any other query params are treated as variable values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AgentResolvePromptRequest{
        Slug: "slug",
    }
client.AgentAPI.AgentResolvePrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**format:** `*promptvmgosdk.AgentResolvePromptRequestFormat` 
    
</dd>
</dl>

<dl>
<dd>

**template:** `*string` — Force the {{placeholder}} form for every caller, including the prompt author. Used by surfaces that must render identically for everyone, such as a marketplace listing.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SkillsPublic
<details><summary><code>client.SkillsPublic.RecordPublicSkillInstall(Slug) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unauthenticated. Atomically increments the tied marketplace listing's downloadCount (never importCount). No-op (still 204) when the public skill has no listing. 404 for a missing/non-public skill.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RecordPublicSkillInstallRequest{
        Slug: "slug",
    }
client.SkillsPublic.RecordPublicSkillInstall(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**creator:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SkillsPublic.GetPublicSkillBySlug(Slug) -> *promptvmgosdk.GetPublicSkillBySlugResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unauthenticated read. Returns the parsed Agent Skills frontmatter as JSON, or the literal SKILL.md bytes when ?format=skill_md.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetPublicSkillBySlugRequest{
        Slug: "slug",
    }
client.SkillsPublic.GetPublicSkillBySlug(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**format:** `*promptvmgosdk.GetPublicSkillBySlugRequestFormat` 
    
</dd>
</dl>

<dl>
<dd>

**creator:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## HooksPublic
<details><summary><code>client.HooksPublic.GetPublicHookBySlug(Slug) -> *promptvmgosdk.GetPublicHookBySlugResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unauthenticated read. Returns the hook configuration as JSON.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetPublicHookBySlugRequest{
        Slug: "slug",
    }
client.HooksPublic.GetPublicHookBySlug(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**creator:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ContextSync
<details><summary><code>client.ContextSync.GetContextSyncManifest() -> *promptvmgosdk.GetContextSyncManifestResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The canonical hook + manifest spec the CLI `promptvm sync init` and the MCP setup prompt render from. Public, CORS-enabled, cached 300s.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ContextSync.GetContextSyncManifest(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContextSync.ListCapturedSessions() -> *promptvmgosdk.ListCapturedSessionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListCapturedSessionsRequest{
        WorkspaceID: "workspaceId",
    }
client.ContextSync.ListCapturedSessions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.ListCapturedSessionsRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**project:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**repoSlug:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**repo:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**branch:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**authorID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*promptvmgosdk.ListCapturedSessionsRequestSort` 
    
</dd>
</dl>

<dl>
<dd>

**order:** `*promptvmgosdk.ListCapturedSessionsRequestOrder` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContextSync.IngestCapturedSession(request) -> *promptvmgosdk.IngestCapturedSessionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.IngestCapturedSessionRequest{
        WorkspaceID: "workspaceId",
        ClaudeSessionID: "claudeSessionId",
        Source: "source",
        CaptureMode: promptvmgosdk.IngestCapturedSessionRequestCaptureModeSummary,
    }
client.ContextSync.IngestCapturedSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**claudeSessionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**source:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**captureMode:** `*promptvmgosdk.IngestCapturedSessionRequestCaptureMode` 
    
</dd>
</dl>

<dl>
<dd>

**summary:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `*promptvmgosdk.IngestCapturedSessionRequestMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**contentHash:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**occurredAt:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**redactionApplied:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**lowSignal:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContextSync.GetContextSyncHealth() -> *promptvmgosdk.GetContextSyncHealthResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetContextSyncHealthRequest{
        WorkspaceID: "workspaceId",
    }
client.ContextSync.GetContextSyncHealth(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContextSync.ListCaptureRepos() -> *promptvmgosdk.ListCaptureReposResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListCaptureReposRequest{
        WorkspaceID: "workspaceId",
    }
client.ContextSync.ListCaptureRepos(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContextSync.BulkPromoteCapturedSessions(request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.BulkPromoteCapturedSessionsRequest{
        IDs: []string{
            "ids",
        },
    }
client.ContextSync.BulkPromoteCapturedSessions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContextSync.BulkDiscardCapturedSessions(request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.BulkDiscardCapturedSessionsRequest{
        IDs: []string{
            "ids",
        },
    }
client.ContextSync.BulkDiscardCapturedSessions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContextSync.PromoteCapturedSession(ID) -> *promptvmgosdk.PromoteCapturedSessionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.PromoteCapturedSessionRequest{
        ID: "id",
    }
client.ContextSync.PromoteCapturedSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContextSync.DiscardCapturedSession(ID) -> *promptvmgosdk.DiscardCapturedSessionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DiscardCapturedSessionRequest{
        ID: "id",
    }
client.ContextSync.DiscardCapturedSession(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Contexts
<details><summary><code>client.Contexts.ListContextKinds() -> *promptvmgosdk.ListContextKindsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the catalogue of supported content kinds (prompt, skill, ...) with their full KindDefinition. Stable shape — agents can build typed adapters from this response without consulting docs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Contexts.ListContextKinds(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## MarketplaceSkillsFeed
<details><summary><code>client.MarketplaceSkillsFeed.GetMarketplaceSkillsFeed() -> *promptvmgosdk.GetMarketplaceSkillsFeedResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unauthenticated, paginated feed of all public skills. Use ?since=<ISO> to fetch only skills updated after a timestamp; use ?cursor= for continuation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetMarketplaceSkillsFeedRequest{}
client.MarketplaceSkillsFeed.GetMarketplaceSkillsFeed(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**since:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Marketplace - Browse
<details><summary><code>client.MarketplaceBrowse.ListMarketplaceListings() -> *promptvmgosdk.ListMarketplaceListingsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Paginated browse with optional search, category filter, and sort.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListMarketplaceListingsRequest{}
client.MarketplaceBrowse.ListMarketplaceListings(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` — Search query — matches title/description plus the underlying prompt description and SKILL.md when_to_use.
    
</dd>
</dl>

<dl>
<dd>

**categoryID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sellerID:** `*string` — Filter to a single creator's own listings (listings.seller_id). Powers the public creator profile.
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*promptvmgosdk.ListMarketplaceListingsRequestKind` — Filter listings by their underlying content_kind.
    
</dd>
</dl>

<dl>
<dd>

**tag:** `*string` — Match a single listing tag (case-insensitive).
    
</dd>
</dl>

<dl>
<dd>

**model:** `*string` — Narrow to listings recommended for one model, as `provider/model` (e.g. `anthropic/claude-opus-5`). Slug rather than id so the link is portable between environments.
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*promptvmgosdk.ListMarketplaceListingsRequestSort` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ranking:** `*promptvmgosdk.ListMarketplaceListingsRequestRanking` — Search ranking for `q`. `semantic` ranks by multilingual embedding similarity (503 when the embedding backend is down); `hybrid` fuses keyword + semantic via RRF and degrades to keyword with `meta.degraded: true`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceBrowse.GetMarketplaceListing(ListingID) -> *promptvmgosdk.GetMarketplaceListingResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetMarketplaceListingRequest{
        ListingID: "listingId",
    }
client.MarketplaceBrowse.GetMarketplaceListing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceBrowse.ListFeaturedMarketplaceListings() -> *promptvmgosdk.ListFeaturedMarketplaceListingsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.MarketplaceBrowse.ListFeaturedMarketplaceListings(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceBrowse.ListMarketplaceCategories() -> *promptvmgosdk.ListMarketplaceCategoriesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.MarketplaceBrowse.ListMarketplaceCategories(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceBrowse.ListMarketplaceAiModels() -> *promptvmgosdk.ListMarketplaceAiModelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The catalog creators pick from and buyers filter by. Returns active providers with their active models nested; a model without its own logo inherits its provider's.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.MarketplaceBrowse.ListMarketplaceAiModels(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceBrowse.ListMarketplaceContentTypes() -> *promptvmgosdk.ListMarketplaceContentTypesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns ENABLED marketplace content types ordered by sort order (public storefront nav). Disabled types are never exposed here; the full set lives at the platform-admin surface /api/v1/admin/content-types.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.MarketplaceBrowse.ListMarketplaceContentTypes(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceBrowse.GetMarketplaceFacets() -> *promptvmgosdk.GetMarketplaceFacetsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unauthenticated. Returns { total, kinds, categories } counted over ACTIVE listings using the same predicate as browse. Optional ?sellerId scopes to one creator. kinds includes every enabled content_type (0-count included) so the frontend can hide empty types.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetMarketplaceFacetsRequest{}
client.MarketplaceBrowse.GetMarketplaceFacets(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sellerID:** `*string` — Scope counts to a single creator's active listings (listings.seller_id).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Marketplace - Listings
<details><summary><code>client.MarketplaceListings.CreateMarketplaceListing(request) -> *promptvmgosdk.CreateMarketplaceListingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Publishes a prompt, skill, hook, collection, or directory to the marketplace. Skills/hooks are sourced via promptId (or the skillId/hookId aliases). Listings are free-only — priceCents must be 0. Requires a creator profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateMarketplaceListingRequest{
        Title: "title",
        Description: "description",
    }
client.MarketplaceListings.CreateMarketplaceListing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**skillID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**hookID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**collectionID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**title:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**categoryIDs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**priceCents:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**accessType:** `*promptvmgosdk.CreateMarketplaceListingRequestAccessType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceListings.ArchiveMarketplaceListing(ListingID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ArchiveMarketplaceListingRequest{
        ListingID: "listingId",
    }
client.MarketplaceListings.ArchiveMarketplaceListing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceListings.UpdateMarketplaceListing(ListingID, request) -> *promptvmgosdk.UpdateMarketplaceListingResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateMarketplaceListingRequest{
        ListingID: "listingId",
    }
client.MarketplaceListings.UpdateMarketplaceListing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**title:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**categoryIDs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.UpdateMarketplaceListingRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**priceCents:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**accessType:** `*promptvmgosdk.UpdateMarketplaceListingRequestAccessType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceListings.InitiateListingMediaUpload(ListingID, request) -> *promptvmgosdk.InitiateListingMediaUploadResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Owner-only. Validates the MIME type and size against the featured-media allowlist (images ≤ 5 MB: png/jpeg/webp/gif/svg; videos ≤ 50 MB: mp4/webm) and returns a presigned S3 PUT URL plus the deterministic storage key. Upload the bytes to the URL, then call the confirm endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.InitiateListingMediaUploadRequest{
        ListingID: "listingId",
        Kind: promptvmgosdk.InitiateListingMediaUploadRequestKindImage,
        ContentType: "contentType",
        SizeBytes: 1,
    }
client.MarketplaceListings.InitiateListingMediaUpload(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*promptvmgosdk.InitiateListingMediaUploadRequestKind` 
    
</dd>
</dl>

<dl>
<dd>

**contentType:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**sizeBytes:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceListings.ConfirmListingMediaUpload(ListingID, request) -> *promptvmgosdk.ConfirmListingMediaUploadResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Owner-only. Verifies the uploaded object exists in S3 with a valid size/content-type, then persists the key to the listing. Idempotent — re-confirming the same key returns the updated listing.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ConfirmListingMediaUploadRequest{
        ListingID: "listingId",
        Kind: promptvmgosdk.ConfirmListingMediaUploadRequestKindImage,
        Key: "key",
    }
client.MarketplaceListings.ConfirmListingMediaUpload(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*promptvmgosdk.ConfirmListingMediaUploadRequestKind` 
    
</dd>
</dl>

<dl>
<dd>

**key:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceListings.DeleteListingMedia(ListingID, Kind) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Owner-only. Clears the stored key for the given kind and best-effort deletes the S3 object. Deleting a video also clears its poster. Idempotent.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteListingMediaRequest{
        ListingID: "listingId",
        Kind: promptvmgosdk.DeleteListingMediaRequestKindImage.Ptr(),
    }
client.MarketplaceListings.DeleteListingMedia(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*promptvmgosdk.DeleteListingMediaRequestKind` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceListings.ClaimMarketplaceListing(ListingID, request) -> *promptvmgosdk.ClaimMarketplaceListingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Claims a free listing and copies the prompt/collection into the user workspace.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ClaimMarketplaceListingRequest{
        ListingID: "listingId",
        WorkspaceID: "workspaceId",
    }
client.MarketplaceListings.ClaimMarketplaceListing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Marketplace - Creator
<details><summary><code>client.MarketplaceCreator.GetMarketplaceCreatorProfile(UserID) -> *promptvmgosdk.GetMarketplaceCreatorProfileResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetMarketplaceCreatorProfileRequest{
        UserID: "userId",
    }
client.MarketplaceCreator.GetMarketplaceCreatorProfile(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceCreator.GetMarketplaceCreatorProfileByHandle(Handle) -> *promptvmgosdk.GetMarketplaceCreatorProfileByHandleResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resolves a cosmetic profile handle. Falls back to handles the creator previously held, in which case `meta.redirectTo` carries the current handle so the caller can issue a 301. Unrelated to `users.username`, which remains the install namespace for `creator/name` refs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetMarketplaceCreatorProfileByHandleRequest{
        Handle: "handle",
    }
client.MarketplaceCreator.GetMarketplaceCreatorProfileByHandle(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**handle:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceCreator.CheckMarketplaceCreatorHandleAvailability() -> *promptvmgosdk.CheckMarketplaceCreatorHandleAvailabilityResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CheckMarketplaceCreatorHandleAvailabilityRequest{
        Handle: "handle",
    }
client.MarketplaceCreator.CheckMarketplaceCreatorHandleAvailability(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**handle:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceCreator.BrowseMarketplaceCreators() -> *promptvmgosdk.BrowseMarketplaceCreatorsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.BrowseMarketplaceCreatorsRequest{}
client.MarketplaceCreator.BrowseMarketplaceCreators(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*promptvmgosdk.BrowseMarketplaceCreatorsRequestSort` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceCreator.GetMyMarketplaceCreatorProfile() -> *promptvmgosdk.GetMyMarketplaceCreatorProfileResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.MarketplaceCreator.GetMyMarketplaceCreatorProfile(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceCreator.UpdateMyMarketplaceCreatorProfile(request) -> *promptvmgosdk.UpdateMyMarketplaceCreatorProfileResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateMyMarketplaceCreatorProfileRequest{}
client.MarketplaceCreator.UpdateMyMarketplaceCreatorProfile(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bio:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**website:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**socialLinks:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**avatarURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**profileHandle:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceCreator.CreateMarketplaceCreatorProfile(request) -> *promptvmgosdk.CreateMarketplaceCreatorProfileResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateMarketplaceCreatorProfileRequest{
        Bio: "bio",
    }
client.MarketplaceCreator.CreateMarketplaceCreatorProfile(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**bio:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**website:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**socialLinks:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**avatarURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**profileHandle:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceCreator.ClaimMarketplaceCreatorProfile(CreatorUserID, request) -> *promptvmgosdk.ClaimMarketplaceCreatorProfileResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ClaimMarketplaceCreatorProfileRequest{
        CreatorUserID: "creatorUserId",
        Reason: "reason",
        ProofURL: "proofUrl",
    }
client.MarketplaceCreator.ClaimMarketplaceCreatorProfile(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creatorUserID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**reason:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**proofURL:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**message:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## MarketplaceStats
<details><summary><code>client.MarketplaceStats.GetMarketplaceStats() -> *promptvmgosdk.GetMarketplaceStatsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unauthenticated. Returns { totalListings, totalInstalls, totalCreators } over active listings. totalInstalls is the live SUM(downloadCount).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.MarketplaceStats.GetMarketplaceStats(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## MarketplaceResolve
<details><summary><code>client.MarketplaceResolve.ResolveInstallRef() -> *promptvmgosdk.ResolveInstallRefResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unauthenticated. Resolves `creator/name`, a legacy vanity slug, a `name-<uuid8>` file slug, or an unambiguous bare name. Returns { ref, kind, content, … } for skill/agent/command/hook/mcp/settings/prompt. 404 for missing/non-public; 409 AMBIGUOUS_REF + candidates when a bare name is owned by multiple creators.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ResolveInstallRefRequest{
        Ref: "ref",
    }
client.MarketplaceResolve.ResolveInstallRef(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ref:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceResolve.RecordResolveInstall() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unauthenticated. Resolves the ref and atomically increments the tied active listing's downloadCount (never importCount). 204 even when the resolved artifact has no active listing. 404 for missing/non-public; 409 AMBIGUOUS_REF.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RecordResolveInstallRequest{
        Ref: "ref",
    }
client.MarketplaceResolve.RecordResolveInstall(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ref:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sharing
<details><summary><code>client.Sharing.AccessSharedPrompt(Token) -> *promptvmgosdk.AccessSharedPromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Public endpoint — no authentication required. Each successful access increments the link's useCount (counted against maxUses). Pass meta=1 (or meta=true) for a metadata-only read: the response and all state checks (revoked / expired / password / max-use) are identical, but useCount is NOT incremented — intended for server-side metadata/OG/share-image fetches so they never burn maxUses.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AccessSharedPromptRequest{
        Token: "token",
    }
client.Sharing.AccessSharedPrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**token:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**password:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**meta:** `*promptvmgosdk.AccessSharedPromptRequestMeta` — When "1" or "true", performs a metadata-only read that does not increment useCount.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sharing.SharePrompt(PromptID, request) -> *promptvmgosdk.SharePromptResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.SharePromptRequest{
        PromptID: "promptId",
        Permission: promptvmgosdk.SharePromptRequestPermissionView,
    }
client.Sharing.SharePrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**groupID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**permission:** `*promptvmgosdk.SharePromptRequestPermission` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sharing.ListPromptCollaborators(PromptID) -> *promptvmgosdk.ListPromptCollaboratorsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListPromptCollaboratorsRequest{
        PromptID: "promptId",
    }
client.Sharing.ListPromptCollaborators(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sharing.RevokePromptCollaborator(PromptID, CollaboratorID) -> *promptvmgosdk.RevokePromptCollaboratorResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RevokePromptCollaboratorRequest{
        PromptID: "promptId",
        CollaboratorID: "collaboratorId",
    }
client.Sharing.RevokePromptCollaborator(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**collaboratorID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sharing.ListPromptShareLinks(PromptID) -> *promptvmgosdk.ListPromptShareLinksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns non-revoked share links for the prompt. Requires share permission.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListPromptShareLinksRequest{
        PromptID: "promptId",
    }
client.Sharing.ListPromptShareLinks(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sharing.RevokePromptShareLink(PromptID, LinkID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Sets revokedAt/revokedBy. Idempotent — revoking an already-revoked link returns 204.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RevokePromptShareLinkRequest{
        PromptID: "promptId",
        LinkID: "linkId",
    }
client.Sharing.RevokePromptShareLink(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**linkID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sharing.CreatePromptShareLink(PromptID, request) -> *promptvmgosdk.CreatePromptShareLinkResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreatePromptShareLinkRequest{
        PromptID: "promptId",
    }
client.Sharing.CreatePromptShareLink(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` — Optional. Replays the original 2xx response for 24h on retry with the same key + same body. A different body with the same key returns 422 idempotency_conflict.
    
</dd>
</dl>

<dl>
<dd>

**permission:** `*promptvmgosdk.CreatePromptShareLinkRequestPermission` 
    
</dd>
</dl>

<dl>
<dd>

**password:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**expiresInHours:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**maxUses:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**versionNumber:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**label:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*promptvmgosdk.CreatePromptShareLinkRequestKind` 
    
</dd>
</dl>

<dl>
<dd>

**version:** `*promptvmgosdk.CreatePromptShareLinkRequestVersion` 
    
</dd>
</dl>

<dl>
<dd>

**values:** `map[string]string` 
    
</dd>
</dl>

<dl>
<dd>

**resolveVariables:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sharing.ImportSharedPrompt(Token, request) -> *promptvmgosdk.ImportSharedPromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Copies the prompt behind a share link into the caller's workspace (body.workspaceId, or the caller's default/first writable workspace when omitted). Imports the SHARED version — the pinned version's content for pinned links, else the latest. Link-state semantics match GET /share/:token (404/410/401); the link's useCount is incremented only when the import succeeds, never on a 402/403/404 failure.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ImportSharedPromptRequest{
        Token: "token",
    }
client.Sharing.ImportSharedPrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**token:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` — Optional. Replays the original 2xx response for 24h on retry with the same key + same body. A different body with the same key returns 422 idempotency_conflict.
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**password:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## API Keys
<details><summary><code>client.APIKeys.ListAPIKeys() -> *promptvmgosdk.ListAPIKeysResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a paginated list of API keys for the authenticated user. Results include metadata but never expose the secret key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListAPIKeysRequest{
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.APIKeys.ListAPIKeys(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — Cursor for pagination (base64-encoded)
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of items per page
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.ListAPIKeysRequestStatus` — Filter by API key status
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.CreateAPIKey(request) -> *promptvmgosdk.CreateAPIKeyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Generates a new API key for the authenticated user. Requires Pro or Enterprise tier. The secret key is only returned once during creation - store it securely.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateAPIKeyRequest{
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
        Name: "Production API Key",
        Scopes: []promptvmgosdk.CreateAPIKeyRequestScopesItem{
            promptvmgosdk.CreateAPIKeyRequestScopesItemRead,
            promptvmgosdk.CreateAPIKeyRequestScopesItemWrite,
        },
    }
client.APIKeys.CreateAPIKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Human-readable name for the API key
    
</dd>
</dl>

<dl>
<dd>

**scopes:** `[]*promptvmgosdk.CreateAPIKeyRequestScopesItem` — Permission scopes for the API key. Requires Pro or Enterprise tier.
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*time.Time` — Optional expiration date. If set, the key becomes invalid after this time.
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `*string` — Workspace binding (UUID) for capture-scoped keys. REQUIRED when scopes is ["capture"]; not accepted for other scopes. The minted key may only ingest Context Sync captures for this workspace.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.GetAPIKey(APIKeyID) -> *promptvmgosdk.GetAPIKeyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves detailed information about a specific API key. Does not include the secret key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetAPIKeyRequest{
        APIKeyID: "apiKeyId",
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.APIKeys.GetAPIKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — UUID of the API key
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.RevokeAPIKey(APIKeyID, request) -> *promptvmgosdk.RevokeAPIKeyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently revokes an API key. Revoked keys cannot be used for API access but remain in the system for audit purposes. This action cannot be undone.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RevokeAPIKeyRequest{
        APIKeyID: "apiKeyId",
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.APIKeys.RevokeAPIKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — UUID of the API key to revoke
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*string` — Optional reason for revocation
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.UpdateAPIKey(APIKeyID, request) -> *promptvmgosdk.UpdateAPIKeyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the metadata (name) of an existing API key. Cannot modify scopes after creation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateAPIKeyRequest{
        APIKeyID: "apiKeyId",
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.APIKeys.UpdateAPIKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — UUID of the API key
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — New name for the API key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.GetAPIKeyUsage(APIKeyID) -> *promptvmgosdk.GetAPIKeyUsageResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves usage statistics for a specific API key including request counts, error rates, and latency metrics.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetAPIKeyUsageRequest{
        APIKeyID: "apiKeyId",
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.APIKeys.GetAPIKeyUsage(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — UUID of the API key
    
</dd>
</dl>

<dl>
<dd>

**period:** `*promptvmgosdk.GetAPIKeyUsageRequestPeriod` — Time period for statistics aggregation
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.RotateAPIKey(APIKeyID, request) -> *promptvmgosdk.RotateAPIKeyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Generates a new secret for an existing API key. The public key remains unchanged. The old secret remains valid for a configurable grace period (default: 24 hours). Requires JWT auth — cannot rotate using the key itself.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RotateAPIKeyRequest{
        APIKeyID: "apiKeyId",
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.APIKeys.RotateAPIKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**apiKeyID:** `string` — UUID of the API key to rotate
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>

<dl>
<dd>

**gracePeriodHours:** `*int` — Hours the old secret remains valid (0-168, default 24). Public key is unchanged.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## MarketplaceSocial
<details><summary><code>client.MarketplaceSocial.FollowACreator(CreatorUserID) -> *promptvmgosdk.PostAPIV1MarketplaceCreatorCreatorUserIDFollowResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.PostAPIV1MarketplaceCreatorCreatorUserIDFollowRequest{
        CreatorUserID: "creatorUserId",
    }
client.MarketplaceSocial.FollowACreator(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creatorUserID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceSocial.UnfollowACreator(CreatorUserID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteAPIV1MarketplaceCreatorCreatorUserIDFollowRequest{
        CreatorUserID: "creatorUserId",
    }
client.MarketplaceSocial.UnfollowACreator(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creatorUserID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceSocial.ListCreatorsIFollow() -> *promptvmgosdk.GetAPIV1MarketplaceCreatorMeFollowingResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetAPIV1MarketplaceCreatorMeFollowingRequest{}
client.MarketplaceSocial.ListCreatorsIFollow(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceSocial.FollowedCreatorListingFeed() -> *promptvmgosdk.GetAPIV1MarketplaceCreatorMeFeedResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetAPIV1MarketplaceCreatorMeFeedRequest{}
client.MarketplaceSocial.FollowedCreatorListingFeed(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceSocial.ListACreatorsFollowers(CreatorUserID) -> *promptvmgosdk.GetAPIV1MarketplaceCreatorCreatorUserIDFollowersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetAPIV1MarketplaceCreatorCreatorUserIDFollowersRequest{
        CreatorUserID: "creatorUserId",
    }
client.MarketplaceSocial.ListACreatorsFollowers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creatorUserID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Marketplace
<details><summary><code>client.Marketplace.SetListingRecommendedModels(ListingID, request) -> *promptvmgosdk.SetListingRecommendedModelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replace-all, in the order given. Up to 10. Independent of the resource version this listing was published from — editing the version never reaches back into a live listing.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.SetListingRecommendedModelsRequest{
        ListingID: "listingId",
        ModelIDs: []string{
            "modelIds",
        },
    }
client.Marketplace.SetListingRecommendedModels(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**modelIDs:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminListContentTypes() -> *promptvmgosdk.AdminListContentTypesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns ALL marketplace content types including disabled ones, ordered by sort order. Requires internal_marketplace:manage.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Marketplace.AdminListContentTypes(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminCreateContentType(request) -> *promptvmgosdk.AdminCreateContentTypeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new marketplace content type. Slug must be unique and match ^[a-z0-9][a-z0-9-]{0,31}$. Setting enabled=true requires a registered code kind (422 otherwise). Requires internal_marketplace:manage.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminCreateContentTypeRequest{
        Slug: "slug",
        Label: "label",
        PluralLabel: "pluralLabel",
        Icon: "icon",
        Color: "color",
    }
client.Marketplace.AdminCreateContentType(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — Stable key; matches the code kind-registry slug.
    
</dd>
</dl>

<dl>
<dd>

**label:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**pluralLabel:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**icon:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**color:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**mediaAllowed:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**installVerb:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**descriptionOneLiner:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**settingsSchema:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminReorderContentTypes(request) -> *promptvmgosdk.AdminReorderContentTypesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Bulk-updates sort_order for a set of slugs in one transaction. Every slug must exist. Requires internal_marketplace:manage.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminReorderContentTypesRequest{
        Items: []*promptvmgosdk.AdminReorderContentTypesRequestItemsItem{
            &promptvmgosdk.AdminReorderContentTypesRequestItemsItem{
                Slug: "slug",
                SortOrder: 1,
            },
        },
    }
client.Marketplace.AdminReorderContentTypes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**items:** `[]*promptvmgosdk.AdminReorderContentTypesRequestItemsItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminDeleteContentType(Slug) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a content type. Allowed ONLY for non-built-in slugs with no dependent content — otherwise 409 (disable instead). Requires internal_marketplace:manage.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminDeleteContentTypeRequest{
        Slug: "slug",
    }
client.Marketplace.AdminDeleteContentType(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminUpdateContentType(Slug, request) -> *promptvmgosdk.AdminUpdateContentTypeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Partially updates a content type's presentation/config fields. The slug is immutable. Flipping enabled=true requires a registered code kind (422 otherwise). Requires internal_marketplace:manage.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminUpdateContentTypeRequest{
        Slug: "slug",
    }
client.Marketplace.AdminUpdateContentType(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**label:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**pluralLabel:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**color:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**mediaAllowed:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**installVerb:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**descriptionOneLiner:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**settingsSchema:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminListAiModelProviders() -> *promptvmgosdk.AdminListAiModelProvidersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

All providers including retired ones, ordered by sort order.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Marketplace.AdminListAiModelProviders(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminCreateAiModelProvider(request) -> *promptvmgosdk.AdminCreateAiModelProviderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Logo and website URLs are validated (https, no credentials, no private hosts) and are never fetched by the server.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminCreateAiModelProviderRequest{
        Slug: "slug",
        Name: "name",
    }
client.Marketplace.AdminCreateAiModelProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**logoURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**websiteURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminReorderAiModelProviders(request) -> *promptvmgosdk.AdminReorderAiModelProvidersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminReorderAiModelProvidersRequest{
        Items: []*promptvmgosdk.AdminReorderAiModelProvidersRequestItemsItem{
            &promptvmgosdk.AdminReorderAiModelProvidersRequestItemsItem{
                ID: "id",
                SortOrder: 1,
            },
        },
    }
client.Marketplace.AdminReorderAiModelProviders(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**items:** `[]*promptvmgosdk.AdminReorderAiModelProvidersRequestItemsItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminDeleteAiModelProvider(Slug) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Refused with 409 while the provider still has models. Retire it instead — retiring hides it from pickers without disturbing existing recommendations.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminDeleteAiModelProviderRequest{
        Slug: "slug",
    }
client.Marketplace.AdminDeleteAiModelProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminUpdateAiModelProvider(Slug, request) -> *promptvmgosdk.AdminUpdateAiModelProviderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Slug is immutable. Set `isActive: false` to retire a provider.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminUpdateAiModelProviderRequest{
        Slug: "slug",
    }
client.Marketplace.AdminUpdateAiModelProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**logoURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**websiteURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminListAiModels() -> *promptvmgosdk.AdminListAiModelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

All models including retired ones. Optionally scoped to one provider.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminListAiModelsRequest{}
client.Marketplace.AdminListAiModels(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**providerID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminCreateAiModel(request) -> *promptvmgosdk.AdminCreateAiModelResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminCreateAiModelRequest{
        ProviderID: "providerId",
        Slug: "slug",
        Name: "name",
    }
client.Marketplace.AdminCreateAiModel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**providerID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**logoURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**modality:** `*promptvmgosdk.AdminCreateAiModelRequestModality` 
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminReorderAiModels(request) -> *promptvmgosdk.AdminReorderAiModelsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminReorderAiModelsRequest{
        Items: []*promptvmgosdk.AdminReorderAiModelsRequestItemsItem{
            &promptvmgosdk.AdminReorderAiModelsRequestItemsItem{
                ID: "id",
                SortOrder: 1,
            },
        },
    }
client.Marketplace.AdminReorderAiModels(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**items:** `[]*promptvmgosdk.AdminReorderAiModelsRequestItemsItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminDeleteAiModel(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Refused with 409 while any listing OR version recommends it — a private draft breaks as surely as a live listing, and less visibly. Retire it instead.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminDeleteAiModelRequest{
        ID: "id",
    }
client.Marketplace.AdminDeleteAiModel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminUpdateAiModel(ID, request) -> *promptvmgosdk.AdminUpdateAiModelResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Slug is immutable. Set `isActive: false` to retire — the model disappears from pickers while every listing and version that already recommends it keeps rendering it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminUpdateAiModelRequest{
        ID: "id",
    }
client.Marketplace.AdminUpdateAiModel(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**providerID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**logoURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**modality:** `*promptvmgosdk.AdminUpdateAiModelRequestModality` 
    
</dd>
</dl>

<dl>
<dd>

**sortOrder:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminListCategories() -> *promptvmgosdk.AdminListCategoriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

All categories including inactive ones, with their parent relationships.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Marketplace.AdminListCategories(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminCreateCategory(request) -> *promptvmgosdk.AdminCreateCategoryResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminCreateCategoryRequest{
        Name: "name",
        Slug: "slug",
    }
client.Marketplace.AdminCreateCategory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**slug:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**parentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**displayOrder:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminReorderCategories(request) -> *promptvmgosdk.AdminReorderCategoriesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminReorderCategoriesRequest{
        Items: []*promptvmgosdk.AdminReorderCategoriesRequestItemsItem{
            &promptvmgosdk.AdminReorderCategoriesRequestItemsItem{
                ID: "id",
                DisplayOrder: 1,
            },
        },
    }
client.Marketplace.AdminReorderCategories(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**items:** `[]*promptvmgosdk.AdminReorderCategoriesRequestItemsItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminDeleteCategory(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Refused with 409 while the category has subcategories or listings. The join is ON DELETE CASCADE, so an unguarded delete would silently strip the category from every listing using it — deactivate instead.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminDeleteCategoryRequest{
        ID: "id",
    }
client.Marketplace.AdminDeleteCategory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Marketplace.AdminUpdateCategory(ID, request) -> *promptvmgosdk.AdminUpdateCategoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Slug is immutable — it is the stable half of a shared browse link. A parent that is a descendant of this category is rejected, so the tree cannot cycle.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AdminUpdateCategoryRequest{
        ID: "id",
    }
client.Marketplace.AdminUpdateCategory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**parentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**displayOrder:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Marketplace - Creator Dashboard
<details><summary><code>client.MarketplaceCreatorDashboard.ListMarketplaceCreatorListings() -> *promptvmgosdk.ListMarketplaceCreatorListingsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListMarketplaceCreatorListingsRequest{}
client.MarketplaceCreatorDashboard.ListMarketplaceCreatorListings(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**status:** `*promptvmgosdk.ListMarketplaceCreatorListingsRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceCreatorDashboard.ListMarketplaceCreatorSubscribers() -> *promptvmgosdk.ListMarketplaceCreatorSubscribersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListMarketplaceCreatorSubscribersRequest{}
client.MarketplaceCreatorDashboard.ListMarketplaceCreatorSubscribers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Marketplace - Subscriptions
<details><summary><code>client.MarketplaceSubscriptions.SubscribeToMarketplaceCreator(CreatorUserID) -> *promptvmgosdk.SubscribeToMarketplaceCreatorResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.SubscribeToMarketplaceCreatorRequest{
        CreatorUserID: "creatorUserId",
    }
client.MarketplaceSubscriptions.SubscribeToMarketplaceCreator(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creatorUserID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceSubscriptions.UnsubscribeFromMarketplaceCreator(CreatorUserID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UnsubscribeFromMarketplaceCreatorRequest{
        CreatorUserID: "creatorUserId",
    }
client.MarketplaceSubscriptions.UnsubscribeFromMarketplaceCreator(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**creatorUserID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## MarketplaceRatings
<details><summary><code>client.MarketplaceRatings.ListRatingsForAListing(ListingID) -> *promptvmgosdk.GetAPIV1MarketplaceListingsListingIDRatingsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetAPIV1MarketplaceListingsListingIDRatingsRequest{
        ListingID: "listingId",
    }
client.MarketplaceRatings.ListRatingsForAListing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*promptvmgosdk.GetAPIV1MarketplaceListingsListingIDRatingsRequestSort` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceRatings.CreateARatingOnAPurchasedListing(ListingID, request) -> *promptvmgosdk.PostAPIV1MarketplaceListingsListingIDRatingsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.PostAPIV1MarketplaceListingsListingIDRatingsRequest{
        ListingID: "listingId",
        Score: 1,
    }
client.MarketplaceRatings.CreateARatingOnAPurchasedListing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**score:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**review:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceRatings.GetTheCallersRatingOnAListing(ListingID) -> *promptvmgosdk.GetAPIV1MarketplaceListingsListingIDRatingsMeResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetAPIV1MarketplaceListingsListingIDRatingsMeRequest{
        ListingID: "listingId",
    }
client.MarketplaceRatings.GetTheCallersRatingOnAListing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceRatings.DeleteMyRating(RatingID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteAPIV1MarketplaceRatingsRatingIDRequest{
        RatingID: "ratingId",
    }
client.MarketplaceRatings.DeleteMyRating(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ratingID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceRatings.UpdateMyRating(RatingID, request) -> *promptvmgosdk.PatchAPIV1MarketplaceRatingsRatingIDResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.PatchAPIV1MarketplaceRatingsRatingIDRequest{
        RatingID: "ratingId",
    }
client.MarketplaceRatings.UpdateMyRating(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ratingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**score:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**review:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceRatings.VoteHelpfulUnhelpfulOnARating(RatingID, request) -> *promptvmgosdk.PostAPIV1MarketplaceRatingsRatingIDVoteResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.PostAPIV1MarketplaceRatingsRatingIDVoteRequest{
        RatingID: "ratingId",
        VoteType: promptvmgosdk.PostAPIV1MarketplaceRatingsRatingIDVoteRequestVoteTypeUpvote,
    }
client.MarketplaceRatings.VoteHelpfulUnhelpfulOnARating(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ratingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**voteType:** `*promptvmgosdk.PostAPIV1MarketplaceRatingsRatingIDVoteRequestVoteType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceRatings.RemoveMyVoteFromARating(RatingID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteAPIV1MarketplaceRatingsRatingIDVoteRequest{
        RatingID: "ratingId",
    }
client.MarketplaceRatings.RemoveMyVoteFromARating(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ratingID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## MarketplaceComments
<details><summary><code>client.MarketplaceComments.ListThreadedCommentsOnAListing(ListingID) -> *promptvmgosdk.GetAPIV1MarketplaceListingsListingIDCommentsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetAPIV1MarketplaceListingsListingIDCommentsRequest{
        ListingID: "listingId",
    }
client.MarketplaceComments.ListThreadedCommentsOnAListing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceComments.CreateACommentOrReplyOnAListing(ListingID, request) -> *promptvmgosdk.PostAPIV1MarketplaceListingsListingIDCommentsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.PostAPIV1MarketplaceListingsListingIDCommentsRequest{
        ListingID: "listingId",
        Content: "content",
    }
client.MarketplaceComments.CreateACommentOrReplyOnAListing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**listingID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**content:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**parentID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MarketplaceComments.SoftDeleteMyComment(CommentID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteAPIV1MarketplaceCommentsCommentIDRequest{
        CommentID: "commentId",
    }
client.MarketplaceComments.SoftDeleteMyComment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Onboarding
<details><summary><code>client.Onboarding.GetOnboardingStatus() -> *promptvmgosdk.GetOnboardingStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return the dashboard onboarding checklist status for the current user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Onboarding.GetOnboardingStatus(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Onboarding.UpdateOnboardingStep(request) -> *promptvmgosdk.UpdateOnboardingStepResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resolve or unresolve an onboarding checklist step.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateOnboardingStepRequest{
        Step: promptvmgosdk.UpdateOnboardingStepRequestStepFirstPrompt,
        Completed: true,
    }
client.Onboarding.UpdateOnboardingStep(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**step:** `*promptvmgosdk.UpdateOnboardingStepRequestStep` 
    
</dd>
</dl>

<dl>
<dd>

**completed:** `bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Onboarding.DismissOnboarding() -> *promptvmgosdk.DismissOnboardingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Dismiss the dashboard onboarding checklist.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Onboarding.DismissOnboarding(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users
<details><summary><code>client.Users.DeleteAccount() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Soft-deletes the caller: every organization the user owns is deleted (workspaces and files soft-deleted, members removed, API keys and sessions revoked, Stripe subscription cancelled immediately), memberships in other organizations are removed, and all of the user's sessions and API keys are revoked. Irreversible from the client perspective; data is retained for 30 days for compliance before purge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.DeleteAccount(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Settings
<details><summary><code>client.Settings.GetSettingsOverview() -> *promptvmgosdk.GetSettingsOverviewResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns workspace + API-key counts, derived admin-task status, and a top-5 recent-activity slice for the active organization. Recent activity is server-scoped by role: owners/admins see all org events, members/viewers see only their own.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetSettingsOverviewRequest{
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.Settings.GetSettingsOverview(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.GetAdminTasks() -> *promptvmgosdk.GetAdminTasksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the unified admin-task list (billing method on file, first API key created). Tasks are derived from observable state, not stored — there is no toggle endpoint.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetAdminTasksRequest{
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.Settings.GetAdminTasks(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.ListAuditLogs() -> *promptvmgosdk.ListAuditLogsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Server-scoped by role: owners/admins see all org events; members/viewers see only events where they are the actor.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListAuditLogsRequest{
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.Settings.ListAuditLogs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.ListOrgTags() -> *promptvmgosdk.ListOrgTagsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListOrgTagsRequest{
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.Settings.ListOrgTags(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.CreateOrgTag(request) -> *promptvmgosdk.CreateOrgTagResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateOrgTagRequest{
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
        Name: "name",
        Category: promptvmgosdk.CreateOrgTagRequestCategoryEngineering,
        Protected: true,
    }
client.Settings.CreateOrgTag(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**alias:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**category:** `*promptvmgosdk.CreateOrgTagRequestCategory` 
    
</dd>
</dl>

<dl>
<dd>

**protected:** `bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.DeleteOrgTag(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteOrgTagRequest{
        ID: "id",
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.Settings.DeleteOrgTag(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.UpdateOrgTag(ID, request) -> *promptvmgosdk.UpdateOrgTagResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateOrgTagRequest{
        ID: "id",
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.Settings.UpdateOrgTag(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**alias:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**category:** `*promptvmgosdk.UpdateOrgTagRequestCategory` 
    
</dd>
</dl>

<dl>
<dd>

**protected:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.GetSecurityStatus() -> *promptvmgosdk.GetSecurityStatusResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetSecurityStatusRequest{
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
    }
client.Settings.GetSecurityStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.SetupTotp() -> *promptvmgosdk.SetupTotpResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Settings.SetupTotp(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.VerifyTotp(request) -> *promptvmgosdk.VerifyTotpResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.VerifyTotpRequest{
        Code: "code",
    }
client.Settings.VerifyTotp(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**code:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.DisableTotp(request) -> *promptvmgosdk.DisableTotpResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DisableTotpRequest{
        Code: "code",
    }
client.Settings.DisableTotp(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**code:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.RegenerateRecoveryCodes(request) -> *promptvmgosdk.RegenerateRecoveryCodesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RegenerateRecoveryCodesRequest{
        Code: "code",
    }
client.Settings.RegenerateRecoveryCodes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**code:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.ChangePassword(request) -> *promptvmgosdk.ChangePasswordResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ChangePasswordRequest{
        CurrentPassword: "currentPassword",
        NewPassword: "newPassword",
    }
client.Settings.ChangePassword(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**currentPassword:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**newPassword:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.UpdateOrgMfaPolicy(request) -> *promptvmgosdk.UpdateOrgMfaPolicyResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateOrgMfaPolicyRequest{
        OrgID: promptvmgosdk.String(
            "018e4a3b-0000-0000-0000-000000000001",
        ),
        RequireMfa: true,
    }
client.Settings.UpdateOrgMfaPolicy(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `*string` — Active organization identifier. UUID is the primary form; slug is accepted as a deprecated legacy fallback (logs `billing.org_id.legacy_slug`). Frontend resolves slug → UUID locally via /auth/me and SHOULD send the UUID.
    
</dd>
</dl>

<dl>
<dd>

**requireMfa:** `bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.SendEmailOtp() -> *promptvmgosdk.SendEmailOtpResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Settings.SendEmailOtp(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.VerifyEmailOtp(request) -> *promptvmgosdk.VerifyEmailOtpResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.VerifyEmailOtpRequest{
        Code: "code",
    }
client.Settings.VerifyEmailOtp(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**code:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Settings.DisableEmailOtp(request) -> *promptvmgosdk.DisableEmailOtpResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DisableEmailOtpRequest{
        Code: "code",
    }
client.Settings.DisableEmailOtp(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**code:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organizations
<details><summary><code>client.Organizations.ListOrganizations() -> *promptvmgosdk.ListOrganizationsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.ListOrganizations(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.CreateOrganization(request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateOrganizationRequest{
        Name: "name",
    }
client.Organizations.CreateOrganization(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.DeleteOrganization(OrgID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Soft-deletes the organization and cascades: workspaces and files are soft-deleted, members removed, API keys and org-scoped sessions revoked, the active Stripe subscription is cancelled immediately, and an org.deleted audit log is written. Owner only; JWT session only — API-key and CLI-token callers are rejected with 403 regardless of role. Personal organizations cannot be deleted standalone (422) — they are deleted together with the account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteOrganizationRequest{
        OrgID: "orgId",
    }
client.Organizations.DeleteOrganization(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.AcceptOrganizationInvitation(Token) -> *promptvmgosdk.AcceptOrganizationInvitationResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AcceptOrganizationInvitationRequest{
        Token: "token",
    }
client.Organizations.AcceptOrganizationInvitation(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**token:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.ListOrganizationWorkspaces(OrgID) -> *promptvmgosdk.ListOrganizationWorkspacesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all workspaces in the org visible to the caller. Org owners and admins see every workspace; regular members see workspaces they own, are a member of, or that have public/internal visibility.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListOrganizationWorkspacesRequest{
        OrgID: "orgId",
    }
client.Organizations.ListOrganizationWorkspaces(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.ListOrganizationMembers(OrgID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListOrganizationMembersRequest{
        OrgID: "orgId",
    }
client.Organizations.ListOrganizationMembers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `*promptvmgosdk.ListOrganizationMembersRequestRole` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.ListOrganizationMembersRequestStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.GetOrganizationPermissions(OrgID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetOrganizationPermissionsRequest{
        OrgID: "orgId",
    }
client.Organizations.GetOrganizationPermissions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.UpdateOrganizationPermissions(OrgID, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateOrganizationPermissionsRequest{
        OrgID: "orgId",
        Updates: []*promptvmgosdk.UpdateOrganizationPermissionsRequestUpdatesItem{
            &promptvmgosdk.UpdateOrganizationPermissionsRequestUpdatesItem{
                Permission: "permission",
                Role: promptvmgosdk.UpdateOrganizationPermissionsRequestUpdatesItemRoleOwner,
                Enabled: true,
            },
        },
    }
client.Organizations.UpdateOrganizationPermissions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**updates:** `[]*promptvmgosdk.UpdateOrganizationPermissionsRequestUpdatesItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.ListOrganizationRoles(OrgID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListOrganizationRolesRequest{
        OrgID: "orgId",
    }
client.Organizations.ListOrganizationRoles(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.CreateOrganizationRole(OrgID, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateOrganizationRoleRequest{
        OrgID: "orgId",
        Name: "name",
        BaseRole: promptvmgosdk.CreateOrganizationRoleRequestBaseRoleOwner,
    }
client.Organizations.CreateOrganizationRole(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**baseRole:** `*promptvmgosdk.CreateOrganizationRoleRequestBaseRole` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**color:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.RemoveOrganizationMember(OrgID, MemberID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RemoveOrganizationMemberRequest{
        OrgID: "orgId",
        MemberID: "memberId",
    }
client.Organizations.RemoveOrganizationMember(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.UpdateOrganizationMemberRole(OrgID, MemberID, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateOrganizationMemberRoleRequest{
        OrgID: "orgId",
        MemberID: "memberId",
        Role: promptvmgosdk.UpdateOrganizationMemberRoleRequestRoleOwner,
    }
client.Organizations.UpdateOrganizationMemberRole(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `*promptvmgosdk.UpdateOrganizationMemberRoleRequestRole` 
    
</dd>
</dl>

<dl>
<dd>

**customRoleID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.ListOrganizationInvitations(OrgID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListOrganizationInvitationsRequest{
        OrgID: "orgId",
    }
client.Organizations.ListOrganizationInvitations(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.ListOrganizationInvitationsRequestStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.CreateOrganizationInvitation(OrgID, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateOrganizationInvitationRequest{
        OrgID: "orgId",
        Email: "email",
        Role: promptvmgosdk.CreateOrganizationInvitationRequestRoleOwner,
    }
client.Organizations.CreateOrganizationInvitation(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `*promptvmgosdk.CreateOrganizationInvitationRequestRole` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.RevokeOrganizationInvitation(OrgID, InvitationID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RevokeOrganizationInvitationRequest{
        OrgID: "orgId",
        InvitationID: "invitationId",
    }
client.Organizations.RevokeOrganizationInvitation(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**invitationID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.ResendOrganizationInvitation(OrgID, InvitationID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ResendOrganizationInvitationRequest{
        OrgID: "orgId",
        InvitationID: "invitationId",
    }
client.Organizations.ResendOrganizationInvitation(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**invitationID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.DeleteOrganizationRole(OrgID, RoleID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteOrganizationRoleRequest{
        OrgID: "orgId",
        RoleID: "roleId",
    }
client.Organizations.DeleteOrganizationRole(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**roleID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.UpdateOrganizationRole(OrgID, RoleID, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateOrganizationRoleRequest{
        OrgID: "orgId",
        RoleID: "roleId",
    }
client.Organizations.UpdateOrganizationRole(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orgID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**roleID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**baseRole:** `*promptvmgosdk.UpdateOrganizationRoleRequestBaseRole` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**color:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Prompts
<details><summary><code>client.Prompts.ListPrompts() -> *promptvmgosdk.ListPromptsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a paginated list of prompts in a workspace.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListPromptsRequest{
        WorkspaceID: "workspaceId",
    }
client.Prompts.ListPrompts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.ListPromptsRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*promptvmgosdk.ListPromptsRequestKind` 
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Prompts.CreatePrompt(request) -> *promptvmgosdk.CreatePromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new prompt with an initial version (v1). Requires edit permission on the target workspace.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreatePromptRequest{
        Name: "name",
        WorkspaceID: "workspaceId",
        Content: "content",
    }
client.Prompts.CreatePrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` — Optional. Replays the original 2xx response for 24h on retry with the same key + same body. A different body with the same key returns 422 idempotency_conflict.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**content:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**systemPrompt:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**kind:** `*promptvmgosdk.CreatePromptRequestKind` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.CreatePromptRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**variablesSchema:** `map[string]*promptvmgosdk.CreatePromptRequestVariablesSchemaValue` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Prompts.GetPrompt(PromptID) -> *promptvmgosdk.GetPromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a prompt with its current version content.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetPromptRequest{
        PromptID: "promptId",
    }
client.Prompts.GetPrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Prompts.DeletePrompt(PromptID) -> *promptvmgosdk.DeletePromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Soft-deletes a prompt by setting deletedAt on its file record.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeletePromptRequest{
        PromptID: "promptId",
    }
client.Prompts.DeletePrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Prompts.UpdatePrompt(PromptID, request) -> *promptvmgosdk.UpdatePromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates name, description, status, tags, or isPublic. Does not create a new version.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdatePromptRequest{
        PromptID: "promptId",
    }
client.Prompts.UpdatePrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` — Optional. Replays the original 2xx response for 24h on retry with the same key + same body. A different body with the same key returns 422 idempotency_conflict.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.UpdatePromptRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Prompt Versions
<details><summary><code>client.PromptVersions.ListPromptVersions(PromptID) -> *promptvmgosdk.ListPromptVersionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListPromptVersionsRequest{
        PromptID: "promptId",
    }
client.PromptVersions.ListPromptVersions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptVersions.CreatePromptVersion(PromptID, request) -> *promptvmgosdk.CreatePromptVersionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreatePromptVersionRequest{
        PromptID: "promptId",
        Content: "content",
    }
client.PromptVersions.CreatePromptVersion(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` — Optional. Replays the original 2xx response for 24h on retry with the same key + same body. A different body with the same key returns 422 idempotency_conflict.
    
</dd>
</dl>

<dl>
<dd>

**content:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**systemPrompt:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**changeNote:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**versionLabel:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**variablesSchema:** `map[string]*promptvmgosdk.CreatePromptVersionRequestVariablesSchemaValue` 
    
</dd>
</dl>

<dl>
<dd>

**modelIDs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**baseVersion:** `*int` — Optional optimistic-concurrency guard. Must equal the current head versionNumber (0 for a prompt with no versions). A stale value yields 409 version_conflict with no mutation; omit for append-only (last writer wins).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptVersions.RollbackPrompt(PromptID, request) -> *promptvmgosdk.RollbackPromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new version from an older version's content.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RollbackPromptRequest{
        PromptID: "promptId",
        TargetVersion: 1,
    }
client.PromptVersions.RollbackPrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` — Optional. Replays the original 2xx response for 24h on retry with the same key + same body. A different body with the same key returns 422 idempotency_conflict.
    
</dd>
</dl>

<dl>
<dd>

**targetVersion:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptVersions.GetPromptVersion(PromptID, VersionID) -> *promptvmgosdk.GetPromptVersionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetPromptVersionRequest{
        PromptID: "promptId",
        VersionID: "versionId",
    }
client.PromptVersions.GetPromptVersion(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptVersions.UpdatePromptVersion(PromptID, VersionID, request) -> *promptvmgosdk.UpdatePromptVersionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdatePromptVersionRequest{
        PromptID: "promptId",
        VersionID: "versionId",
    }
client.PromptVersions.UpdatePromptVersion(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionLabel:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**changeNote:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**content:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**systemPrompt:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**variablesSchema:** `map[string]*promptvmgosdk.UpdatePromptVersionRequestVariablesSchemaValue` 
    
</dd>
</dl>

<dl>
<dd>

**baseVersion:** `*int` — Optional optimistic-concurrency guard. Must equal the targeted version’s current versionNumber; a mismatch yields 409 version_conflict with no columns written.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptVersions.DiffPromptVersions(PromptID) -> *promptvmgosdk.DiffPromptVersionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a unified diff between two version numbers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DiffPromptVersionsRequest{
        PromptID: "promptId",
        From: "from",
        To: "to",
    }
client.PromptVersions.DiffPromptVersions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptVersions.GetVersionRecommendedModels(PromptID, VersionID) -> *promptvmgosdk.GetVersionRecommendedModelsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetVersionRecommendedModelsRequest{
        PromptID: "promptId",
        VersionID: "versionId",
    }
client.PromptVersions.GetVersionRecommendedModels(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptVersions.SetVersionRecommendedModels(PromptID, VersionID, request) -> *promptvmgosdk.SetVersionRecommendedModelsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replace-all, in the order given. Up to 10. An unknown or retired model id is rejected rather than dropped, so a partially-invalid selection never silently succeeds.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.SetVersionRecommendedModelsRequest{
        PromptID: "promptId",
        VersionID: "versionId",
        ModelIDs: []string{
            "modelIds",
        },
    }
client.PromptVersions.SetVersionRecommendedModels(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**modelIDs:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## PromptDeployments
<details><summary><code>client.PromptDeployments.DeployPromptVersion(PromptID, request) -> *promptvmgosdk.DeployPromptVersionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Points the named environment at an existing published version. Idempotent — redeploying the same version updates deployedAt only.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeployPromptVersionRequest{
        PromptID: "promptId",
        Environment: promptvmgosdk.DeployPromptVersionRequestEnvironmentDevelopment,
        VersionID: "versionId",
    }
client.PromptDeployments.DeployPromptVersion(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**environment:** `*promptvmgosdk.DeployPromptVersionRequestEnvironment` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Prompt Resolution
<details><summary><code>client.PromptResolution.ResolvePrompt(PromptID) -> *promptvmgosdk.ResolvePromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resolves [[include:]] references and {{variable}} substitutions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ResolvePromptRequest{
        PromptID: "promptId",
    }
client.PromptResolution.ResolvePrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptResolution.ResolvePromptPost(PromptID, request) -> *promptvmgosdk.ResolvePromptPostResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resolves [[include:]] references and {{variable}} substitutions. Read-safe: identical semantics to GET /prompts/{promptId}/resolve, but accepts a JSON body so callers can pass a Record<string,string> `variables` map cleanly. Idempotent — no `Idempotency-Key` header is required or honoured. Unknown variables remain literal in the output (e.g. `{{unknown}}`); the endpoint never throws on missing keys. Gated by the `RESOLVE_POST_ENABLED` env flag (default true); when disabled the endpoint returns 503 and SDK/CLI clients should fall back to the GET path with `variables` ignored.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ResolvePromptPostRequest{
        PromptID: "promptId",
    }
client.PromptResolution.ResolvePromptPost(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**variables:** `map[string]string` — Optional caller-supplied values for `{{name}}` tokens in the prompt. Unknown variables remain literal in the response. Never throws.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptResolution.ResolvePromptStream(PromptID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Same as /resolve but delivers content as Server-Sent Events for large payloads.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ResolvePromptStreamRequest{
        PromptID: "promptId",
    }
client.PromptResolution.ResolvePromptStream(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Collections
<details><summary><code>client.Collections.ListCollections() -> *promptvmgosdk.ListCollectionsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListCollectionsRequest{}
client.Collections.ListCollections(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Collections.CreateCollection(request) -> *promptvmgosdk.CreateCollectionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateCollectionRequest{
        Name: "name",
    }
client.Collections.CreateCollection(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Collections.GetCollection(CollectionID) -> *promptvmgosdk.GetCollectionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetCollectionRequest{
        CollectionID: "collectionId",
    }
client.Collections.GetCollection(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**collectionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Collections.DeleteCollection(CollectionID) -> *promptvmgosdk.DeleteCollectionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteCollectionRequest{
        CollectionID: "collectionId",
    }
client.Collections.DeleteCollection(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**collectionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Collections.UpdateCollection(CollectionID, request) -> *promptvmgosdk.UpdateCollectionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateCollectionRequest{
        CollectionID: "collectionId",
    }
client.Collections.UpdateCollection(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**collectionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Collections.AddCollectionItem(CollectionID, request) -> *promptvmgosdk.AddCollectionItemResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AddCollectionItemRequest{
        CollectionID: "collectionId",
        FileID: "fileId",
    }
client.Collections.AddCollectionItem(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**collectionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**fileID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Collections.RemoveCollectionItem(CollectionID, ItemID) -> *promptvmgosdk.RemoveCollectionItemResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RemoveCollectionItemRequest{
        CollectionID: "collectionId",
        ItemID: "itemId",
    }
client.Collections.RemoveCollectionItem(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**collectionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**itemID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Collections.ReorderCollectionItems(CollectionID, request) -> *promptvmgosdk.ReorderCollectionItemsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Rewrites each member's position to its index in orderedItemIds. The id set must match the collection's current members exactly.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ReorderCollectionItemsRequest{
        CollectionID: "collectionId",
        OrderedItemIDs: []string{
            "orderedItemIds",
        },
    }
client.Collections.ReorderCollectionItems(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**collectionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**orderedItemIDs:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Templates
<details><summary><code>client.Templates.ConvertPromptToTemplate(PromptID) -> *promptvmgosdk.ConvertPromptToTemplateResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ConvertPromptToTemplateRequest{
        PromptID: "promptId",
    }
client.Templates.ConvertPromptToTemplate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Templates.ListTemplates() -> *promptvmgosdk.ListTemplatesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists templates visible org-wide.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListTemplatesRequest{
        WorkspaceID: "workspaceId",
    }
client.Templates.ListTemplates(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Templates.CreatePromptFromTemplate(request) -> *promptvmgosdk.CreatePromptFromTemplateResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreatePromptFromTemplateRequest{
        TemplateID: "templateId",
        Name: "name",
        WorkspaceID: "workspaceId",
    }
client.Templates.CreatePromptFromTemplate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**variableValues:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Prompt Organization
<details><summary><code>client.PromptOrganization.MovePrompt(PromptID, request) -> *promptvmgosdk.MovePromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Moves a prompt to a different directory and/or workspace.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.MovePromptRequest{
        PromptID: "promptId",
    }
client.PromptOrganization.MovePrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptOrganization.ForkPrompt(PromptID, request) -> *promptvmgosdk.ForkPromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a copy in the specified workspace with forkedFromPromptId set.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ForkPromptRequest{
        PromptID: "promptId",
        WorkspaceID: "workspaceId",
    }
client.PromptOrganization.ForkPrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptOrganization.ListPromptReferences(PromptID) -> *promptvmgosdk.ListPromptReferencesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all [[include:]] references in the current version.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListPromptReferencesRequest{
        PromptID: "promptId",
    }
client.PromptOrganization.ListPromptReferences(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PromptOrganization.ListPromptDependents(PromptID) -> *promptvmgosdk.ListPromptDependentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all prompts that reference this one.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListPromptDependentsRequest{
        PromptID: "promptId",
    }
client.PromptOrganization.ListPromptDependents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Prompt Export
<details><summary><code>client.PromptExport.ExportPrompt(PromptID, request) -> *promptvmgosdk.ExportPromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Exports a prompt in Markdown, JSON, or XML format.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ExportPromptRequest{
        PromptID: "promptId",
        Format: promptvmgosdk.ExportPromptRequestFormatMd,
    }
client.PromptExport.ExportPrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**format:** `*promptvmgosdk.ExportPromptRequestFormat` 
    
</dd>
</dl>

<dl>
<dd>

**resolveReferences:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**includeMetadata:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Directories
<details><summary><code>client.Directories.ListDirectories() -> *promptvmgosdk.ListDirectoriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all active directories in a workspace.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListDirectoriesRequest{
        WorkspaceID: "workspaceId",
    }
client.Directories.ListDirectories(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Directories.CreateDirectory(request) -> *promptvmgosdk.CreateDirectoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a persisted directory inside a workspace.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateDirectoryRequest{
        WorkspaceID: "workspaceId",
        Name: "name",
    }
client.Directories.CreateDirectory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**parentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Directories.DeleteDirectory(DirectoryID) -> *promptvmgosdk.DeleteDirectoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes an empty directory.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteDirectoryRequest{
        DirectoryID: "directoryId",
    }
client.Directories.DeleteDirectory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**directoryID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Directories.UpdateDirectory(DirectoryID, request) -> *promptvmgosdk.UpdateDirectoryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates persisted directory metadata.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateDirectoryRequest{
        DirectoryID: "directoryId",
    }
client.Directories.UpdateDirectory(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**directoryID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**parentID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Resources
<details><summary><code>client.Resources.ListWorkspaceResources() -> *promptvmgosdk.ListWorkspaceResourcesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all confirmed, non-deleted resources belonging to a workspace. By default, resources bundled into the current version of a skill are excluded — they are internal to the skill and surface through its files manifest. Resources attached to ordinary prompts remain included. Pass includeBundled=true to return everything, including skill-bundled resources. Pass orphansOnly=true to return ONLY resources not bound to any prompt/skill version — typically leftovers from failed CLI skill uploads (the CLI cleanup-orphans command uses this).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListWorkspaceResourcesRequest{
        WorkspaceID: "workspaceId",
    }
client.Resources.ListWorkspaceResources(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeBundled:** `*bool` — Include resources bundled into the current version of a skill. Default false.
    
</dd>
</dl>

<dl>
<dd>

**orphansOnly:** `*bool` — Return only orphan resources (no binding to any prompt/skill version). Default false. Overrides includeBundled.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.InitiateResourceUpload(request) -> *promptvmgosdk.InitiateResourceUploadResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a file record and returns a presigned S3 PUT URL. Upload your file bytes directly to the presigned URL within 15 minutes. Then call POST /resources/:resourceId/confirm to complete the process.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.InitiateResourceUploadRequest{
        WorkspaceID: "workspaceId",
        Name: "name",
        ContentType: "contentType",
        SizeBytes: 1,
    }
client.Resources.InitiateResourceUpload(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` — Workspace to upload the resource into
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Display name for the resource
    
</dd>
</dl>

<dl>
<dd>

**contentType:** `string` — MIME type (e.g. text/markdown, image/png)
    
</dd>
</dl>

<dl>
<dd>

**sizeBytes:** `int` — File size in bytes (max 100MB)
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` — Directory to place the resource in (null = workspace root)
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**categories:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.ConfirmResourceUpload(ResourceID, request) -> *promptvmgosdk.ConfirmResourceUploadResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Verifies the file was uploaded to S3 and marks the resource as available for use.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ConfirmResourceUploadRequest{
        ResourceID: "resourceId",
    }
client.Resources.ConfirmResourceUpload(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resourceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.GetResource(ResourceID) -> *promptvmgosdk.GetResourceResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetResourceRequest{
        ResourceID: "resourceId",
    }
client.Resources.GetResource(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resourceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.DeleteResource(ResourceID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Soft-deletes the resource and removes its S3 object. Also detaches it from all prompt versions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteResourceRequest{
        ResourceID: "resourceId",
    }
client.Resources.DeleteResource(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resourceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.UpdateResource(ResourceID, request) -> *promptvmgosdk.UpdateResourceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update resource metadata (name, directory placement).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateResourceRequest{
        ResourceID: "resourceId",
    }
client.Resources.UpdateResource(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resourceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.GetResourceDownloadURL(ResourceID) -> *promptvmgosdk.GetResourceDownloadURLResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a time-limited presigned URL (1 hour) for downloading the resource from S3. Disposition defaults to attachment; inline is honored only for sniff-verified preview-safe types (png, jpeg, gif, webp, pdf) and otherwise falls back to attachment, reported in the response.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetResourceDownloadURLRequest{
        ResourceID: "resourceId",
    }
client.Resources.GetResourceDownloadURL(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resourceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**disposition:** `*promptvmgosdk.GetResourceDownloadURLRequestDisposition` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.GetResourceRagStatus(ResourceID) -> *promptvmgosdk.GetResourceRagStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the document-processing status for the resource's current version: ingestion job status, failure stage/message, chunk count, whether a markdown render exists, and the number of current embeddings. Powers the "RAG Status" metadata row in the resource detail panel.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetResourceRagStatusRequest{
        ResourceID: "resourceId",
    }
client.Resources.GetResourceRagStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resourceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.SetResourceSearchIndexing(ResourceID, request) -> *promptvmgosdk.SetResourceSearchIndexingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Toggle whether this resource is indexed for Knowledge Search (RAG). Disabling removes it from search and purges its embeddings; enabling re-triggers processing. Returns the refreshed RAG status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.SetResourceSearchIndexingRequest{
        ResourceID: "resourceId",
        Enabled: true,
    }
client.Resources.SetResourceSearchIndexing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resourceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `bool` — Include this item in Knowledge Search.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.GetResourceMarkdownURL(ResourceID) -> *promptvmgosdk.GetResourceMarkdownURLResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a time-limited presigned URL (1 hour) for the markdown render of the resource's current version (e.g. a PDF converted to markdown by docproc). Returns 404 when no markdown render is available. Powers the Original/Markdown toggle.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetResourceMarkdownURLRequest{
        ResourceID: "resourceId",
    }
client.Resources.GetResourceMarkdownURL(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resourceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.GetPromptRagStatus(PromptID) -> *promptvmgosdk.GetPromptRagStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the document-processing status for the prompt's current version: ingestion job status, chunk count, and current embedding count. Prompts are embedded like files, so this powers the RAG status indicator in the editor.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetPromptRagStatusRequest{
        PromptID: "promptId",
    }
client.Resources.GetPromptRagStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.SetPromptSearchIndexing(PromptID, request) -> *promptvmgosdk.SetPromptSearchIndexingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Toggle whether this prompt is indexed for Knowledge Search (RAG). Disabling removes it from search and purges its embeddings; enabling re-triggers processing. Returns the refreshed RAG status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.SetPromptSearchIndexingRequest{
        PromptID: "promptId",
        Enabled: true,
    }
client.Resources.SetPromptSearchIndexing(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `bool` — Include this item in Knowledge Search.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.ListPromptResources(PromptID) -> *promptvmgosdk.ListPromptResourcesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListPromptResourcesRequest{
        PromptID: "promptId",
    }
client.Resources.ListPromptResources(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.AttachPromptResource(PromptID, request) -> *promptvmgosdk.AttachPromptResourceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The resource must be confirmed (upload completed) and belong to the same workspace.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.AttachPromptResourceRequest{
        PromptID: "promptId",
        ResourceID: "resourceId",
    }
client.Resources.AttachPromptResource(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**resourceID:** `string` — ID of a confirmed resource to attach
    
</dd>
</dl>

<dl>
<dd>

**usageContext:** `*promptvmgosdk.AttachPromptResourceRequestUsageContext` — How the resource is used by the LLM at execution time
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Resources.DetachPromptResource(PromptID, ResourceID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes the association only — the resource file is not deleted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DetachPromptResourceRequest{
        PromptID: "promptId",
        ResourceID: "resourceId",
    }
client.Resources.DetachPromptResource(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**promptID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**resourceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Skills
<details><summary><code>client.Skills.ListSkills() -> *promptvmgosdk.ListSkillsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns paginated skills. Filter `q` matches name, description, AND when_to_use (the Claude/Agent Skills documented match basis). Filter `tag` matches the tags array exactly (case-insensitive).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListSkillsRequest{
        WorkspaceID: "workspaceId",
    }
client.Skills.ListSkills(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tag:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*promptvmgosdk.ListSkillsRequestIsPublic` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Skills.CreateSkill(request) -> *promptvmgosdk.CreateSkillResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a skill artifact (content_kind=skill). The full SKILL.md bytes (frontmatter + body) are stored verbatim. Frontmatter is parsed into a derived metadata cache used by the typed read shape.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateSkillRequest{
        SkillMd: "skill_md",
        WorkspaceID: "workspaceId",
    }
client.Skills.CreateSkill(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**skillMd:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**files:** `[]*promptvmgosdk.CreateSkillRequestFilesItem` 
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.CreateSkillRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Skills.GetSkill(SkillID) -> *promptvmgosdk.GetSkillResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetSkillRequest{
        SkillID: "skillId",
    }
client.Skills.GetSkill(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**skillID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Skills.DeleteSkill(SkillID) -> *promptvmgosdk.DeleteSkillResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteSkillRequest{
        SkillID: "skillId",
    }
client.Skills.DeleteSkill(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**skillID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Skills.UpdateSkill(SkillID, request) -> *promptvmgosdk.UpdateSkillResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replaces any subset of skill fields. Replacing skill_md re-runs ingest, regenerates the metadata cache, and creates a new version. Replacing files atomically rebinds resource references.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateSkillRequest{
        SkillID: "skillId",
    }
client.Skills.UpdateSkill(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**skillID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**skillMd:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**files:** `[]*promptvmgosdk.UpdateSkillRequestFilesItem` 
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.UpdateSkillRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**baseVersion:** `*int` — Optional optimistic-concurrency guard for the skill_md re-version path. Must equal the current head versionNumber; a stale value yields 409 version_conflict with no new version committed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Skills.ListSkillFiles(SkillID) -> *promptvmgosdk.ListSkillFilesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the manifest of files bundled into the skill’s current version — path, sizeBytes, contentType — without any file bytes. Pull an individual file with GET /skills/:skillId/files/content?path=…
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListSkillFilesRequest{
        SkillID: "skillId",
    }
client.Skills.ListSkillFiles(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**skillID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Skills.GetSkillFileContent(SkillID) -> *promptvmgosdk.GetSkillFileContentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a single bundled file identified by its manifest path. Text files are returned inline (encoding=inline, content is UTF-8); binary files (and any file over the inline cap) return a short-lived presigned download URL (encoding=url, downloadUrl).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetSkillFileContentRequest{
        SkillID: "skillId",
        Path: "path",
    }
client.Skills.GetSkillFileContent(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**skillID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**path:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hooks
<details><summary><code>client.Hooks.ListHooks() -> *promptvmgosdk.ListHooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns paginated hooks. Supports filtering by event, handler type, tag, and text search.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListHooksRequest{
        WorkspaceID: "workspaceId",
    }
client.Hooks.ListHooks(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**event:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**handlerType:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tag:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*promptvmgosdk.ListHooksRequestIsPublic` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.CreateHook(request) -> *promptvmgosdk.CreateHookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a hook artifact (content_kind=hook). The JSON configuration is validated against the Claude Code hooks schema.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateHookRequest{
        HookJSON: "hook_json",
        WorkspaceID: "workspaceId",
    }
client.Hooks.CreateHook(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**hookJSON:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**files:** `[]*promptvmgosdk.CreateHookRequestFilesItem` 
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.CreateHookRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.GetHook(HookID) -> *promptvmgosdk.GetHookResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetHookRequest{
        HookID: "hookId",
    }
client.Hooks.GetHook(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**hookID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.DeleteHook(HookID) -> *promptvmgosdk.DeleteHookResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteHookRequest{
        HookID: "hookId",
    }
client.Hooks.DeleteHook(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**hookID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.UpdateHook(HookID, request) -> *promptvmgosdk.UpdateHookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replaces any subset of hook fields. Replacing hook_json re-runs ingest, regenerates the metadata cache, and creates a new version.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateHookRequest{
        HookID: "hookId",
    }
client.Hooks.UpdateHook(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**hookID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**hookJSON:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**files:** `[]*promptvmgosdk.UpdateHookRequestFilesItem` 
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*promptvmgosdk.UpdateHookRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**directoryID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.ExportHook(HookID) -> *promptvmgosdk.ExportHookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the hook configuration as a Claude Code settings.json fragment ready to merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ExportHookRequest{
        HookID: "hookId",
    }
client.Hooks.ExportHook(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**hookID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.ValidateHook(request) -> *promptvmgosdk.ValidateHookResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Validates the hook JSON without saving. Returns errors and warnings.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ValidateHookRequest{
        HookJSON: "hook_json",
    }
client.Hooks.ValidateHook(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**hookJSON:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.ImportHooks(request) -> *promptvmgosdk.ImportHooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Parses a Claude Code settings.json and creates hook artifacts from the hooks block.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ImportHooksRequest{
        Source: promptvmgosdk.ImportHooksRequestSourceSettingsJSON,
        Content: "content",
        WorkspaceID: "workspaceId",
    }
client.Hooks.ImportHooks(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**source:** `*promptvmgosdk.ImportHooksRequestSource` 
    
</dd>
</dl>

<dl>
<dd>

**content:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isPublic:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**tags:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Workspaces
<details><summary><code>client.Workspaces.ListWorkspaces() -> *promptvmgosdk.ListWorkspacesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns workspaces the authenticated user can access within an organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.ListWorkspacesRequest{
        OrganizationID: "organizationId",
    }
client.Workspaces.ListWorkspaces(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**organizationID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Workspaces.CreateWorkspace(request) -> *promptvmgosdk.CreateWorkspaceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new workspace within an organization. The creator is added as owner.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.CreateWorkspaceRequest{
        Name: "name",
        OrganizationID: "organizationId",
    }
client.Workspaces.CreateWorkspace(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*promptvmgosdk.CreateWorkspaceRequestVisibility` 
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**organizationID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Workspaces.GetWorkspace(WorkspaceID) -> *promptvmgosdk.GetWorkspaceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns full workspace metadata including owner, members, and content counts. PRIVATE workspaces are readable only by org owners/admins, the workspace owner, or workspace members; other org members receive 404.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetWorkspaceRequest{
        WorkspaceID: "workspaceId",
    }
client.Workspaces.GetWorkspace(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Workspaces.DeleteWorkspace(WorkspaceID) -> *promptvmgosdk.DeleteWorkspaceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Soft-deletes a workspace. Cannot delete default workspace. Use cascade=true to also delete prompts/directories.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.DeleteWorkspaceRequest{
        WorkspaceID: "workspaceId",
    }
client.Workspaces.DeleteWorkspace(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cascade:** `*promptvmgosdk.DeleteWorkspaceRequestCascade` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Workspaces.UpdateWorkspace(WorkspaceID, request) -> *promptvmgosdk.UpdateWorkspaceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates workspace metadata. Only owner or admin can update. Re-generates slug if name changes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateWorkspaceRequest{
        WorkspaceID: "workspaceId",
    }
client.Workspaces.UpdateWorkspace(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**visibility:** `*promptvmgosdk.UpdateWorkspaceRequestVisibility` 
    
</dd>
</dl>

<dl>
<dd>

**icon:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Workspaces.UpdateWorkspacePin(WorkspaceID, request) -> *promptvmgosdk.UpdateWorkspacePinResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pins or unpins a workspace for the authenticated user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.UpdateWorkspacePinRequest{
        WorkspaceID: "workspaceId",
        Pinned: true,
    }
client.Workspaces.UpdateWorkspacePin(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**pinned:** `bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Workspaces.TransferWorkspaceOwnership(WorkspaceID, request) -> *promptvmgosdk.TransferWorkspaceOwnershipResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Transfers workspace ownership to another org member. Only current owner can transfer. Cannot transfer default workspace.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.TransferWorkspaceOwnershipRequest{
        WorkspaceID: "workspaceId",
        NewOwnerID: "newOwnerId",
    }
client.Workspaces.TransferWorkspaceOwnership(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**workspaceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**newOwnerID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Search
<details><summary><code>client.Search.Organization() -> *promptvmgosdk.SearchOrganizationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cross-workspace, ranked search returning prompts and files the requester can read. Optional scoping by workspace or directory. Cursor pagination, normalized relevance score, ts_headline highlights with <mark> markers. The contract is forward-compat: future ranking modes (`semantic`, `hybrid`) and result kinds (`collection`, `team`) are additive — existing clients ignore unknown enum values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.SearchOrganizationRequest{
        Q: "q",
        OrganizationID: "organizationId",
    }
client.Search.Organization(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `string` — Search query. Less than 2 chars after trim short-circuits to an empty result set (no error).
    
</dd>
</dl>

<dl>
<dd>

**organizationID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**workspaceIDs:** `*string` — Restrict results to these workspaces. Workspaces from a different org or that the requester cannot access are silently ignored.
    
</dd>
</dl>

<dl>
<dd>

**directoryIDs:** `*string` — Restrict file results to these directories.
    
</dd>
</dl>

<dl>
<dd>

**kinds:** `*promptvmgosdk.SearchOrganizationRequestKindsItem` — Subset of [prompt, file]. Default both. Unknown kinds return 400 UNSUPPORTED_KIND.
    
</dd>
</dl>

<dl>
<dd>

**contentKind:** `*promptvmgosdk.SearchOrganizationRequestContentKind` — Filter prompt results by content_kind. Omit for all content kinds.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — Opaque base64url cursor returned in a previous response.
    
</dd>
</dl>

<dl>
<dd>

**ranking:** `*promptvmgosdk.SearchOrganizationRequestRanking` — Ranking mode. `semantic` returns 503 SEARCH_BACKEND_UNAVAILABLE when the embedding backend is down; `hybrid` degrades to keyword (response flag `degraded: true`). Semantic hits carry a plain-text `content` highlight (the best-matching chunk). Hybrid pagination is bounded by the fused candidate pool (60 keyword + 60 semantic).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Identity
<details><summary><code>client.Identity.GetMyIdentity() -> *promptvmgosdk.GetMyIdentityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the authenticated caller's identity context including user info, organization, and default workspace. Works with both JWT session auth and API key auth.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Identity.GetMyIdentity(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Identity.GetMyWorkspaces() -> *promptvmgosdk.GetMyWorkspacesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns all workspaces the authenticated caller can access within their organization. Works with both JWT session auth and API key auth.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Identity.GetMyWorkspaces(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ai
<details><summary><code>client.Ai.EnhancePrompt(request) -> *promptvmgosdk.EnhancePromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Rewrites the submitted prompt to follow good prompting practices. Backed by the platform-internal AI middleware; the OpenRouter API key never leaves the backend. Accepts either JWT session or org-scoped pk:sk (write scope required for the pk:sk path).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.EnhancePromptRequest{
        Prompt: "prompt",
    }
client.Ai.EnhancePrompt(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `*string` — Optional client-supplied key. A retry within 24h with the same key and body replays the cached response without re-calling OpenRouter. Different body with same key → 422.
    
</dd>
</dl>

<dl>
<dd>

**prompt:** `string` — The user-authored prompt to be improved.
    
</dd>
</dl>

<dl>
<dd>

**instructions:** `*string` — Optional hint to the enhancer — e.g. "make it shorter", "target a junior engineer". Substituted into a server-controlled named slot in the system prompt; never concatenated raw.
    
</dd>
</dl>

<dl>
<dd>

**targetModel:** `*string` — Optional OpenRouter-style model id (e.g. `anthropic/claude-sonnet-4`) the enhanced prompt will be sent to. Lightly tunes the rewrite for that model family.
    
</dd>
</dl>

<dl>
<dd>

**preset:** `*promptvmgosdk.EnhancePromptRequestPreset` — Optional optimization preset. When provided, the handler routes to a specialized system prompt for that optimization type (e.g. `shorten` → task slug `prompt-optimization-shorten`). When absent, the default `prompt-enhancement` task is used. Mutually exclusive with `instructions` — if both are provided, `preset` takes precedence.
    
</dd>
</dl>

<dl>
<dd>

**workspaceID:** `*string` — Optional workspace id for cost attribution. Caller must be a member; otherwise 403.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Mcp
<details><summary><code>client.Mcp.ListMcpSessionTokens() -> *promptvmgosdk.ListMcpSessionTokensResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Mcp.ListMcpSessionTokens(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Mcp.MintMcpSessionToken(request) -> *promptvmgosdk.MintMcpSessionTokenResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a short-lived opaque session token bound to the caller. The plaintext token is returned exactly once. See contracts/session-token.md.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.MintMcpSessionTokenRequest{}
client.Mcp.MintMcpSessionToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**label:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ttlSeconds:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**defaultWorkspaceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**scopes:** `[]*promptvmgosdk.MintMcpSessionTokenRequestScopesItem` 
    
</dd>
</dl>

<dl>
<dd>

**client:** `*promptvmgosdk.MintMcpSessionTokenRequestClient` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Mcp.RevokeMcpSessionToken(TokenID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.RevokeMcpSessionTokenRequest{
        TokenID: "tokenId",
    }
client.Mcp.RevokeMcpSessionToken(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**tokenID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## OAuth
<details><summary><code>client.OAuth.OAuth21AuthorizationServerMetadata() -> map[string]any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OAuth.OAuth21AuthorizationServerMetadata(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OAuth.DynamicClientRegistrationRfc7591(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.PostAPIV1OauthRegisterRequest{
        ClientName: "client_name",
        RedirectURIs: []string{
            "redirect_uris",
        },
    }
client.OAuth.DynamicClientRegistrationRfc7591(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**redirectURIs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**grantTypes:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**responseTypes:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**tokenEndpointAuthMethod:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OAuth.OAuth21AuthorizationEndpoint() -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &promptvmgosdk.GetAPIV1OauthAuthorizeRequest{
        ResponseType: "response_type",
        ClientID: "client_id",
        RedirectURI: "redirect_uri",
        State: "state",
        CodeChallenge: "code_challenge",
        CodeChallengeMethod: "code_challenge_method",
    }
client.OAuth.OAuth21AuthorizationEndpoint(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**responseType:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**redirectURI:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**state:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**scope:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**codeChallenge:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**codeChallengeMethod:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**resource:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OAuth.OAuth21TokenEndpoint() -> map[string]any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OAuth.OAuth21TokenEndpoint(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OAuth.OAuth21JSONWebKeySet() -> map[string]any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OAuth.OAuth21JSONWebKeySet(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.OAuth.OAuth21TokenRevocationRfc7009() -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.OAuth.OAuth21TokenRevocationRfc7009(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
