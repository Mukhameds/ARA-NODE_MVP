// core/bootstrap_morphology.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMorphologyRules инициализирует базовые правила словообразования
func BootstrapMorphologyRules(mem *MemoryEngine) {
	// === Приставки (prefixes) ===
	prefixes := []string{
		"un",  // not
		"re",  // again
		"pre", // before
		"dis", // opposite
		"sub", // under
		"inter", // between
	}

	for _, p := range prefixes {
		q := QBit{
			ID:        "prefix_" + p,
			Content:   p,
			Tags:      []string{"morphology", "prefix", "core", "bootstrap"},
			Phase:     0.75,
			Weight:    0.85,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	// === Суффиксы (suffixes) ===
	suffixes := []string{
		"ing",  // action/state
		"ed",   // past tense
		"er",   // one who
		"ness", // state of being
		"ly",   // in manner
		"able", // can be
	}

	for _, s := range suffixes {
		q := QBit{
			ID:        "suffix_" + s,
			Content:   s,
			Tags:      []string{"morphology", "suffix", "core", "bootstrap"},
			Phase:     0.75,
			Weight:    0.85,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	// === Общие правила морфологии ===
	rules := []struct {
		ID      string
		Content string
	}{
		{"rule_prefix_root", "prefix + root → derived meaning"},
		{"rule_root_suffix", "root + suffix → derived word"},
		{"rule_wordchain", "prefix + root + suffix → full word"},
		{"rule_repeat_suffix", "verb + ing → ongoing action"},
		{"rule_past", "verb + ed → past form"},
	}

	for _, r := range rules {
		q := QBit{
			ID:        r.ID,
			Content:   r.Content,
			Tags:      []string{"morphology", "rule", "core", "bootstrap"},
			Phase:     0.9,
			Weight:    1.0,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("🔠 [Bootstrap] Morphology rules, prefixes, suffixes loaded.")
}
