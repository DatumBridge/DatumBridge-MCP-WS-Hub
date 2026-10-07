# web_fetch

Fetch a web page and return text content for LLM consumption. Strips navigation, scripts, and boilerplate before extraction. Providers: nanohtml2text (default), firecrawl, tavily. Deprecated alias: fast_html2md. Security: allowlist-only domains, blocked_domains, and no local/private hosts.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `url` | yes | The HTTP or HTTPS URL to fetch |

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
