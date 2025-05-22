// internal/consciousness_capture.go
package internal

import (
	"fmt"
	"time"

	"ara-node/core"
)

type ConsciousnessCaptureEngine struct {
	Memory     *core.MemoryEngine
	LogEnabled bool
	LastLogged time.Time
}

func NewConsciousnessCaptureEngine(mem *core.MemoryEngine) *ConsciousnessCaptureEngine {
	return &ConsciousnessCaptureEngine{
		Memory:     mem,
		LogEnabled: true,
		LastLogged: time.Now(),
	}
}

func (cce *ConsciousnessCaptureEngine) StartConsciousnessLoop() {
	go func() {
		for {
			cce.ScanConsciousMoments()
			time.Sleep(5 * time.Second)
		}
	}()
}

func (cce *ConsciousnessCaptureEngine) ScanConsciousMoments() {
	qbits := cce.Memory.FindByTag("standard")

	for _, std := range qbits {
		if std.Weight*std.Phase < 0.7 {
			continue
		}

		matches := cce.Memory.FindAll(func(q core.QBit) bool {
			return q.Type == "user" && core.PhaseClose(q.Phase, std.Phase, 0.05)
		})

		if len(matches) == 0 {
			continue
		}

		top := matches[0]
		fmt.Printf("[Consciousness] ⚡ Match to standard (%s) => %s\n", std.ID, top.Content)

		// пометка как осознанный отклик
		top.Tags = core.AddUniqueTag(top.Tags, "conscious")
		std.Tags = core.AddUniqueTag(std.Tags, "evoked")
		top.Weight += 0.1
		std.Weight += 0.05

		cce.Memory.StoreQBit(top)
		cce.Memory.StoreQBit(std)

		if cce.LogEnabled {
			fmt.Printf("[ConsciousnessLog] 🔷 Conscious identity shift: %s\n", top.Content)
		}
	}
}

func (cce *ConsciousnessCaptureEngine) IsConscious(q core.QBit) bool {
	return core.Contains(q.Tags, "conscious") || core.Contains(q.Tags, "self-related")
}
