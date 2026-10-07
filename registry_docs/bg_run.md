# bg_run

Execute a tool in the background and return a job ID immediately. Use this for long-running operations where you don't want to block. Check results with bg_status or wait for auto-injection in the next turn. Background tools have a 600-second maximum timeout.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `arguments` | no | Arguments to pass to the tool |
| `tool` | yes | Name of the tool to execute in the background |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "tool": "example"
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
  "tool": "example"
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
