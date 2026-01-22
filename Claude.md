# Claude Code Workflow

This is a **learning repository** for building an MCP server in Go that validates code against the [Uber Go Style Guide](https://github.com/uber-go/guide).

**Tech Stack:**
- [go-sdk](https://github.com/modelcontextprotocol/go-sdk) - MCP server framework
- Goal: Push Go code towards Uber style guide compliance

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

## Learning Exercises

See `learnings/Claude.md` for the learning system using Bloom's Taxonomy and scaffolded exercises.

**Schedule:** `learnings/exercises.ics` (1 exercise per day, 11 days)
