package core

import (
	"fmt"
	"time"
)

type WillBlock struct {
	ID        string
	Sources   []string // Источники: emotion:joy, instinct:order, pattern:symmetric
	Phase     float64
	Weight    float64
	LastUsed  time.Time
	Tags      []string
}

type WillEngine struct {
	Memory  *MemoryEngine
	Engine  *SignalEngine
	Ghost   *GhostField
	Fantom  FanthomInterface
	Blocks  []WillBlock
}

func NewWillEngine(mem *MemoryEngine, se *SignalEngine, gf *GhostField, pe FanthomInterface) *WillEngine {
	return &WillEngine{
		Memory:  mem,
		Engine:  se,
		Ghost:   gf,
		Fantom:  pe,
		Blocks:  []WillBlock{},
	}
}

func (we *WillEngine) Evaluate(q QBit) bool {
	score := 0.0
	for _, wb := range we.Blocks {
		if matchesWill(q, wb.Sources) {
			score += wb.Weight * wb.Phase
		}
	}
	return score >= 0.6
}

func (we *WillEngine) GenerateWillBlock(id string, sources []string, tags []string) {
	for _, wb := range we.Blocks {
		if equalSources(wb.Sources, sources) {
			fmt.Println("[WillEngine] 🔁 WillBlock already exists:", id)
			return
		}
	}
	wb := WillBlock{
		ID:       id,
		Sources:  sources,
		Phase:    0.7,
		Weight:   1.0,
		Tags:     tags,
		LastUsed: time.Now(),
	}
	we.Blocks = append(we.Blocks, wb)
	fmt.Println("[WillEngine] 💡 New WillBlock formed:", id)
}

func (we *WillEngine) Decay() {
	now := time.Now()
	for i := range we.Blocks {
		age := now.Sub(we.Blocks[i].LastUsed).Seconds()
		if age > 300 {
			we.Blocks[i].Weight *= 0.95
			we.Blocks[i].Phase *= 0.97
			if we.Blocks[i].Weight < 0.1 {
				fmt.Println("[WillEngine] ⚠️ WillBlock faded:", we.Blocks[i].ID)
				we.Blocks = append(we.Blocks[:i], we.Blocks[i+1:]...)
				break
			}
		}
	}
}

// DesireLoop — фоновый цикл воли, активирует мысли
func (we *WillEngine) DesireLoop() {
	go func() {
		for {
			qbits := we.Memory.FindTopRelevant("user", 0.6)
			for _, q := range qbits {
				if q.Archived {
					continue
				}

				reasons := []string{}
				accepted := we.Evaluate(q)

				ok, std, stdReason := ShouldTriggerStandard(q.Content, nil)
				if ok && std != nil {
					TriggerStandard(std.ID, we.Engine, we.Ghost, we.Fantom)
					fmt.Println("[WillEngine] 🎯 Triggered Standard:", std.ID, "→", stdReason)
				}

				if accepted {
					fmt.Println("[WillEngine] ✅ Accepted:", q.ID)

					sig := Signal{
						ID:        "will_" + q.ID,
						Content:   q.Content,
						Tags:      append(q.Tags, "will"),
						Phase:     q.Phase,
						Weight:    q.Weight,
						Origin:    "will",
						Type:      "will",
						Timestamp: time.Now(),
					}

					we.Engine.ProcessSignal(sig)
					we.Ghost.Propagate(sig)
					we.Fantom.TriggerFromMatch(sig)
				} else {
					reasons = append(reasons, "adaptive will rejected")
					fmt.Printf("[WillEngine] ❌ Rejected: %s (%v)\n", q.ID, reasons)
					q.Weight *= 0.9
					if q.Weight < 0.4 {
						q.Archived = true
					}
					we.Memory.UpdateQBit(q)
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func matchesWill(q QBit, sources []string) bool {
	for _, src := range sources {
		if Contains(q.Tags, src) {
			return true
		}
	}
	return false
}

func equalSources(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	match := 0
	for _, x := range a {
		for _, y := range b {
			if x == y {
				match++
				break
			}
		}
	}
	return match >= len(a)*80/100
}
