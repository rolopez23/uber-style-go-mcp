package styleguide

import (
	"regexp"
	"strings"
)

var (
	// Remove backticks and their content's surrounding quotes.
	backtickRe = regexp.MustCompile("`[^`]*`")
	// Keep only alphanumeric, spaces, hyphens, and underscores.
	stripRe = regexp.MustCompile(`[^\w\s-]`)
	// Collapse multiple spaces/hyphens into a single hyphen.
	collapseRe = regexp.MustCompile(`[\s-]+`)
	// Remove leading/trailing hyphens.
	trimHyphenRe = regexp.MustCompile(`^-+|-+$`)
)

// slug converts a heading string into a URL-safe slug.
// Example: "Error Wrapping" → "error-wrapping"
func slug(heading string) string {
	s := heading

	// Extract text from backtick-quoted content (strip backticks and quotes).
	s = backtickRe.ReplaceAllStringFunc(s, func(match string) string {
		inner := strings.Trim(match, "`")
		inner = strings.Trim(inner, `"`)
		return inner
	})

	// Remove special characters except underscores and hyphens.
	s = stripRe.ReplaceAllString(s, "")

	// Lowercase.
	s = strings.ToLower(s)

	// Collapse whitespace/hyphens to single hyphen.
	s = collapseRe.ReplaceAllString(s, "-")

	// Trim leading/trailing hyphens.
	s = trimHyphenRe.ReplaceAllString(s, "")

	return s
}
