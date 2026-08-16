# Database Design

## Hub

This service does **not** use a relational or document database.

| Store | Mechanism | Contents |
|-------|-----------|----------|
| Device credentials | JSON file (`HUB_CREDENTIALS_FILE`) | bcrypt password hashes, device metadata |
| Pairing codes | In-memory (TTL) | ephemeral codes |
| MCP sessions | In-memory (TTL) | session ids |
| Pending RPC | In-memory map | correlation waiters |
| WS connections | In-memory map | live `Conn`s |

**Backup/restore:** treat the credentials file as sensitive state; do not commit it. Prefer volume mounts in K8s when tokens must survive restarts.

## New MCP tools

- Use a database only when the product requires durable domain data (example: social-listening MongoDB).
- Document schema in that service’s own `docs/technical/database-design.md`.
- Do not add a database to a pure relay or stateless API wrapper without a concrete need.
