# Business Rules — MCP Platform Contracts

Normative rules for any DatumBridge MCP service (tool-server or relay).

1. **Class declaration** — Every MCP `design.md` states tool-server vs relay.
2. **Health** — `GET /health` is mandatory for deploy probes.
3. **MCP endpoint** — `POST /mcp` speaks JSON-RPC 2.0 Streamable HTTP compatible with platform publish/approve clients.
4. **Session** — `initialize` issues `Mcp-Session-Id`; `tools/list` and `tools/call` require a valid session. Session is continuity, not caller identity.
5. **Tool identity** — Tool `name` is stable; list and call use the same names.
6. **Error channels** — REST uses `APIError`; MCP protocol uses JSON-RPC errors; tool domain failures use `isError` results.
7. **Capability honesty** — Do not advertise unimplemented resources/prompts.
8. **Config** — Secrets from environment; `.env.example` documents all knobs.
9. **Container** — Non-root + HEALTHCHECK.
10. **Relay exception** — Only relays implement device WS, pairing, edge catalog sync, and HTTP↔WS correlation.
11. **Fail fast** — Missing required auth must not look like successful empty data.
12. **Registry alignment** — `mcpServer` + deploy base URL must route execution to the advertising MCP.
13. **Data ≠ decisions** — MCP tools return data/facts/provider results; they must not authorize, approve, refuse, or choose agent next actions. Policy stays in workflows/agents.
14. **Production auth (relay)** — `HUB_REGISTER_API_KEY` and `HUB_ALLOWED_ORIGINS` required; `/mcp` behind platform/gateway auth.
