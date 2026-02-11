// Run: go test ./learnings/uber-style-guide/exercises/... -run TestKeywordMatch -v

package exercises

// Rule maps a set of keywords to a label.
type Rule struct {
	Keywords []string
	Label    string
}

// TODO: Implement KeywordMatch
// Given a list of rules and input text, return matching labels.
//
// Rules:
// - For each rule, check if ANY of its keywords appear in the text
// - Use strings.Contains for matching (case-sensitive)
// - If a rule matches, include its label in the result
// - Deduplicate labels (each label should appear at most once)
// - Return nil if no rules match
//
// Hints:
// - Loop over rules, then over keywords within each rule
// - Use a map[string]bool to track seen labels
// - Break out of the keyword loop once you find a match
//
// Example:
//   rules := []Rule{
//       {Keywords: []string{"error", "err"}, Label: "error-handling"},
//       {Keywords: []string{"defer"}, Label: "resource-cleanup"},
//   }
//   KeywordMatch(rules, "if err != nil { defer cleanup() }")
//   // Returns: ["error-handling", "resource-cleanup"]

func KeywordMatch(rules []Rule, text string) []string {
	// TODO: Implement this function
	return nil
}
