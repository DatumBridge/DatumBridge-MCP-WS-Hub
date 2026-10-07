# http_request

Make HTTP requests to external APIs. Supports GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS methods. Security constraints: allowlist-only domains, no local/private hosts, configurable timeout/response size limits, and optional env-backed credential profiles.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `body` | no | Optional request body (for POST, PUT, PATCH requests) |
| `credential_profile` | no | Optional profile name from [http_request.credential_profiles]. Lets the harness inject credentials from environment variables without passing raw tokens in tool arguments. |
| `headers` | no | Optional HTTP headers as key-value pairs (e.g., {"Authorization": "Bearer token", "Content-Type": "application/json"}) |
| `method` | no | HTTP method (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS) |
| `url` | yes | HTTP or HTTPS URL to request |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "body": "Hello from Weaver",
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
  "body": "Hello from Weaver",
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
