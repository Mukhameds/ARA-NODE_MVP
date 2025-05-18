package core

import (
	"fmt"
	"strings"
	"time"
)

// Intent — внутренняя цель агента
type Intent struct {
	Tag     string
	Phase   float64
	Urgency float64
}

// WillEngine — движок воли агента
type WillEngine struct {
	Memory    *MemoryEngine
	Delay     time.Duration
	Active    bool
	lastTried map[string]time.Time
}

// NewWillEngine — создать движок воли
func NewWillEngine(mem *MemoryEngine) *WillEngine {
	return &WillEngine{
		Memory:    mem,
		Delay:     8 * time.Second,
		Active:    true,
		lastTried: make(map[string]time.Time),
	}
}

// isAlignedWithStandards — сравнение с эталонными блоками
func isAlignedWithStandards(content string) bool {
	for _, std := range StandardLibrary {
		for _, kw := range std.Keywords {
			if strings.Contains(strings.ToLower(content), strings.ToLower(kw)) {
				return true
			}
		}
	}
	return false
}

// DesireLoop — постоянный фон для самозапуска
func (we *WillEngine) DesireLoop() {
	go func() {
		for we.Active {
			time.Sleep(we.Delay)

			now := time.Now()
			intent := Intent{Tag: "user", Phase: 0.85, Urgency: 1.0}
			qbits := we.Memory.FindByTag(intent.Tag)

			for _, q := range qbits {
				if q.Archived || q.Phase < intent.Phase {
					continue
				}

				// задержка между сверками
				if t, seen := we.lastTried[q.ID]; seen && now.Sub(t) < we.Delay {
					continue
				}
				we.lastTried[q.ID] = now

				if isAlignedWithStandards(q.Content) {
					fmt.Println("[WillEngine] ⚡ Intent triggered:", q.ID)
					sig := Signal{
						ID:        fmt.Sprintf("will_%d", time.Now().UnixNano()),
						Content:   "[WILL] " + q.Content,
						Tags:      []string{"phantom"},
						Timestamp: now,
						Phase:     q.Phase,
						Weight:    q.Weight,
						Origin:    "will",
					}
					fmt.Println("←", sig.Content)
					we.Delay = 8 * time.Second
				} else {
					fmt.Println("[WillEngine] ❌ Rejected:", q.ID)
					we.Memory.AdjustWeight(q.ID, -0.2)
					adjusted := we.Memory.QBits[q.ID]
					if adjusted.Weight < 0.1 {
						we.Memory.AddTag(q.ID, "archived")
						fmt.Println("[WillEngine] 🗃 Archived:", q.ID)
					}
					we.Delay *= 2
					if we.Delay > 120*time.Second {
						we.Delay = 120 * time.Second
					}
				}
			}
		}
	}()
}
