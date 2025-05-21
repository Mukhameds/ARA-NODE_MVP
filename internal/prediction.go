package internal

import (
	"fmt"
	"time"

	"ara-node/core"
)

// PredictionRule — правило предсказания
type PredictionRule struct {
	IfTag       string
	ThenContent string
	MinPhase    float64
	Confidence  float64
}

// PredictionEngine — реактивный генератор предсказаний
type PredictionEngine struct {
	Memory          *core.MemoryEngine
	Rules           []PredictionRule
	Engine          *core.SignalEngine
	Ghost           *core.GhostField
	Enabled         bool
	LastPredictions map[string]time.Time
	Pause           time.Duration
}

func NewPredictionEngine(mem *core.MemoryEngine, eng *core.SignalEngine, ghost *core.GhostField) *PredictionEngine {
	return &PredictionEngine{
		Memory:          mem,
		Engine:          eng,
		Ghost:           ghost,
		Enabled:         true,
		LastPredictions: make(map[string]time.Time),
		Pause:           5 * time.Second,
		Rules: []PredictionRule{
			{IfTag: "user", ThenContent: "followup", MinPhase: 0.6, Confidence: 0.8},
		},
	}
}

// Tick — проверка и генерация предсказаний
func (pe *PredictionEngine) Tick() {
	if !pe.Enabled {
		return
	}

	now := time.Now()

	for _, rule := range pe.Rules {
		candidates := pe.Memory.FindTopRelevant(rule.IfTag, rule.MinPhase)
		if len(candidates) == 0 {
			continue
		}

		best := candidates[0]
		confidence := rule.Confidence * best.Weight * best.Phase

		if confidence < 0.5 {
			continue
		}

		// Спам-фильтр: не предсказывать одно и то же слишком часто
		lastTime, seen := pe.LastPredictions[rule.ThenContent]
		if seen && now.Sub(lastTime) < pe.Pause {
			continue
		}
		pe.LastPredictions[rule.ThenContent] = now

		sig := core.Signal{
			ID:        fmt.Sprintf("pred_%d", time.Now().UnixNano()),
			Content:   rule.ThenContent,
			Tags:      []string{"predicted", rule.IfTag},
			Type:      "prediction",
			Origin:    "prediction_engine",
			Phase:     best.Phase,
			Weight:    confidence,
			Timestamp: time.Now(),
		}

		fmt.Printf("[PredictionEngine] 🔮 Predict: '%s' (from %s) with confidence %.2f\n", sig.Content, best.ID, confidence)

		pe.Engine.ProcessSignal(sig)
		pe.Ghost.Propagate(sig)
	}
}
