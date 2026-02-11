package styleguide

import "strings"

// matchRule maps code patterns to relevant style guide section IDs.
type matchRule struct {
	patterns   []string
	sectionIDs []string
}

// matchRules defines the keyword→section mapping table.
// Over-matching is preferred over under-matching.
var matchRules = []matchRule{
	{
		patterns:   []string{"interface{}", "any"},
		sectionIDs: []string{"pointers-to-interfaces", "verify-interface-compliance"},
	},
	{
		patterns:   []string{"error", "fmt.Errorf", "errors.New", "errors.Is", "errors.As"},
		sectionIDs: []string{"errors", "error-types", "error-wrapping", "error-naming", "handle-errors-once"},
	},
	{
		patterns:   []string{"go func", "goroutine"},
		sectionIDs: []string{"dont-fire-and-forget-goroutines", "wait-for-goroutines-to-exit"},
	},
	{
		patterns:   []string{"sync.Mutex", "sync.RWMutex"},
		sectionIDs: []string{"zero-value-mutexes-are-valid"},
	},
	{
		patterns:   []string{"map["},
		sectionIDs: []string{"initializing-maps", "copy-slices-and-maps-at-boundaries", "specifying-map-capacity-hints"},
	},
	{
		patterns:   []string{"chan "},
		sectionIDs: []string{"channel-size-is-one-or-none"},
	},
	{
		patterns:   []string{"defer "},
		sectionIDs: []string{"defer-to-clean-up"},
	},
	{
		patterns:   []string{"func init()", "init()"},
		sectionIDs: []string{"avoid-init", "no-goroutines-in-init"},
	},
	{
		patterns:   []string{"struct {"},
		sectionIDs: []string{"initializing-structs", "use-field-tags-in-marshaled-structs", "embedding-in-structs", "avoid-embedding-types-in-public-structs"},
	},
	{
		patterns:   []string{"make([]"},
		sectionIDs: []string{"prefer-specifying-container-capacity", "specifying-slice-capacity", "nil-is-a-valid-slice"},
	},
	{
		patterns:   []string{"import ("},
		sectionIDs: []string{"import-group-ordering", "import-aliasing"},
	},
	{
		patterns:   []string{"atomic."},
		sectionIDs: []string{"use-gouberorgatomic"},
	},
	{
		patterns:   []string{"panic("},
		sectionIDs: []string{"dont-panic"},
	},
	{
		patterns:   []string{"_test.go", "func Test"},
		sectionIDs: []string{"test-tables"},
	},
	{
		patterns:   []string{"strconv."},
		sectionIDs: []string{"prefer-strconv-over-fmt"},
	},
	{
		patterns:   []string{"[]byte("},
		sectionIDs: []string{"avoid-repeated-string-to-byte-conversions"},
	},
	{
		patterns:   []string{"func(opts", "Option)", "WithTimeout", "functional option"},
		sectionIDs: []string{"functional-options"},
	},
	{
		patterns:   []string{"iota"},
		sectionIDs: []string{"start-enums-at-one"},
	},
	{
		patterns:   []string{"time.Time", "time.Duration", "time.Now"},
		sectionIDs: []string{"use-time-to-handle-time"},
	},
}

// MatchSections scans the code for keyword patterns and returns matching
// style guide sections. Deduplicates results.
func (sg *StyleGuide) MatchSections(code string) []Section {
	matched := make(map[string]bool)

	for _, rule := range matchRules {
		for _, pattern := range rule.patterns {
			if strings.Contains(code, pattern) {
				for _, id := range rule.sectionIDs {
					matched[id] = true
				}
				break // One pattern match per rule is enough.
			}
		}
	}

	if len(matched) == 0 {
		return nil
	}

	ids := make([]string, 0, len(matched))
	for id := range matched {
		ids = append(ids, id)
	}

	sections, _ := sg.GetSections(ids)
	return sections
}
