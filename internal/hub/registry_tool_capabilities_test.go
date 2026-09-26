package hub

import (
	"testing"
)

func TestEdgeCapabilitiesForTool_siblingsDiffer(t *testing.T) {
	shell := edgeCapabilitiesForTool("shell", "dtbclaw-default")
	file := edgeCapabilitiesForTool("file_read", "dtbclaw-default")
	if len(shell) == 0 || len(file) == 0 {
		t.Fatalf("expected extras: shell=%v file=%v", shell, file)
	}
	for _, c := range shell {
		if c == "edge_device" || c == "mcp_ws_hub" || c == "dtbclaw-default" {
			t.Fatalf("shared server stamp on shell: %#v", shell)
		}
	}
	same := true
	if len(shell) != len(file) {
		same = false
	} else {
		for i := range shell {
			if shell[i] != file[i] {
				same = false
				break
			}
		}
	}
	if same {
		t.Fatalf("siblings identical: %#v", shell)
	}
	foundShell := false
	for _, c := range shell {
		if c == "shell" {
			foundShell = true
		}
	}
	if !foundShell {
		t.Fatalf("missing shell cap: %#v", shell)
	}
}
