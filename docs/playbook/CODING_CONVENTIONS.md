# Coding Conventions — DatumBridge MCP

Normative style for new MCP services. Language sections apply only to that stack; shared rules apply to all.

## Shared (all MCP services)

### Naming

| Kind | Convention | Examples |
|------|------------|----------|
| Repo / service id | `kebab-case`, prefix optional `datumbridge-` | `searxng-web-search-mcp` |
| Registry `mcpServer` | `snake_case` stable id | `datumbridge_mcp_ws_hub` |
| MCP tool `name` | `snake_case`, capability-oriented | `search_web`, `list_files` |
| REST paths | `/api/v1/...` for admin; `/mcp` for protocol; `/health` | |
| Env vars | `SCREAMING_SNAKE`, service prefix | `HUB_PORT`, `APIFY_API_TOKEN` |
| JSON fields | `snake_case` | `error_code`, `device_id` |
| Log fields | `snake_case` structured keys | `device_id`, `mcp_session` |

### Entrypoint

- Keep process bootstrap **thin**: load env, configure logging, wire router/server, graceful shutdown.
- Put domain logic in an internal package (`internal/<domain>/` in Go, `app/` in Python).
- Do not put tool implementations in `main`.

### Configuration

- Env-first; ship `.env.example` with comments (no secrets).
- Fail fast on missing **required** credentials at startup or first use with explicit errors — never silent empty results that look like “no data.”
- Sensible defaults only for non-secret knobs (port, log level, timeouts).
- Docker image: do **not** bake secrets into `ENV`; inject at runtime.

### Logging

- Structured logs (Go: `zerolog`; Python: structured JSON or equivalent).
- Levels: request noise → Debug; lifecycle/sync → Info; failures → Error.
- **Never** log tokens, API keys, OAuth secrets, pairing codes, or full auth headers.
- Prefer correlation ids (`X-Correlation-ID` / `traceparent`) on outbound/inbound hops when present.

### Errors

| Surface | Shape |
|---------|-------|
| REST (non-MCP) | `{ "error_code", "error_message", "retryable" }` |
| MCP protocol | JSON-RPC `error` with standard codes (see [api-specification](../technical/api-specification.md)) |
| Tool domain failure | JSON-RPC **result** with `isError: true` + structured text (embed `error_code` / `error_message` / `retryable` when useful) |

Encode JSON via encoders — do not concatenate untrusted strings into raw response bodies.

### HTTP hygiene

- `GET /health` returns 200 when process is ready.
- Body size limits on ingest endpoints (hub reference: 1 MB).
- Explicit timeouts on HTTP servers and upstream clients.
- CORS allowlists in production; document empty-allowlist = dev-only.

### Tests

- Co-locate unit tests with the package under test.
- Prefer table-driven cases for middleware/parsers.
- Mock upstream HTTP with `httptest` / FastAPI TestClient — no live credentials in CI.
- Test names describe behavior: `tools_call_rejects_missing_session`.

### Docs in each new MCP repo

Minimum: `README.md`, `design.md` (or ADR), `.env.example`, Dockerfile notes. Link this playbook if reusing patterns.

---

## Go conventions (from this hub)

### Layout

```text
cmd/api/main.go          # bootstrap only
internal/<domain>/       # all business logic; package name short (e.g. hub)
```

### Handlers

- Use method receivers on the aggregate root (`func (h *Hub) HandleXxx`).
- Name HTTP handlers `HandleXxx`; keep unexported helpers lowercase.
- Middleware: pure `func(http.Handler) http.Handler`; document wrap order (e.g. CORS → Logging → PathStrip → Router).
- If wrapping ResponseWriter, preserve `Hijacker`/`Flusher` when WebSockets may upgrade.

### Types

- Exported types are API/domain contracts (`Hub`, `APIError`, `DeviceInfo`).
- Stores/helpers unexported (`credentialStore`, `pendingReq`).
- Prefer `json.RawMessage` for opaque JSON-RPC payloads (avoid free-form `interface{}` decode of untrusted nested objects).
- Concurrency: `sync.RWMutex` for maps; never share `Conn` without documented ownership.

### Dependencies (keep lean)

Reference set: `gorilla/mux`, `gorilla/websocket` (relay only), `zerolog`, `godotenv`, `x/crypto/bcrypt` (device tokens). Do not add heavy frameworks without a concrete need.

### MCP in Go

- Hand-rolled Streamable HTTP is acceptable if it matches platform client (`initialize` → `Mcp-Session-Id` → tools).
- Centralize protocol version constant (e.g. `2024-11-05`).
- Catalog-driven tools: embed JSON once; build descriptors + registry bodies from one source (`sync.Once` load).

---

## Python conventions (tool-server siblings)

### Layout

```text
app/mcp_server.py          # tool surface (FastMCP / MCP SDK)
app/services/              # upstream API clients
app/schemas/               # Pydantic / typed models
app/core/exceptions.py     # typed errors → tool isError mapping
```

### Tools

- One tool = one clear capability; schemas via Pydantic or explicit JSON Schema.
- Return structured JSON in text content; map provider errors to typed exceptions → `isError` payloads.
- Keep OAuth / API credential resolution in services, not in workflow callers.
- Stateless request handling (no workflow state inside the MCP).

### Dependencies

Prefer FastMCP / official MCP Python SDK used by siblings; pin versions; keep Apify/Google/etc. clients behind service modules.

---

## Anti-conventions

- Silent catch-all that returns empty tool results on auth failure.
- Advertising `resources` / `prompts` capabilities without implementing them.
- Putting secrets in images, git, or example env files with real values.
- Copying hub `internal/hub/ws.go` correlation into a Drive/SearXNG-style tool server.
