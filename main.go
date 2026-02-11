package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/robertlopez/go-mcp/styleguide"
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

	// Load Uber Go Style Guide
	sg, err := styleguide.LoadStyleGuide("vendor/uber-go-guide/style.md")
	if err != nil {
		log.Fatal(err)
	}

	// Register tools
	mcp.AddTool(server, tools.PingTool, tools.Ping)
	mcp.AddTool(server, tools.GetUberStylesTool, tools.NewGetUberStylesHandler(sg))

	// Run server over stdio transport
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}
