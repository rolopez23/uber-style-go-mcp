// Run: go test ./learnings/go-testing/exercises/... -run TestMultiply -v

package exercises

import "testing"

// TODO: Implement TestMultiply_TableDriven
// Use table-driven tests to test multiple cases:
// - 2 * 3 = 6
// - 0 * 5 = 0
// - -2 * 3 = -6
// - -2 * -3 = 6
//
// Example table-driven test structure:
//
//	tests := []struct {
//	    name     string
//	    a, b     int
//	    expected int
//	}{
//	    {"positive", 2, 3, 6},
//	    // add more cases...
//	}
//
//	for _, tt := range tests {
//	    t.Run(tt.name, func(t *testing.T) {
//	        result := Multiply(tt.a, tt.b)
//	        if result != tt.expected {
//	            t.Errorf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
//	        }
//	    })
//	}
func TestMultiply_TableDriven(t *testing.T) {
	// TODO: Implement this test using table-driven approach
	t.Skip("TODO: Implement this test")
}
