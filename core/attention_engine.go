// core/attention_engine.go
package core

import (
	"fmt"
	"time"
)

// AttentionEngine — генератор внутренних возбуждений
// Отвечает за фокусировку мысли и фоновое возбуждение

type AttentionEngine struct {
	Memory          *MemoryEngine
	
	Fanthom         FanthomInterface
	Engine          *SignalEngine
	SuppressedUntil time.Time
	Ghost 			GhostLike

}

func NewAttentionEngine(mem *MemoryEngine, ghost GhostLike, fant FanthomInterface, eng *SignalEngine) *AttentionEngine {

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
	fmt.Println("[Attention] ⏸️ Suppressed for", d)
}

// StartBackgroundThinking — фоновое мышление по резонансу
func (ae *AttentionEngine) StartBackgroundThinking() {
	go func() {
		for {
			time.Sleep(1 * time.Second)

			// Отбор кандидатов с высокой фазой и весом
			candidates := ae.Memory.FindAll(func(q QBit) bool {
				if Contains(q.Tags, "phantom") || Contains(q.Tags, "contradiction") {
					return false
				}
				return q.Weight*q.Phase > 0.5
			})

			if len(candidates) == 0 {
				continue
			}

			var top QBit
			topScore := -1.0

			for _, q := range candidates {
				score := q.Weight * q.Phase

				if Contains(q.Tags, "self-related") {
					score += 0.2
				}
				if Contains(q.Tags, "conscious") {
					score += 0.1
				}
				if Contains(q.Tags, "standard") {
					score -= 0.3
				}

				if score > topScore {
					top = q
					topScore = score
				}
			}

			if top.Content == "" {
				continue
			}

			sig := Signal{
				ID:        fmt.Sprintf("echo_%d", time.Now().UnixNano()),
				Content:   top.Content,
				Tags:      append(top.Tags, "echo", "background"),
				Type:      "echo",
				Origin:    "echo_loop",
				Phase:     top.Phase * 0.95,
				Weight:    top.Weight * 0.9,
				Timestamp: time.Now(),
			}

			fmt.Printf("[EchoMode] 🌀 Internal thought: %s\n", sig.Content)

			ae.Engine.ProcessSignal(sig)
			ae.Ghost.Propagate(sig)
			ae.Fanthom.TriggerFromMatch(sig)
		}
	}()
}
