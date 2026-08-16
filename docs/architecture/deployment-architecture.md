# Deployment Architecture

## Hub

| Aspect | Practice |
|--------|----------|
| Build | Multi-stage Go build, `CGO_ENABLED=0` |
| Runtime user | Non-root (`appuser`) |
| Port | `HUB_PORT` (default `8000`) |
| Health | `GET /health` (Docker HEALTHCHECK) |
| Secrets | Runtime env / K8s Secret — not image `ENV` |
| State | Credentials file path via `HUB_CREDENTIALS_FILE` (PVC or emptyDir as appropriate) |
| Platform deploy | `datumbridge-deploy-service` builds from this tree |
| Studio ingress | `WS_HUB_UPSTREAM` → public `/api/ws-hub` (hub may strip prefix) |

Details and kubectl examples: root [README.md](../../README.md).

## New MCP tools

Same ops shape:

1. Container listens on documented port.
2. `/health` for probes.
3. `/mcp` reachable from `datumbridge-mcp` / Studio.
4. Env for provider credentials.
5. Optional registry sync if not using publish flow alone.

Do not require WebSocket Services/Ingresses unless building a relay.
