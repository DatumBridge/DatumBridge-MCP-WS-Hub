# content_search

Search file contents by regex pattern within the workspace. Supports ripgrep (rg) with grep fallback. Output modes: 'content' (matching lines with context), 'files_with_matches' (file paths only), 'count' (match counts per file). Example: pattern='fn main', include='*.rs', output_mode='content'.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `case_sensitive` | no | Case-sensitive matching. Defaults to true |
| `context_after` | no | Lines of context after each match (content mode only) |
| `context_before` | no | Lines of context before each match (content mode only) |
| `include` | no | File glob filter, e.g. '*.rs', '*.{ts,tsx}' |
| `max_results` | no | Maximum number of results to return. Defaults to 1000 |
| `multiline` | no | Enable multiline matching (ripgrep only, errors on grep fallback) |
| `output_mode` | no | Output format: 'content' (matching lines), 'files_with_matches' (paths only), 'count' (match counts) |
| `path` | no | Directory to search in, relative to workspace root. Defaults to '.' |
| `pattern` | yes | Regular expression pattern to search for |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "pattern": "example"
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
  "pattern": "example"
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
