// Run: go test ./learnings/go-testing/exercises/... -run TestSqrt -v

package exercises

import (
	"errors"
	"testing"
)

// TODO: Implement TestSqrt_Positive
// Test that Sqrt(16) returns 4 with no error
func TestSqrt_Positive(t *testing.T) {
	// TODO: Implement this test
	// Hint: result, err := Sqrt(16)
	// Check err is nil, then check result == 4
	t.Skip("TODO: Implement this test")
}

// TODO: Implement TestSqrt_Zero
// Test that Sqrt(0) returns 0 with no error
func TestSqrt_Zero(t *testing.T) {
	// TODO: Implement this test
	t.Skip("TODO: Implement this test")
}

// TODO: Implement TestSqrt_Negative
// Test that Sqrt(-1) returns ErrNegativeInput
// Use errors.Is() to check the error type
func TestSqrt_Negative(t *testing.T) {
	// TODO: Implement this test
	// Hint: _, err := Sqrt(-1)
	// Use: if !errors.Is(err, ErrNegativeInput) { ... }
	_ = errors.Is // silence unused import
	t.Skip("TODO: Implement this test")
}
