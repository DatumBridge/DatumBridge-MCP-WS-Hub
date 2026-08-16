# Workflows — Scaffold to Publish

## New tool-server

```text
1. Pick Python template
2. Implement tools + schemas + services
3. Verify /health and /mcp initialize→list→call
4. Containerize (non-root)
5. Deploy via deploy-service / K8s
6. Publish/approve in datumbridge-mcp (discovers tools/list)
7. Wire Studio / LangGraph nodes to registry tool ids
```

## New relay (rare)

```text
1. Confirm topology needs WS/device auth
2. Use Go relay template
3. Implement auth + connect + correlation + forward
4. Add /mcp façade if Studio must list tools
5. Optionally sync catalog to registry
6. Configure Studio upstream prefix
7. Pair devices; smoke forward + tools/call with device_id
```

## Change an existing edge tool

```text
1. Implement on device (DTBClaw/OctoClaw)
2. Regenerate/export catalog
3. Update hub embedded catalog (if hub advertises it)
4. Redeploy hub; confirm registry sync / tools/list
```

Checklist detail: [NEW_MCP_CHECKLIST](../playbook/NEW_MCP_CHECKLIST.md).
