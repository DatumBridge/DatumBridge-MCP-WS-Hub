# Project Structure Templates

Two allowed trees. Do not invent a hybrid.

---

## A. Python tool-server (default)

```text
my-mcp-tool/
├── README.md
├── design.md
├── .env.example
├── .gitignore
├── Dockerfile
├── requirements.txt          # or pyproject.toml
├── mcp_server_entrypoint.sh  # optional, matches siblings
├── app/
│   ├── __init__.py
│   ├── mcp_server.py         # FastMCP / tools surface
│   ├── core/
│   │   ├── exceptions.py
│   │   └── ...
│   ├── schemas/
│   │   └── ...
│   └── services/
│       └── ...
├── scripts/                  # local smoke tests only
└── docs/                     # product docs as needed
```

**Must include:** `/health` (or platform equivalent), Streamable HTTP MCP entry, env example, non-root Dockerfile.

**Must not include:** `internal/hub/ws.go`-style device registry, pairing stores, edge catalogs.

Reference siblings: `google-drive-mcp`, `searxng-web-search-mcp`, `social-listening`.

---

## B. Go relay / hub-style (exception)

```text
my-mcp-relay/
├── README.md
├── design.md
├── .env.example
├── Dockerfile
├── go.mod
├── go.sum
├── cmd/
│   └── api/
│       ├── main.go           # bootstrap, routes, shutdown
│       └── web/              # optional embedded UI (//go:embed)
├── internal/
│   └── <domain>/             # e.g. hub
│       ├── http.go           # REST + APIError
│       ├── mcp_http.go       # Streamable HTTP subset
│       ├── mcp_session.go
│       ├── ws.go             # only if WS required
│       ├── auth.go
│       ├── middleware.go
│       ├── *_test.go
│       └── ...
└── data/                     # runtime state; gitignored
```

**Keep `cmd` thin.** All domain logic under `internal/`.

When **not** building a device relay, omit `ws.go`, pairing, correlation, and catalog sync — or prefer Template A instead.

Reference: this repository (`datumbridge-mcp-ws-hub`).

---

## Shared operational files

| File | Purpose |
|------|---------|
| `README.md` | Architecture diagram, endpoints, env table, security bullets |
| `design.md` | Design decisions / informal ADR |
| `.env.example` | Config contract |
| `Dockerfile` | Multi-stage, non-root, HEALTHCHECK |
| `.dockerignore` | Exclude `.env`, `data/`, secrets |

---

## Mapping hub files → new services

| Hub path | Tool-server? | Relay? |
|----------|--------------|--------|
| `cmd/api/main.go` | Adapt as thin entry | Yes |
| `internal/hub/mcp_http.go` | Port session+tools pattern | Yes |
| `internal/hub/http.go` (`APIError`) | Yes (REST if any) | Yes |
| `internal/hub/middleware.go` | CORS/logging yes; path strip if Studio-proxied | Yes |
| `internal/hub/ws.go` | **No** | If WS needed |
| `internal/hub/auth.go` (device bcrypt) | **No** | If device auth |
| `internal/hub/dtbclaw_*.json` | **No** | Hub/edge only |
| `internal/hub/registry_sync.go` | Optional thin publish path | Optional |
