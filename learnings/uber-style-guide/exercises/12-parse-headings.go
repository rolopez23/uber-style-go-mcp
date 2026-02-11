// Run: go test ./learnings/uber-style-guide/exercises/... -run TestParseHeadings -v

package exercises

// Heading represents a parsed markdown heading.
type Heading struct {
	Level int    // 2 for ##, 3 for ###
	Text  string // The heading text (without # prefix)
}

// TODO: Implement ParseHeadings
// Given a markdown string, extract all ## and ### headings.
// Return a slice of Heading structs with the level and text.
//
// Rules:
// - Only match lines starting with "## " or "### "
// - Strip the "#" prefix and leading/trailing whitespace from the text
// - Ignore # (h1) and #### (h4) or deeper headings
//
// Hints:
// - Use strings.Split to split by newlines
// - Use strings.HasPrefix to check for "## " and "### "
// - Use strings.TrimPrefix and strings.TrimSpace
//
// Example:
//   input := "## Hello\nsome text\n### World"
//   result := ParseHeadings(input)
//   // result = [{Level: 2, Text: "Hello"}, {Level: 3, Text: "World"}]

func ParseHeadings(markdown string) []Heading {
	// TODO: Implement this function
	return nil
}
