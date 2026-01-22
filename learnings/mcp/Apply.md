# MCP - Level 3: Apply

Use knowledge to solve problems. Complete the exercises in `exercises/`.

## Exercises

### Exercise 1: Define a Tool

**File:** `exercises/tool_definition.go`
**Test:** `exercises/tool_definition_test.go`

Create a tool definition for an "echo" tool.

```bash
go test ./learnings/mcp/exercises/... -run TestEchoTool -v
```

### Exercise 2: Implement a Tool Handler

**File:** `exercises/echo_handler.go`
**Test:** `exercises/echo_handler_test.go`

Implement the Echo tool handler that returns the input message.

```bash
go test ./learnings/mcp/exercises/... -run TestEcho -v
```

### Exercise 3: Tool with Validation

**File:** `exercises/validated_tool.go`
**Test:** `exercises/validated_tool_test.go`

Create a tool that validates its input and returns an error for invalid data.

```bash
go test ./learnings/mcp/exercises/... -run TestValidated -v
```

## Run All Exercises

```bash
go test ./learnings/mcp/exercises/... -v
```
