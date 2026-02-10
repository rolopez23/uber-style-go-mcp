// Run: go test ./learnings/go-testing/exercises/... -bench=. -v

package exercises

import "testing"

// TODO: Implement BenchmarkFibonacci
// Benchmark the Fibonacci function with n=20
//
// Benchmark functions:
// - Start with "Benchmark" prefix
// - Take *testing.B parameter
// - Run the code b.N times in a loop
//
// Example:
//
//	func BenchmarkSomething(b *testing.B) {
//	    for i := 0; i < b.N; i++ {
//	        // code to benchmark
//	    }
//	}
//
// Run with: go test -bench=.
func BenchmarkFibonacci(b *testing.B) {
	// TODO: Implement this benchmark
	// Call Fibonacci(20) in a loop b.N times
	b.Skip("TODO: Implement this benchmark")
}

// TODO: Implement BenchmarkFibonacci_Small
// Benchmark Fibonacci with n=10 for comparison
func BenchmarkFibonacci_Small(b *testing.B) {
	// TODO: Implement this benchmark
	b.Skip("TODO: Implement this benchmark")
}
