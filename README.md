# Go Style Guide MCP Server

An MCP server that validates Go code against the [Uber Go Style Guide](https://github.com/uber-go/guide).

## Prerequisites

- Go 1.21+

## Build

```bash
go mod tidy
go build -o go-style-guide
```

## Run

The server communicates over stdio and is designed to be launched by an MCP client.

### With Claude Code

Add to your Claude Code MCP configuration (`~/.claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "go-style-guide": {
      "command": "/path/to/go-mcp/go-style-guide"
    }
  }
}
```

Replace `/path/to/go-mcp` with the actual path to this repository.

### Manual Testing

```bash
./go-style-guide
```

The server will start and wait for MCP protocol messages on stdin. Press `Ctrl+C` to stop.

## Development

See the `Claude.md` files for the development workflow:

- `spec/` - Feature specifications
- `plan/` - Implementation plans
- `learnings/` - Captured insights
