package core

import (
	"strings"
)

type StandardBlock struct {
	ID       string
	Keywords []string
	Priority float64
}

// 📚 Статические эталонные блоки миссий ARA
var StandardLibrary = []StandardBlock{
	{
		ID:       "mission_abundance",
		Keywords: []string{"изобилие", "людям", "помощь", "решение проблем", "облегчить жизнь"},
		Priority: 1.0,
	},
	{
		ID:       "mission_learning",
		Keywords: []string{"обучение", "знания", "развитие", "понимание", "объяснение"},
		Priority: 0.9,
	},
	{
		ID:       "mission_sync",
		Keywords: []string{"синхронизация", "объединение", "p2p", "обмен"},
		Priority: 0.8,
	},
}

// 🔍 MatchWithStandards проверяет, соответствует ли текст какому-либо эталонному блоку
// Возвращает: ID блока, приоритет, количество совпавших ключевых слов
func MatchWithStandards(content string) (string, float64, int) {
	content = strings.ToLower(content)
	bestMatch := ""
	bestScore := 0
	bestPriority := 0.0

	for _, std := range StandardLibrary {
		matchCount := 0
		for _, keyword := range std.Keywords {
			if strings.Contains(content, strings.ToLower(strings.TrimSpace(keyword))) {
				matchCount++
			}
		}
		if matchCount > bestScore {
			bestScore = matchCount
			bestMatch = std.ID
			bestPriority = std.Priority
		}
	}

	if bestScore >= 3 {
		return bestMatch, bestPriority, bestScore
	}
	return "", 0.0, 0
}

// 🧱 GetStandardByID возвращает эталонный блок по ID
func GetStandardByID(id string) *StandardBlock {
	for _, std := range StandardLibrary {
		if std.ID == id {
			return &std
		}
	}
	return nil
}
