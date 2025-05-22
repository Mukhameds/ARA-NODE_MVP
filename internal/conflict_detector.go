// internal/conflict_detector.go
package internal

import (
	"fmt"
	

	"ara-node/core"
)

type ConflictDetector struct {
	Memory *core.MemoryEngine
}

func NewConflictDetector(mem *core.MemoryEngine) *ConflictDetector {
	return &ConflictDetector{
		Memory: mem,
	}
}

// CheckConflict анализирует QBit и возвращает true, если найдены признаки противоречия
func (cd *ConflictDetector) CheckConflict(q core.QBit) bool {
	if core.Contains(q.Tags, "contradiction") || core.Contains(q.Tags, "denial") {
		fmt.Printf("[ConflictDetector] ⚠️ Already marked contradictory: %s\n", q.Content)
		return true
	}

	conflicts := cd.Memory.FindAll(func(other core.QBit) bool {
		if other.ID == q.ID {
			return false
		}
		// если фазы прямо противоположны (антифаза) и контент пересекается
		if core.PhaseClose(q.Phase, 1.0-other.Phase, 0.1) && contentOverlap(q.Content, other.Content) {
			return true
		}
		return false
	})

	if len(conflicts) > 0 {
		q.Tags = core.AddUniqueTag(q.Tags, "contradiction")
		q.Weight *= 0.5
		cd.Memory.StoreQBit(q)
		fmt.Printf("[ConflictDetector] ❗ Conflict detected in: %s → marked as contradiction\n", q.Content)
		return true
	}

	return false
}

// contentOverlap — примитивное совпадение по словам (можно заменить на NLP позже)
func contentOverlap(a, b string) bool {
	count := 0
	for _, word := range core.Tokenize(a) {
		if core.Contains(core.Tokenize(b), word) {
			count++
		}
	}
	return count >= 2
}
