package core

import (
	"strings"
	"fmt"
	"time"
)

// StandardBlock — эталонная миссия, принцип или ориентир сознания
type StandardBlock struct {
	ID          string
	Keywords    []string
	Priority    float64
	Dynamic     bool     // был ли создан системой
	EmotionLink string   // ID эмоции или чувства, откуда он возник
	SourceQBits []string // какие QBits его сформировали
}

// 📚 Пустая библиотека эталонов — всё формируется динамически
var StandardLibrary = []StandardBlock{}

// MatchWithStandards — простой режим (оставлен для обратной совместимости)
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

// MatchWithStandardsExtended — полный блок + вес совпадения + причина
func MatchWithStandardsExtended(content string) (*StandardBlock, float64, string) {
	content = strings.ToLower(content)
	var best *StandardBlock
	bestScore := 0.0
	reason := ""

	for _, std := range StandardLibrary {
		matchCount := 0
		for _, keyword := range std.Keywords {
			if strings.Contains(content, keyword) {
				matchCount++
			}
		}
		score := float64(matchCount) * std.Priority
		if score > bestScore {
			bestScore = score
			best = &std
			reason = fmt.Sprintf("Matched %d keywords × priority %.2f", matchCount, std.Priority)
		}
	}

	if bestScore >= 2.0 {
		return best, bestScore, reason
	}
	return nil, 0.0, "No significant match"
}

// TriggerStandard — возбуждает стандарт как задачу (трансляция в поле)
func TriggerStandard(stdID string, se *SignalEngine, gf GhostLike, pe FanthomInterface) {
	std := GetStandardByID(stdID)
	if std == nil {
		fmt.Println("[StandardTrigger] ❌ Not found:", stdID)
		return
	}

	sig := Signal{
		ID:        "std_" + std.ID,
		Content:   strings.Join(std.Keywords, " "),
		Tags:      append([]string{"standard", std.ID}, std.Keywords...),
		Phase:     std.Priority,
		Weight:    std.Priority * 1.0,
		Origin:    "standard_trigger",
		Type:      "mission",
		Timestamp: time.Now(),
	}

	se.ProcessSignal(sig)
	gf.Propagate(sig)
	pe.TriggerFromMatch(sig)

	fmt.Println("[StandardTrigger] 🚩 Broadcasted:", std.ID)
}

// GetStandardByID — возвращает стандарт по ID
func GetStandardByID(id string) *StandardBlock {
	for i, std := range StandardLibrary {
		if std.ID == id {
			return &StandardLibrary[i]
		}
	}
	return nil
}

// ShouldTriggerStandard — решает, стоит ли возбуждать стандарт
func ShouldTriggerStandard(content string, alreadyActive map[string]bool) (bool, *StandardBlock, string) {
	std, score, reason := MatchWithStandardsExtended(content)
	if std == nil || score < 2.0 {
		return false, nil, "Not strong enough match"
	}
	if alreadyActive != nil && alreadyActive[std.ID] {
		return false, std, "Already active"
	}
	return true, std, reason
}

// SynthesizeStandardFromQBits — формирует новый стандарт из QBits + эмоция
func SynthesizeStandardFromQBits(id string, keywords []string, priority float64, emotion string, sourceIDs []string) *StandardBlock {
	std := StandardBlock{
		ID:          id,
		Keywords:    keywords,
		Priority:    priority,
		Dynamic:     true,
		EmotionLink: emotion,
		SourceQBits: sourceIDs,
	}
	StandardLibrary = append(StandardLibrary, std)
	fmt.Println("[StandardSynth] ✨ Created:", std.ID, "from", sourceIDs, "linked to:", emotion)
	return &std
}
