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

<details><summary><code>client.Billing.ListBillingPlans() -> []*promptvmgosdk.ListBillingPlansResponseItem</code></summary>
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

Returns a Stripe-hosted Checkout URL for the caller to complete payment. Owner/admin only. Rate-limited 5/min/org. The frontend redirects the browser to the returned URL; the actual subscription state lands in our DB via the `customer.subscription.created` webhook (US-01-4), not via the success URL.
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

Returns a Stripe-hosted Customer Portal URL the caller can use to update their payment method, view invoices, edit billing details, or cancel their subscription. Owner/admin only. Rate-limited 5/min/org. Plan switching is intentionally disabled in the Portal — use `/billing/change-plan` for that.
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

Routes through the Phase 02 change-plan algorithm: detects upgrade/downgrade/seat-adjust direction, releases any pending `subscription_schedule` first (FR-02-5a), then writes the change to Stripe (immediate proration on upgrade, scheduled phase on downgrade). Owner/admin only. Rate-limited per org via Redis SET NX with TTL `BILLING_CHANGE_PLAN_COOLDOWN_SECONDS` (default 60, 0 in tests). The authoritative state lands in our DB via the `customer.subscription.updated` webhook (US-01-4) — clients should re-read `/billing/status` after a 200.
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

Owner/admin only. Pushes a `quantity` change to the Stripe subscription item with `proration_behavior=create_prorations`. Increases generate a prorated charge on the next invoice; decreases (still ≥ used_seats) generate a prorated credit on the next invoice (no card refund). Per-seat plans only.
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

Owner/admin only. Sets `cancel_at_period_end = true` on the Stripe subscription. Releases any pending downgrade schedule first (FR-02-15) so the update can apply cleanly. The cancellation is scheduled, not immediate — billing continues through the current period and access ends at `cancelAt`. Re-invoking with the same period is a no-op (FR-02-14 idempotency).
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

Owner/admin only. Sets `cancel_at_period_end = false` on the Stripe subscription. Does NOT recreate any previously scheduled downgrade — the user must re-request that via `/change-plan` (FR-02-15). Re-invoking with the same period is a no-op (FR-02-14 idempotency).
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

Owner/admin only. Atomically locks the token, increments the offer counter, and returns a Stripe Checkout URL configured with `trial_period_days` and `payment_method_collection: 'always'`. On Stripe failure runs a compensating reversal so the token is reusable (FR-06-8). One promotional trial per org per lifetime — enforced at the DB layer via the `promo_one_trial_per_org` partial UNIQUE index (FR-06-9).
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

Returns the FR-05-11 dead-letter view: `stripe_webhook_events` rows where `error IS NOT NULL`. Ordered `received_at DESC, id DESC` with opaque cursor pagination. Default `limit=50`, max `200`. Supports `?since=<ISO8601>` and `?kind=<error-keyword>` filters. The `payload` jsonb column is INTENTIONALLY EXCLUDED — Stripe payloads can carry PII (customer email, billing address); ops reads the full payload from the worker logs or the Stripe Dashboard, not from this endpoint. Currently gated by `requireOwnerOrAdmin` (org-scoped) as a placeholder — a future PR will swap in a platform-admin middleware.
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
</dd>
</dl>


</dd>
</dl>
</details>

## SkillsPublic
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

## Sharing
<details><summary><code>client.Sharing.AccessSharedPrompt(Token) -> *promptvmgosdk.AccessSharedPromptResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Public endpoint — no authentication required.
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
request := &promptvmgosdk.ListAPIKeysRequest{}
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

**gracePeriodHours:** `*int` — Hours the old secret remains valid (0-168, default 24). Public key is unchanged.
    
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

## Marketplace - Creator
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
</dd>
</dl>


</dd>
</dl>
</details>

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

## Marketplace - Listings
<details><summary><code>client.MarketplaceListings.CreateMarketplaceListing(request) -> *promptvmgosdk.CreateMarketplaceListingResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Publishes a prompt or collection to the marketplace. Requires a creator profile.
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

**variablesSchema:** `map[string]any` 
    
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

Returns all confirmed, non-deleted resources belonging to a workspace.
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

Returns a time-limited presigned URL (1 hour) for downloading the resource from S3.
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

Returns full workspace metadata including owner, members, and content counts.
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

**ranking:** `*string` — MVP only honours `keyword`. Other values return 400 UNSUPPORTED_RANKING. Reserved future values: `semantic`, `hybrid`.
    
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
