package tools

import (
	"context"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type PingInput struct{}

type PingOutput struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

var PingTool = &mcp.Tool{
	Name:        "ping",
	Description: "Check if the go-style-guide server is running",
}

func Ping(ctx context.Context, sess *mcp.ServerSession, params *mcp.CallToolParamsFor[PingInput]) (*mcp.CallToolResultFor[PingOutput], error) {
	output := PingOutput{
		Message:   "pong",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	return &mcp.CallToolResultFor[PingOutput]{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "pong - " + output.Timestamp},
		},
		StructuredContent: output,
	}, nil
}
