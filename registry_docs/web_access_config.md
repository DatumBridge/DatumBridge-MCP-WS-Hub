# web_access_config

Inspect and update shared network URL access policy ([security.url_access]) including first-visit approval, global allowlist/blocklist, and approved domains.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | yes | Operation to perform |
| `add_approved_domains` | no | string |
| `add_domain_allowlist` | no | string |
| `add_domain_blocklist` | no | string |
| `allow_cidrs` | no | string |
| `allow_domains` | no | string |
| `allow_loopback` | no | boolean |
| `approved_domains` | no | string |
| `block_private_ip` | no | boolean |
| `domain_allowlist` | no | string |
| `domain_blocklist` | no | string |
| `enforce_domain_allowlist` | no | boolean |
| `remove_approved_domains` | no | string |
| `remove_domain_allowlist` | no | string |
| `remove_domain_blocklist` | no | string |
| `require_first_visit_approval` | no | boolean |
| `url` | no | string |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "action": "example",
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
  "action": "example",
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
