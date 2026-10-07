# apply_patch

Safely check/apply a unified diff to the current git repository, optionally staging and committing.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `commit_message` | no | If provided (and dry_run=false), stage all changes and create a git commit with this message. |
| `dry_run` | no | If true, only checks whether the patch would apply cleanly (no changes made). |
| `patch` | yes | Unified diff text (e.g. output of `git diff`). |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "patch": "diff --git a/README.md b/README.md\n"
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
  "patch": "diff --git a/README.md b/README.md\n"
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
