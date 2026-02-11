// Run: go test ./learnings/uber-style-guide/exercises/... -run TestParseHeadings -v

package exercises

import "testing"

func TestParseHeadings_Level2(t *testing.T) {
	input := "## Introduction\nSome text here."
	headings := ParseHeadings(input)
	if len(headings) != 1 {
		t.Fatalf("expected 1 heading, got %d", len(headings))
	}
	if headings[0].Level != 2 {
		t.Errorf("expected level 2, got %d", headings[0].Level)
	}
	if headings[0].Text != "Introduction" {
		t.Errorf("expected text 'Introduction', got %q", headings[0].Text)
	}
}

func TestParseHeadings_Level3(t *testing.T) {
	input := "### Error Wrapping\nDetails about errors."
	headings := ParseHeadings(input)
	if len(headings) != 1 {
		t.Fatalf("expected 1 heading, got %d", len(headings))
	}
	if headings[0].Level != 3 {
		t.Errorf("expected level 3, got %d", headings[0].Level)
	}
	if headings[0].Text != "Error Wrapping" {
		t.Errorf("expected text 'Error Wrapping', got %q", headings[0].Text)
	}
}

func TestParseHeadings_Multiple(t *testing.T) {
	input := "## Guidelines\nText.\n### Pointers\nMore text.\n### Errors\nEven more."
	headings := ParseHeadings(input)
	if len(headings) != 3 {
		t.Fatalf("expected 3 headings, got %d", len(headings))
	}
}

func TestParseHeadings_IgnoresH1(t *testing.T) {
	input := "# Title\n## Real Heading"
	headings := ParseHeadings(input)
	if len(headings) != 1 {
		t.Fatalf("expected 1 heading (ignoring h1), got %d", len(headings))
	}
	if headings[0].Text != "Real Heading" {
		t.Errorf("expected 'Real Heading', got %q", headings[0].Text)
	}
}

func TestParseHeadings_IgnoresH4(t *testing.T) {
	input := "## Section\n#### Too Deep"
	headings := ParseHeadings(input)
	if len(headings) != 1 {
		t.Fatalf("expected 1 heading (ignoring h4), got %d", len(headings))
	}
}

func TestParseHeadings_Empty(t *testing.T) {
	headings := ParseHeadings("Just some text without headings.")
	if len(headings) != 0 {
		t.Errorf("expected 0 headings, got %d", len(headings))
	}
}
