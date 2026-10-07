# file_write

Write contents to a file in the workspace. Sensitive files (for example .env and key material) are blocked by default.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `content` | yes | Content to write to the file |
| `path` | yes | Path to the file. Relative paths resolve from workspace; outside paths require policy allowlist. |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "content": "hello",
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
  "content": "hello",
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
