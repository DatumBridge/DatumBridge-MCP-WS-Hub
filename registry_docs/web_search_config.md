# web_search_config

Inspect and update [web_search] configuration at runtime (providers, fallbacks, retries, provider-specific keys/options).

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | yes | Operation to perform |
| `api_key` | no | ['string', 'null'] |
| `brave_api_key` | no | ['string', 'null'] |
| `country` | no | ['string', 'null'] |
| `domain_filter` | no | string |
| `enabled` | no | boolean |
| `exa_api_key` | no | ['string', 'null'] |
| `exa_include_text` | no | boolean |
| `exa_search_type` | no | string |
| `fallback_providers` | no | string |
| `jina_api_key` | no | ['string', 'null'] |
| `jina_site_filters` | no | string |
| `language_filter` | no | string |
| `max_results` | no | integer |
| `max_tokens` | no | ['integer', 'null'] |
| `max_tokens_per_page` | no | ['integer', 'null'] |
| `perplexity_api_key` | no | ['string', 'null'] |
| `provider` | no | string |
| `recency_filter` | no | ['string', 'null'] |
| `retries_per_provider` | no | integer |
| `retry_backoff_ms` | no | integer |
| `timeout_secs` | no | integer |

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
