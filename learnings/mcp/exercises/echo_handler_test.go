// Run: go test ./learnings/mcp/exercises/... -run TestEcho -v

package exercises

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestEcho_ReturnsMessage(t *testing.T) {
	ctx := context.Background()
	params := &mcp.CallToolParamsFor[EchoInput]{
		Arguments: EchoInput{Message: "hello world"},
	}

	result, err := Echo(ctx, nil, params)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result, got nil")
	}
	if result.StructuredContent.Echo != "hello world" {
		t.Errorf("expected echo 'hello world', got %q", result.StructuredContent.Echo)
	}
}

func TestEcho_EmptyMessage(t *testing.T) {
	ctx := context.Background()
	params := &mcp.CallToolParamsFor[EchoInput]{
		Arguments: EchoInput{Message: ""},
	}

	result, err := Echo(ctx, nil, params)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result, got nil")
	}
	if result.StructuredContent.Echo != "" {
		t.Errorf("expected empty echo, got %q", result.StructuredContent.Echo)
	}
}
