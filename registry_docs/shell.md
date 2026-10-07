# shell

Execute a shell command in the workspace directory

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `approved` | no | Set true to explicitly approve medium/high-risk commands in supervised mode |
| `command` | yes | The shell command to execute |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "command": "example"
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
  "command": "example"
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
