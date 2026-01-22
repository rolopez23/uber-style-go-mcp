# Claude Code Workflow

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
