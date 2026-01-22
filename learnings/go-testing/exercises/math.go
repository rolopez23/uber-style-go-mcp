package exercises

import (
	"errors"
	"math"
)

// ErrNegativeInput is returned when input is negative
var ErrNegativeInput = errors.New("input must be non-negative")

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Sqrt returns the square root of a non-negative number
// Returns ErrNegativeInput if n is negative
func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	return math.Sqrt(n), nil
}

// Fibonacci returns the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
