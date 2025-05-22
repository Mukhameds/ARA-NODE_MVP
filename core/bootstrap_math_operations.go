// core/bootstrap_math_operations.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMathOperations загружает базовые операции и арифметические правила
func BootstrapMathOperations(mem *MemoryEngine) {
	operations := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Simple calculations
		{"math_op_add_1_2", "1 + 2 = 3", []string{"math", "operation", "addition", "core", "bootstrap"}},
		{"math_op_add_2_3", "2 + 3 = 5", []string{"math", "operation", "addition", "core", "bootstrap"}},
		{"math_op_sub_3_1", "3 - 1 = 2", []string{"math", "operation", "subtraction", "core", "bootstrap"}},
		{"math_op_mul_2_2", "2 * 2 = 4", []string{"math", "operation", "multiplication", "core", "bootstrap"}},
		{"math_op_div_4_2", "4 / 2 = 2", []string{"math", "operation", "division", "core", "bootstrap"}},

		// Properties
		{"math_prop_comm_add", "a + b = b + a", []string{"math", "property", "commutative", "addition", "core", "bootstrap"}},
		{"math_prop_comm_mul", "a * b = b * a", []string{"math", "property", "commutative", "multiplication", "core", "bootstrap"}},
		{"math_prop_zero_add", "a + 0 = a", []string{"math", "property", "identity", "addition", "core", "bootstrap"}},
		{"math_prop_one_mul", "a * 1 = a", []string{"math", "property", "identity", "multiplication", "core", "bootstrap"}},

		// Comparison/equality
		{"math_eq_5", "2 + 3 = 5", []string{"math", "comparison", "equality", "core", "bootstrap"}},
		{"math_eq_6", "3 * 2 = 6", []string{"math", "comparison", "equality", "core", "bootstrap"}},
	}

	for _, op := range operations {
		q := QBit{
			ID:        op.ID,
			Content:   op.Content,
			Tags:      op.Tags,
			Phase:     0.84,
			Weight:    0.95,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("➗ [Bootstrap] Math operations and arithmetic rules loaded.")
}
