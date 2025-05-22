// core/bootstrap_math_concepts.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMathConcepts загружает понятия числа, счёта, нуля и величины
func BootstrapMathConcepts(mem *MemoryEngine) {
	concepts := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Core number concepts
		{"math_concept_number", "number is a concept of quantity", []string{"math", "concept", "core", "bootstrap"}},
		{"math_concept_zero", "zero means nothing", []string{"math", "concept", "zero", "core", "bootstrap"}},
		{"math_concept_one", "one means a single unit", []string{"math", "concept", "one", "core", "bootstrap"}},
		{"math_concept_two", "two means a pair", []string{"math", "concept", "two", "core", "bootstrap"}},
		{"math_concept_three", "three means three units", []string{"math", "concept", "three", "core", "bootstrap"}},

		// Quantity comparison
		{"math_comparison_gt", "three is more than two", []string{"math", "comparison", "greater-than", "core", "bootstrap"}},
		{"math_comparison_lt", "one is less than two", []string{"math", "comparison", "less-than", "core", "bootstrap"}},
		{"math_comparison_eq", "one plus one equals two", []string{"math", "comparison", "equality", "core", "bootstrap"}},

		// Counting and size
		{"math_counting", "counting means assigning numbers to items", []string{"math", "process", "counting", "core", "bootstrap"}},
		{"math_quantity", "quantity means how much of something there is", []string{"math", "definition", "quantity", "core", "bootstrap"}},
	}

	for _, c := range concepts {
		q := QBit{
			ID:        c.ID,
			Content:   c.Content,
			Tags:      c.Tags,
			Phase:     0.83,
			Weight:    0.94,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("🔢 [Bootstrap] Math quantity and number concepts loaded.")
}
