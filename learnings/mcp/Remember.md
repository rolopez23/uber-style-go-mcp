# MCP - Level 1: Remember

Recall facts and basic concepts about MCP (Model Context Protocol).

## Questions

1. What does MCP stand for?
2. Name the two parameters passed to `mcp.NewServer()`.
3. What transport does our server use for communication?
4. What function registers a tool with the server?
5. What two fields are required in the `Implementation` struct?

## Answers

<details>
<summary>Click to reveal answers</summary>

1. Model Context Protocol
2. `*Implementation` (server info) and `*ServerOptions` (configuration, can be nil)
3. Stdio (standard input/output)
4. `mcp.AddTool()`
5. `Name` and `Version`

</details>
