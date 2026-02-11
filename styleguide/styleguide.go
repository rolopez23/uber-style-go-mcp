package styleguide

import (
	"fmt"
	"os"
	"strings"
)

// Section represents a parsed section of the style guide.
type Section struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Summary  string `json:"summary"`
	Content  string `json:"content"`
	Level    int    `json:"level"`    // Heading level (2=##, 3=###, 4=####)
	ParentID string `json:"parentID"` // ID of the parent section
}

// SectionSummary is a lightweight view of a section without full content.
type SectionSummary struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Summary  string `json:"summary"`
}

// StyleGuide holds all parsed sections and an index for fast lookup.
type StyleGuide struct {
	Sections []Section
	index    map[string]*Section
}

// LoadStyleGuide reads and parses the style guide markdown file.
func LoadStyleGuide(path string) (*StyleGuide, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("style.md not found — run `git submodule update --init`: %w", err)
	}

	sections := parseSections(string(data))

	sg := &StyleGuide{
		Sections: sections,
		index:    make(map[string]*Section, len(sections)),
	}
	for i := range sg.Sections {
		sg.index[sg.Sections[i].ID] = &sg.Sections[i]
	}

	return sg, nil
}

// ListSections returns a summary of all sections.
func (sg *StyleGuide) ListSections() []SectionSummary {
	summaries := make([]SectionSummary, len(sg.Sections))
	for i, s := range sg.Sections {
		summaries[i] = SectionSummary{
			ID:       s.ID,
			Title:    s.Title,
			Category: s.Category,
			Summary:  s.Summary,
		}
	}
	return summaries
}

// GetSections returns full sections for the given IDs.
// It returns the found sections and a list of IDs that were not found.
func (sg *StyleGuide) GetSections(ids []string) ([]Section, []string) {
	var found []Section
	var notFound []string
	for _, id := range ids {
		if s, ok := sg.index[id]; ok {
			found = append(found, *s)
		} else {
			notFound = append(notFound, id)
		}
	}
	return found, notFound
}

// headingLevel returns the heading level (number of # characters) and the
// heading text. Returns 0 if the line is not a heading.
func headingLevel(line string) (int, string) {
	trimmed := strings.TrimLeft(line, "#")
	level := len(line) - len(trimmed)
	if level < 2 || level > 4 {
		return 0, ""
	}
	// Must have a space after the # characters.
	if len(trimmed) == 0 || trimmed[0] != ' ' {
		return 0, ""
	}
	return level, strings.TrimSpace(trimmed)
}

// extractSummary returns the first paragraph of content, trimmed to ~200 chars.
func extractSummary(content string) string {
	// Skip the heading line itself.
	lines := strings.SplitN(content, "\n", 2)
	if len(lines) < 2 {
		return strings.TrimSpace(content)
	}
	body := strings.TrimSpace(lines[1])

	// Find the first non-empty paragraph.
	paragraphs := strings.SplitN(body, "\n\n", 2)
	summary := strings.TrimSpace(paragraphs[0])

	// Collapse newlines within the paragraph.
	summary = strings.ReplaceAll(summary, "\n", " ")

	// Truncate to ~200 chars at a word boundary.
	if len(summary) > 200 {
		cut := strings.LastIndex(summary[:200], " ")
		if cut > 100 {
			summary = summary[:cut] + "..."
		} else {
			summary = summary[:200] + "..."
		}
	}

	return summary
}

// parseSections splits the markdown into sections based on ## / ### / #### headings.
func parseSections(content string) []Section {
	lines := strings.Split(content, "\n")

	type rawSection struct {
		level    int
		title    string
		startIdx int
	}

	var raws []rawSection

	for i, line := range lines {
		level, title := headingLevel(line)
		if level > 0 {
			raws = append(raws, rawSection{level: level, title: title, startIdx: i})
		}
	}

	if len(raws) == 0 {
		return nil
	}

	// Determine categories: ## headings are category boundaries.
	// Track the current category and parent at each level.
	var sections []Section
	currentCategory := ""
	// parentStack tracks the most recent heading ID at each level.
	parentAtLevel := map[int]string{}

	for i, raw := range raws {
		// Determine end line.
		endIdx := len(lines)
		if i+1 < len(raws) {
			endIdx = raws[i+1].startIdx
		}

		sectionContent := strings.Join(lines[raw.startIdx:endIdx], "\n")
		sectionContent = strings.TrimRight(sectionContent, "\n ")

		id := slug(raw.title)

		// ## headings set the category.
		if raw.level == 2 {
			currentCategory = raw.title
		}

		// Determine parent.
		parentID := ""
		if raw.level > 2 {
			parentID = parentAtLevel[raw.level-1]
		}

		parentAtLevel[raw.level] = id

		summary := extractSummary(sectionContent)

		sections = append(sections, Section{
			ID:       id,
			Title:    raw.title,
			Category: currentCategory,
			Summary:  summary,
			Content:  sectionContent,
			Level:    raw.level,
			ParentID: parentID,
		})
	}

	return sections
}
