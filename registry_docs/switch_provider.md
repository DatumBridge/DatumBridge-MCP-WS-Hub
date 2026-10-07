# switch_provider

Switch to a different AI provider/model by updating config.toml. Use when current provider is rate-limited or when user explicitly requests a specific provider for a task. The change persists across requests.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `model` | no | Specific model (optional, e.g., 'gemini-2.5-flash', 'claude-opus-4') |
| `provider` | yes | Provider name (e.g., 'gemini', 'openai', 'anthropic') |
| `reason` | no | Reason for switching (for logging and user notification) |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "provider": "example"
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
  "provider": "example"
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
