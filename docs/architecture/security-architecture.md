# Security Architecture

Canonical checklist: [SECURITY_RULES](../playbook/SECURITY_RULES.md).

## Hub trust boundaries

```text
Internet / Studio (platform auth / network policy REQUIRED in production)
       │  API key REQUIRED in production on register (+ protect admin GETs)
       ▼
   Hub REST + /mcp
       │  Mcp-Session-Id = continuity only (not caller identity)
       │
       ├──► Tool Registry (server-side sync, API key)
       │
       ▼  bcrypt-validated token on /ws
   Device WebSocket
       │
       ▼
   Edge MCP (executes tools)
```

## Controls summary

| Boundary | Control |
|----------|---------|
| Caller → hub (production) | Studio/gateway auth, mTLS, and/or network policy — not MCP session alone |
| Device identity | Token once + bcrypt at rest |
| Registration | `HUB_REGISTER_API_KEY` **required in production**; empty = local-dev only. **Known debt:** confirm + some admin GETs are still open in current hub code — new relays must enforce admin key on all admin routes |
| WS CSRF-ish | `HUB_ALLOWED_ORIGINS` **required in production** (exact match; never `*`) |
| MCP | Random session + TTL (continuity, not authn) |
| Payload | Max body size |
| Container | Non-root |
| Data plane | Relays only — no local execution of edge tools |
| Tool semantics | Tools return data; decisions stay in workflows/agents/policy |

## Tool-server boundary (contrast)

```text
Studio (platform auth) → /mcp (session continuity) → tool handler → provider API (scoped credentials)
```

No device WS plane.
