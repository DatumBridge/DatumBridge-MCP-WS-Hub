# cron_add

Create a scheduled cron job (shell or agent) with cron/at/every schedules. Use job_type='agent' with a prompt to run the AI agent on schedule. Use schedule.kind='at' for one-time reminders/delayed sends (recommended). Agent jobs with schedule.kind='cron' or schedule.kind='every' are recurring and require explicit recurring confirmation. To deliver output to a channel (Discord, Telegram, Slack, Mattermost, QQ, Napcat, Lark, Feishu, Email), set delivery={"mode":"announce","channel":"discord","to":"<channel_id_or_chat_id>"}. This is the preferred tool for sending scheduled/delayed messages to users via channels.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `approved` | no | Set true to explicitly approve medium/high-risk shell commands in supervised mode |
| `command` | no | string |
| `delete_after_run` | no | boolean |
| `delivery` | no | Delivery config to send job output to a channel. Example: {"mode":"announce","channel":"discord","to":"<channel_id>"} |
| `job_type` | no | string |
| `model` | no | string |
| `name` | no | string |
| `prompt` | no | string |
| `recurring_confirmed` | no | Required for agent recurring schedules (schedule.kind='cron' or 'every'). Set true only when recurring behavior is intentional. |
| `schedule` | yes | Schedule object: {kind:'cron',expr,tz?} recurring / {kind:'at',at} one-time / {kind:'every',every_ms} recurring interval |
| `session_target` | no | string |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "schedule": {
    "key": "value"
  }
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
  "schedule": {
    "key": "value"
  }
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
