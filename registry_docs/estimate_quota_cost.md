# estimate_quota_cost

Estimate quota cost (tokens, requests) for an operation before executing it. Useful for warning user if operation may exhaust quota or when planning parallel tool calls.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `estimated_tokens` | no | Estimated input+output tokens (optional, default: 1000) |
| `operation` | yes | Operation type |
| `parallel_count` | no | Number of parallel operations (if applicable, default: 1) |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "operation": "example"
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
  "operation": "example"
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
