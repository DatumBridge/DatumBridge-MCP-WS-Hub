# bg_status

Query the status of a background job by ID, or list all jobs if no ID provided. Returns job status (running/complete/failed), result output, and elapsed time.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `job_id` | no | Optional job ID to query. If omitted, returns all jobs. |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1"
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
{}
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
