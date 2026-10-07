# docx_read

Extract plain text from a DOCX (Word) file in the workspace. Returns all readable text content. No formatting, images, or charts.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `max_chars` | no | Maximum characters to return (default: 50000, max: 200000) |
| `path` | yes | Path to the DOCX file. Relative paths resolve from workspace. |

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
