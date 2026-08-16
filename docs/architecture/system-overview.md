# System Overview

This repository is the **MCP WebSocket Hub**: a DatumBridge **relay/transport** that bridges cloud HTTP clients to edge MCP servers (DTBClaw) over WebSocket, and exposes a Streamable HTTP `/mcp` façade for Tool Registry publish/approve.

It is also the **reference repo** for the reusable MCP scaffolding playbooks under `docs/playbook/`.

## What changed (docs pack)

| Area | Change | Impact |
|------|--------|--------|
| Docs | Added playbook + taxonomy + ADR-0001 | New MCP tools have a copyable contract |
| Runtime | Unchanged by the docs pack | Hub behavior same |

## Impacted components

- Authors of new MCP tools (primary audience)
- Reviewers validating scaffolds against rules
- Hub maintainers (taxonomy clarifies dual role: proxy + façade)

## Dependencies

- Platform MCP client expects Streamable HTTP session header
- Optional `TOOL_REGISTRY_*` for catalog sync
- Studio nginx `/api/ws-hub` path strip

## Risks

- Authors may still copy hub WS/correlation into tool-servers — playbook anti-checklists mitigate this; reinforce in PR review.
- `/mcp` session is continuity, not caller identity — production must front the hub with Studio/gateway auth or network policy (see [SECURITY_RULES](../playbook/SECURITY_RULES.md)).

## Related

- [mcp-taxonomy.md](mcp-taxonomy.md)
- [docs/README.md](../README.md)
- Root [README.md](../../README.md) for ops
