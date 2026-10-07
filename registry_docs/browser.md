# browser

Web/browser automation with pluggable backends (agent-browser, rust-native, computer_use). Supports DOM actions plus optional OS-level actions (mouse_move, mouse_click, mouse_drag, key_type, key_press, screen_capture) through a computer-use sidecar. Use 'snapshot' to map interactive elements to refs (@e1, @e2). Enforces browser.allowed_domains for open actions.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | yes | Browser action to perform (OS-level actions require backend=computer_use) |
| `button` | no | Mouse button for computer_use mouse_click |
| `by` | no | For find: semantic locator type |
| `compact` | no | For snapshot: remove empty structural elements |
| `depth` | no | For snapshot: limit tree depth |
| `direction` | no | Scroll direction |
| `fill_value` | no | For find with fill action: value to fill |
| `find_action` | no | For find: action to perform on found element |
| `from_x` | no | Drag source X coordinate (computer_use: mouse_drag) |
| `from_y` | no | Drag source Y coordinate (computer_use: mouse_drag) |
| `full_page` | no | For screenshot: capture full page |
| `interactive_only` | no | For snapshot: only show interactive elements |
| `key` | no | Key to press (Enter, Tab, Escape, etc.) |
| `ms` | no | Milliseconds to wait |
| `path` | no | File path for screenshot |
| `pixels` | no | Pixels to scroll |
| `selector` | no | Element selector: @ref (e.g. @e1), CSS (#id, .class), or text=... |
| `text` | no | Text to type or wait for |
| `to_x` | no | Drag target X coordinate (computer_use: mouse_drag) |
| `to_y` | no | Drag target Y coordinate (computer_use: mouse_drag) |
| `url` | no | URL to navigate to (for 'open' action) |
| `value` | no | Value to fill or type |
| `x` | no | Screen X coordinate (computer_use: mouse_move/mouse_click) |
| `y` | no | Screen Y coordinate (computer_use: mouse_move/mouse_click) |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "action": "example",
  "text": "hello",
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
  "text": "hello",
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
