// core/bootstrap_math_axioms.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMathAxioms загружает базовые арифметические аксиомы и свойства операций
func BootstrapMathAxioms(mem *MemoryEngine) {
	axioms := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Ассоциативность
		{"axiom_assoc_add", "(a + b) + c = a + (b + c)", []string{"math", "axiom", "associative", "addition", "core", "bootstrap"}},
		{"axiom_assoc_mul", "(a * b) * c = a * (b * c)", []string{"math", "axiom", "associative", "multiplication", "core", "bootstrap"}},

		// Коммутативность (повтор для аксиоматики)
		{"axiom_comm_add", "a + b = b + a", []string{"math", "axiom", "commutative", "addition", "core", "bootstrap"}},
		{"axiom_comm_mul", "a * b = b * a", []string{"math", "axiom", "commutative", "multiplication", "core", "bootstrap"}},

		// Дистрибутивность
		{"axiom_distr", "a * (b + c) = a * b + a * c", []string{"math", "axiom", "distributive", "core", "bootstrap"}},

		// Нейтральные элементы
		{"axiom_zero_add", "a + 0 = a", []string{"math", "axiom", "identity", "addition", "core", "bootstrap"}},
		{"axiom_one_mul", "a * 1 = a", []string{"math", "axiom", "identity", "multiplication", "core", "bootstrap"}},

		// Инверсии
		{"axiom_add_inverse", "a + (-a) = 0", []string{"math", "axiom", "inverse", "addition", "core", "bootstrap"}},
		{"axiom_mul_inverse", "a ≠ 0 → a * (1/a) = 1", []string{"math", "axiom", "inverse", "multiplication", "core", "bootstrap"}},

		// Равенство и замена
		{"axiom_eq_subst", "if a = b then a can be replaced by b", []string{"math", "axiom", "equality", "substitution", "core", "bootstrap"}},
	}

	for _, a := range axioms {
		q := QBit{
			ID:        a.ID,
			Content:   a.Content,
			Tags:      a.Tags,
			Phase:     0.85,
			Weight:    0.96,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("📐 [Bootstrap] Math axioms and operational laws loaded.")
}
