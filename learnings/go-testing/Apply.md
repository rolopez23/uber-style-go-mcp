# Go Testing - Level 3: Apply

Use knowledge to solve problems. Complete the exercises in `exercises/`.

## Exercises

### Exercise 1: Write a Basic Test

**File:** `exercises/basic_test.go`

Write a test for an `Add` function.

```bash
go test ./learnings/go-testing/exercises/... -run TestAdd -v
```

### Exercise 2: Table-Driven Tests

**File:** `exercises/table_test.go`

Write table-driven tests for a `Multiply` function.

```bash
go test ./learnings/go-testing/exercises/... -run TestMultiply -v
```

### Exercise 3: Testing Errors

**File:** `exercises/error_test.go`

Write tests that verify error conditions.

```bash
go test ./learnings/go-testing/exercises/... -run TestSqrt -v
```

### Exercise 4: Benchmark Test

**File:** `exercises/benchmark_test.go`

Write a benchmark for a `Fibonacci` function.

```bash
go test ./learnings/go-testing/exercises/... -bench=. -v
```

## Run All Exercises

```bash
go test ./learnings/go-testing/exercises/... -v
```
