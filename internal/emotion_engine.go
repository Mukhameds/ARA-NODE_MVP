package internal

import (
	"fmt"
	"strings"
	

	"ara-node/core"
)

// EmotionEngine — управляет внутренними эмоциями ARA
// Эмоции усиливают важные фантомы, помогают воле и ориентируют внимание

type EmotionEngine struct {
	Memory         *core.MemoryEngine
	Instincts      *InstinctEngine
	emotionState   map[string]float64
	emotionDecay   float64
	registered     []EmotionTrigger
}

type EmotionTrigger struct {
	Tag   string
	Phase float64
	Name  string
}

func NewEmotionEngine(mem *core.MemoryEngine) *EmotionEngine {
	return &EmotionEngine{
		Memory:       mem,
		emotionState: make(map[string]float64),
		emotionDecay: 0.98,
		registered:  []EmotionTrigger{},
	}
}

func (e *EmotionEngine) AddTrigger(name, tag string, minPhase float64) {
	e.registered = append(e.registered, EmotionTrigger{
		Tag:   tag,
		Phase: minPhase,
		Name:  name,
	})
}

func (e *EmotionEngine) React(sig core.Signal) {
	for _, rule := range e.registered {
		if core.Contains(sig.Tags, rule.Tag) && sig.Phase >= rule.Phase {
			e.emotionState[rule.Name] += 1.0
			fmt.Println("[Emotion] ❤️ +", rule.Name)
		}
	}

	if e.Instincts != nil {
		instinctBoost := e.Instincts.GetInstinctBoost(sig.Tags)
		if instinctBoost > 0.05 {
			e.emotionState["hope"] += instinctBoost * 0.5
			fmt.Println("[Emotion] 🧬 boosted by instinct +hope")
		}
	}

	e.DecayEmotionStates()
}

func (e *EmotionEngine) DecayEmotionStates() {
	for name, val := range e.emotionState {
		e.emotionState[name] = val * e.emotionDecay
		if e.emotionState[name] < 0.05 {
			delete(e.emotionState, name)
		}
	}
}

func (e *EmotionEngine) CurrentEmotions() []string {
	var active []string
	for name := range e.emotionState {
		active = append(active, name)
	}
	return active
}

func (e *EmotionEngine) GetPhaseBoost(tags []string) float64 {
	boost := 0.0
	if containsAny(tags, []string{"standard", "instinct", "mission"}) {
		if e.emotionState["joy"] > 0.5 {
			boost += 0.1
		}
		if e.emotionState["hope"] > 0.3 {
			boost += 0.05
		}
	}
	if containsAny(tags, []string{"fail", "risk", "conflict"}) {
		if e.emotionState["fear"] > 0.5 || e.emotionState["frustration"] > 0.5 {
			boost -= 0.1
		}
	}
	return boost
}

func containsAny(tags []string, keys []string) bool {
	for _, t := range tags {
		for _, k := range keys {
			if strings.Contains(t, k) {
				return true
			}
		}
	}
	return false
}

func DefaultEmotionSet(e *EmotionEngine) {
	e.AddTrigger("joy", "success", 0.6)
	e.AddTrigger("frustration", "fail", 0.5)
	e.AddTrigger("fear", "risk", 0.8)
	e.AddTrigger("anger", "conflict", 0.9)
	e.AddTrigger("hope", "mission", 0.6)
}
