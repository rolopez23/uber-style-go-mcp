// Run: go test ./learnings/uber-style-guide/exercises/... -run TestClosureHandler -v

package exercises

import "testing"

func TestClosureHandler_Hello(t *testing.T) {
	handler := NewGreetHandler("Hello")
	if handler == nil {
		t.Fatal("NewGreetHandler must return a non-nil function")
	}
	got := handler("World")
	if got != "Hello, World!" {
		t.Errorf("expected %q, got %q", "Hello, World!", got)
	}
}

func TestClosureHandler_Hi(t *testing.T) {
	handler := NewGreetHandler("Hi")
	if handler == nil {
		t.Fatal("NewGreetHandler must return a non-nil function")
	}
	got := handler("Go")
	if got != "Hi, Go!" {
		t.Errorf("expected %q, got %q", "Hi, Go!", got)
	}
}

func TestClosureHandler_DifferentPrefixes(t *testing.T) {
	hello := NewGreetHandler("Hello")
	hey := NewGreetHandler("Hey")

	if hello == nil || hey == nil {
		t.Fatal("handlers must not be nil")
	}

	// Each closure captures its own prefix.
	r1 := hello("Alice")
	r2 := hey("Alice")

	if r1 == r2 {
		t.Errorf("different prefixes should produce different results, both got %q", r1)
	}
	if r1 != "Hello, Alice!" {
		t.Errorf("expected %q, got %q", "Hello, Alice!", r1)
	}
	if r2 != "Hey, Alice!" {
		t.Errorf("expected %q, got %q", "Hey, Alice!", r2)
	}
}

func TestClosureHandler_EmptyName(t *testing.T) {
	handler := NewGreetHandler("Greetings")
	if handler == nil {
		t.Fatal("NewGreetHandler must return a non-nil function")
	}
	got := handler("")
	if got != "Greetings, !" {
		t.Errorf("expected %q, got %q", "Greetings, !", got)
	}
}
