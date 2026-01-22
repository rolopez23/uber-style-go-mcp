# Learnings Files

This is a **learning repository** for Go and MCP. Learnings use Bloom's Taxonomy to structure knowledge and practice.

**Learning Focus:**
- Go programming language
- MCP (Model Context Protocol)
- Go testing patterns

## Directory Structure

```
learnings/
├── Claude.md                    # This file
├── mcp-server-bootstrap.md      # Feature-specific learnings
├── go/                          # Go language learnings
│   ├── Remember.md              # Level 1: Recall facts
│   ├── Understand.md            # Level 2: Explain concepts
│   ├── Apply.md                 # Level 3: Exercises
│   └── exercises/               # Scaffolded code + failing tests
├── mcp/                         # MCP protocol learnings
│   ├── Remember.md
│   ├── Understand.md
│   ├── Apply.md
│   └── exercises/
└── go-testing/                  # Go testing learnings
    ├── Remember.md
    ├── Understand.md
    ├── Apply.md
    └── exercises/
```

## Bloom's Taxonomy Levels

| Level | Name | Description | File |
|-------|------|-------------|------|
| 1 | Remember | Recall facts and basic concepts | `Remember.md` |
| 2 | Understand | Explain ideas or concepts | `Understand.md` |
| 3 | Apply | Use information in new situations | `Apply.md` |

## Schedule

Import `exercises.ics` into your calendar for a daily exercise schedule (11 days, including weekends).

| Days 1-4 | Go fundamentals |
|----------|-----------------|
| Days 5-7 | MCP protocol |
| Days 8-11 | Go testing |

## How to Use

### 1. Remember (Level 1)
Read the questions and try to answer from memory. Check answers when done.

### 2. Understand (Level 2)
Read questions and write your explanation. Compare with provided answers.

### 3. Apply (Level 3)
Complete the coding exercises:

```bash
# Run all exercises for a topic
go test ./learnings/go/exercises/... -v
go test ./learnings/mcp/exercises/... -v
go test ./learnings/go-testing/exercises/... -v
```

## Exercise Scaffolding

Each exercise follows a test-driven pattern:

**Structure:**
- `.go` file - Scaffolded code with TODOs and hints
- `_test.go` file - Failing tests that define expected behavior

**Workflow:**
1. Read the TODO comments in the `.go` file
2. Run the test to see it fail
3. Implement the code
4. Run the test until it passes

**Run comments:** Each test file has a `// Run:` comment at the top showing the exact command

## Adding New Topics

Create a new directory with:
1. `Remember.md` - 3-5 factual recall questions
2. `Understand.md` - 3-5 conceptual questions
3. `Apply.md` - Exercise descriptions
4. `exercises/` - Scaffolded code + tests
