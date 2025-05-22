// core/bootstrap_grammar.go
package core

import (
	"fmt"
	"time"
)

// BootstrapGrammarStructures инициализирует базовые грамматические категории и структуры
func BootstrapGrammarStructures(mem *MemoryEngine) {
	// === Части речи ===
	partsOfSpeech := []struct {
		ID      string
		Word    string
		RoleTag string
	}{
		{"noun_person", "person", "noun"},
		{"noun_idea", "idea", "noun"},
		{"verb_run", "run", "verb"},
		{"verb_know", "know", "verb"},
		{"adj_happy", "happy", "adjective"},
		{"adj_large", "large", "adjective"},
		{"adv_quickly", "quickly", "adverb"},
		{"prep_with", "with", "preposition"},
		{"pron_he", "he", "pronoun"},
		{"conj_and", "and", "conjunction"},
	}

	for _, item := range partsOfSpeech {
		q := QBit{
			ID:        item.ID,
			Content:   item.Word,
			Tags:      []string{"grammar", "part-of-speech", item.RoleTag, "core", "bootstrap"},
			Phase:     0.78,
			Weight:    0.9,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	// === Грамматические структуры / шаблоны предложения ===
	structures := []string{
		"subject + verb + object",
		"noun + verb + noun",
		"adjective + noun",
		"pronoun + verb",
		"noun + verb + preposition + noun",
		"if + condition + then + result",
	}

	for i, pattern := range structures {
		q := QBit{
			ID:        fmt.Sprintf("sentence_structure_%d", i),
			Content:   pattern,
			Tags:      []string{"grammar", "structure", "pattern", "core", "bootstrap"},
			Phase:     0.85,
			Weight:    1.0,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("🧠 [Bootstrap] Grammar roles and sentence structures loaded.")
}
