# Design: DatumBridge MCP WebSocket Hub

## Purpose

The MCP WebSocket Hub acts as a relay between the DatumBridge cloud platform (which speaks HTTP) and DTBClaw edge devices (which run local MCP servers over WebSocket). The **core transport** is a transparent HTTP↔WebSocket proxy that forwards JSON-RPC and correlates request/response pairs.

> **Role update:** The shipped service also exposes Streamable HTTP `POST /mcp` (session + `tools/list` / `tools/call`) plus an embedded edge catalog for DatumBridge publish/approve. Treat this as a **relay with an MCP façade**, not a generic tool-server template. See [`docs/adr/ADR-0001-transport-vs-tool-server.md`](docs/adr/ADR-0001-transport-vs-tool-server.md) and the playbook at [`docs/README.md`](docs/README.md).

## Architecture

```
┌─────────────────┐         ┌─────────────────────┐         ┌─────────────────┐
│   DatumBridge   │  HTTP   │   MCP WebSocket Hub │   WS    │   DTBClaw       │
│   Platform      │────────▶│                     │◀═══════▶│   Device        │
│   (Cloud)       │◀────────│   - Auth (bcrypt)   │         │   (Edge MCP)    │
│                 │         │   - Correlation     │         │                 │
│                 │         │   - Timeout (60s)   │         │                 │
└─────────────────┘         └─────────────────────┘         └─────────────────┘
```

## Key Design Decisions

### 1. Dual surface: opaque proxy + MCP façade

**Core transport (opaque proxy):** The hub remains an HTTP↔WebSocket bridge for device MCP traffic:
- Cloud sends `POST /api/v1/devices/{id}/mcp` with a JSON-RPC body
- Hub forwards the body to the device over WebSocket
- Device processes the request and returns a JSON-RPC response
- Hub correlates using `deviceID|rpcID` and returns the response to the HTTP caller
- On this path the hub does **not** interpret tool semantics

**Additive MCP façade:** Separately, `POST /mcp` implements Streamable HTTP (`initialize` → `Mcp-Session-Id`, `tools/list`, `tools/call`) so DatumBridge publish/approve can discover hub builtins and catalog-backed edge relay tools. That façade is a platform integration, not permission to treat this repo as the default template for SaaS tool-servers. See `docs/adr/ADR-0001-transport-vs-tool-server.md`.

### 2. Device Authentication

Devices authenticate with a server-generated token:
1. Registration creates a random 256-bit token
2. Plain-text token is returned once to the registrant
3. Token is stored as a bcrypt hash on disk
4. Device presents the plain token when connecting via WebSocket
5. Hub validates using `bcrypt.CompareHashAndPassword`

Two registration flows are supported:
- **Direct**: `POST /register` → immediate `{device_id, token}` response
- **Pairing**: `POST /register {pairing:true}` → 6-digit code displayed on hub → device enters code → `POST /register/confirm {code}` → `{device_id, token}`

### 3. Request/Response Correlation

Pending HTTP requests are tracked in `map[string]*pendingReq` keyed by `"deviceID|rpcID"`:
- JSON-RPC `id` field is extracted from the outgoing request
- A buffered channel is created and stored in the pending map
- When a response arrives on the WebSocket, its `id` is matched to the pending entry
- The response is delivered via the channel, and the HTTP handler returns it
- A timer fires after 60s to prevent leaks if the device never responds

### 4. Connection Health

WebSocket connections use gorilla/websocket ping/pong:
- Hub sends pings every 54 seconds
- If pong is not received within 60 seconds, the connection is considered dead
- Dead connections are automatically unregistered, and pending requests are canceled

### 5. Security Layers

| Layer | Mechanism |
|-------|-----------|
| Token storage | bcrypt (cost 10) |
| Registration | `HUB_REGISTER_API_KEY` — **required in production**; empty allows open register (local-dev only) |
| WebSocket | Token validated before upgrade; Origin allowlist via `HUB_ALLOWED_ORIGINS` (**required in production**; never `*`) |
| CORS | Same origin allowlist as WS when set |
| Body size | 1 MB limit via `http.MaxBytesReader` |
| Container | Non-root user in Docker |
| MCP façade | `Mcp-Session-Id` continuity; production callers must use Studio/gateway auth (session ≠ identity) |

**Known hub debt (do not copy into new relays):** some admin/pairing GETs and confirm are not API-key gated in the current code. Production deployments should front the hub with network policy / Studio auth and treat open admin routes as a gap. Playbook target: enforce admin key on all admin routes from day one — see `docs/playbook/SECURITY_RULES.md`.

### 6. Error Response Format

REST endpoints use a standardized error structure matching other DatumBridge MCP servers:
```json
{
  "error_code": "VALIDATION_ERROR",
  "error_message": "device_id required",
  "retryable": false
}
```

The MCP proxy endpoint (`/mcp`) returns JSON-RPC error format:
```json
{
  "jsonrpc": "2.0",
  "id": null,
  "error": {"code": -32000, "message": "device not connected or request timeout"}
}
```

## Alignment with DatumBridge Platform

This hub follows the same conventions as other MCP services (`google-drive-mcp`, `social-listening`):
- Health endpoint at `/health`
- Docker with multi-stage build, non-root user, HEALTHCHECK
- Standardized error responses
- Environment-based configuration with `.env.example`
- Structured logging (zerolog)

The key difference is that this service is a **transport layer** rather than a tool provider. It enables the DatumBridge Execution Engine to reach MCP servers running on edge devices without requiring direct network access.
