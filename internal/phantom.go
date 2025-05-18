package internal

import (
	"fmt"
	"strings"
	"time"

	"ara-node/core"
)

// PhantomEngine — генератор фантомов
type PhantomEngine struct {
	Memory    *core.MemoryEngine
	Instincts *core.InstinctEngine
	Emotions  *core.EmotionEngine
}

func NewPhantomEngine(mem *core.MemoryEngine, inst *core.InstinctEngine, emo *core.EmotionEngine) *PhantomEngine {
	return &PhantomEngine{
		Memory:    mem,
		Instincts: inst,
		Emotions:  emo,
	}
}

func (pe *PhantomEngine) TriggerFromMatch(sig core.Signal) {
	if sig.Weight < 0.5 {
		fmt.Println("[PhantomEngine] ❌ Signal weight too low, skip phantom generation")
		return
	}
	qbits := pe.Memory.FindByPhase(sig.Phase, 0.05)
	if len(qbits) < 2 {
		return
	}
	pe.GeneratePhantomChain(qbits)
}

func (pe *PhantomEngine) GeneratePhantomChain(chain []core.QBit) {
	var summary string
	var sources []string
	var signalMass float64
	seen := map[string]bool{}
	allPhantom := true

	for _, q := range chain {
		if seen[q.ID] {
			fmt.Println("[PhantomEngine] ❌ Cycle detected, abort phantom generation")
			return
		}
		seen[q.ID] = true

		// Исключаем фантомные QBit
		if q.Type == "phantom" || strings.HasPrefix(q.Content, "[phantom]") {
			continue
		}

		allPhantom = false

		inf := 1.0
		if q.Type == "standard" {
			inf += 1.5
		}
		if core.Contains(q.Tags, "instinct") {
			inf += 1.2
		}
		if core.Contains(q.Tags, "emotion") {
			inf += 1.1
		}

		signalMass += q.Phase * q.Weight * inf
		summary += q.Content + " + "
		sources = append(sources, q.ID)
	}

	summary = strings.TrimSuffix(summary, " + ")

	if allPhantom {
		fmt.Println("[PhantomEngine] ❌ All QBits are phantom, abort generation")
		return
	}
	if signalMass < 5.0 {
		fmt.Println("[PhantomEngine] ❌ Signal mass too low:", signalMass)
		return
	}
	if len(summary) > 256 {
		fmt.Println("[PhantomEngine] ❌ Summary too long, abort generation")
		return
	}

	for _, existing := range pe.Memory.FindByTag("phantom") {
		if existing.Content == "[phantom] "+summary {
			fmt.Println("[PhantomEngine] ❌ Duplicate phantom exists, skip")
			return
		}
	}

	if !pe.CheckInstinctEmotionAlignment(signalMass, summary) {
		fmt.Println("[PhantomEngine] ❌ Phantom rejected by instinct/emotion filter")
		return
	}

	fmt.Println("[PhantomChain] 🧩 Related QBits:")
	for _, q := range chain {
		fmt.Printf("• %s | %.2f | %s\n", q.ID, q.Phase, q.Content)
	}
	fmt.Println("[PhantomChain] → Hypothesis: something meaningful links these signals.")

	// Создание фантома
	newQ := pe.Memory.CreateQBit("[phantom] " + summary)
	newQ.Tags = []string{"phantom"}
	newQ.Type = "phantom"
	newQ.Phase = chain[0].Phase
	newQ.Weight = signalMass / float64(len(chain))
	pe.Memory.StoreQBit(*newQ)

	go pe.DecayPhantom(newQ.ID, newQ.Weight)

	pe.Memory.PhantomTree = append(pe.Memory.PhantomTree, core.PhantomLog{
		PhantomID: newQ.ID,
		SourceIDs: sources,
	})

	fmt.Println("[PhantomEngine] 🔮 Phantom QBit:", newQ.ID)
	fmt.Println("[PhantomEngine] ↪ Sources:", strings.Join(sources, ","))
}

func (pe *PhantomEngine) CheckInstinctEmotionAlignment(signalMass float64, summary string) bool {
	instincts := pe.Instincts.Tick(time.Now(), summary)
	emotions := pe.Emotions.CurrentEmotions()

	allowedInstincts := []string{"instinct_think", "instinct_repeat"}
	blockedEmotions := []string{"fear", "anger", "disgust"}

	allow := false

	for _, inst := range instincts {
		for _, ai := range allowedInstincts {
			if inst == ai {
				allow = true
				break
			}
		}
		if allow {
			break
		}
	}

	for _, emo := range emotions {
		for _, be := range blockedEmotions {
			if emo == be {
				allow = false
				break
			}
		}
		if !allow {
			break
		}
	}

	if signalMass < 5.0 {
		allow = false
	}

	return allow
}

func (pe *PhantomEngine) DecayPhantom(id string, weight float64) {
	if weight < 0.1 {
		pe.Memory.DeleteQBit(id)
		fmt.Println("[PhantomEngine] ⬇️ Phantom deleted due to low mass:", id)
	}
}
