# Architecture Rules — DatumBridge MCP

Mandatory boundaries when creating or reviewing a new MCP service.

## 1. Choose class first

1. **Tool-server (default):** tools execute in this process.
2. **Relay:** this process forwards to another MCP (usually edge over WS).
3. Do **not** invent a third hybrid without an ADR.

See [mcp-taxonomy.md](../architecture/mcp-taxonomy.md).

## 2. Layering

```text
Transport (HTTP/WS) → Protocol (MCP JSON-RPC / REST) → Domain (tools/services) → Upstream (APIs, DB, devices)
```

- Dependencies point **inward**: concrete clients depend on domain types, not on other MCP repos.
- Workflows / LangGraph must **not** embed provider SDK calls; they call MCP tools.
- Gateway/policy concerns stay outside tool handlers unless the service **is** the policy plane.

## 3. Single responsibility per module

| Module | Responsibility |
|--------|----------------|
| Entrypoint | Listen, signal, wire |
| MCP handler | Session + method dispatch |
| Tool handlers | Validate + invoke domain |
| Services | Upstream I/O |
| Auth store | Credentials only (relay) |
| Catalog | Tool descriptors (if catalog-driven) |

## 4. MCP surface honesty

- Advertise only capabilities you implement (`tools` typical).
- Unknown methods → `-32601`.
- Tool names in `tools/list` **must** equal `tools/call` names.
- Resources/prompts: opt-in only.

## 5. Hub-specific rules (relay)

- Least privilege: hub **relays**; it does not execute shell/browser/etc. itself.
- Correlation key includes both device and RPC id.
- Fail closed on edge capability flags.
- Studio path prefixes (`/api/ws-hub`) are handled by strip middleware or ingress — document either way.

## 6. Tool-server must not implement

- Device WebSocket registry / pairing / bcrypt device file
- `pendingReq` maps / HTTP↔WS timeout bridges
- Hub builtins (`forward_jsonrpc_to_device`) or mandatory `device_id` on every tool
- Embedded DTBClaw edge catalog sync **unless this service is the hub**

## 7. YAGNI / KISS

- No speculative config flags “for later.”
- No fat shared abstractions across MCP repos until rule-of-three.
- Prefer explicit switch/handlers over plugin frameworks unless already standard in that language (FastMCP is fine for Python).

## 8. Reversibility

- One concern per PR/service scaffold.
- Keep Docker/deploy contracts (`/health`, port env) stable.
- Document rollback: redeploy previous image; registry tool entries may need deactivation.

## 9. Validation before merge

New MCP PR should show:

- [ ] Taxonomy class declared in `design.md`
- [ ] `/health` + `/mcp` initialize session + tools list/call smoke
- [ ] Errors on correct channel (REST vs JSON-RPC vs `isError`)
- [ ] No hub-only files copied into tool-servers
- [ ] `.env.example` complete; secrets not committed
