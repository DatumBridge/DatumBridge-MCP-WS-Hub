# pushover

Send a Pushover notification to your device. Uses PUSHOVER_TOKEN/PUSHOVER_USER_KEY from process environment first, then falls back to .env.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `message` | yes | The notification message to send |
| `priority` | no | Message priority: -2 (lowest/silent), -1 (low/no sound), 0 (normal), 1 (high), 2 (emergency/repeating) |
| `sound` | no | Notification sound override (e.g., 'pushover', 'bike', 'bugle', 'cashregister', etc.) |
| `title` | no | Optional notification title |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "message": "example"
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
  "message": "example"
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
