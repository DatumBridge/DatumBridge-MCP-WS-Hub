# New MCP Tool Checklist

Use this when opening a new MCP service. Default path is **Python tool-server**.

## 0. Decide class

- [ ] Read [mcp-taxonomy.md](../architecture/mcp-taxonomy.md)
- [ ] Class is **tool-server** OR **relay** (document in `design.md`)
- [ ] If tool-server: confirm you will **not** copy WS correlation / pairing / edge catalog

## 1. Scaffold

- [ ] Create repo tree from [PROJECT_STRUCTURE](PROJECT_STRUCTURE.md)
- [ ] Apply template: [Python](TEMPLATE_PYTHON_TOOL_SERVER.md) or [Go relay](TEMPLATE_GO_RELAY.md)
- [ ] Add `README.md`, `design.md`, `.env.example`, Dockerfile

## 2. Platform contracts

- [ ] `GET /health` works
- [ ] `POST /mcp` implements:
  - [ ] `initialize` → sets `Mcp-Session-Id`
  - [ ] `tools/list` (session required)
  - [ ] `tools/call` (session required)
- [ ] Protocol version compatible with platform baseline (`2024-11-05` unless upgraded deliberately)
- [ ] Capabilities advertise only what you implement
- [ ] REST admin routes (if any) use `APIError` shape — never as MCP body

## 3. Tools

- [ ] Each tool: stable `snake_case` name, description, `inputSchema`
- [ ] List/call names identical
- [ ] Registry schemas (if POSTed) match `tools/list`
- [ ] Domain failures → `result.isError` (not misused `-32603`)
- [ ] Auth/config failures are explicit and testable
- [ ] No decision/policy tools (authorize/approve/refuse) — data only

## 4. Security & ops

- [ ] Secrets via env; none in git/image
- [ ] Body limits + timeouts
- [ ] Structured logs without secrets
- [ ] Non-root container + HEALTHCHECK
- [ ] Production: `/mcp` behind platform/gateway auth (session ≠ identity)
- [ ] **If relay:** admin API key on register **and** confirm **and** all device/pairing admin routes (including GET) — do not copy hub open-admin gaps
- [ ] **If relay:** pairing start/confirm rate-limit (or accepted residual risk with owner)
- [ ] **If relay:** `HUB_REGISTER_API_KEY` + `HUB_ALLOWED_ORIGINS` set in production
- [ ] Follow [SECURITY_RULES](SECURITY_RULES.md)

## 5. Registry / Studio

- [ ] Choose `mcpServer` id and document it
- [ ] Confirm execute base URL will match registry routing
- [ ] Smoke: publish/approve or `tools/list` from MCP client
- [ ] If startup registry POST used: `TOOL_REGISTRY_BASE_URL` includes `/api/v1`

## 6. Quality

- [ ] Unit/smoke tests for initialize + session rejection + one tool call
- [ ] Prefer also covering parse error / invalid params / unknown tool (`-32700` / `-32602` / `-32601`)
- [ ] Conventions: [CODING_CONVENTIONS](CODING_CONVENTIONS.md)
- [ ] Patterns: [DESIGN_PATTERNS](DESIGN_PATTERNS.md)
- [ ] Architecture review against [ARCHITECTURE_RULES](ARCHITECTURE_RULES.md)
- [ ] Failure channels match [api-specification](../technical/api-specification.md) matrix

## Anti-checklist (tool-server)

Do **not** mark these done unless building a relay:

- [ ] ~~WebSocket `/ws` device registry~~
- [ ] ~~Pairing codes / bcrypt device file~~
- [ ] ~~`deviceID|rpcID` pending map~~
- [ ] ~~Embedded DTBClaw edge catalog~~
- [ ] ~~`forward_jsonrpc_to_device` hub tool~~
