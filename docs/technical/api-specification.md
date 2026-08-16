# API Specification — Shared MCP Contracts

Platform-facing HTTP/MCP surface every new DatumBridge MCP should honor. Hub-specific admin/WS routes are marked **[hub]**.

## Common endpoints

| Method | Path | Purpose |
|--------|------|---------|
| GET | `/health` | Liveness/readiness |
| POST | `/mcp` | Streamable HTTP MCP (JSON-RPC 2.0) |
| OPTIONS | `/mcp` | CORS preflight when applicable |

### [hub] Additional

| Method | Path | Auth today (hub code) | Production target for scaffolds |
|--------|------|----------------------|--------------------------------|
| WS | `/ws?device_id=&token=` | Device token (query; treat as secret) | Same; prefer not logging query strings |
| POST | `/api/v1/devices/register` | `X-API-Key` / `Bearer` when `HUB_REGISTER_API_KEY` set; open if empty | Key **required**; refuse empty key outside local-dev |
| POST | `/api/v1/devices/register/confirm` | **No API key check in current hub** (pairing code only) | Same key middleware as register |
| GET | `/api/v1/devices` | **Open in current hub** | Same admin key (network policy = defense-in-depth only) |
| POST | `/api/v1/devices/{device_id}/mcp` | **Open in current hub** (device must be connected) | Platform/gateway auth + admin key |
| DELETE | `/api/v1/devices/{device_id}` | **Open in current hub** | Same admin key |
| GET | `/api/v1/pairing/pending` | **Open in current hub** | Same admin key |

When an API key is configured and enforced, unauthorized calls should return HTTP 401/403 with REST `APIError` `UNAUTHORIZED`.

Empty `HUB_REGISTER_API_KEY` is **local-dev only**. Production `/mcp` must not be an unauthenticated public internet path that can relay to devices — place Studio/gateway auth or network policy in front (MCP session is continuity only).

**Playbook rule for new relays:** do not copy the hub’s current open admin GETs/confirm gap — enforce admin key on **all** device/pairing admin routes from day one.

---

## Streamable HTTP MCP (required)

### Envelope

```json
{ "jsonrpc": "2.0", "id": 1, "method": "…", "params": { } }
```

### Minimal compliant sequence

```text
initialize
  → HTTP 200
  → Header Mcp-Session-Id: <opaque>
  → result.protocolVersion, capabilities, serverInfo

notifications/initialized   (optional; may return 202, no JSON-RPC body)

tools/list   (+ Mcp-Session-Id)
  → result.tools[]

tools/call   (+ Mcp-Session-Id)
  → result.content[] , result.isError
```

### `initialize` result (shape)

- `protocolVersion` — platform baseline `"2024-11-05"` unless upgraded with clients
- `capabilities.tools` — present even if empty object
- `serverInfo.name`, `serverInfo.version`
- **Must** set response header `Mcp-Session-Id`

### `tools/list` tool object

| Field | Required | Notes |
|-------|----------|-------|
| `name` | yes | `snake_case`, stable |
| `description` | yes | planner-facing |
| `inputSchema` | yes | JSON Schema object |
| `outputSchema` | no | optional |

### `tools/call` params

```json
{ "name": "tool_name", "arguments": { } }
```

Missing `arguments` → `{}`.

### `tools/call` success result

```json
{
  "content": [ { "type": "text", "text": "…" } ],
  "isError": false
}
```

Domain/provider failures: still JSON-RPC **result** with `isError: true` (prefer structured JSON text embedding `error_code`, `error_message`, `retryable`).

### JSON-RPC error codes (`/mcp` only)

| Code | When |
|------|------|
| `-32700` | Parse error |
| `-32600` | Invalid request / not JSON-RPC 2.0 |
| `-32601` | Method not found **or** unknown tool |
| `-32602` | Invalid params / schema |
| `-32603` | Internal protocol/server failure (not routine business errors) |
| `-32000` | Session missing/invalid/expired |

Error body:

```json
{ "jsonrpc": "2.0", "id": null, "error": { "code": -32601, "message": "…" } }
```

### Resources / prompts

Optional. Default: **do not implement**. If absent, return `-32601`. Never advertise in `initialize` capabilities unless implemented.

---

## REST `APIError` (non-MCP)

```json
{
  "error_code": "VALIDATION_ERROR",
  "error_message": "device_id required",
  "retryable": false
}
```

Common codes: `VALIDATION_ERROR`, `UNAUTHORIZED`, `METHOD_NOT_ALLOWED`, `PAYLOAD_TOO_LARGE`, `NOT_FOUND`, `INTERNAL_ERROR`.

**Never** return this shape as the body of `/mcp` JSON-RPC responses.

---

## Failure matrix

| Problem | Channel |
|---------|---------|
| Bad JSON on `/mcp` | JSON-RPC `-32700` |
| Missing MCP session | JSON-RPC `-32000` |
| Unknown tool | JSON-RPC `-32601` |
| Invalid tool args | JSON-RPC `-32602` |
| Provider 401 / config missing | Tool `isError: true` |
| REST validation | `APIError` |
| HTTP method / body too large | HTTP status (+ optional REST error) |
