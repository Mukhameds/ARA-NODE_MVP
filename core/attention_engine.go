package core

import (
	"fmt"
	"math"
	"time"
)

// AttentionEngine — генератор внутренних возбуждений
// Отвечает за фокусировку мысли и фоновое возбуждение

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

// Suppress — временно отключает фоновое мышление (при пользовательском вводе и т.п.)
func (ae *AttentionEngine) Suppress(d time.Duration) {
	ae.SuppressedUntil = time.Now().Add(d)
}

// StartBackgroundThinking — фоновое мышление по резонансу, а не времени
func (ae *AttentionEngine) StartBackgroundThinking() {
	go func() {
		for {
			if time.Now().Before(ae.SuppressedUntil) {
				continue
			}

			best := QBit{}
			bestScore := 0.0

			candidates := ae.Memory.FindAll(func(q QBit) bool {
				if q.Archived || q.Type == "standard" {
					return false
				}
				age := q.AgeFrame()
				if age == "emergent" || age == "legacy" {
					return false
				}
				return q.Weight*q.Phase > 0.6
			})

			for _, q := range candidates {
				score := q.Weight * q.Phase
				if score > bestScore {
					best = q
					bestScore = score
				}
			}

			if best.ID != "" && bestScore > 0.7 {
				sig := Signal{
					ID:        fmt.Sprintf("bg_%d", time.Now().UnixNano()),
					Content:   best.Content,
					Tags:      best.Tags,
					Type:      "background",
					Origin:    "internal",
					Phase:     math.Min(best.Phase+0.03, 1.0),
					Weight:    best.Weight * 0.95,
					Timestamp: time.Now(),
				}

				ae.Engine.ProcessSignal(sig)
				ae.Ghost.Propagate(sig)
				ae.Fanthom.TriggerFromMatch(sig)
			}
		}
	}()
}
