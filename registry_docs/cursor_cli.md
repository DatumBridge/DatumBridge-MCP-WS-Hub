# cursor_cli

Control Cursor via its CLI. Key operation: 'agent' runs Cursor Agent headlessly in terminal (no IDE needed) — it reads, writes, and runs shell commands on the workspace autonomously. Other operations: open, goto, diff, prompt, install_extension, uninstall_extension, list_extensions, new_window, status, help. Requires the `cursor` command on PATH. Call with {"operation":"help"} for full usage.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `column` | no | Column number (for 'goto', default: 1) |
| `extension_id` | no | Extension marketplace ID, e.g. 'ms-python.python' (for install/uninstall) |
| `file` | no | File path (for 'goto') |
| `file1` | no | First file path (for 'diff') |
| `file2` | no | Second file path (for 'diff') |
| `force` | no | Auto-approve all agent actions without prompting (for 'agent', default: true) |
| `headless` | no | Run agent in headless mode (--print, no terminal window). Default: false (opens a visible terminal) |
| `line` | no | Line number (for 'goto', default: 1) |
| `message` | no | Prompt or instruction text to send to Cursor AI (for 'prompt') |
| `mode` | no | How to use the prompt in Cursor (for 'prompt'): 'composer' = Cmd+I multi-file edits, 'chat' = Cmd+L ask questions, 'edit' = Cmd+K inline code edits, 'rules' = save as project rule. Default: 'composer' |
| `model` | no | AI model to use (for 'agent'), e.g. 'gpt-5', 'sonnet-4', 'sonnet-4-thinking' |
| `operation` | yes | Cursor CLI operation. 'agent' starts Cursor Agent in a terminal. 'agent_status' polls agent progress and reads output. Use 'help' for full guide. |
| `output_format` | no | Output format for agent responses (for 'agent', default: 'text') |
| `path` | no | File or folder path (for 'open' and 'new_window') |
| `reuse_window` | no | Reuse existing window instead of opening a new one (for 'open', default: true) |
| `tail` | no | Number of output lines to return (for 'agent_status', default: 50) |
| `wait` | no | Wait for the file to be closed before returning (for 'open', default: false) |
| `workspace` | no | Workspace directory for agent to operate on (for 'agent'/'agent_status', defaults to current workspace) |

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
