// core/bootstrap_math_sets.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMathSets загружает основные сигнальные понятия теории множеств
func BootstrapMathSets(mem *MemoryEngine) {
	sets := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Определения и базовые объекты
		{"set_def", "a set is a collection of distinct elements", []string{"math", "set", "definition", "core", "bootstrap"}},
		{"set_element_def", "an element is a single member of a set", []string{"math", "set", "element", "definition", "core", "bootstrap"}},

		// Принадлежность и включение
		{"set_in", "x ∈ A means x is an element of set A", []string{"math", "set", "in", "relation", "core", "bootstrap"}},
		{"set_notin", "x ∉ A means x is not in set A", []string{"math", "set", "notin", "relation", "core", "bootstrap"}},

		// Операции над множествами
		{"set_union", "A ∪ B is the union of sets A and B", []string{"math", "set", "union", "operation", "core", "bootstrap"}},
		{"set_intersection", "A ∩ B is the intersection of A and B", []string{"math", "set", "intersection", "operation", "core", "bootstrap"}},
		{"set_difference", "A − B is the set of elements in A but not in B", []string{"math", "set", "difference", "operation", "core", "bootstrap"}},
		{"set_subset", "A ⊆ B means A is a subset of B", []string{"math", "set", "subset", "relation", "core", "bootstrap"}},
		{"set_proper_subset", "A ⊂ B means A is a proper subset of B", []string{"math", "set", "subset", "relation", "core", "bootstrap"}},

		// Особые множества
		{"set_empty", "∅ is the empty set with no elements", []string{"math", "set", "empty", "core", "bootstrap"}},
		{"set_universal", "U is the universal set containing all elements", []string{"math", "set", "universal", "core", "bootstrap"}},

		// Мощность и размер
		{"set_cardinality", "|A| is the number of elements in set A", []string{"math", "set", "cardinality", "core", "bootstrap"}},
	}

	for _, s := range sets {
		q := QBit{
			ID:        s.ID,
			Content:   s.Content,
			Tags:      s.Tags,
			Phase:     0.87,
			Weight:    0.95,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("📦 [Bootstrap] Mathematical sets and relations loaded.")
}
