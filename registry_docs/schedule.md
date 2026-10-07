# schedule

Manage scheduled shell-only tasks. Actions: create/add/once/list/get/cancel/remove/pause/resume. WARNING: This tool creates shell jobs whose output is only logged, NOT delivered to any channel. To send a scheduled message to Discord/Telegram/Slack/Mattermost/QQ/Lark/Feishu/Email, use the cron_add tool with job_type='agent' and a delivery config like {"mode":"announce","channel":"discord","to":"<channel_id>"}.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | yes | Action to perform |
| `approved` | no | Set true to explicitly approve medium/high-risk shell commands in supervised mode |
| `command` | no | Shell command to execute. Required for create/add/once. |
| `delay` | no | Delay for one-shot tasks (e.g. '30m', '2h', '1d'). |
| `expression` | no | Cron expression for recurring tasks (e.g. '*/5 * * * *'). |
| `id` | no | Task ID. Required for get/cancel/remove/pause/resume. |
| `run_at` | no | Absolute RFC3339 time for one-shot tasks (e.g. '2030-01-01T00:00:00Z'). |

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
