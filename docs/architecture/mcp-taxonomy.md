# MCP Service Taxonomy

DatumBridge has two MCP service classes. Mixing them creates control-plane duplication and security bugs.

## Diagram

```text
┌─────────────────────┐     Streamable HTTP      ┌──────────────────────────┐
│ Studio / LangGraph  │ ───────────────────────▶ │ MCP Tool-Server (default)│
│ Execution Engine    │                          │ tools run in-process     │
└──────────┬──────────┘                          │ Drive, SearXNG, social…  │
           │                                     └──────────────────────────┘
           │ Streamable HTTP + REST
           ▼
┌─────────────────────┐     WebSocket            ┌──────────────────────────┐
│ MCP WS Hub (relay)  │ ═══════════════════════▶ │ Edge MCP (DTBClaw device)│
│ transport + façade  │                          │ shell, browser, …        │
└─────────────────────┘                          └──────────────────────────┘
```

## Classes

| Class | Role | Examples | When to choose |
|-------|------|----------|----------------|
| **Tool-server** | Implements tools; calls upstream APIs/DBs | `google-drive-mcp`, `searxng-web-search-mcp`, `social-listening` | Almost always |
| **Relay / transport** | Auth devices, correlate HTTP↔WS, optional `/mcp` façade | `datumbridge-mcp-ws-hub` | Only when cloud cannot reach edge MCP directly |
| **Edge MCP** | Runs on device; tools execute locally | DTBClaw / OctoClaw tools | Device-side capability, not a cloud MCP repo |

## Shared contracts (both cloud classes)

- `GET /health`
- `POST /mcp` Streamable HTTP JSON-RPC (`initialize`, `tools/list`, `tools/call`)
- REST errors as `{error_code, error_message, retryable}` outside `/mcp`
- Env-based config + `.env.example`
- Multi-stage Docker, non-root user, `HEALTHCHECK` on `/health`
- Structured logging; never log secrets

## Forbidden cross-copy

| Pattern | Allowed in |
|---------|------------|
| WS `/ws`, pairing, bcrypt device tokens | Relay only |
| Pending map `deviceID\|rpcID` | Relay only |
| Embedded edge catalog + registry bulk sync | Relay (hub) only |
| Domain API clients + typed tool handlers | Tool-server |
| Opaque JSON-RPC forward without tool semantics | Relay only |

## Ownership

| Concern | Owner |
|---------|-------|
| Tool semantics & upstream credentials | Tool-server team |
| Device identity, WS health, edge catalog sync | Hub / platform |
| Registry routing (`mcpServer`, base URL) | Deploy + `datumbridge-mcp` publish |

See [ADR-0001](../adr/ADR-0001-transport-vs-tool-server.md).
