// Run: go test ./learnings/mcp/exercises/... -run TestEchoTool -v

package exercises

import "testing"

func TestEchoTool_Name(t *testing.T) {
	if EchoTool == nil {
		t.Fatal("EchoTool must not be nil")
	}
	if EchoTool.Name != "echo" {
		t.Errorf("expected tool name 'echo', got %q", EchoTool.Name)
	}
}

func TestEchoTool_Description(t *testing.T) {
	if EchoTool == nil {
		t.Fatal("EchoTool must not be nil")
	}
	if EchoTool.Description == "" {
		t.Error("EchoTool must have a description")
	}
}
