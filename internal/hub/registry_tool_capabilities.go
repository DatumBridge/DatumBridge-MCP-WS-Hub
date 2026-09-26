package hub

import "strings"

// registryToolCapabilityExtras maps MCP tool names to that tool's own
// Tool Registry capability tags. A server-wide stamp is not added.
var registryToolCapabilityExtras = map[string][]string{
	"shell":                    {"shell", "device_control"},
	"process":                  {"shell", "device_control"},
	"bg_run":                   {"shell", "device_control"},
	"bg_status":                {"shell", "observability"},
	"file_read":                {"file_operations"},
	"file_write":               {"file_operations"},
	"file_edit":                {"file_operations"},
	"apply_patch":              {"file_operations", "git"},
	"glob_search":              {"file_operations", "search"},
	"content_search":           {"file_operations", "search"},
	"pdf_read":                 {"file_operations", "document"},
	"xlsx_read":                {"file_operations", "spreadsheet"},
	"docx_read":                {"file_operations", "document"},
	"pptx_read":                {"file_operations", "presentation"},
	"image_info":               {"file_operations"},
	"browser":                  {"browser_automation"},
	"browser_open":             {"browser_automation"},
	"screenshot":               {"browser_automation"},
	"memory_store":             {"memory"},
	"memory_recall":            {"memory"},
	"memory_observe":           {"memory"},
	"memory_forget":            {"memory"},
	"web_fetch":                {"network", "web_browse"},
	"http_request":             {"network"},
	"cron_add":                 {"scheduling"},
	"cron_list":                {"scheduling"},
	"cron_remove":              {"scheduling"},
	"cron_run":                 {"scheduling"},
	"cron_runs":                {"scheduling"},
	"cron_update":              {"scheduling"},
	"schedule":                 {"scheduling"},
	"git_operations":           {"git"},
	"task_plan":                {"planning"},
	"edge_agent_run":           {"agent"},
	"cursor_cli":               {"agent"},
	"pushover":                 {"communication", "notification"},
	"manage_auth_profile":      {"authentication"},
	"check_provider_quota":     {"observability"},
	"estimate_quota_cost":      {"observability"},
	"switch_provider":          {"ai"},
	"model_routing_config":     {"ai"},
	"proxy_config":             {"network"},
	"web_access_config":        {"network"},
	"web_search_config":        {"search"},
	"channel_ack_config":       {"communication"},
	"openclaw_migration":       {"migration"},
	"hub_info":                 {"observability"},
	"forward_jsonrpc_to_device": {"network"},
}

func edgeCapabilitiesForTool(mcpToolName, _ string) []string {
	if extras, ok := registryToolCapabilityExtras[mcpToolName]; ok && len(extras) > 0 {
		return append([]string{}, extras...)
	}
	if mcpToolName == "" {
		return nil
	}
	return []string{mcpToolName}
}

func annotateDeclaredCapabilities(schema map[string]interface{}, desc, name, profile string) (map[string]interface{}, string) {
	caps := edgeCapabilitiesForTool(name, profile)
	if schema == nil {
		schema = map[string]interface{}{"type": "object", "properties": map[string]interface{}{}}
	}
	schema["x-datumbridge-capabilities"] = caps
	if len(caps) > 0 && !strings.Contains(strings.ToLower(desc), "capabilities:") {
		desc = strings.TrimSpace(desc) + "\nCapabilities: " + strings.Join(caps, ", ")
	}
	return schema, desc
}

func capabilityMeta(name, profile string) map[string]interface{} {
	caps := edgeCapabilitiesForTool(name, profile)
	return map[string]interface{}{
		"capabilities": caps,
		"datumbridge":  map[string]interface{}{"capabilities": caps},
	}
}
