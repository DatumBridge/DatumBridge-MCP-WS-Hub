# openclaw_migration

Preview or execute merge-first migration from OpenClaw (memory + config + agents) without overwriting existing OctoClaw data.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | no | preview runs a dry-run report; migrate applies merge changes |
| `include_config` | no | Whether to migrate provider/channels/agents config (default true) |
| `include_memory` | no | Whether to migrate memory entries (default true) |
| `source_config` | no | Optional OpenClaw config path (default ~/.openclaw/openclaw.json) |
| `source_workspace` | no | Optional OpenClaw workspace path (default ~/.openclaw/workspace) |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1"
}
```

Output:

```json
{
  "success": true,
  "result": "completed"
}
```

### Missing `device_id`

The tool rejects the call and does not guess the missing value.

Input:

```json
{}
```

Output:

```json
{
  "success": false,
  "error": {
    "error_code": "invalid_argument",
    "error_message": "device_id is required",
    "retryable": false
  }
}
```
