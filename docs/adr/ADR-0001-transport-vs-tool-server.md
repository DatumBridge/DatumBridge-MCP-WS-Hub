# ADR-0001 — Transport (relay) vs Tool-Server

## Context

Historically, root `design.md` described `datumbridge-mcp-ws-hub` as a **transparent HTTP↔WebSocket proxy that is not an MCP server**. The shipped code also implements:

- Streamable HTTP `POST /mcp` (`initialize`, `tools/list`, `tools/call`)
- Built-in hub tools (`hub_info`, `forward_jsonrpc_to_device`)
- Embedded DTBClaw edge catalog (~45 tools) advertised for publish/approve
- Optional startup sync to DatumBridge Tool Registry

Root `design.md` now documents this **dual role** (core transport + MCP façade) and links here. Without that recorded decision, new MCP authors may either (a) treat the hub as “proxy only” and omit `/mcp`, or (b) clone WS/correlation into every tool-server.

## Decision

1. **Default new MCP** = **tool-server**: implement tools locally; expose Streamable HTTP `/mcp`; do not implement device WS transport.
2. **This hub** is a **relay-first** service with an **optional MCP Streamable HTTP façade** so Studio/`datumbridge-mcp` can publish/approve and invoke relayed edge tools.
3. Original `design.md` “proxy thesis” remains valid for the **core transport path** (`POST /api/v1/devices/{id}/mcp` + `/ws`). The `/mcp` façade is an **additive platform integration**, not a license to copy relay internals into tool-servers.
4. Edge tool **execution** stays on the device; the hub **relays** and **advertises**; it does not re-implement edge tool logic.

## Alternatives Considered

| Alternative | Why rejected |
|-------------|--------------|
| Keep hub proxy-only (remove `/mcp`) | Breaks DatumBridge publish/approve client expectations |
| Make every MCP a mini-hub (WS + correlation) | Explosion of auth surfaces; wrong trust boundary |
| Edge tools only via opaque forward (no catalog) | Studio cannot list/approve tools consistently |

## Consequences

- Playbooks split **Python/Go tool-server** vs **Go relay** templates.
- Hub-only patterns (pairing, correlation, edge catalog sync) are documented under “relay exception,” not coding style defaults.
- Product ops docs stay in root `README.md`; this ADR is the role-truth for scaffolders.
- Root `design.md` annotation (dual-role callout + ADR link): **done**.

## Trade-offs

- Hub codebase is denser than a pure proxy (MCP session store + catalog).
- Readers must read taxonomy before copying structure.

## Risks

- Doc/code drift if `/mcp` grows without updating playbooks.
- Mis-sync of edge catalog versions vs device binaries.
