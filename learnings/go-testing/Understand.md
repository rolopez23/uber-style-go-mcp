# Go Testing - Level 2: Understand

Explain ideas and concepts in your own words.

## Questions

1. Explain the difference between `t.Error()` and `t.Fatal()`.

2. Why do we use the Arrange-Act-Assert pattern in tests?

3. Describe what the `-v` flag does when running `go test`.

4. Why did we need to truncate time in `TestPing_ReturnsValidTimestamp`?

5. Explain why we can pass `nil` for the `*mcp.ServerSession` parameter in tests.

## Answers

<details>
<summary>Click to reveal answers</summary>

1. **Error vs Fatal** - `t.Error()` marks the test as failed but continues running. `t.Fatal()` marks as failed and stops immediately. Use `Fatal` when continuing would cause panics or meaningless results (e.g., nil pointer).

2. **Arrange-Act-Assert** - A pattern that structures tests clearly:
   - **Arrange**: Set up test data and preconditions
   - **Act**: Execute the code being tested
   - **Assert**: Verify the results
   This makes tests readable and maintainable.

3. **Verbose flag** - `-v` shows each test name as it runs and whether it passed/failed. Without it, only failures are shown. Useful for seeing progress and which tests ran.

4. **Time truncation** - `time.RFC3339` format only has second precision. If we capture `before` time with nanoseconds, the parsed timestamp (with only seconds) could appear to be "before" our before time. Truncating to seconds ensures fair comparison.

5. **Nil session** - The `Ping` handler doesn't use the session parameter for its logic. In unit tests, we can pass `nil` for unused dependencies. This is acceptable for simple tests but integration tests should use real sessions.

</details>
