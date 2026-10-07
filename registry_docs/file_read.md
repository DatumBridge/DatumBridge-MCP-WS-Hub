# file_read

Read file contents with line numbers. Supports partial reading via offset and limit. Extracts text from PDF; other binary files are read with lossy UTF-8 conversion. Sensitive files (for example .env and key material) are blocked by default.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `limit` | no | Maximum number of lines to return (default: all) |
| `offset` | no | Starting line number (1-based, default: 1) |
| `path` | yes | Path to the file. Relative paths resolve from workspace; outside paths require policy allowlist. |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "path": "example"
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
{
  "path": "example"
}
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
