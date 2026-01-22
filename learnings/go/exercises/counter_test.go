// Run: go test ./learnings/go/exercises/... -run TestCounter -v

package exercises

import "testing"

func TestCounter_StartsAtZero(t *testing.T) {
	c := &Counter{}
	if c.Value() != 0 {
		t.Errorf("new Counter should start at 0, got %d", c.Value())
	}
}

func TestCounter_Increment(t *testing.T) {
	c := &Counter{}
	c.Increment()
	if c.Value() != 1 {
		t.Errorf("after one Increment, Value should be 1, got %d", c.Value())
	}
}

func TestCounter_MultipleIncrements(t *testing.T) {
	c := &Counter{}
	c.Increment()
	c.Increment()
	c.Increment()
	if c.Value() != 3 {
		t.Errorf("after three Increments, Value should be 3, got %d", c.Value())
	}
}
