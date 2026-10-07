# cron_run

Force-run a cron job immediately and record run history

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `approved` | no | Set true to explicitly approve medium/high-risk shell commands in supervised mode |
| `job_id` | yes | string |

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
