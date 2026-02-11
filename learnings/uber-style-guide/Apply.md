# Uber Style Guide Tool - Level 3: Apply

Use knowledge to solve problems. Complete the exercises in `exercises/`.

## Exercises

### Exercise 12: Parse Markdown Headings

**File:** `exercises/12-parse-headings.go`

Extract all `##` and `###` headings from a markdown string, returning their level and text.

```bash
go test ./learnings/uber-style-guide/exercises/... -run TestParseHeadings -v
```

### Exercise 13: Generate Slugs

**File:** `exercises/13-slugify.go`

Convert heading text to URL-safe slugs (lowercase, spaces to hyphens, strip special chars).

```bash
go test ./learnings/uber-style-guide/exercises/... -run TestSlugify -v
```

### Exercise 14: Keyword Matcher

**File:** `exercises/14-keyword-matcher.go`

Given keyword-to-label rules and input text, return matching labels (deduplicated).

```bash
go test ./learnings/uber-style-guide/exercises/... -run TestKeywordMatch -v
```

### Exercise 15: Closure Handler

**File:** `exercises/15-closure-handler.go`

Create a function that returns a handler closure capturing a dependency.

```bash
go test ./learnings/uber-style-guide/exercises/... -run TestClosureHandler -v
```

## Run All Exercises

```bash
go test ./learnings/uber-style-guide/exercises/... -v
```
