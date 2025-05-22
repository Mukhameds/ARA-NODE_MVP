// core/bootstrap_logic_axioms.go
package core

import (
	"fmt"
	"time"
)

// BootstrapLogicAxioms загружает законы формальной логики как сигнальные QBits
func BootstrapLogicAxioms(mem *MemoryEngine) {
	logicLaws := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Основные законы логики
		{"logic_identity", "A = A", []string{"logic", "axiom", "identity", "core", "bootstrap"}},
		{"logic_noncontradiction", "¬(A ∧ ¬A)", []string{"logic", "axiom", "noncontradiction", "core", "bootstrap"}},
		{"logic_excluded_middle", "A ∨ ¬A", []string{"logic", "axiom", "excluded-middle", "core", "bootstrap"}},
		{"logic_double_negation", "¬(¬A) = A", []string{"logic", "axiom", "negation", "core", "bootstrap"}},

		// Законы распределения
		{"logic_distributive_and_over_or", "A ∧ (B ∨ C) = (A ∧ B) ∨ (A ∧ C)", []string{"logic", "axiom", "distributive", "core", "bootstrap"}},
		{"logic_distributive_or_over_and", "A ∨ (B ∧ C) = (A ∨ B) ∧ (A ∨ C)", []string{"logic", "axiom", "distributive", "core", "bootstrap"}},

		// Де Моргана
		{"logic_demorgan_1", "¬(A ∧ B) = ¬A ∨ ¬B", []string{"logic", "axiom", "demorgan", "core", "bootstrap"}},
		{"logic_demorgan_2", "¬(A ∨ B) = ¬A ∧ ¬B", []string{"logic", "axiom", "demorgan", "core", "bootstrap"}},

		// Контрапозиция
		{"logic_contrapositive", "A → B = ¬B → ¬A", []string{"logic", "axiom", "contrapositive", "core", "bootstrap"}},
	}

	for _, law := range logicLaws {
		q := QBit{
			ID:        law.ID,
			Content:   law.Content,
			Tags:      law.Tags,
			Phase:     0.86,
			Weight:    0.96,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("🔁 [Bootstrap] Formal logic axioms loaded.")
}
