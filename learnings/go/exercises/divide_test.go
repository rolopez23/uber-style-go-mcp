// Run: go test ./learnings/go/exercises/... -run TestDivide -v

package exercises

import (
	"errors"
	"testing"
)

func TestDivide_Basic(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d, want 5", result)
	}
}

func TestDivide_ByZero(t *testing.T) {
	result, err := Divide(10, 0)
	if err == nil {
		t.Fatal("expected error for division by zero")
	}
	if !errors.Is(err, ErrDivideByZero) {
		t.Errorf("expected ErrDivideByZero, got %v", err)
	}
	if result != 0 {
		t.Errorf("result should be 0 on error, got %d", result)
	}
}

func TestDivide_Negative(t *testing.T) {
	result, err := Divide(-10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != -5 {
		t.Errorf("Divide(-10, 2) = %d, want -5", result)
	}
}
