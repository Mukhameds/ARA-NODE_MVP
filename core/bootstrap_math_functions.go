// core/bootstrap_math_functions.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMathFunctions загружает понятия математических функций, отображений и их свойств
func BootstrapMathFunctions(mem *MemoryEngine) {
	functions := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Определения
		{"func_def", "a function maps each input to exactly one output", []string{"math", "function", "definition", "core", "bootstrap"}},
		{"func_symbol", "f(x) represents a function named f applied to x", []string{"math", "function", "notation", "core", "bootstrap"}},

		// Область и множество значений
		{"func_domain", "the domain of a function is the set of valid inputs", []string{"math", "function", "domain", "core", "bootstrap"}},
		{"func_range", "the range of a function is the set of possible outputs", []string{"math", "function", "range", "core", "bootstrap"}},

		// Свойства отображений
		{"func_injective", "a function is injective if it maps distinct inputs to distinct outputs", []string{"math", "function", "injective", "core", "bootstrap"}},
		{"func_surjective", "a function is surjective if every element in the range is covered", []string{"math", "function", "surjective", "core", "bootstrap"}},
		{"func_bijective", "a bijective function is both injective and surjective", []string{"math", "function", "bijective", "core", "bootstrap"}},

		// Композиция
		{"func_composition", "composition of functions: (f ∘ g)(x) = f(g(x))", []string{"math", "function", "composition", "core", "bootstrap"}},
		{"func_identity", "identity function: id(x) = x", []string{"math", "function", "identity", "core", "bootstrap"}},

		// Специальные примеры
		{"func_square", "f(x) = x² is a function that squares its input", []string{"math", "function", "example", "core", "bootstrap"}},
		{"func_abs", "f(x) = |x| returns the absolute value of x", []string{"math", "function", "example", "core", "bootstrap"}},
	}

	for _, f := range functions {
		q := QBit{
			ID:        f.ID,
			Content:   f.Content,
			Tags:      f.Tags,
			Phase:     0.88,
			Weight:    0.96,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("🔁 [Bootstrap] Mathematical functions and mappings loaded.")
}
