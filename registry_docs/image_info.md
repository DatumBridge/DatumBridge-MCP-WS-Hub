# image_info

Read image file metadata (format, dimensions, size) and optionally return base64-encoded data.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `include_base64` | no | Include base64-encoded image data in output (default: false) |
| `path` | yes | Path to the image file (absolute or relative to workspace) |

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
