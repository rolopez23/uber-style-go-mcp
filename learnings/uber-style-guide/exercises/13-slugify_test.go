// Run: go test ./learnings/uber-style-guide/exercises/... -run TestSlugify -v

package exercises

import "testing"

func TestSlugify_Simple(t *testing.T) {
	got := Slugify("Error Wrapping")
	if got != "error-wrapping" {
		t.Errorf("Slugify(%q) = %q, want %q", "Error Wrapping", got, "error-wrapping")
	}
}

func TestSlugify_Apostrophe(t *testing.T) {
	got := Slugify("Don't Panic")
	if got != "dont-panic" {
		t.Errorf("Slugify(%q) = %q, want %q", "Don't Panic", got, "dont-panic")
	}
}

func TestSlugify_MultipleSpaces(t *testing.T) {
	got := Slugify("Too   Many   Spaces")
	if got != "too-many-spaces" {
		t.Errorf("Slugify(%q) = %q, want %q", "Too   Many   Spaces", got, "too-many-spaces")
	}
}

func TestSlugify_Hyphenated(t *testing.T) {
	got := Slugify("fire-and-forget goroutines")
	if got != "fire-and-forget-goroutines" {
		t.Errorf("Slugify(%q) = %q, want %q", "fire-and-forget goroutines", got, "fire-and-forget-goroutines")
	}
}

func TestSlugify_SingleWord(t *testing.T) {
	got := Slugify("Errors")
	if got != "errors" {
		t.Errorf("Slugify(%q) = %q, want %q", "Errors", got, "errors")
	}
}
