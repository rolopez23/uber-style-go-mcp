# Plan Files

Plan files define implementation strategy with orchestration and agent coordination.

## Top-Level Plan (`<feature>.md`)

The main plan focuses on orchestration and coordination:

- **Overview**: Summary of the implementation approach
- **Orchestration**: How work is divided and sequenced across agents
- **Agent Plan Links**: References to individual agent plans in `<feature>/`
- **Coordination Points**: Where agents hand off or depend on each other
- **Common Patterns**: Shared approaches agents should follow
- **Dependencies**: Prerequisites and blockers

## Agent Plans (`<feature>/<sub-feature>.md`)

Each agent plan is scoped for a single agent to implement:

- **Objective**: Clear, focused goal for this agent
- **Context**: What this agent needs to know from the top-level plan
- **Implementation Steps**: Concrete tasks to complete
- **Inputs/Outputs**: What this agent receives and produces
- **Success Criteria**: How to know when done

## Naming

Top-level plans are named `<feature>.md` and correspond to:
- `spec/<feature>.md` (specification)
- `learnings/<feature>.md` (learnings)

Agent plans live in `<feature>/<sub-feature>.md` subdirectories.
