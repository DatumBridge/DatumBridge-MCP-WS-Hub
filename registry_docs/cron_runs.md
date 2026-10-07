# cron_runs

List recent run history for a cron job

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `job_id` | yes | string |
| `limit` | no | integer |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "job_id": "example-id"
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
  "job_id": "example-id"
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
