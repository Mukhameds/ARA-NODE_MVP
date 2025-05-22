// core/bootstrap_math_discrete.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMathDiscrete загружает понятия дискретной математики: логика, графы, булевы функции, комбинаторика
func BootstrapMathDiscrete(mem *MemoryEngine) {
	discrete := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Булева логика
		{"disc_boolean_def", "Boolean logic uses true and false values", []string{"math", "discrete", "logic", "boolean", "core", "bootstrap"}},
		{"disc_bool_and", "A ∧ B is true only if both A and B are true", []string{"math", "discrete", "boolean", "and", "core", "bootstrap"}},
		{"disc_bool_or", "A ∨ B is true if at least one of A or B is true", []string{"math", "discrete", "boolean", "or", "core", "bootstrap"}},
		{"disc_bool_not", "¬A is true if A is false", []string{"math", "discrete", "boolean", "not", "core", "bootstrap"}},
		{"disc_bool_xor", "A ⊕ B is true if A and B are different", []string{"math", "discrete", "boolean", "xor", "core", "bootstrap"}},

		// Графы
		{"disc_graph_def", "a graph is a set of nodes connected by edges", []string{"math", "discrete", "graph", "definition", "core", "bootstrap"}},
		{"disc_node", "a node is a point in a graph", []string{"math", "discrete", "graph", "node", "core", "bootstrap"}},
		{"disc_edge", "an edge connects two nodes", []string{"math", "discrete", "graph", "edge", "core", "bootstrap"}},
		{"disc_tree", "a tree is an acyclic connected graph", []string{"math", "discrete", "graph", "tree", "core", "bootstrap"}},
		{"disc_cycle", "a cycle is a path that starts and ends at the same node", []string{"math", "discrete", "graph", "cycle", "core", "bootstrap"}},

		// Множества и отношения
		{"disc_relation", "a relation connects elements of two sets", []string{"math", "discrete", "relation", "core", "bootstrap"}},
		{"disc_equiv", "an equivalence relation is reflexive, symmetric, and transitive", []string{"math", "discrete", "relation", "equivalence", "core", "bootstrap"}},

		// Комбинаторика
		{"disc_permutation", "a permutation is an ordered arrangement of elements", []string{"math", "discrete", "combinatorics", "permutation", "core", "bootstrap"}},
		{"disc_combination", "a combination is a selection of elements without order", []string{"math", "discrete", "combinatorics", "combination", "core", "bootstrap"}},
		{"disc_factorial", "n! is the product of all positive integers up to n", []string{"math", "discrete", "combinatorics", "factorial", "core", "bootstrap"}},
	}

	for _, d := range discrete {
		q := QBit{
			ID:        d.ID,
			Content:   d.Content,
			Tags:      d.Tags,
			Phase:     0.88,
			Weight:    0.96,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("🔂 [Bootstrap] Discrete mathematics concepts loaded.")
}
