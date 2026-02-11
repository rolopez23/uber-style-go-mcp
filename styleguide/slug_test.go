package styleguide

import "testing"

func TestSlug_Simple(t *testing.T) {
	got := slug("Error Wrapping")
	want := "error-wrapping"
	if got != want {
		t.Errorf("slug(%q) = %q, want %q", "Error Wrapping", got, want)
	}
}

func TestSlug_Backticks(t *testing.T) {
	got := slug("Use `\"time\"` to handle time")
	want := "use-time-to-handle-time"
	if got != want {
		t.Errorf("slug(%q) = %q, want %q", "Use `time`", got, want)
	}
}

func TestSlug_SpecialChars(t *testing.T) {
	got := slug("Don't Panic")
	want := "dont-panic"
	if got != want {
		t.Errorf("slug(%q) = %q, want %q", "Don't Panic", got, want)
	}
}

func TestSlug_Underscores(t *testing.T) {
	got := slug("Prefix Unexported Globals with _")
	want := "prefix-unexported-globals-with-_"
	if got != want {
		t.Errorf("slug(%q) = %q, want %q", "Prefix Unexported Globals with _", got, want)
	}
}

func TestSlug_MultipleSpaces(t *testing.T) {
	got := slug("Too   Many   Spaces")
	want := "too-many-spaces"
	if got != want {
		t.Errorf("slug(%q) = %q, want %q", "Too   Many   Spaces", got, want)
	}
}

func TestSlug_AvoidInit(t *testing.T) {
	got := slug("Avoid `init()`")
	want := "avoid-init"
	if got != want {
		t.Errorf("slug(%q) = %q, want %q", "Avoid `init()`", got, want)
	}
}

func TestSlug_DontFireAndForget(t *testing.T) {
	got := slug("Don't fire-and-forget goroutines")
	want := "dont-fire-and-forget-goroutines"
	if got != want {
		t.Errorf("slug(%q) = %q, want %q", "Don't fire-and-forget goroutines", got, want)
	}
}
