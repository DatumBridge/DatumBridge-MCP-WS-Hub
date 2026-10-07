# process

Manage background processes: spawn long-running commands, check output, and terminate them

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | yes | Action to perform: spawn a process, list all, get output, or kill |
| `approved` | no | Approve medium/high-risk commands (for 'spawn') |
| `command` | no | Shell command to run in background (required for 'spawn') |
| `id` | no | Process ID returned by spawn (required for 'output' and 'kill') |

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
