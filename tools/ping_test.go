package tools

import (
	"context"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestPing(t *testing.T) {
	// Arrange
	ctx := context.Background()
	params := &mcp.CallToolParamsFor[PingInput]{
		Arguments: PingInput{},
	}

	// Act
	result, err := Ping(ctx, nil, params)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	if result.StructuredContent.Message != "pong" {
		t.Errorf("expected message 'pong', got %q", result.StructuredContent.Message)
	}
}

func TestPing_ReturnsValidTimestamp(t *testing.T) {
	// Arrange
	ctx := context.Background()
	params := &mcp.CallToolParamsFor[PingInput]{
		Arguments: PingInput{},
	}
	// Truncate to seconds since RFC3339 doesn't include sub-second precision
	before := time.Now().UTC().Truncate(time.Second)

	// Act
	result, err := Ping(ctx, nil, params)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	timestamp, err := time.Parse(time.RFC3339, result.StructuredContent.Timestamp)
	if err != nil {
		t.Fatalf("timestamp not in RFC3339 format: %v", err)
	}

	after := time.Now().UTC().Add(time.Second) // Add buffer for test execution
	if timestamp.Before(before) || timestamp.After(after) {
		t.Errorf("timestamp %v not between %v and %v", timestamp, before, after)
	}
}

func TestPingTool_Definition(t *testing.T) {
	if PingTool.Name != "ping" {
		t.Errorf("expected tool name 'ping', got %q", PingTool.Name)
	}

	if PingTool.Description == "" {
		t.Error("expected tool description, got empty string")
	}
}
