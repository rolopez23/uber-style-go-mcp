# MCP Server Bootstrap Learnings

## Implementation Insights

### What Worked

- Using go-sdk's `mcp.NewServer()` with `Implementation` struct for server identity
- Separating tools into a `tools/` package for organization
- Generic `ToolHandlerFor[In, Out]` pattern keeps handlers type-safe
- `CallToolParamsFor[T]` and `CallToolResultFor[T]` for typed request/response

### What Didn't

- Initial attempts used wrong function signature for tool handlers
- `mcp.CallToolRequest` doesn't exist - the correct type is `*mcp.ServerSession` + `*mcp.CallToolParamsFor[T]`
- `mcp.TextContent{}` needs to be a pointer to satisfy the `Content` interface

### Insights

- MCP tool handlers have a specific signature: `func(context.Context, *ServerSession, *CallToolParamsFor[In]) (*CallToolResultFor[Out], error)`
- The go-sdk uses generics heavily - understanding Go generics is essential
- Tools are registered with `mcp.AddTool(server, toolDef, handler)` - not as server methods

---

## Human Understanding Section

### Code Walkthrough

| File | Purpose |
|------|---------|
| `main.go` | Server initialization, tool registration, transport setup |
| `tools/ping.go` | Example tool with input/output types and handler |
| `tools/ping_test.go` | Unit tests demonstrating Go testing patterns |
| `go.mod` | Module definition and dependencies |

### Architecture Decisions

- **Stdio transport**: Simplest option for local MCP server, matches how Claude Code connects
- **go-sdk**: Official MCP SDK for Go, maintained by protocol authors
- **tools/ package**: Separates tool implementations from server setup

### Reading Order

1. `go.mod` - Understand dependencies
2. `tools/ping.go` - See how a tool is structured (types + handler)
3. `main.go` - See how server is created and tools are registered

### Concepts to Study

- [MCP Protocol](https://modelcontextprotocol.io/) - Protocol basics
- [go-sdk](https://github.com/modelcontextprotocol/go-sdk) - SDK documentation
- Go generics (type parameters)
- Go interfaces and type assertions
- Go testing (`go test`, `testing` package)

---

## Practice Questions - Go

### Level 1 - Remember

1. What keyword is used to define a package in Go?
2. What is the file called that defines Go module dependencies?
3. What does `context.Context` represent?
4. What symbol is used to export a function or type in Go?
5. What command builds a Go binary?

### Level 2 - Understand

1. Explain why we use a separate `tools/` package instead of putting everything in `main`.
2. Describe what `CallToolParamsFor[PingInput]` means - what is the `[PingInput]` part?
3. Why does the `Ping` function return two values `(*CallToolResultFor[PingOutput], error)`?
4. Explain the difference between `PingInput struct{}` and `PingInput struct{ Name string }`.
5. Why do we use pointers (`*mcp.ServerSession`) instead of values in function parameters?

### Level 3 - Apply

1. Write a new struct `EchoInput` with a single field `Message string` including a JSON tag.
2. Modify `PingOutput` to include a new field `ServerName string` and update the handler.
3. Create a new file `tools/echo.go` with an echo tool that returns the input message.
4. Add error handling to the Ping handler that returns an error if the context is cancelled.

---

## Practice Questions - MCP

### Level 1 - Remember

1. What does MCP stand for?
2. Name the two parameters passed to `mcp.NewServer()`.
3. What transport does our server use for communication?
4. What function registers a tool with the server?
5. What two fields are required in the `Implementation` struct?

### Level 2 - Understand

1. Explain how stdio transport enables bidirectional communication.
2. Why does `mcp.NewServer()` need an `Implementation` struct with Name and Version?
3. Describe the flow: What happens when Claude Code calls the `ping` tool?
4. Why do we define both `PingTool` (the definition) and `Ping` (the handler) separately?
5. What is the purpose of the `StructuredContent` field in `CallToolResultFor`?

### Level 3 - Apply

1. Register a second tool in `main.go` (after creating echo.go).
2. Change the server name from "go-style-guide" to "uber-style-validator" and rebuild.
3. Add a description field to `PingTool` that explains what arguments it accepts.
4. Implement a tool that accepts a `Filename string` input and returns whether the file exists.
5. Create a tool that lists all files in a given directory path.

---

## Practice Questions - Go Testing

### Level 1 - Remember

1. What is the naming convention for Go test files?
2. What package provides testing utilities in Go?
3. What command runs all tests in a module?
4. What prefix must test function names have?
5. What parameter type do test functions receive?

### Level 2 - Understand

1. Explain the difference between `t.Error()` and `t.Fatal()`.
2. Why do we use the Arrange-Act-Assert pattern in tests?
3. Describe what the `-v` flag does when running `go test`.
4. Why did we need to truncate time in `TestPing_ReturnsValidTimestamp`?
5. Explain why we can pass `nil` for the `*mcp.ServerSession` parameter in tests.

### Level 3 - Apply

1. Write a test that verifies `PingOutput.Message` is never empty.
2. Add a table-driven test for the Ping function with multiple scenarios.
3. Create a test that checks behavior when context is cancelled.
4. Write a benchmark test for the Ping function using `testing.B`.
5. Add test coverage reporting and identify untested code paths.
