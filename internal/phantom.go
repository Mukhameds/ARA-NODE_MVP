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

func NewFanthomEngine(mem *core.MemoryEngine) *FanthomEngine {
	return &FanthomEngine{Memory: mem}
}

func (fe *FanthomEngine) TriggerFromMatch(sig core.Signal) {
	qbits := fe.Memory.FindByPhase(sig.Phase, 0.05)
	if len(qbits) < 2 {
		return
	}
	fe.GeneratePhantomChain(qbits)
}

func (fe *FanthomEngine) GeneratePhantomChain(chain []core.QBit) {
	var summary string
	var sources []string
	var signalMass float64
	seen := map[string]bool{}
	phantomCount := 0
	allPhantom := true

	for _, q := range chain {
		if seen[q.ID] {
			return // цикл
		}
		seen[q.ID] = true

		inf := 1.0
		if q.Type == "standard" {
			inf += 1.5
		}
		if contains(q.Tags, "instinct") {
			inf += 1.2
		}
		if contains(q.Tags, "emotion") {
			inf += 1.1
		}
		if !strings.HasPrefix(q.Content, "[phantom]") {
			allPhantom = false
		} else {
			phantomCount++
		}
		signalMass += q.Phase * q.Weight * inf
		summary += q.Content + " + "
		sources = append(sources, q.ID)
	}

	summary = strings.TrimSuffix(summary, " + ")
	if allPhantom || signalMass < 3.0 {
		return // отклонить слабый или фантомный только фантом
	}

	for _, existing := range fe.Memory.FindByTag("phantom") {
		if existing.Content == "[phantom] "+summary {
			return // уже существует
		}
	}

	fmt.Println("[FanthomChain] 🧩 Related QBits:")
	for _, q := range chain {
		fmt.Printf("• %s | %.2f | %s\n", q.ID, q.Phase, q.Content)
	}
	fmt.Println("[FanthomChain] → Hypothesis: something meaningful links these signals.")

	// Создаём фантом
	newQ := fe.Memory.CreateQBit("[phantom] " + summary)
	newQ.Tags = []string{"phantom"}
	newQ.Type = "phantom"
	newQ.Phase = chain[0].Phase
	newQ.Weight = signalMass / float64(len(chain))
	fe.Memory.StoreQBit(*newQ)

	fe.Memory.PhantomTree = append(fe.Memory.PhantomTree, core.PhantomLog{
		PhantomID: newQ.ID,
		SourceIDs: sources,
	})

	fmt.Println("[FanthomEngine] 🔮 Phantom QBit:", newQ.ID)
	fmt.Println("[FanthomEngine] ↪ Sources:", strings.Join(sources, ","))
}

func contains(tags []string, key string) bool {
	for _, t := range tags {
		if t == key {
			return true
		}
	}
	return false
}
