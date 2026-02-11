// Run: go test ./learnings/uber-style-guide/exercises/... -run TestKeywordMatch -v

package exercises

import "testing"

var testRules = []Rule{
	{Keywords: []string{"error", "fmt.Errorf"}, Label: "error-handling"},
	{Keywords: []string{"go func"}, Label: "goroutines"},
	{Keywords: []string{"defer"}, Label: "defer"},
	{Keywords: []string{"map["}, Label: "maps"},
}

func TestKeywordMatch_SingleMatch(t *testing.T) {
	result := KeywordMatch(testRules, "defer f.Close()")
	if len(result) != 1 {
		t.Fatalf("expected 1 match, got %d", len(result))
	}
	if result[0] != "defer" {
		t.Errorf("expected label 'defer', got %q", result[0])
	}
}

func TestKeywordMatch_MultipleMatches(t *testing.T) {
	result := KeywordMatch(testRules, `
		if error != nil {
			go func() { process() }()
		}
	`)
	if len(result) < 2 {
		t.Fatalf("expected at least 2 matches, got %d", len(result))
	}
	labels := make(map[string]bool)
	for _, l := range result {
		labels[l] = true
	}
	if !labels["error-handling"] {
		t.Error("expected 'error-handling' label")
	}
	if !labels["goroutines"] {
		t.Error("expected 'goroutines' label")
	}
}

func TestKeywordMatch_NoMatch(t *testing.T) {
	result := KeywordMatch(testRules, "x := 1 + 2")
	if result != nil {
		t.Errorf("expected nil for no matches, got %v", result)
	}
}

func TestKeywordMatch_NoDuplicates(t *testing.T) {
	// Both "error" and "fmt.Errorf" are in the same rule.
	result := KeywordMatch(testRules, `fmt.Errorf("error: %w", error)`)
	count := 0
	for _, l := range result {
		if l == "error-handling" {
			count++
		}
	}
	if count > 1 {
		t.Errorf("expected 'error-handling' at most once, got %d times", count)
	}
}

func TestKeywordMatch_MapPattern(t *testing.T) {
	result := KeywordMatch(testRules, `m := map[string]int{}`)
	if len(result) != 1 {
		t.Fatalf("expected 1 match, got %d", len(result))
	}
	if result[0] != "maps" {
		t.Errorf("expected label 'maps', got %q", result[0])
	}
}
