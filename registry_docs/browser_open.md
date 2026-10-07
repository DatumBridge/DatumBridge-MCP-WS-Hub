# browser_open

Open an approved HTTPS URL in a browser. Security constraints: allowlist-only domains, no local/private hosts, no scraping.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `url` | yes | HTTPS URL to open |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "url": "https://example.com/page"
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
  "url": "https://example.com/page"
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
