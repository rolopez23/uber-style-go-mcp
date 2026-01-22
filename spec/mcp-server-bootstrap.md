# MCP Server Bootstrap Spec

## Why

We need a foundation to build an MCP server that validates Go code against the Uber style guide. Before adding validation logic, we need a working MCP server that can boot up and respond to basic requests.

This is a learning project - understanding the MCP protocol and go-sdk is as important as the end result.

## High-Level Goal

A minimal MCP server that:
- Boots up successfully
- Communicates over stdio transport
- Can be connected to from Claude Code or another MCP client
- Returns server info/capabilities

## Scope

- Initialize Go module
- Implement minimal MCP server using go-sdk
- Basic project structure
- Verify server boots and responds

## Non-Goals

- Style guide validation logic (future work)
- Multiple transport options (stdio only for now)
- Production-ready error handling
- Configuration management
