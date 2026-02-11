package tools

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/robertlopez/go-mcp/styleguide"
)

func loadTestStyleGuide(t *testing.T) *styleguide.StyleGuide {
	t.Helper()
	path := filepath.Join("..", "vendor", "uber-go-guide", "style.md")
	if _, err := os.Stat(path); err != nil {
		t.Skipf("style.md not found at %s — run git submodule update --init", path)
	}
	sg, err := styleguide.LoadStyleGuide(path)
	if err != nil {
		t.Fatalf("LoadStyleGuide failed: %v", err)
	}
	return sg
}

// --- TDD Cycle 1: Tool definition ---

func TestGetUberStylesTool_Name(t *testing.T) {
	if GetUberStylesTool.Name != "get_uber_styles" {
		t.Errorf("expected tool name 'get_uber_styles', got %q", GetUberStylesTool.Name)
	}
}

func TestGetUberStylesTool_Description(t *testing.T) {
	if GetUberStylesTool.Description == "" {
		t.Error("expected non-empty tool description")
	}
}

func TestGetUberStylesTool_ReadOnly(t *testing.T) {
	if GetUberStylesTool.Annotations == nil {
		t.Fatal("expected non-nil annotations")
	}
	if !GetUberStylesTool.Annotations.ReadOnlyHint {
		t.Error("expected ReadOnlyHint to be true")
	}
}

// --- TDD Cycle 2: Handler behavior ---

func TestGetUberStylesHandler_ReturnsMatchingSections(t *testing.T) {
	sg := loadTestStyleGuide(t)
	handler := NewGetUberStylesHandler(sg)

	params := &mcp.CallToolParamsFor[GetUberStylesInput]{
		Arguments: GetUberStylesInput{
			Code: `func doSomething() error {
	return fmt.Errorf("failed: %w", err)
}`,
		},
	}

	result, err := handler(context.Background(), nil, params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.StructuredContent.Count == 0 {
		t.Error("expected matching sections, got count 0")
	}
	if len(result.StructuredContent.Sections) == 0 {
		t.Error("expected non-empty sections slice")
	}
}

func TestGetUberStylesHandler_EmptyCode_ReturnsError(t *testing.T) {
	sg := loadTestStyleGuide(t)
	handler := NewGetUberStylesHandler(sg)

	params := &mcp.CallToolParamsFor[GetUberStylesInput]{
		Arguments: GetUberStylesInput{Code: ""},
	}

	result, err := handler(context.Background(), nil, params)
	if err != nil {
		t.Fatalf("expected no error (MCP error in result), got %v", err)
	}
	if !result.IsError {
		t.Error("expected IsError to be true for empty code")
	}
}

func TestGetUberStylesHandler_NoMatch_ReturnsEmptyList(t *testing.T) {
	sg := loadTestStyleGuide(t)
	handler := NewGetUberStylesHandler(sg)

	params := &mcp.CallToolParamsFor[GetUberStylesInput]{
		Arguments: GetUberStylesInput{Code: "x := 1 + 2"},
	}

	result, err := handler(context.Background(), nil, params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.IsError {
		t.Error("expected IsError to be false for no matches")
	}
	if result.StructuredContent.Count != 0 {
		t.Errorf("expected count 0, got %d", result.StructuredContent.Count)
	}
}

func TestGetUberStylesHandler_ResponseHasContent(t *testing.T) {
	sg := loadTestStyleGuide(t)
	handler := NewGetUberStylesHandler(sg)

	params := &mcp.CallToolParamsFor[GetUberStylesInput]{
		Arguments: GetUberStylesInput{
			Code: `defer f.Close()`,
		},
	}

	result, err := handler(context.Background(), nil, params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(result.Content) == 0 {
		t.Error("expected Content array to be non-empty")
	}
}
