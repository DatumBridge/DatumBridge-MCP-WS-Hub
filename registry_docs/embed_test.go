package registrydocs

import (
	"strings"
	"testing"
)

func TestMarkdown_hubInfo(t *testing.T) {
	text := Markdown("hub_info")
	if !strings.Contains(text, "hub_info") {
		t.Fatal("hub guide missing")
	}
	if Markdown("../hub_info") != "" {
		t.Fatal("path escape returned a guide")
	}
}
