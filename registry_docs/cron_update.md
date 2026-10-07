# cron_update

Patch an existing cron job (schedule, command, prompt, enabled, delivery, model, etc.)

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `approved` | no | Set true to explicitly approve medium/high-risk shell commands in supervised mode |
| `job_id` | yes | string |
| `patch` | yes | object |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "job_id": "example-id",
  "patch": "diff --git a/README.md b/README.md\n"
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
  "job_id": "example-id",
  "patch": "diff --git a/README.md b/README.md\n"
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
