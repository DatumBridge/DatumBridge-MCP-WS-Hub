# proxy_config

Manage OctoClaw proxy settings (scope: environment | octoclaw | services), including runtime and process env application

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | no | string |
| `all_proxy` | no | Fallback proxy URL for all protocols |
| `clear_env` | no | When action=disable, clear process proxy environment variables |
| `enabled` | no | Enable or disable proxy |
| `http_proxy` | no | HTTP proxy URL |
| `https_proxy` | no | HTTPS proxy URL |
| `no_proxy` | no | Comma-separated string or array of NO_PROXY entries |
| `scope` | no | Proxy scope: environment / octoclaw / services |
| `services` | no | Comma-separated string or array of service selectors used when scope=services |

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
