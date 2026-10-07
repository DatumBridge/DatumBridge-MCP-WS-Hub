# edge_agent_run

Bounded edge mission: run a high-level objective on this gateway with strict max steps and wall time. Enforces operator caps; returns an acceptance envelope with negotiated budgets (Approach 2/3 contract).

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `correlation_id` | no | Optional id to join master run ↔ edge mission in logs. |
| `max_steps` | no | Requested maximum tool/agent steps (clamped by gateway). |
| `max_wall_seconds` | no | Requested wall-clock budget in seconds (clamped by gateway). |
| `mission_protocol` | no | Expected mission protocol version from master (must match edge). |
| `objective` | yes | High-level task for the edge agent loop. |
| `tool_allowlist` | no | Optional subset of MCP tool names the mission may invoke (gateway may further restrict). |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "objective": "example"
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
  "objective": "example"
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
