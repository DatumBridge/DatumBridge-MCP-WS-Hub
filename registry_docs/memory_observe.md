# memory_observe

Store an observation entry in observation memory for long-horizon context continuity.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `category` | no | Optional category override. Defaults to 'observation'. |
| `confidence` | no | Optional confidence score in [0.0, 1.0]. |
| `key` | no | Optional custom key. Auto-generated when omitted. |
| `observation` | yes | Observation to capture (fact, pattern, or running context signal) |
| `source` | no | Optional source label for traceability (e.g. 'chat', 'tool_result'). |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "observation": "example"
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
  "observation": "example"
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
