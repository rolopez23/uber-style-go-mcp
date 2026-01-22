# MCP - Level 2: Understand

Explain ideas and concepts in your own words.

## Questions

1. Explain how stdio transport enables bidirectional communication.

2. Why does `mcp.NewServer()` need an `Implementation` struct with Name and Version?

3. Describe the flow: What happens when Claude Code calls the `ping` tool?

4. Why do we define both `PingTool` (the definition) and `Ping` (the handler) separately?

5. What is the purpose of the `StructuredContent` field in `CallToolResultFor`?

## Answers

<details>
<summary>Click to reveal answers</summary>

1. **Stdio transport** - The server reads JSON-RPC requests from stdin and writes responses to stdout. The MCP client (Claude Code) spawns the server as a subprocess and communicates through these streams. It's bidirectional because both sides can initiate messages.

2. **Server identity** - The Implementation struct identifies the server to clients. Name helps users identify which server they're connected to. Version enables compatibility checking and debugging. Clients may display this info or use it for feature negotiation.

3. **Tool call flow**:
   - Claude Code sends a JSON-RPC `tools/call` request with tool name "ping"
   - Server receives request via stdin
   - Server looks up registered handler for "ping"
   - Handler function `Ping()` executes
   - Handler returns `CallToolResultFor[PingOutput]`
   - Server serializes result and writes to stdout
   - Claude Code receives and displays result

4. **Separation of definition and implementation** - `PingTool` describes the tool (name, description, schema) for clients to discover. `Ping` is the actual code that runs. This allows the SDK to auto-generate schemas, validate inputs, and handle the tool registry separately from business logic.

5. **Typed responses** - `StructuredContent` holds the strongly-typed output value. The SDK serializes it to JSON for the client. Using generics (`CallToolResultFor[PingOutput]`) ensures type safety at compile time.

</details>
