// Run: go test ./learnings/go/exercises/... -run TestGreet -v

package exercises

import "testing"

func TestGreet_Basic(t *testing.T) {
	result := Greet("World")
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Greet(\"World\") = %q, want %q", result, expected)
	}
}

func TestGreet_Name(t *testing.T) {
	result := Greet("Alice")
	expected := "Hello, Alice!"
	if result != expected {
		t.Errorf("Greet(\"Alice\") = %q, want %q", result, expected)
	}
}

func TestGreet_Empty(t *testing.T) {
	result := Greet("")
	expected := "Hello, !"
	if result != expected {
		t.Errorf("Greet(\"\") = %q, want %q", result, expected)
	}
}
