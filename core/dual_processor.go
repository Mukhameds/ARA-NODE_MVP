// core/dual_processor.go
package core

import (
	"fmt"
)

type DualProcessor struct {
	Cortex   []*SignalEngine
	Ghost GhostLike

	Memory   *MemoryEngine
}

func NewDualProcessor(mem *MemoryEngine, ghost GhostLike) *DualProcessor {
	return &DualProcessor{
		Cortex: []*SignalEngine{
			NewSignalEngine(mem, ghost), // LeftCortex
			NewSignalEngine(mem, ghost), // RightCortex
		},
		Ghost:  ghost,
		Memory: mem,
	}
}

// ProcessDual — запускает суперпозиционное мышление через оба Cortex'а
func (dp *DualProcessor) ProcessDual(sig Signal) {
	if len(dp.Cortex) < 2 {
		fmt.Println("[DualProcessor] ❌ Not enough Cortexes initialized.")
		return
	}

	fmt.Println("[DualProcessor] 🧠 L→R Superposition Start:", sig.Content)

	// Шаг 1: Левый Cortex (реактивный)
	left := dp.Cortex[0]
	leftReact := left.ProcessSignal(sig)
	fmt.Println("[LeftCortex] ➡", leftReact.Response)

	// Шаг 2: Правый Cortex (осмысление реакции левого)
	right := dp.Cortex[1]
	interpretSignal := Signal{
		ID:        "interp_" + sig.ID,
		Content:   leftReact.Response,
		Type:      "internal",
		Tags:      []string{"interpreted", "self"},
		Phase:     sig.Phase * 0.95,
		Weight:    sig.Weight * 0.95,
		Timestamp: sig.Timestamp,
	}

	rightReact := right.ProcessSignal(interpretSignal)
	fmt.Println("[RightCortex] ➡", rightReact.Response)

	// Финальный отклик
	finalQBit := QBit{
		ID:        "qbit_" + sig.ID,
		Content:   rightReact.Response,
		Tags:      []string{"superposed", "final"},
		Phase:     sig.Phase,
		Weight:    sig.Weight,
		CreatedAt: sig.Timestamp,
	}

	dp.Memory.StoreQBit(finalQBit)
	fmt.Println("[DualProcessor] ✅ Stored superposed QBit:", finalQBit.Content)
}
