# Uber Style Guide Tool - Level 2: Understand

Explain ideas or concepts. Write your explanation, then check the provided answer.

## Questions

### 1. Why parse the style guide once at startup instead of on every tool call?

<details>
<summary>Answer</summary>

Parsing a ~2,000-line markdown file involves file I/O, string splitting, regex processing, and building an index. Doing this once at startup amortizes the cost across all requests. Each tool call then only needs to run lightweight string matching against the pre-built index, keeping response times fast and predictable.
</details>

### 2. Why use a closure to construct the handler instead of a global variable?

<details>
<summary>Answer</summary>

A closure (`NewGetUberStylesHandler(sg)`) injects the `*StyleGuide` dependency explicitly, making the handler testable — tests can pass different StyleGuide instances. A global variable creates hidden coupling: the handler depends on something being initialized elsewhere, making it harder to test in isolation and easier to introduce init-order bugs.
</details>

### 3. Why is over-matching better than under-matching for the keyword matcher?

<details>
<summary>Answer</summary>

The tool's output is consumed by an LLM, which can filter irrelevant sections. Returning a few extra sections is low-cost — the LLM ignores them. But missing a relevant section means the user gets incomplete guidance with no way to recover. Over-matching trades a small cost (extra context) for a large benefit (no missed recommendations).
</details>

### 4. How does slug generation from headings create stable section IDs?

<details>
<summary>Answer</summary>

Slugs are deterministic transformations of heading text: lowercase, spaces to hyphens, strip special characters. As long as the heading text doesn't change, the slug stays the same. This means section IDs are stable across parses and can be reliably used in match rules, even if the guide content around the heading changes.
</details>

### 5. Why return `StructuredContent` alongside the `Content` array?

<details>
<summary>Answer</summary>

`Content` (text) is the universal fallback — all MCP clients can display it. `StructuredContent` (typed JSON) enables programmatic consumers to process the result without parsing text. Returning both ensures the tool works for human-readable display and machine consumption.
</details>
