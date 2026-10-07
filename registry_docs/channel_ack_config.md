# channel_ack_config

Inspect and update configurable ACK emoji reaction policies for Telegram/Discord/Lark/Feishu under [channels_config.ack_reaction]. Supports enabling/disabling reactions, setting emoji pools, and rule-based conditions.

## Parameters

| Name | Required | Meaning |
|---|---|---|
| `device_id` | yes | Hub-registered device with an active WebSocket |
| `action` | yes | Operation to perform |
| `channel` | no | string |
| `chat_id` | no | ['string', 'null'] |
| `chat_type` | no | string |
| `defaults` | no | string |
| `emojis` | no | string |
| `enabled` | no | boolean |
| `index` | no | integer |
| `locale_hint` | no | ['string', 'null'] |
| `rule` | no | object |
| `rules` | no | ['array', 'null'] |
| `runs` | no | integer |
| `sample_rate` | no | ['number', 'null'] |
| `sender_id` | no | ['string', 'null'] |
| `strategy` | no | ['string', 'null'] |
| `text` | no | string |

## Cases

### Typical call

Input:

```json
{
  "device_id": "edge-device-1",
  "action": "example",
  "text": "hello"
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
  "text": "hello"
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
