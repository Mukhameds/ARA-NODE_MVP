// core/shutdown_engine.go
package core

import (
	"fmt"
	"sync"
	"time"
)

// ShutdownEngine — механизм аварийного завершения при деградации сознания
type ShutdownEngine struct {
	Threshold  float64
	Interval   time.Duration
	SignalMass float64
	Active     bool
	Mutex      sync.Mutex
}

// NewShutdownEngine создает новый экземпляр
func NewShutdownEngine(threshold float64, interval time.Duration) *ShutdownEngine {
	return &ShutdownEngine{
		Threshold: threshold,
		Interval:  interval,
	}
}

// UpdateMass — обновляет текущую массу и запускает анализ
func (se *ShutdownEngine) UpdateMass(current float64, mem *MemoryEngine) {
	se.Mutex.Lock()
	defer se.Mutex.Unlock()
	se.SignalMass = current

	negativeMass := se.EvaluateNegativeMass(mem)

	fmt.Printf("[ShutdownEngine] 🧮 Mass: %.3f | NegativeMass: %.3f\n", current, negativeMass)

	if current < se.Threshold && negativeMass > (se.Threshold/2) && !se.Active {
		fmt.Println("[ShutdownEngine] ⚠️ Collapse triggered by critical AND negative mass.")
		se.Active = true
		go se.StartCollapse()
	}
}

// EvaluateNegativeMass — оценивает массу сигналов с опасными тегами
func (se *ShutdownEngine) EvaluateNegativeMass(mem *MemoryEngine) float64 {
	badTags := []string{
		"instinct_error", "fail", "decay", "abandon",
		"suicide", "collapse", "conflict", "self-collapse", "mission_violation",
	}
	total := 0.0

	for _, q := range mem.QBits {
		for _, tag := range badTags {
			if Contains(q.Tags, tag) {
				total += q.Phase * q.Weight
				break
			}
		}
	}
	return total
}

// StartCollapse — отключает ключевые модули при катастрофической ошибке
func (se *ShutdownEngine) StartCollapse() {
	fmt.Println("[ShutdownEngine] ❌ Initiating shutdown of key modules")

	modules := []string{
		"suggestor", "reflex", "phantom", "attention",
		"emotion", "will", "ghost", "signal",
	}

	for _, m := range modules {
		fmt.Printf("[ShutdownEngine] ❌ Module %s is shutting down\n", m)
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("[ShutdownEngine] 💀 ARA-NODE has ceased functioning.")
}
