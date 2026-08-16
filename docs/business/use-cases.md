# Use Cases — Choosing an MCP Shape

| Use case | Choose | Why |
|----------|--------|-----|
| Wrap Google Drive / SaaS API | Python tool-server | OAuth + API in-process |
| Web search / crawl / enrich | Python tool-server | Upstream HTTP APIs |
| Store + analyze social posts | Python tool-server | Domain DB + LLM in service |
| Reach MCP on laptop/Pi/edge | Go relay (hub) | NAT / no inbound to device |
| Advertise many edge tools in Studio | Hub + edge catalog | Catalog sync + relay call |
| Transparent JSON-RPC to one device | Hub forward path | `POST .../devices/{id}/mcp` |

If unsure: start with tool-server. Promote to relay only when network topology requires it.
