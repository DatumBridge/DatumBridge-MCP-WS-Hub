package hub

import "github.com/datumbridge/mcp-ws-hub/registry_docs"

func attachRegistryDocs(schema map[string]interface{}, name string) {
	if schema == nil {
		return
	}
	if doc := registrydocs.Markdown(name); doc != "" {
		schema["x-datumbridge-docs"] = doc
	}
}
