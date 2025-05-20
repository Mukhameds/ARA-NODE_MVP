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
	return &SuggestorEngine{Memory: mem}
}

// SuggestFromQBits — ищет цепочки и предлагает мысль
func (s *SuggestorEngine) SuggestFromQBits() {
	relevant := s.FindRecentRelevant(50)
	filtered := []core.QBit{}

	// Фильтрация по зрелости и фазе
	for _, q := range relevant {
		age := q.AgeFrame()
		if age == "emergent" || age == "legacy" {
			continue
		}
		if q.Phase < 0.5 {
			continue
		}
		filtered = append(filtered, q)
	}

	if len(filtered) < 3 {
		return
	}

	groups := groupBySimilarity(filtered)
	for _, group := range groups {
		if len(group) < 3 {
			continue
		}

		idea := mergeSummary(group)
		signalMass := 0.0
		for _, q := range group {
			signalMass += q.Weight * q.Phase
		}

		if signalMass < 2.0 {
			continue
		}

		fmt.Println("[Suggestor] 💡", idea)

		q := s.Memory.CreateQBit("[suggestion] " + idea)
		q.Tags = []string{"suggestion", "phantom", "standard_candidate"}
		q.Type = "phantom"
		q.Phase = group[0].Phase
		q.Weight = signalMass / float64(len(group))
		s.Memory.StoreQBit(*q)
	}
}

// FindRecentRelevant — выбирает последние значимые QBits
func (s *SuggestorEngine) FindRecentRelevant(n int) []core.QBit {
	all := s.Memory.FindAll(func(q core.QBit) bool {
		if q.Archived {
			return false
		}
		tags := q.Tags
		return core.Contains(tags, "user") ||
			core.Contains(tags, "instinct") ||
			core.Contains(tags, "emotion") ||
			core.Contains(tags, "predict")
	})

	if len(all) <= n {
		return all
	}

	return all[len(all)-n:]
}

// groupBySimilarity — группирует по содержательному совпадению
func groupBySimilarity(qbits []core.QBit) [][]core.QBit {
	clusters := [][]core.QBit{}
	for _, q := range qbits {
		found := false
		for i, group := range clusters {
			if isSimilar(q.Content, group[0].Content) {
				clusters[i] = append(clusters[i], q)
				found = true
				break
			}
		}
		if !found {
			clusters = append(clusters, []core.QBit{q})
		}
	}
	return clusters
}

// mergeSummary — объединяет содержимое в одну идею
func mergeSummary(group []core.QBit) string {
	parts := []string{}
	seen := map[string]bool{}
	for _, q := range group {
		t := strings.TrimSpace(q.Content)
		if t == "" || seen[t] {
			continue
		}
		parts = append(parts, t)
		seen[t] = true
		if len(parts) >= 5 {
			break
		}
	}
	return strings.Join(parts, " + ")
}

// isSimilar — грубая проверка похожести по словам
func isSimilar(a, b string) bool {
	wa := strings.Fields(strings.ToLower(a))
	wb := strings.Fields(strings.ToLower(b))
	match := 0
	for _, x := range wa {
		for _, y := range wb {
			if x == y {
				match++
			}
		}
	}
	return match >= 2
}

// GenerateSuggestion — (сохранили старый интерфейс для обратной совместимости)
func (s *SuggestorEngine) GenerateSuggestion(ideas []string) string {
	if len(ideas) == 0 {
		return "No suggestion available."
	}
	return fmt.Sprintf("Would you like to explore the idea: \"%s\" + ...?", strings.Join(ideas, " + "))
}
