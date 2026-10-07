# task_plan

Manage a task checklist for the current session. Use to break complex work into steps and track progress.
Actions: create (batch), add (single), update (change status), list (view all), delete (clear all).

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | yes | Operation to perform |
| `id` | no | For 'update': ID of the task to update |
| `status` | no | For 'update': new status |
| `tasks` | no | For 'create': list of tasks to create (replaces existing list) |
| `title` | no | For 'add': title of the new task |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "action": "example"
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
  "action": "example"
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
