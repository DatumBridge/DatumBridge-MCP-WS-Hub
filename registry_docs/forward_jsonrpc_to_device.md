# forward_jsonrpc_to_device

Forward one JSON-RPC payload to a connected edge device.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Registered device with an active WebSocket |
| `jsonrpc_request` | yes | Full JSON-RPC object, for example a tools/list request |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "jsonrpc_request": {
    "key": "value"
  }
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
  "jsonrpc_request": {
    "key": "value"
  }
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
