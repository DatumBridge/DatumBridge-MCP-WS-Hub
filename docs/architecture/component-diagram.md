# Component Diagram

## Hub runtime (reference implementation)

```text
                    ┌─────────────────────────────────────────┐
                    │              cmd/api (main)             │
                    │  env, logging, mux routes, shutdown     │
                    └───────────────────┬─────────────────────┘
                                        │
          ┌─────────────────────────────┼─────────────────────────────┐
          ▼                             ▼                             ▼
   Middleware stack              Hub aggregate                 Background
   CORS → Log → Strip            (internal/hub)              registry sync
          │                             │
          ▼                             ├─ REST handlers (http.go)
     http.Server                        ├─ MCP Streamable HTTP (mcp_http.go)
                                        ├─ MCP sessions (mcp_session.go)
                                        ├─ WS + auth (ws.go, auth.go)
                                        ├─ Correlation (hub.go)
                                        └─ Catalog (dtbclaw_*.go/json)
```

## New MCP tool (default Python)

```text
entrypoint → mcp_server (tools) → services (providers) → upstream APIs
                 ↑
            schemas / exceptions
```

See [PROJECT_STRUCTURE](../playbook/PROJECT_STRUCTURE.md) for file trees.
