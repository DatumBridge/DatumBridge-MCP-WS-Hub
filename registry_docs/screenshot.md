# screenshot

Capture a screenshot of the current screen. Returns the file path and base64-encoded PNG data.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `filename` | no | Optional filename (default: screenshot_<timestamp>.png). Saved in workspace. |
| `region` | no | Optional region for macOS: 'selection' for interactive crop, 'window' for front window. Ignored on Linux. |

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
