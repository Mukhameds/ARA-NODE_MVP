package core

import (
	"fmt"
	"time"
	"math"
)

// AttentionEngine — генератор внутренних возбуждений
type AttentionEngine struct {
	Memory   *MemoryEngine
	Ghost    *GhostField
	Fanthom  FanthomInterface
	Engine   *SignalEngine
	SuppressedUntil time.Time
}

func NewAttentionEngine(mem *MemoryEngine, ghost *GhostField, fant FanthomInterface, eng *SignalEngine) *AttentionEngine {
	return &AttentionEngine{
		Memory:  mem,
		Ghost:   ghost,
		Fanthom: fant,
		Engine:  eng,
	}
}

// Suppress временно приостанавливает внутреннее мышление (например, при вводе пользователя)
func (ae *AttentionEngine) Suppress(d time.Duration) {
	ae.SuppressedUntil = time.Now().Add(d)
}

// StartBackgroundThinking запускает постоянное самовозбуждение
func (ae *AttentionEngine) StartBackgroundThinking() {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			if time.Now().Before(ae.SuppressedUntil) {
				continue
			}

			active := ae.Memory.FindAll(func(q QBit) bool {
				return q.Weight*q.Phase > 0.6 && !q.Archived && q.Type != "standard"
			})

			for _, q := range active {
				sig := Signal{
					ID:        fmt.Sprintf("bg_%d", time.Now().UnixNano()),
					Content:   q.Content,
					Tags:      q.Tags,
					Type:      "background",
					Origin:    "internal",
					Phase:     math.Min(q.Phase+0.05, 1.0),
					Weight:    q.Weight * 0.9,
					Timestamp: time.Now(),
				}

				ae.Engine.ProcessSignal(sig)
				ae.Ghost.Propagate(sig)
				ae.Fanthom.TriggerFromMatch(sig)
			}
		}
	}()
}
