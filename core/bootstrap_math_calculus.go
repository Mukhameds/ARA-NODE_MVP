// core/bootstrap_math_calculus.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMathCalculus загружает базовые сигнальные QBits анализа: пределы, производные, интегралы
func BootstrapMathCalculus(mem *MemoryEngine) {
	calculus := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Пределы
		{"calc_limit_def", "a limit describes the value a function approaches as input approaches a point", []string{"math", "calculus", "limit", "core", "bootstrap"}},
		{"calc_limit_symbol", "lim f(x) as x → a", []string{"math", "calculus", "limit", "notation", "core", "bootstrap"}},

		// Производные
		{"calc_derivative_def", "a derivative measures how a function changes at a point", []string{"math", "calculus", "derivative", "definition", "core", "bootstrap"}},
		{"calc_derivative_symbol", "f'(x) or df/dx", []string{"math", "calculus", "derivative", "notation", "core", "bootstrap"}},
		{"calc_velocity", "velocity is the derivative of position with respect to time", []string{"math", "calculus", "application", "core", "bootstrap"}},

		// Непрерывность
		{"calc_continuity", "a function is continuous if its graph has no breaks", []string{"math", "calculus", "continuity", "core", "bootstrap"}},

		// Интегралы
		{"calc_integral_def", "an integral calculates accumulated area or total change", []string{"math", "calculus", "integral", "definition", "core", "bootstrap"}},
		{"calc_integral_symbol", "∫ f(x) dx", []string{"math", "calculus", "integral", "notation", "core", "bootstrap"}},
		{"calc_area", "the integral of a function represents area under the curve", []string{"math", "calculus", "application", "area", "core", "bootstrap"}},

		// Основная теорема
		{"calc_fundamental", "the fundamental theorem links derivative and integral", []string{"math", "calculus", "theorem", "core", "bootstrap"}},
	}

	for _, c := range calculus {
		q := QBit{
			ID:        c.ID,
			Content:   c.Content,
			Tags:      c.Tags,
			Phase:     0.91,
			Weight:    0.97,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("📉 [Bootstrap] Calculus concepts (limits, derivatives, integrals) loaded.")
}
