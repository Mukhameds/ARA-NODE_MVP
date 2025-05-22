// core/bootstrap_math_symbols.go
package core

import (
	"fmt"
	"strconv"
	"time"
)

// BootstrapMathSymbols загружает математические символы: цифры, операторы, логические и группирующие знаки
func BootstrapMathSymbols(mem *MemoryEngine) {
	// === Цифры 0–9 ===
	for i := 0; i <= 9; i++ {
		s := strconv.Itoa(i)
		q := QBit{
			ID:        "math_digit_" + s,
			Content:   s,
			Tags:      []string{"symbol", "digit", "math", "core", "bootstrap"},
			Phase:     0.8,
			Weight:    0.9,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	// === Арифметические операторы
	operators := []struct {
		Symbol string
		ID     string
		Desc   string
	}{
		{"+", "plus", "addition operator"},
		{"-", "minus", "subtraction operator"},
		{"*", "multiply", "multiplication operator"},
		{"/", "divide", "division operator"},
		{"=", "equals", "equality operator"},
	}

	for _, op := range operators {
		q := QBit{
			ID:        "math_op_" + op.ID,
			Content:   op.Symbol,
			Tags:      []string{"symbol", "operator", "math", "core", "bootstrap"},
			Phase:     0.85,
			Weight:    0.95,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)

		desc := QBit{
			ID:        "math_op_" + op.ID + "_desc",
			Content:   op.Symbol + " is " + op.Desc,
			Tags:      []string{"description", "operator", "math", "core", "bootstrap"},
			Phase:     0.82,
			Weight:    0.88,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(desc)
	}

	// === Логические операторы
	logic := []struct {
		Symbol string
		ID     string
		Desc   string
	}{
		{">", "greater", "greater than"},
		{"<", "less", "less than"},
		{"≠", "not_equal", "not equal to"},
	}

	for _, lg := range logic {
		q := QBit{
			ID:        "math_rel_" + lg.ID,
			Content:   lg.Symbol,
			Tags:      []string{"symbol", "relation", "math", "core", "bootstrap"},
			Phase:     0.82,
			Weight:    0.92,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)

		desc := QBit{
			ID:        "math_rel_" + lg.ID + "_desc",
			Content:   lg.Symbol + " means " + lg.Desc,
			Tags:      []string{"description", "relation", "math", "core", "bootstrap"},
			Phase:     0.8,
			Weight:    0.88,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(desc)
	}

	// === Скобки и группирующие символы
	groupers := []struct {
		Symbol string
		ID     string
		Desc   string
	}{
		{"(", "left_paren", "opening bracket"},
		{")", "right_paren", "closing bracket"},
	}

	for _, g := range groupers {
		q := QBit{
			ID:        "math_group_" + g.ID,
			Content:   g.Symbol,
			Tags:      []string{"symbol", "grouping", "math", "core", "bootstrap"},
			Phase:     0.75,
			Weight:    0.85,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)

		desc := QBit{
			ID:        "math_group_" + g.ID + "_desc",
			Content:   g.Symbol + " is " + g.Desc,
			Tags:      []string{"description", "grouping", "math", "core", "bootstrap"},
			Phase:     0.72,
			Weight:    0.8,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(desc)
	}

	fmt.Println("🔢 [Bootstrap] Math symbols, digits, and operators loaded.")
}
