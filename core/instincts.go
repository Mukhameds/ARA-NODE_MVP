package core

import (
	"strings"
	"sync"
	"time"
)

type InstinctEngine struct {
	LastInputTime time.Time
	mu            sync.Mutex
	recentSignals []string
	maxHistory    int
}

func NewInstinctEngine() *InstinctEngine {
	return &InstinctEngine{
		LastInputTime: time.Now(),
		maxHistory:    100,
		recentSignals: make([]string, 0, 100),
	}
}

func (ie *InstinctEngine) Tick(currentTime time.Time, signal string) (instincts []string) {
	ie.mu.Lock()
	defer ie.mu.Unlock()

	var results []string

	// 1. Silence trigger (instinct_think)
	if currentTime.Sub(ie.LastInputTime) > 10*time.Second {
		results = append(results, "instinct_think")
		ie.LastInputTime = currentTime
	}

	// 2. Repeat detection (instinct_repeat)
	if ie.isRepeat(signal) {
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

	// Добавляем сигнал в историю
	ie.addSignal(signal)

	return results
}

func (ie *InstinctEngine) isRepeat(signal string) bool {
	// Проверяем, встречался ли сигнал в истории недавно
	for _, s := range ie.recentSignals {
		if s == signal {
			return true
		}
	}
	return false
}

func (ie *InstinctEngine) addSignal(signal string) {
	if signal == "" {
		return
	}
	// Добавляем в историю, с ограничением длины
	if len(ie.recentSignals) >= ie.maxHistory {
		ie.recentSignals = ie.recentSignals[1:]
	}
	ie.recentSignals = append(ie.recentSignals, signal)
}

// ClearHistory очищает историю сигналов
func (ie *InstinctEngine) ClearHistory() {
	ie.mu.Lock()
	defer ie.mu.Unlock()
	ie.recentSignals = make([]string, 0, ie.maxHistory)
}
