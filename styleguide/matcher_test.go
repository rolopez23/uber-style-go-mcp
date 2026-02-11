package styleguide

import (
	"testing"
)

func hasSectionID(sections []Section, id string) bool {
	for _, s := range sections {
		if s.ID == id {
			return true
		}
	}
	return false
}

func TestMatchSections_ErrorPatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
func doSomething() error {
	err := someFunc()
	if err != nil {
		return fmt.Errorf("failed: %w", err)
	}
	return nil
}
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected error-related sections, got 0")
	}
	expectedIDs := []string{"errors", "error-wrapping", "error-types", "handle-errors-once"}
	for _, id := range expectedIDs {
		if !hasSectionID(sections, id) {
			t.Errorf("expected section %q in results", id)
		}
	}
}

func TestMatchSections_InterfacePatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
func process(v interface{}) {
	// do something
}
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected interface-related sections, got 0")
	}
	if !hasSectionID(sections, "pointers-to-interfaces") {
		t.Error("expected section 'pointers-to-interfaces'")
	}
	if !hasSectionID(sections, "verify-interface-compliance") {
		t.Error("expected section 'verify-interface-compliance'")
	}
}

func TestMatchSections_GoroutinePatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
func start() {
	go func() {
		doWork()
	}()
}
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected goroutine-related sections, got 0")
	}
	if !hasSectionID(sections, "dont-fire-and-forget-goroutines") {
		t.Error("expected section 'dont-fire-and-forget-goroutines'")
	}
}

func TestMatchSections_MutexPatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected mutex-related sections, got 0")
	}
	if !hasSectionID(sections, "zero-value-mutexes-are-valid") {
		t.Error("expected section 'zero-value-mutexes-are-valid'")
	}
}

func TestMatchSections_MapPatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
m := map[string]int{
	"foo": 1,
}
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected map-related sections, got 0")
	}
	if !hasSectionID(sections, "initializing-maps") {
		t.Error("expected section 'initializing-maps'")
	}
}

func TestMatchSections_DeferPatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
func readFile(path string) {
	f, _ := os.Open(path)
	defer f.Close()
}
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected defer-related sections, got 0")
	}
	if !hasSectionID(sections, "defer-to-clean-up") {
		t.Error("expected section 'defer-to-clean-up'")
	}
}

func TestMatchSections_InitPatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
func init() {
	registerPlugin()
}
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected init-related sections, got 0")
	}
	if !hasSectionID(sections, "avoid-init") {
		t.Error("expected section 'avoid-init'")
	}
}

func TestMatchSections_StructPatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
type Config struct {
	Name    string
	Timeout int
}
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected struct-related sections, got 0")
	}
	if !hasSectionID(sections, "initializing-structs") {
		t.Error("expected section 'initializing-structs'")
	}
	if !hasSectionID(sections, "use-field-tags-in-marshaled-structs") {
		t.Error("expected section 'use-field-tags-in-marshaled-structs'")
	}
}

func TestMatchSections_SlicePatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
items := make([]string, 0, 10)
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected slice-related sections, got 0")
	}
	if !hasSectionID(sections, "prefer-specifying-container-capacity") {
		t.Error("expected section 'prefer-specifying-container-capacity'")
	}
}

func TestMatchSections_ChannelPatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
ch := make(chan int, 1)
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected channel-related sections, got 0")
	}
	if !hasSectionID(sections, "channel-size-is-one-or-none") {
		t.Error("expected section 'channel-size-is-one-or-none'")
	}
}

func TestMatchSections_ImportPatterns(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)
`
	sections := sg.MatchSections(code)
	if len(sections) == 0 {
		t.Fatal("expected import-related sections, got 0")
	}
	if !hasSectionID(sections, "import-group-ordering") {
		t.Error("expected section 'import-group-ordering'")
	}
}

func TestMatchSections_NoMatch(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
x := 1 + 2
`
	sections := sg.MatchSections(code)
	if len(sections) != 0 {
		t.Errorf("expected 0 sections for vanilla code, got %d", len(sections))
	}
}

func TestMatchSections_MultipleMatches(t *testing.T) {
	sg := loadTestStyleGuide(t)

	code := `
func doWork() error {
	go func() {
		processItem()
	}()
	return fmt.Errorf("something failed: %w", err)
}
`
	sections := sg.MatchSections(code)
	if len(sections) < 2 {
		t.Errorf("expected multiple matches for error + goroutine code, got %d", len(sections))
	}
	// Should have both error and goroutine sections.
	hasError := hasSectionID(sections, "errors") || hasSectionID(sections, "error-wrapping")
	hasGoroutine := hasSectionID(sections, "dont-fire-and-forget-goroutines")
	if !hasError {
		t.Error("expected error-related section")
	}
	if !hasGoroutine {
		t.Error("expected goroutine-related section")
	}
}

func TestMatchSections_NoDuplicates(t *testing.T) {
	sg := loadTestStyleGuide(t)

	// Code that triggers the same section from multiple rules.
	code := `
err := fmt.Errorf("wrap: %w", errors.New("base"))
`
	sections := sg.MatchSections(code)
	seen := make(map[string]bool)
	for _, s := range sections {
		if seen[s.ID] {
			t.Errorf("duplicate section ID in results: %q", s.ID)
		}
		seen[s.ID] = true
	}
}
