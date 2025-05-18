package internal

import (
	"fmt"
	"strings"
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
	var summary string
	var sources []string

	for _, q := range chain {
		fmt.Printf("• %s | %.2f | %s\n", q.ID, q.Phase, q.Content)
		summary += q.Content + " + "
		sources = append(sources, q.ID)
	}

	summary = strings.TrimSuffix(summary, " + ")
	fmt.Println("[FanthomChain] → Hypothesis: something meaningful links these signals.")

	// Создать и сохранить QBit-фантом
	newQ := fe.Memory.CreateQBit("[phantom] " + summary)
	newQ.Tags = []string{"phantom"}
	newQ.Type = "phantom"
	newQ.Phase = chain[0].Phase
	newQ.Weight = 0.8
	fe.Memory.StoreQBit(*newQ)

	// Лог
	fmt.Println("[FanthomEngine] 🔮 Phantom QBit:", newQ.ID)
	fmt.Println("[FanthomEngine] ↪ Sources:", strings.Join(sources, ","))
}
