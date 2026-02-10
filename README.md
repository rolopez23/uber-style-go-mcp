# Go Style Guide MCP Server

An MCP server that validates Go code against the [Uber Go Style Guide](https://github.com/uber-go/guide). Built with the [MCP Go SDK](https://github.com/modelcontextprotocol/go-sdk).

## Prerequisites

- Go 1.23+
- Node.js (for MCP Inspector)

## Build

```bash
go mod tidy
go build -o go-style-guide
```

## Testing

Run all tests:

```bash
go test ./...
```

Run tests for a specific package:

```bash
go test ./tools/
```

Run a single test by name:

```bash
go test ./tools/ -run TestPing
```

Run tests with verbose output:

```bash
go test -v ./tools/
```

Run benchmarks:

```bash
go test -bench=. ./...
```

## Run

The server communicates over stdio and is designed to be launched by an MCP client.

### MCP Inspector

The [MCP Inspector](https://github.com/modelcontextprotocol/inspector) provides a web UI for testing your server interactively.

```bash
npx @modelcontextprotocol/inspector ./go-style-guide
```

This opens a browser at `http://localhost:6274` where you can:
- Connect to the server
- List available tools
- Call tools with custom arguments
- Inspect request/response payloads

### Connecting to Claude Desktop

Add to your Claude Desktop config file:

- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "go-style-guide": {
      "command": "/absolute/path/to/go-mcp/go-style-guide"
    }
  }
}
```

Replace `/absolute/path/to/go-mcp` with the actual path to the built binary. Restart Claude Desktop after updating the config.

### Connecting to Claude Code

Add the server to your Claude Code MCP settings (`~/.claude.json`):

```json
{
  "mcpServers": {
    "go-style-guide": {
      "command": "/absolute/path/to/go-mcp/go-style-guide"
    }
  }
}
```

Or add it via the CLI:

```bash
claude mcp add go-style-guide /absolute/path/to/go-mcp/go-style-guide
```

### Manual Testing

```bash
./go-style-guide
```

The server starts and waits for MCP protocol messages on stdin. Press `Ctrl+C` to stop.

## Development

See the `Claude.md` files for the development workflow:

- `spec/` - Feature specifications
- `plan/` - Implementation plans
- `learnings/` - Captured insights
