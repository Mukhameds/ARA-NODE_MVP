// core/bootstrap_language.go
package core

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// BootstrapCoreKnowledge инициализирует базовый алфавит, цифры и грамматические правила
func BootstrapCoreKnowledge(mem *MemoryEngine) {
	// === Алфавит (английский) ===
	letters := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z",
	}

	for _, letter := range letters {
		q := QBit{
			ID:        "alpha_" + strings.ToLower(letter),
			Content:   letter,
			Tags:      []string{"symbol", "letter", "alphabet", "english", "core", "bootstrap"},
			Phase:     0.75,
			Weight:    0.85,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	// === Цифры 0–9 ===
	for i := 0; i <= 9; i++ {
		s := strconv.Itoa(i)
		q := QBit{
			ID:        "digit_" + s,
			Content:   s,
			Tags:      []string{"symbol", "digit", "number", "core", "bootstrap"},
			Phase:     0.8,
			Weight:    0.9,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	// === Базовые грамматические правила ===
	grammarRules := []string{
		"subject → verb → object",
		"noun + verb + noun",
		"adjective + noun",
		"if + condition → then + action",
		"question → answer",
	}

	for i, rule := range grammarRules {
		q := QBit{
			ID:        fmt.Sprintf("grammar_rule_%d", i),
			Content:   rule,
			Tags:      []string{"grammar", "rule", "structure", "core", "bootstrap"},
			Phase:     0.9,
			Weight:    1.0,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("📘 [Bootstrap] Core language knowledge loaded: alphabet, digits, grammar rules.")
}
