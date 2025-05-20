package internal

import (
	"fmt"
	"time"

	"ara-node/core"
)

// DecayAnalysisEngine — удаляет старые или слабые узлы

type DecayAnalysisEngine struct {
	Memory *core.MemoryEngine
}

func NewDecayAnalysisEngine(mem *core.MemoryEngine) *DecayAnalysisEngine {
	return &DecayAnalysisEngine{Memory: mem}
}

// StartDecayLoop — фоновая проверка каждые 30 секунд
func (d *DecayAnalysisEngine) StartDecayLoop() {
	go func() {
		for {
			d.RunDecayCheck()
			time.Sleep(30 * time.Second)
		}
	}()
}

// RunDecayCheck — удаляет архив или устаревшие фантомы с малым весом
func (d *DecayAnalysisEngine) RunDecayCheck() {
	count := 0
	for id, q := range d.Memory.QBits {
		if q.Archived && q.Weight < 0.05 {
			d.Memory.DeleteQBit(id)
			fmt.Println("[DecayEngine] ❌ Archived deleted:", id)
			count++
			continue
		}
		if q.AgeFrame() == "legacy" && q.Weight < 0.2 {
			if q.Type == "phantom" || q.Type == "suggestion" {
				d.Memory.DeleteQBit(id)
				fmt.Println("[DecayEngine] 🧹 Old phantom removed:", id)
				count++
			}
		}
	}
	if count > 0 {
		fmt.Printf("[DecayEngine] → Total removed: %d\n", count)
	}
}
