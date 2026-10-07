# memory_forget

Remove a memory by key. Use to delete outdated facts or sensitive data. Returns whether the memory was found and removed.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `key` | yes | The key of the memory to forget |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "key": "example"
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
  "key": "example"
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
