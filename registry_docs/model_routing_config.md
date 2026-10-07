# model_routing_config

Manage default model settings, scenario routes, classification rules, delegate profiles, and agent team/subagent orchestration controls. Designed for natural-language runtime reconfiguration (enable/disable, strategy, and capacity tuning).

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | no | string |
| `agentic` | no | Enable tool-call loop mode for delegate agent |
| `allowed_tools` | no | Allowed tools for agentic delegate mode (string or string array) |
| `api_key` | no | Optional API key override for scenario route or delegate agent |
| `capabilities` | no | Capability tags for automatic agent selection (string or string array) |
| `classification_enabled` | no | When true, upsert classification rule for this hint; false removes it |
| `enabled` | no | Enable or disable a delegate profile for selection/invocation |
| `hint` | no | Scenario hint name (for example: conversation, coding, reasoning) |
| `keywords` | no | Classification keywords for upsert_scenario (string or string array) |
| `max_concurrent_subagents` | no | Maximum number of concurrently running background sub-agents (positive integer, no hard-coded upper cap) |
| `max_depth` | no | Delegate max recursion depth |
| `max_iterations` | no | Maximum tool-call iterations for agentic delegate mode |
| `max_length` | no | Optional maximum message length matcher |
| `max_team_agents` | no | Maximum number of delegate profiles activated for teams (positive integer, no hard-coded upper cap) |
| `min_length` | no | Optional minimum message length matcher |
| `model` | no | Model for set_default/upsert_scenario/upsert_agent |
| `name` | no | Delegate sub-agent name for upsert_agent/remove_agent |
| `patterns` | no | Classification literal patterns for upsert_scenario (string or string array) |
| `priority` | no | Priority value. For scenarios: classifier order (higher runs first). For upsert_agent: delegate selection priority. |
| `provider` | no | Provider for set_default/upsert_scenario/upsert_agent |
| `remove_classification` | no | When remove_scenario, whether to remove matching classification rule (default true) |
| `subagents_auto_activate` | no | Enable/disable automatic sub-agent selection when agent is omitted or 'auto' |
| `subagents_enabled` | no | Enable/disable background sub-agent tools |
| `subagents_inflight_penalty` | no | Sub-agent score penalty per in-flight task |
| `subagents_load_window_secs` | no | Recent-event window for sub-agent load balancing (seconds) |
| `subagents_queue_poll_ms` | no | Poll interval while waiting for sub-agent capacity (milliseconds) |
| `subagents_queue_wait_ms` | no | How long to wait for sub-agent capacity before failing (milliseconds) |
| `subagents_recent_failure_penalty` | no | Sub-agent score penalty per recent failure in the load window |
| `subagents_recent_selection_penalty` | no | Sub-agent score penalty per recent assignment in the load window |
| `subagents_strategy` | no | Sub-agent auto-selection strategy |
| `system_prompt` | no | Optional system prompt override for delegate agent |
| `teams_auto_activate` | no | Enable/disable automatic team-agent selection when agent is omitted or 'auto' |
| `teams_enabled` | no | Enable/disable synchronous agent-team delegation tools |
| `teams_inflight_penalty` | no | Team score penalty per in-flight task |
| `teams_load_window_secs` | no | Recent-event window for team load balancing (seconds) |
| `teams_recent_failure_penalty` | no | Team score penalty per recent failure in the load window |
| `teams_recent_selection_penalty` | no | Team score penalty per recent assignment in the load window |
| `teams_strategy` | no | Team auto-selection strategy |
| `temperature` | no | Optional temperature override (0.0-2.0) |
| `transport` | no | Optional route transport override for upsert_scenario (auto, websocket, sse) |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1"
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
{}
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
