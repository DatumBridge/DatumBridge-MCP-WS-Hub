# Design Patterns — DatumBridge MCP

Patterns proven in `datumbridge-mcp-ws-hub` and sibling tool-servers. Use the **tool-server** set by default. Hub/relay patterns are reference-only.

---

## A. Tool-server patterns (default)

### A1. Thin entrypoint + domain package

**Problem:** Bootstrap mixed with tools becomes untestable.  
**Solution:** `main` / ASGI entry wires config + server; tools live in `app/` or `internal/`.  
**Reference:** `cmd/api/main.go` (bootstrap) vs `internal/hub/*` (domain).

### A2. Typed tool façade

**Problem:** Agents need stable names and JSON Schema.  
**Solution:** Each tool has `name`, `description`, `inputSchema`; `tools/list` and `tools/call` share one registry.  
**Implementation tips:**

- Single function/list that builds descriptors used by both list and call validation.
- Validate arguments before calling upstream; return `-32602` for schema failures.

### A3. Session gate (Streamable HTTP)

**Problem:** Platform clients require session continuity after `initialize`.  
**Solution:** Create opaque session id on `initialize`; set `Mcp-Session-Id`; gate `tools/*` on validity (TTL/sliding).  
**Reference:** `internal/hub/mcp_session.go`, `HandleMCPStreamableHTTP`.

### A4. Dual error channel

**Problem:** Protocol failures and business failures must not look the same.  
**Solution:**

| Kind | Mechanism |
|------|-----------|
| Parse / unknown method / bad session | JSON-RPC `error` |
| Tool business / upstream failure | `result.isError = true` + structured text |
| Non-MCP REST | `APIError` |

### A5. Fail-fast upstream auth

**Problem:** Missing API token returned as “empty search.”  
**Solution:** Typed config/auth errors → tool error with explicit message; optional `check_*_status` tool.  
**Reference sibling:** `social-listening` Apify failure model.

### A6. Service adapter

**Problem:** Provider SDKs leak into tool handlers.  
**Solution:** Tool handler validates args → service client → normalize result/errors.  
**Reference siblings:** `app/services/drive_service.py`, Apify crawler modules.

### A7. Middleware / ASGI middleware chain

**Problem:** CORS, logging, path prefixes cross-cut every route.  
**Solution:** Ordered wrappers; outermost = CORS (so errors still get headers).  
**Reference:** `CORSMiddleware(LoggingMiddleware(OptionalStudioProxyStripMiddleware(r)))`.

---

## B. Hub / relay patterns (exception only)

Copy these **only** when building another transport relay.

### B1. Connection registry

In-memory `device_id → Conn` with mutex; register on WS auth success; unregister on close; revoke disconnects.

### B2. Request/response correlation

Key pending HTTP waiters by `"deviceID|rpcID"`; deliver WS JSON-RPC responses to the matching channel; timer cancels leaks.  
**Do not** use this inside a normal tool-server.

### B3. Opaque forward façade

`POST /api/v1/devices/{id}/mcp` forwards raw JSON-RPC bytes; hub does not interpret tool semantics on that path.

### B4. Catalog-driven edge tools

Embed JSON catalog → `tools/list` descriptors (inject routing args like `device_id`) → optional registry POST bodies. Adding an edge tool = update catalog + device binary, not a new Go `case` for every tool (relay path uses generic `handleMCPEdgeRelayTool`).

### B5. Credential lifecycle

Generate high-entropy token once → return plaintext once → store bcrypt hash (mode `0600`) → validate before WS upgrade → revoke deletes hash and drops connection.

### B6. Fail-closed capabilities

Missing optional edge capability fields default to **false/empty**, never “assume allowed.”

### B7. Control-plane handshake

Edge hello / hub handshake messages separate from MCP JSON-RPC responses; classify inbound WS frames before correlation.

---

## C. Pattern selection cheat sheet

| Goal | Use pattern |
|------|-------------|
| New SaaS/API MCP tool | A1–A7 |
| Multi-tenant OAuth tool | A2, A5, A6 + security isolation |
| Publish to Tool Registry | A2, A3 + registry fields ([integrations](../technical/integrations.md)) |
| Cloud → edge device MCP | B1–B7 (Go relay template) |
| List many edge tools without code churn | B4 |

---

## D. Anti-patterns

- Hybrid “tool-server that also opens device WS” without a product decision (see [ADR-0001](../adr/ADR-0001-transport-vs-tool-server.md)).
- Per-tool hard-coded registry sync with different schemas than `tools/list`.
- Using JSON-RPC `-32603` for ordinary business validation.
- Wildcards (`*`) for WebSocket Origin checks.
