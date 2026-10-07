# git_operations

Perform structured Git operations (status, diff, log, branch, commit, add, checkout, stash). Provides parsed JSON output and integrates with security policy for autonomy controls.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | no | Stash action (for 'stash' operation) |
| `branch` | no | Branch name (for 'checkout' operation) |
| `cached` | no | Show staged changes (for 'diff' operation) |
| `files` | no | File or path to diff (for 'diff' operation, default: '.') |
| `index` | no | Stash index (for 'stash' with 'drop' action) |
| `limit` | no | Number of log entries (for 'log' operation, default: 10) |
| `message` | no | Commit message (for 'commit' operation) |
| `operation` | yes | Git operation to perform |
| `paths` | no | File paths to stage (for 'add' operation) |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "operation": "example"
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
  "operation": "example"
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
