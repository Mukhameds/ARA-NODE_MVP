package internal

import (
	"fmt"
	"ara-node/core"
)

// FanthomEngine — генератор фантомов
type FanthomEngine struct {
	Memory *core.MemoryEngine
}

// NewFanthomEngine — инициализация
func NewFanthomEngine(mem *core.MemoryEngine) *FanthomEngine {
	return &FanthomEngine{
		Memory: mem,
	}
}

// TriggerFromMatch — ищет совпадения и запускает фантом
func (fe *FanthomEngine) TriggerFromMatch(sig core.Signal) {
	qbits := fe.Memory.FindByPhase(sig.Phase, 0.05)
	if len(qbits) >= 2 {
		fmt.Println("[FanthomEngine] ⚡ Phase-match found. Generating phantom...")
		fe.GeneratePhantomChain(qbits)
	}
}

// GeneratePhantomChain — строит фантом из цепочки QBit
func (fe *FanthomEngine) GeneratePhantomChain(chain []core.QBit) {
	fmt.Println("[FanthomChain] 🧩 Related QBits:")
	for _, q := range chain {
		fmt.Printf("• %s | %.2f | %s\n", q.ID, q.Phase, q.Content)
	}
	fmt.Println("[FanthomChain] → Hypothesis: something meaningful links these signals.")
}
