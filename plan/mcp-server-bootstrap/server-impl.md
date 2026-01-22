# Server Implementation Agent Plan

**Beads Issue:** (link to issue)

## Objective

Implement a minimal MCP server that boots and listens over stdio.

## Context

Depends on Project Setup completing. Uses go-sdk patterns from official examples.

## Implementation Steps

1. Create `main.go` in project root
2. Import `github.com/modelcontextprotocol/go-sdk/mcp`
3. Create server with Implementation info (name: "go-style-guide", version: "0.1.0")
4. Run server with StdioTransport

## Inputs/Outputs

**Inputs:**
- `go.mod` with go-sdk dependency

**Outputs:**
- `main.go` - Minimal MCP server

## Success Criteria

- `go build` compiles without errors
- Server binary can be executed
- Server waits for MCP client connection (doesn't crash immediately)
