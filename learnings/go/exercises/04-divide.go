package exercises

import "errors"

// ErrDivideByZero is returned when dividing by zero
var ErrDivideByZero = errors.New("cannot divide by zero")

// TODO: Implement Divide function
// - Takes two integers: a (numerator), b (denominator)
// - Returns (int, error)
// - If b is 0, return (0, ErrDivideByZero)
// - Otherwise return (a / b, nil)
//
// Example:
//   Divide(10, 2) returns (5, nil)
//   Divide(10, 0) returns (0, ErrDivideByZero)
