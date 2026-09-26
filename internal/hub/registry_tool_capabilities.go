package hub

// registryToolCapabilityExtras maps MCP tool names to Tool Registry capability
// tags beyond the shared edge_device / mcp_ws_hub / profile stamps.
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

func edgeCapabilitiesForTool(mcpToolName, profile string) []string {
	caps := []string{"edge_device", "mcp_ws_hub"}
	if profile != "" {
		caps = append(caps, profile)
	}
	if extras, ok := registryToolCapabilityExtras[mcpToolName]; ok {
		caps = append(caps, extras...)
	}
	return caps
}

func capabilityMeta(name, profile string) map[string]interface{} {
	caps := edgeCapabilitiesForTool(name, profile)
	return map[string]interface{}{
		"capabilities": caps,
		"datumbridge":  map[string]interface{}{"capabilities": caps},
	}
}
