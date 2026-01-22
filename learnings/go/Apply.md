# Go - Level 3: Apply

Use knowledge to solve problems. Complete the exercises in `exercises/`.

## Exercises

### Exercise 1: Define a Struct with JSON Tags

**File:** `exercises/echo_types.go`
**Test:** `exercises/echo_types_test.go`

Create an `EchoInput` struct with a `Message` field and proper JSON tag.

```bash
go test ./learnings/go/exercises/... -run TestEchoInput -v
```

### Exercise 2: Implement a Simple Function

**File:** `exercises/greet.go`
**Test:** `exercises/greet_test.go`

Implement a `Greet` function that takes a name and returns a greeting.

```bash
go test ./learnings/go/exercises/... -run TestGreet -v
```

### Exercise 3: Work with Pointers

**File:** `exercises/counter.go`
**Test:** `exercises/counter_test.go`

Implement a `Counter` struct with `Increment` and `Value` methods using pointer receivers.

```bash
go test ./learnings/go/exercises/... -run TestCounter -v
```

### Exercise 4: Error Handling

**File:** `exercises/divide.go`
**Test:** `exercises/divide_test.go`

Implement a `Divide` function that returns an error for division by zero.

```bash
go test ./learnings/go/exercises/... -run TestDivide -v
```

## Run All Exercises

```bash
go test ./learnings/go/exercises/... -v
```
