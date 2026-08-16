# Integrations

## Tool Registry (`datumbridge-mcp`)

Two paths to get tools into Studio:

### Path 1 — Discover via MCP (typical tool-server)

1. Deploy MCP with reachable `/mcp`.
2. Publish/approve client: `initialize` → `tools/list`.
3. Registry stores `name`, `description`, `schema.inputSchema`, `mcpServer`, base URL.

### Path 2 — Explicit POST (hub catalog sync)

1. Set `TOOL_REGISTRY_BASE_URL` (must include `/api/v1`).
2. Optional `TOOL_REGISTRY_API_KEY` (`Bearer` and/or `X-API-Key`).
3. Startup POSTs each tool body to `{base}/tools`.
4. Success HTTP `200`/`201`; upsert by `id` + `version`.

Hub catalog bodies include edge metadata (`edgeDevice.profile`, `edgeDevice.mcpToolName`) — **relay-only**.

**Rule:** Registry `mcpServer` + execute URL must point at the same MCP that advertises the tool.

## DatumBridge Studio

| Integration | Notes |
|-------------|-------|
| Tool nodes | Resolve via registry; call MCP `/mcp` |
| Hub public URL | `/api/ws-hub` → Studio nginx → hub Service |
| Path strip | Hub middleware may remove `/api/ws-hub` prefix |

## Edge / DTBClaw **[hub]**

- Devices connect `/ws?device_id=&token=`
- Control-plane hello / handshake separate from MCP responses
- Edge catalog embedded in hub; regenerate from device export tooling when tools change

## Correlation headers

Accept/propagate `X-Correlation-ID` or W3C `traceparent` when forwarding (**hub** `ForwardRequestOpts`); tool-servers should similarly pass correlation to upstreams when available.

## Out of scope for normal tool-servers

Do not integrate device WS pairing or edge catalog exporters unless you are extending the hub.
