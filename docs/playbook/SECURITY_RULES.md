# Security Rules — DatumBridge MCP

Apply to every new MCP. Hub extras are marked **[relay]**.

## Least privilege

- Tools expose the minimum arguments and side effects needed.
- Deny by default: missing/invalid credentials → explicit error, not empty success.
- **MCP tools return data / facts / provider results.** They must not authorize, approve, refuse, or choose the agent’s next action. Decisions stay in workflows, agents, or policy layers.
- **[relay]** Hub must not execute edge tool logic; only authenticated connected devices receive forwarded calls.
- **[relay]** Registration protection via API key is **required in production** (`HUB_REGISTER_API_KEY`); empty key is local-dev only.

## Secrets

- Never commit `.env`, credential files, OAuth tokens, or plaintext device tokens.
- Persist hashes or encrypted stores — not plaintext secrets (**[relay]** bcrypt for device tokens).
- Runtime env / K8s secrets only; multi-stage Docker without secret `ENV` bake-in.
- File permissions: credential files `0600`, directories `0700` when writing state to disk.
- Do not log Authorization headers, query tokens, pairing codes, or full auth headers.

## Network & HTTP

- Body size limits on all write endpoints.
- Timeouts on server and outbound clients.
- CORS: production allowlist; document wide-open as development-only.
- **[relay]** WebSocket `CheckOrigin`: exact-match allowlist; **never** `*` for WS.
- **[relay]** Production **must** set non-empty `HUB_ALLOWED_ORIGINS` (exact match). Empty allowlist is local-dev only.
- Prefer header API keys over query-string tokens for long-lived secrets (**[relay]** note: WS today uses query `token=` — treat as sensitive in access logs/proxies).

## MCP session vs caller authentication

- Issue cryptographically random session ids on `initialize`.
- Reject `tools/list` and `tools/call` without valid `Mcp-Session-Id` (`-32000`).
- Bound session TTL; do not accept client-supplied session ids as authoritative without server-side store lookup.
- **`Mcp-Session-Id` is session continuity only — not caller identity.**
- Production: `/mcp` (and **[relay]** device forward / admin APIs) must sit behind Studio/gateway authentication, mTLS, and/or cluster network policy. Do **not** expose an unauthenticated public internet path that can relay to edge devices.

## Multi-tenant / workspace isolation (tool-servers)

- Credentials scoped per tenant/workspace; no cross-tenant token reuse.
- Never return another tenant’s data in tool results.
- Encrypt at-rest token stores when persisting OAuth (see Google Drive MCP design principles).

## Input validation & untrusted content

- Validate tool arguments against schema before side effects.
- Treat all upstream/device JSON as untrusted.
- Prefer `json.RawMessage` / typed models over unpacking arbitrary nested `interface{}` (**CWE-502 caution**).
- **Prompt injection:** treat tool/upstream content as untrusted data. Never promote it into system instructions. Bound/sanitize before re-injecting into agent context (caller-side responsibility).

## Container

- Non-root user.
- HEALTHCHECK on `/health`.
- `.dockerignore` excludes `.env`, runtime data, and local secrets.

## Threat notes for authors

| Threat | Mitigation |
|--------|------------|
| Tool SSRF / open URL fetch | Allowlist hosts; block link-local/metadata IPs |
| Prompt injection via tool results | Return data only; do not execute model instructions in-server; callers treat results as untrusted |
| Tools as decision engines | No authorize/approve/refuse tools; policy stays outside MCP |
| Registry spoof | Protect registry write APIs; align `mcpServer` with deploy URL |
| Abandoned device tokens **[relay]** | Revoke endpoint + disconnect |
| Pairing code guessing **[relay]** | Short TTL + single-consume; **implement rate limits** on register/confirm (do not claim rate-limit until coded) |
| Public unauthenticated `/mcp` **[relay]** | Gateway/network auth in front of hub |
| Open admin GETs **[hub debt]** | Current hub leaves some admin/pairing reads unauthenticated — **new relays must not copy this**; enforce admin API key on all admin routes |

## Review checklist

- [ ] No secrets in repo or image
- [ ] Production: authn on **all** device admin + pairing routes (including GET), not only mutating POSTs
- [ ] Production: `/mcp` not publicly callable without platform/gateway auth
- [ ] MCP session enforced
- [ ] No decision/policy tools on the MCP surface (data only)
- [ ] Errors do not leak stack traces or tokens to clients
- [ ] Tool-server did not copy relay WS auth “for completeness”
- [ ] **[relay]** `HUB_REGISTER_API_KEY` and `HUB_ALLOWED_ORIGINS` set in production
- [ ] **[relay]** Admin API key on confirm + all device/pairing admin routes (including GET) — do not copy hub open-admin gaps
- [ ] **[relay]** Pairing start/confirm rate-limit implemented (or documented as accepted residual risk with owner)
