# Go Testing - Level 1: Remember

Recall facts and basic concepts about testing in Go.

## Questions

1. What is the naming convention for Go test files?
2. What package provides testing utilities in Go?
3. What command runs all tests in a module?
4. What prefix must test function names have?
5. What parameter type do test functions receive?

## Answers

<details>
<summary>Click to reveal answers</summary>

1. `*_test.go` (e.g., `ping_test.go`)
2. `testing` (from the standard library)
3. `go test ./...`
4. `Test` (e.g., `TestPing`, `TestDivide`)
5. `*testing.T`

</details>
