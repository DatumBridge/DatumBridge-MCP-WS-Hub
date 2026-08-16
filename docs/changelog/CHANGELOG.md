# Changelog — MCP Tool Playbook Docs

## 2026-07-15

### Added

- Playbook pack under `docs/` for scaffolding new DatumBridge MCP tools:
  - Architecture taxonomy + ADR-0001 (transport vs tool-server)
  - Coding conventions, design patterns, architecture/security rules
  - Project structure trees (Python tool-server vs Go relay)
  - New-MCP checklist + Python / Go templates
  - Technical contracts (API, configuration, integrations)
  - Business rules, use-cases, workflows

### Changed

- Root `README.md`: link to playbook and `docs/playbook/NEW_MCP_CHECKLIST.md`; project structure lists `docs/`
- Root `design.md`: dual-role clarification (opaque proxy + MCP façade) + ADR / playbook links; §1 rewritten away from “not an MCP server”

### Fixed

- N/A

### Removed

- N/A
