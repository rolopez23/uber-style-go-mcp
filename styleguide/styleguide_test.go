package styleguide

import (
	"os"
	"path/filepath"
	"testing"
)

func testStyleGuidePath(t *testing.T) string {
	t.Helper()
	// Walk up to find the repo root with the vendor submodule.
	path := filepath.Join("..", "vendor", "uber-go-guide", "style.md")
	if _, err := os.Stat(path); err != nil {
		t.Skipf("style.md not found at %s — run git submodule update --init", path)
	}
	return path
}

func loadTestStyleGuide(t *testing.T) *StyleGuide {
	t.Helper()
	sg, err := LoadStyleGuide(testStyleGuidePath(t))
	if err != nil {
		t.Fatalf("LoadStyleGuide failed: %v", err)
	}
	return sg
}

// --- TDD Cycle 1: LoadStyleGuide ---

func TestLoadStyleGuide_ReturnsNonNil(t *testing.T) {
	sg := loadTestStyleGuide(t)
	if sg == nil {
		t.Fatal("expected non-nil StyleGuide")
	}
}

func TestLoadStyleGuide_FileNotFound_ReturnsError(t *testing.T) {
	_, err := LoadStyleGuide("nonexistent/path/style.md")
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestLoadStyleGuide_ParsesSections(t *testing.T) {
	sg := loadTestStyleGuide(t)
	if len(sg.Sections) == 0 {
		t.Fatal("expected sections to be parsed, got 0")
	}
}

func TestLoadStyleGuide_SectionHasID(t *testing.T) {
	sg := loadTestStyleGuide(t)
	for _, s := range sg.Sections {
		if s.ID == "" {
			t.Errorf("section %q has empty ID", s.Title)
		}
	}
}

func TestLoadStyleGuide_SectionHasTitle(t *testing.T) {
	sg := loadTestStyleGuide(t)
	for _, s := range sg.Sections {
		if s.Title == "" {
			t.Errorf("section with ID %q has empty Title", s.ID)
		}
	}
}

func TestLoadStyleGuide_SectionHasCategory(t *testing.T) {
	sg := loadTestStyleGuide(t)
	validCategories := map[string]bool{
		"Introduction": true,
		"Guidelines":   true,
		"Performance":  true,
		"Style":        true,
		"Patterns":     true,
		"Linting":      true,
	}
	for _, s := range sg.Sections {
		if s.Category == "" {
			t.Errorf("section %q has empty Category", s.Title)
		}
		if !validCategories[s.Category] {
			t.Errorf("section %q has invalid Category %q", s.Title, s.Category)
		}
	}
}

func TestLoadStyleGuide_SectionHasContent(t *testing.T) {
	sg := loadTestStyleGuide(t)
	for _, s := range sg.Sections {
		if s.Content == "" {
			t.Errorf("section %q has empty Content", s.Title)
		}
	}
}

func TestLoadStyleGuide_SectionHasSummary(t *testing.T) {
	sg := loadTestStyleGuide(t)
	for _, s := range sg.Sections {
		if s.Summary == "" {
			t.Errorf("section %q has empty Summary", s.Title)
		}
		if len(s.Summary) > 250 {
			t.Errorf("section %q summary too long: %d chars", s.Title, len(s.Summary))
		}
	}
}

func TestLoadStyleGuide_KnownSectionExists(t *testing.T) {
	sg := loadTestStyleGuide(t)
	// "Error Wrapping" should exist as a known section.
	found := false
	for _, s := range sg.Sections {
		if s.ID == "error-wrapping" {
			found = true
			if s.Title != "Error Wrapping" {
				t.Errorf("expected title 'Error Wrapping', got %q", s.Title)
			}
			if s.Category != "Guidelines" {
				t.Errorf("expected category 'Guidelines', got %q", s.Category)
			}
			break
		}
	}
	if !found {
		t.Error("expected to find section with ID 'error-wrapping'")
	}
}

// --- TDD Cycle 2: ListSections + GetSections ---

func TestListSections_ReturnsAllSections(t *testing.T) {
	sg := loadTestStyleGuide(t)
	summaries := sg.ListSections()
	if len(summaries) == 0 {
		t.Fatal("expected summaries, got 0")
	}
}

func TestListSections_CountMatchesExpected(t *testing.T) {
	sg := loadTestStyleGuide(t)
	summaries := sg.ListSections()
	if len(summaries) != len(sg.Sections) {
		t.Errorf("ListSections count %d != Sections count %d", len(summaries), len(sg.Sections))
	}
}

func TestGetSections_ByID(t *testing.T) {
	sg := loadTestStyleGuide(t)
	sections, notFound := sg.GetSections([]string{"error-wrapping"})
	if len(notFound) > 0 {
		t.Errorf("unexpected not-found IDs: %v", notFound)
	}
	if len(sections) != 1 {
		t.Fatalf("expected 1 section, got %d", len(sections))
	}
	if sections[0].ID != "error-wrapping" {
		t.Errorf("expected ID 'error-wrapping', got %q", sections[0].ID)
	}
}

func TestGetSections_MultipleIDs(t *testing.T) {
	sg := loadTestStyleGuide(t)
	sections, notFound := sg.GetSections([]string{"error-wrapping", "dont-panic"})
	if len(notFound) > 0 {
		t.Errorf("unexpected not-found IDs: %v", notFound)
	}
	if len(sections) != 2 {
		t.Fatalf("expected 2 sections, got %d", len(sections))
	}
}

func TestGetSections_UnknownID_ReturnsPartial(t *testing.T) {
	sg := loadTestStyleGuide(t)
	sections, notFound := sg.GetSections([]string{"error-wrapping", "nonexistent-section"})
	if len(sections) != 1 {
		t.Errorf("expected 1 found section, got %d", len(sections))
	}
	if len(notFound) != 1 {
		t.Errorf("expected 1 not-found ID, got %d", len(notFound))
	}
	if notFound[0] != "nonexistent-section" {
		t.Errorf("expected not-found ID 'nonexistent-section', got %q", notFound[0])
	}
}
