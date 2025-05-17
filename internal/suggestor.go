package internal

import (
	"fmt"
	"strings"
	"ara-node/core"
)

// SuggestorEngine — генератор предложений/мыслей
type SuggestorEngine struct {
	Memory *core.MemoryEngine
}

// NewSuggestorEngine — инициализация
func NewSuggestorEngine(mem *core.MemoryEngine) *SuggestorEngine {
	return &SuggestorEngine{
		Memory: mem,
	}
}

// SuggestFromQBits — ищет цепочки и предлагает мысль
func (s *SuggestorEngine) SuggestFromQBits() {
	qbits := s.Memory.FindByTag("user")
	if len(qbits) < 2 {
		return
	}

	// Объединение смыслов
	var ideas []string
	for _, q := range qbits {
		ideas = append(ideas, q.Content)
		if len(ideas) >= 3 {
			break
		}
	}

	// Генерация фразы
	suggestion := s.GenerateSuggestion(ideas)
	fmt.Println("[Suggestor] 💡", suggestion)
}

// GenerateSuggestion — строит предложение на основе смыслов
func (s *SuggestorEngine) GenerateSuggestion(ideas []string) string {
	if len(ideas) == 0 {
		return "No suggestion available."
	}
	return fmt.Sprintf("Would you like to explore the idea: \"%s\" + ...?", strings.Join(ideas, " + "))
}
