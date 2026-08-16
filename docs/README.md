# DatumBridge MCP Tool Playbook

Reusable **rules, coding conventions, and design patterns** extracted from this repository (`datumbridge-mcp-ws-hub`) and sibling DatumBridge MCP services. Use this pack when scaffolding **new** MCP tools — do not treat the hub runtime as a default template.

## Start here

1. Read [MCP taxonomy](architecture/mcp-taxonomy.md) — decide **tool-server** (default) vs **relay/transport** (rare).
2. Follow [NEW_MCP_CHECKLIST](playbook/NEW_MCP_CHECKLIST.md).
3. Apply [CODING_CONVENTIONS](playbook/CODING_CONVENTIONS.md) + [DESIGN_PATTERNS](playbook/DESIGN_PATTERNS.md).
4. Pick a template:
   - **Default:** [TEMPLATE_PYTHON_TOOL_SERVER](playbook/TEMPLATE_PYTHON_TOOL_SERVER.md)
   - **Only if building HTTP↔device/WS transport:** [TEMPLATE_GO_RELAY](playbook/TEMPLATE_GO_RELAY.md)

## Document map

| Area | Document |
|------|----------|
| Taxonomy & roles | [architecture/mcp-taxonomy.md](architecture/mcp-taxonomy.md), [ADR-0001](adr/ADR-0001-transport-vs-tool-server.md) |
| Rules | [ARCHITECTURE_RULES](playbook/ARCHITECTURE_RULES.md), [SECURITY_RULES](playbook/SECURITY_RULES.md) |
| Conventions & patterns | [CODING_CONVENTIONS](playbook/CODING_CONVENTIONS.md), [DESIGN_PATTERNS](playbook/DESIGN_PATTERNS.md), [PROJECT_STRUCTURE](playbook/PROJECT_STRUCTURE.md) |
| Scaffold | [NEW_MCP_CHECKLIST](playbook/NEW_MCP_CHECKLIST.md), Python / Go templates |
| Platform contracts | [technical/api-specification.md](technical/api-specification.md), [configuration](technical/configuration.md), [integrations](technical/integrations.md) |
| Business rules | [business/business-rules.md](business/business-rules.md), [workflows](business/workflows.md), [use-cases](business/use-cases.md) |
| Changelog | [changelog/CHANGELOG.md](changelog/CHANGELOG.md) |

## Hard rule

**Most new MCP tools are Python (or language) tool-servers.** Copy `/health`, Streamable HTTP `/mcp`, typed tool schemas, fail-fast config, standardized errors, and non-root Docker.

**Do not copy** from this hub into a normal tool-server: WebSocket device registry, pairing, `deviceID|rpcID` correlation, edge catalogs, or `forward_jsonrpc_to_device`. Those are relay-only.

Hub **ops** (deploy, env tables, Studio `/api/ws-hub`) stay in the repository root [README.md](../README.md).
