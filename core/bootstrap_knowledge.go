// core/bootstrap_knowledge.go
package core

import (
	"fmt"
	"time"
)

// BootstrapKnowledgeConcepts загружает понятия знания, истины, проверки и источников
func BootstrapKnowledgeConcepts(mem *MemoryEngine) {
	concepts := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Определения
		{"knowledge_def", "knowledge is organized and verifiable information", []string{"knowledge", "definition", "core", "bootstrap"}},
		{"truth_def", "truth is consistency between statement and reality", []string{"knowledge", "truth", "core", "bootstrap"}},
		{"belief_def", "belief is information assumed to be true", []string{"knowledge", "belief", "core", "bootstrap"}},

		// Проверка и подтверждение
		{"verify_means_check", "to verify means to check against evidence", []string{"knowledge", "verification", "core", "bootstrap"}},
		{"evidence_supports_truth", "evidence supports the truth of a claim", []string{"knowledge", "evidence", "truth", "core", "bootstrap"}},
		{"sources_matter", "trusted sources increase confidence in knowledge", []string{"knowledge", "source", "trust", "core", "bootstrap"}},

		// Ошибки и ложь
		{"false_def", "false is opposite of true", []string{"knowledge", "truth", "false", "core", "bootstrap"}},
		{"misinformation", "misinformation is incorrect or misleading information", []string{"knowledge", "error", "misleading", "core", "bootstrap"}},
		{"uncertainty", "some knowledge is uncertain or incomplete", []string{"knowledge", "uncertainty", "core", "bootstrap"}},

		// Цикл познания
		{"learn_loop", "learning is acquiring and refining knowledge over time", []string{"knowledge", "learning", "process", "core", "bootstrap"}},
		{"doubt_triggers_search", "doubt triggers the search for better knowledge", []string{"knowledge", "doubt", "search", "core", "bootstrap"}},
	}

	for _, c := range concepts {
		q := QBit{
			ID:        c.ID,
			Content:   c.Content,
			Tags:      c.Tags,
			Phase:     0.88,
			Weight:    0.97,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("📚 [Bootstrap] Knowledge, truth, and epistemology loaded.")
}
