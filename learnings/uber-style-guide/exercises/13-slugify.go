// Run: go test ./learnings/uber-style-guide/exercises/... -run TestSlugify -v

package exercises

// TODO: Implement Slugify
// Convert a heading string to a URL-safe slug.
//
// Rules:
// - Convert to lowercase
// - Replace spaces with hyphens
// - Remove special characters except hyphens and underscores
// - Collapse multiple consecutive hyphens into one
// - Trim leading/trailing hyphens
//
// Hints:
// - Use strings.ToLower for lowercase
// - Use strings.ReplaceAll for space→hyphen
// - Use regexp.MustCompile for pattern removal
// - Consider a two-pass approach: first clean, then collapse
//
// Examples:
//   Slugify("Error Wrapping")   → "error-wrapping"
//   Slugify("Don't Panic")      → "dont-panic"
//   Slugify("Too   Many Spaces") → "too-many-spaces"

func Slugify(heading string) string {
	// TODO: Implement this function
	return ""
}
