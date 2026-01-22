package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/robertlopez/go-mcp/tools"
)

func main() {
	// Create MCP server with implementation info
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "go-style-guide",
			Version: "0.1.0",
		},
		nil,
	)

	// Register tools
	mcp.AddTool(server, tools.PingTool, tools.Ping)

	// Run server over stdio transport
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
