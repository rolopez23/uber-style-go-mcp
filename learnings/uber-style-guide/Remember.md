# Uber Style Guide Tool - Level 1: Remember

Recall facts and basic concepts. Try to answer from memory before checking.

## Questions

### 1. What function parses the style guide at startup?

<details>
<summary>Answer</summary>

`LoadStyleGuide(path string) (*StyleGuide, error)` — reads and parses the markdown file once into indexed sections.
</details>

### 2. What Go function checks if a string contains a substring?

<details>
<summary>Answer</summary>

`strings.Contains(s, substr string) bool` — returns true if substr is within s.
</details>

### 3. What is a closure in Go?

<details>
<summary>Answer</summary>

A function that captures variables from its enclosing scope. The captured variables remain accessible even after the enclosing function returns.
</details>

### 4. What field must every MCP tool response include?

<details>
<summary>Answer</summary>

The `Content` array — a list of `mcp.Content` objects (e.g., `&mcp.TextContent{Text: "..."}`) that represent the unstructured result of the tool call.
</details>

### 5. How do you add a git submodule?

<details>
<summary>Answer</summary>

`git submodule add <url> <path>` — clones the repository at the given path and records it in `.gitmodules`.
</details>
