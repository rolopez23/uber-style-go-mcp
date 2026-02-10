# Claude Code Workflow

## Project Context

This is a **Go + MCP server learning repository**. Primary languages: Go and Markdown. Follow Go conventions (`gofmt`, idiomatic error handling). When creating MCP tools, always validate that response structs include required fields (like the `Content` array) before testing.

**Tech Stack:**
- [go-sdk](https://github.com/modelcontextprotocol/go-sdk) - MCP server framework
- Goal: Push Go code towards [Uber Go Style Guide](https://github.com/uber-go/guide) compliance

This document describes the structured approach for working with Claude on this project.

## Workflow Overview

1. **Spec** → Define the "why" and high-level goals
2. **Plan** → Create implementation strategy with agent orchestration
3. **Learnings** → Capture insights and patterns discovered

## Directory Structure

```
spec/           # Specification files (see spec/Claude.md)
plan/           # Implementation plans (see plan/Claude.md)
learnings/      # Captured insights (see learnings/Claude.md)
```

## File Naming Convention

Related files across directories share the same base name:

```
spec/<feature>.md                  # Specification
plan/<feature>.md                  # Top-level plan
plan/<feature>/<sub-feature>.md    # Agent level plan
learnings/<feature>.md             # Learnings and insights
```

See the `Claude.md` file in each directory for detailed instructions.

## Issue Tracking

Use [Beads](https://beads.dev) for issue tracking.

- Top-level plans (`plan/<feature>.md`) link to a Beads **epic**
- Agent plans (`plan/<feature>/<sub-feature>.md`) link to a Beads **issue**

## Levels of Done

| Level | Name | Requirements |
|-------|------|--------------|
| Done | Agent verified | Agent has tested via automated test or running code |
| Done Done | Human verified | Agent tested + human has verified |
| Done Done Done | Stress tested | Agent tested + human verified + human stress tested |

- **Done**: The agent (Claude) has validated the work through automated tests or manual execution
- **Done Done**: The human has reviewed and confirmed the work functions correctly
- **Done Done Done**: The human has stress tested edge cases, error conditions, and real-world usage

## Testing

When running Go tests, always scope to the specific package or exercise directory (e.g., `go test -v ./exercises/01/...`) rather than running `go test ./...` across the entire repo. Only run broader tests if explicitly asked.

## Commit Policy

Claude is empowered to commit when code reaches **Done** status.

| Commits | Required Check |
|---------|----------------|
| Every commit | Done (agent verified) |
| Every 5 commits | Done Done (human verified) |
| Every 25 commits | Done Done Done (human stress tested) |

Claude should proactively commit when tests pass or code runs successfully.

## Git Workflow

Use conventional commit messages (e.g., `feat:`, `fix:`, `docs:`, `refactor:`) and keep commits atomic — one logical change per commit. Confirm commit scope with the user if ambiguous.

## Learning Exercises

See `learnings/Claude.md` for the learning system using Bloom's Taxonomy and scaffolded exercises.

**Schedule:** `learnings/exercises.ics` (1 exercise per day, 11 days)
