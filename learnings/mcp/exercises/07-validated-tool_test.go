// Run: go test ./learnings/mcp/exercises/... -run TestValidated -v
// Run: go test ./learnings/mcp/exercises/... -run TestGreetTool -v

package exercises

import (
	"context"
	"errors"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestGreetTool_Definition(t *testing.T) {
	if GreetTool == nil {
		t.Fatal("GreetTool must not be nil")
	}
	if GreetTool.Name != "greet" {
		t.Errorf("expected tool name 'greet', got %q", GreetTool.Name)
	}
}

func TestValidatedGreet_Success(t *testing.T) {
	ctx := context.Background()
	params := &mcp.CallToolParamsFor[GreetInput]{
		Arguments: GreetInput{Name: "Alice"},
	}

	result, err := GreetHandler(ctx, nil, params)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result, got nil")
	}
	if result.StructuredContent.Greeting != "Hello, Alice!" {
		t.Errorf("expected 'Hello, Alice!', got %q", result.StructuredContent.Greeting)
	}
}

func TestValidatedGreet_EmptyName(t *testing.T) {
	ctx := context.Background()
	params := &mcp.CallToolParamsFor[GreetInput]{
		Arguments: GreetInput{Name: ""},
	}

	result, err := GreetHandler(ctx, nil, params)

	if err == nil {
		t.Fatal("expected error for empty name")
	}
	if !errors.Is(err, ErrEmptyName) {
		t.Errorf("expected ErrEmptyName, got %v", err)
	}
	if result != nil {
		t.Error("result should be nil on error")
	}
}
