# Template — Go Relay (Hub-Style)

Use **only** when cloud HTTP must reach MCP servers that are not directly reachable (typical: edge devices over WebSocket).

For normal API tools, use [TEMPLATE_PYTHON_TOOL_SERVER](TEMPLATE_PYTHON_TOOL_SERVER.md) instead.

## When justified

- Persistent device connections with auth
- HTTP → multiplexed transport → device MCP
- Optional Streamable HTTP façade so Tool Registry / Studio can discover relayed tools

## What to copy from this hub

| Keep / adapt | Why |
|--------------|-----|
| `cmd/api` thin bootstrap | Process boundary |
| `internal/<domain>` package | Domain isolation |
| `APIError` REST shape | Platform consistency |
| Streamable HTTP `/mcp` session + tools | Publish/approve |
| Middleware (CORS, logging, optional path strip) | Studio ingress |
| Structured zerolog + godotenv | Ops parity |
| Non-root Dockerfile + `/health` | Deploy parity |

## What to copy only if needed

| Optional | When |
|----------|------|
| WS upgrade + ping/pong | Device transport |
| bcrypt credential store + pairing | Device identity |
| Correlation pending map | Sync HTTP waiters to WS responses |
| Embedded catalog + registry sync | Many edge tools / Studio listing |
| Embedded test UI | Operator debugging |

## Minimal route set (relay)

```text
GET  /health
POST /mcp                      # Streamable HTTP subset
WS   /ws                       # device connect (if applicable)
POST /api/v1/...               # register / forward / admin as required
```

## Implementation order

1. Health + logging + config
2. Auth + connection registry (if WS)
3. Opaque forward path + correlation + timeouts
4. Streamable HTTP façade (`initialize` / list / call)
5. Catalog or hub-local tools (if advertising via `/mcp`)
6. Optional registry sync
7. Hardening: origins, **admin API key on all admin routes (including confirm + GETs)**, pairing rate-limit, body limits, non-root image
8. Do **not** ship open admin GETs/confirm “like the hub currently does” — that is documented debt, not a pattern to copy

## Explicit non-goals (for this template’s users)

- Do not use this template for Drive/Sheets/search/crawl tools
- Do not execute edge tool logic inside the relay
- Do not weaken Origin checks with `*`
- Do not treat root `design.md` “proxy-only” as forbidding `/mcp` without reading [ADR-0001](../adr/ADR-0001-transport-vs-tool-server.md)

## Tests to include

- Register / unregister connection
- Correlation success + timeout
- MCP initialize → session → tools/list rejection without session
- Origin allowlist / middleware path strip (if used)
- Auth hash validation (no plaintext round-trip asserts of production secrets)
