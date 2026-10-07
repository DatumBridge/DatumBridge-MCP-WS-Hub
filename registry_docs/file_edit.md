# file_edit

Edit a file by replacing text in a file. Exact matching is preferred; if exact matching fails, whitespace-flexible line matching is used. Sensitive files (for example .env and key material) are blocked by default.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `new_string` | yes | The replacement text (empty string to delete the matched text) |
| `old_string` | yes | The text to find and replace. Exact matching is attempted first; if no exact match is found, whitespace-flexible line matching is attempted. |
| `path` | yes | Path to the file. Relative paths resolve from workspace; outside paths require policy allowlist. |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "new_string": "example",
  "old_string": "example",
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
  "new_string": "example",
  "old_string": "example",
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
