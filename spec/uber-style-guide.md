# Uber Style Guide Lookup Spec

## Why

Developers writing Go code need quick access to relevant Uber style guide recommendations. The style guide is a ~2,000-line monolithic markdown file with 50+ sections. Reading the whole thing for every code review is impractical. An MCP tool that surfaces only the relevant sections — given a piece of code — makes the guide actionable in the flow of work.

## High-Level Goal

A single MCP tool (`get_uber_styles`) that:

1. Accepts Go code as input
2. Internally scans the style guide summary/TOC for relevant sections (keyword/pattern matching)
3. Retrieves the full content of each matched section
4. Returns the matched sections as structured output

The LLM sees one tool. The two-step lookup (scan summary → pull details) is internal.

## Scope

### Git Submodule

- Add `uber-go/guide` as a git submodule at `vendor/uber-go-guide/`
- Pin to a specific commit on `master` for reproducibility
- Updatable via `git submodule update --remote`

### Style Guide Parser (`styleguide` package)

A new Go package at `styleguide/` that parses `style.md` into indexed sections at startup:

- **`LoadStyleGuide(path string) (*StyleGuide, error)`** — Reads and parses the file once
- **Section model**:
  - `ID` — Slug derived from heading anchor (e.g., `"error-wrapping"`)
  - `Title` — Heading text (e.g., `"Error Wrapping"`)
  - `Category` — Parent section (`"Guidelines"`, `"Performance"`, `"Style"`, `"Patterns"`, `"Linting"`)
  - `Summary` — First paragraph or ~200 chars of the section body
  - `Content` — Full markdown content of the section
- **`ListSections() []SectionSummary`** — Returns all sections (used internally for matching)
- **`GetSections(ids []string) []Section`** — Returns full content for given IDs
- **`MatchSections(code string) []Section`** — Scans code for keywords/patterns, matches against section titles and summaries, returns relevant sections with full content

### Matching Strategy (internal to `MatchSections`)

Keyword/pattern matching between the input code and section titles + summaries:

| Code pattern | Matched sections |
|---|---|
| `interface{}` or `any` | Pointers to Interfaces, Interface Compliance |
| `error`, `fmt.Errorf`, `errors.New` | Error Handling, Error Wrapping, Error Types |
| `go func`, `goroutine` | Goroutine Lifetimes |
| `sync.Mutex`, `sync.RWMutex` | Mutexes, Zero-value Mutexes |
| `map[` | Maps, Nil Maps |
| `chan ` | Channels |
| `defer` | Defer |
| `init()` | Init functions |
| `struct {` | Struct definitions, Field tags |
| `import (` | Import organization |
| Function signatures | Naming conventions, Functional Options |
| `_test.go` patterns | Table-driven tests |
| `make([]` | Slice operations, Prefer specifying capacity |

This table is illustrative, not exhaustive. The matcher expands as we learn which patterns matter most. The key principle: **over-match is better than under-match** — returning a few extra sections is fine since the LLM filters the output.

### MCP Tool: `get_uber_styles`

```go
type GetUberStylesInput struct {
    Code string `json:"code"`
}

type StyleSection struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Category string `json:"category"`
    Content  string `json:"content"`
}

type GetUberStylesOutput struct {
    Sections []StyleSection `json:"sections"`
    Count    int            `json:"count"`
    Version  string         `json:"version"` // submodule commit hash
}
```

- **Tool name**: `get_uber_styles`
- **Description**: "Analyze Go code against the Uber Go Style Guide. Returns relevant style guide sections with full explanations and code examples. Pass the Go code you want to validate."
- **ReadOnly**: `true`

### Registration in `main.go`

```go
sg, err := styleguide.LoadStyleGuide("vendor/uber-go-guide/style.md")
if err != nil {
    log.Fatal(err)
}

mcp.AddTool(server, tools.GetUberStylesTool, tools.NewGetUberStylesHandler(sg))
```

The handler constructor returns a closure capturing `*StyleGuide`.

### Error Handling

- Empty `code` input → MCP error with descriptive message
- Missing `style.md` (submodule not initialized) → startup fatal with message: "style.md not found — run `git submodule update --init`"
- No sections matched → return empty `sections` array with `count: 0` (not an error — the code may already be compliant)

## Non-Goals

- **AST parsing of the input code**: Start with keyword/text matching. AST analysis is a future enhancement.
- **Code modification or diff output**: The tool returns style guide content, not fixed code.
- **Full-text search across section bodies**: Matching is against titles, summaries, and code keywords — not a search engine.
- **Multiple guide versions**: Only the checked-out submodule commit is served.
- **Linting or scoring**: No pass/fail judgment. The tool surfaces relevant guidance; the LLM interprets it.
