package core

import (
	"fmt"
	
	"time"
)

// WillEngine — движок воли агента
// Отслеживает сигналы пользователя и проверяет соответствие миссии
// Генерирует волевые сигналы при совпадении с эталонами

type WillEngine struct {
	Memory *MemoryEngine
	Engine *SignalEngine
	Ghost  *GhostField
	Fantom FanthomInterface
}

func NewWillEngine(mem *MemoryEngine, eng *SignalEngine, gf *GhostField, f FanthomInterface) *WillEngine {
	return &WillEngine{
		Memory: mem,
		Engine: eng,
		Ghost:  gf,
		Fantom: f,
	}
}

func (we *WillEngine) DesireLoop() {
	go func() {
		delay := 5 * time.Second
		for {
			candidates := we.Memory.FindByTag("user")
			for _, q := range candidates {
				if q.Phase < 0.85 {
					continue
				}
				if q.AgeFrame() == "emergent" || q.AgeFrame() == "legacy" {
					continue
				}
				if isAlignedWithStandards(q.Content) {
					fmt.Println("[WillEngine] ✅ Accepted:", q.ID)

					sig := Signal{
						ID:        fmt.Sprintf("will_%d", time.Now().UnixNano()),
						Content:   q.Content,
						Tags:      append(q.Tags, "will", "intent"),
						Type:      "will",
						Origin:    "internal",
						Phase:     q.Phase,
						Weight:    q.Weight,
						Timestamp: time.Now(),
					}

					we.Engine.ProcessSignal(sig)
					we.Ghost.Propagate(sig)
					we.Fantom.TriggerFromMatch(sig)
				} else {
					fmt.Println("[WillEngine] ❌ Rejected:", q.ID)
					q.Weight *= 0.9
					if q.Weight < 0.4 {
						q.Archived = true
					}
					we.Memory.UpdateQBit(q)
				}
			}
			time.Sleep(delay)
		}
	}()
}

func isAlignedWithStandards(content string) bool {
	id, _, score := MatchWithStandards(content)
	return id != "" && score >= 3
}
