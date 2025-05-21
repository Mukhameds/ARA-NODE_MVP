package core

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// HeuristicScore — оценка структуры сигнала как потенциального смысла
func HeuristicScore(content string) float64 {
	if content == "" {
		return 0.0
	}

	score := 0.0
	length := utf8.RuneCountInString(content)

	// 1. Симметрия
	runes := []rune(content)
	symmetric := true
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			symmetric = false
			break
		}
	}
	if symmetric {
		score += 0.3
		fmt.Println("[Heuristics] 🔁 Symmetry detected")
	}

	// 2. Простота (короткие сигналы более запоминаемы)
	if length <= 5 {
		score += 0.2
		fmt.Println("[Heuristics] 🧩 Simplicity detected")
	}

	// 3. Завершённость (наличие =, точек, кругов, if/then)
	if strings.Contains(content, "=") || strings.Contains(content, ".") || strings.Contains(content, "→") || strings.Contains(content, "if") {
		score += 0.2
		fmt.Println("[Heuristics] ✅ Completion detected")
	}

	// 4. Повторяемость (двойные слова/символы)
	words := strings.Fields(content)
	seen := map[string]int{}
	for _, w := range words {
		seen[w]++
	}
	for _, v := range seen {
		if v > 1 {
			score += 0.1
			fmt.Println("[Heuristics] 🔁 Repetition detected")
			break
		}
	}

	// 5. Логика (условные структуры, знаки)
	logicKeywords := []string{"and", "or", "if", "then", "not", "cause", "because"}
	for _, kw := range logicKeywords {
		if strings.Contains(content, kw) {
			score += 0.2
			fmt.Println("[Heuristics] 🧠 Logic keyword detected:", kw)
			break
		}
	}

	if score > 1.0 {
		score = 1.0
	}
	return score
}
