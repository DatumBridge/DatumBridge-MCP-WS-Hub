# check_provider_quota

Check current rate limit and quota status for AI providers. Returns available providers, rate-limited providers, quota remaining, and estimated reset time. Use this when user asks about model availability or when you encounter rate limit errors.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `provider` | no | Specific provider to check (optional). Examples: openai, gemini, anthropic. If omitted, checks all providers. |

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
