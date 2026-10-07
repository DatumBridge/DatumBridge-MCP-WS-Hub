# manage_auth_profile

Manage auth profiles: list all profiles with token status, switch active profile for a provider, or refresh expired OAuth tokens. Use when user asks about accounts, tokens, or when you encounter expired/rate-limited credentials.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | yes | Action to perform: 'list' shows all profiles, 'switch' changes active profile, 'refresh' renews OAuth tokens |
| `profile` | no | Profile name to switch to (for 'switch' action). E.g., 'default', 'work', 'personal'. |
| `provider` | no | Provider name (e.g., 'gemini', 'openai-codex', 'anthropic'). Required for switch and refresh. |

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
