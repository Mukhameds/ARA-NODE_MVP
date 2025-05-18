package core

import (
	"strings"
	"time"
)

type InstinctEngine struct {
	LastInputTime time.Time
}

func NewInstinctEngine() *InstinctEngine {
	return &InstinctEngine{
		LastInputTime: time.Now(),
	}
}

func (ie *InstinctEngine) Tick(currentTime time.Time, signal string) (instincts []string) {
	var results []string

	// 1. Silence trigger (instinct_think)
	if currentTime.Sub(ie.LastInputTime) > 10*time.Second {
		results = append(results, "instinct_think")
		ie.LastInputTime = currentTime
	}

	// 2. Repeat detection (instinct_repeat)
	if isRepeat(signal) {
		results = append(results, "instinct_repeat")
	}

	// 3. Error pattern (instinct_error)
	if strings.Contains(strings.ToLower(signal), "error") {
		results = append(results, "instinct_error")
	}

	// 4. Empty (manual trigger)
	if strings.TrimSpace(signal) == "" {
		results = append(results, "instinct_empty")
	}

	return results
}

// isRepeat — проверка на повтор сигнала (заглушка)
func isRepeat(signal string) bool {
	// TODO: в будущем реализовать анализ дубликатов
	return false
}
