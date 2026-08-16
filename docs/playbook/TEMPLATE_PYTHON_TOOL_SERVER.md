# Template — Python FastMCP Tool-Server

**Default template** for new DatumBridge MCP tools. Do **not** start from the WS hub tree.

## When to use

- Tools call SaaS APIs, search engines, DBs, or internal HTTP services.
- Execution happens in this process (no edge device hop).

## Skeleton steps

1. Create layout from [PROJECT_STRUCTURE](PROJECT_STRUCTURE.md) section A.
2. Implement MCP tools in `app/mcp_server.py` (or equivalent FastMCP module).
3. Put provider clients under `app/services/`.
4. Put Pydantic/schema models under `app/schemas/`.
5. Map failures via `app/core/exceptions.py` → tool `isError` payloads.
6. Expose Streamable HTTP so Studio/`datumbridge-mcp` can `initialize` → `tools/list` → `tools/call`.
7. Add `/health` on the same process or sidecar pattern used by siblings.
8. Ship Dockerfile (non-root) + `.env.example` + `README.md` + `design.md`.

## Minimum MCP behavior

```text
initialize  → 200 + Mcp-Session-Id + serverInfo
tools/list  → requires Mcp-Session-Id; returns tools[]
tools/call  → requires session; returns content[] + isError
unknown     → JSON-RPC -32601
```

Details: [api-specification.md](../technical/api-specification.md).

## Env pattern

```bash
# Service
PORT=8000
LOG_LEVEL=INFO

# Provider (examples — use your own names)
PROVIDER_API_TOKEN=

# Optional registry / platform
# TOOL_REGISTRY_BASE_URL=http://datumbridge-mcp:8081/api/v1
# TOOL_REGISTRY_API_KEY=
```

Fail fast when required tokens are missing; do not return empty “success” lists.

## Deploy notes

- Built/deployed via `datumbridge-deploy-service` like other MCP images (see hub README deploy section for the shared pattern).
- Register tools in Tool Registry with stable `mcpServer` and base URL pointing at this service’s `/mcp`.
- If fronted by Studio nginx, document the public prefix (sibling pattern: `/api/mcp/...`).

## Reference implementations

| Repo | Why useful |
|------|------------|
| `mcp/google-drive-mcp` | OAuth multi-tenant tool split |
| `mcp/searxng-web-search-mcp` | Simple search tool MCP |
| `mcp/social-listening` | Domain services + failure model + docs governance |

## Explicit non-goals

- Device WebSocket, pairing, edge catalogs
- Copying `internal/hub` packages
- Implementing resources/prompts unless product requires them
