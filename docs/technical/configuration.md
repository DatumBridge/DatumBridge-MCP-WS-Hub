# Configuration

## Principles

- Env-first configuration; document every variable in `.env.example`.
- No secrets in images or git.
- Fail fast on required missing credentials.
- Service-prefixed names (`HUB_*`, `TOOL_REGISTRY_*`, provider-specific).

## Hub reference variables

| Variable | Role | Default / notes |
|----------|------|-----------------|
| `HUB_PORT` | Listen port | `8000` |
| `HUB_CREDENTIALS_FILE` | Device credential store path | `./data/devices.json` |
| `HUB_REGISTER_API_KEY` | Protect registration | **Required in production**; empty = open local-dev only |
| `HUB_ALLOWED_ORIGINS` | CORS + WS Origin allowlist | **Required in production** (exact match); empty = permissive local-dev only |
| `HUB_PAIRING_TTL` | Pairing code lifetime | Go duration string |
| `HUB_HTTP_STRIP_PREFIX` | Strip `/api/ws-hub` | on unless `0/false/no` |
| `HUB_EXPECTED_EDGE_VERSION` | Drift warning | optional |
| `HUB_FORWARD_SEND_RETRIES` / `_INTERVAL_MS` | WS send buffer retries | optional |
| `LOG_LEVEL` | `DEBUG\|INFO\|WARN\|ERROR` | `INFO` |
| `TOOL_REGISTRY_BASE_URL` | Startup catalog sync | must include `/api/v1` |
| `TOOL_REGISTRY_API_KEY` | Registry auth | optional |
| `WS_HUB_REGISTRY_MCPSERVER_NAME` | Registry `mcpServer` field | `datumbridge_mcp_ws_hub` |

Source of truth for hub: root [`.env.example`](../../.env.example).

## New tool-server pattern

```bash
PORT=8000
LOG_LEVEL=INFO
# Required provider credentials — fail if unset
PROVIDER_API_TOKEN=
# Optional
TOOL_REGISTRY_BASE_URL=
TOOL_REGISTRY_API_KEY=
```

Choose a clear prefix per service (`DRIVE_`, `SEARXNG_`, `APIFY_`, …).
