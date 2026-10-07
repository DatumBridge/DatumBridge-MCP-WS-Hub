# memory_recall

Search long-term memory for relevant facts, preferences, or context. Returns scored results ranked by relevance.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `limit` | no | Max results to return (default: 5) |
| `query` | yes | Keywords or phrase to search for in memory |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "query": "quarterly report"
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
  "query": "quarterly report"
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
