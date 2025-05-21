package internal

import (
	"fmt"
	"strings"

	 "time" 
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
		registered:   []EmotionTrigger{},
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

	// === Heuristic Backpropagation ===
	score := core.HeuristicScore(sig.Content)
	if score > 0.6 {
		qbits := e.Memory.FindByTag("user")
		e.BackPropagate(qbits, "satisfaction")
		fmt.Println("[Emotion] 🌀 Heuristic resonance → BackPropagate (satisfaction)")

		if len(qbits) >= 3 {
			var ids []string
			var tags []string
			for _, q := range qbits {
				ids = append(ids, q.ID)
				tags = append(tags, q.Content)
			}
			core.SynthesizeStandardFromQBits("std_"+qbits[0].ID, tags, 0.8, "satisfaction", ids)
		}
	}

	e.DecayEmotionStates()
}

func (e *EmotionEngine) BackPropagate(sourceQBits []core.QBit, emotion string) {
	for _, q := range sourceQBits {
		q.Phase += 0.1
		q.Weight += 0.2
		if !core.Contains(q.Tags, "emotionally_bound") {
			q.Tags = append(q.Tags, "emotionally_bound")
		}
		q.LastAccessed = time.Now()

		e.Memory.StoreQBit(q)
		fmt.Printf("[EmotionBackProp] ↑ Phase/Weight for %s via %s\n", q.Content, emotion)
	}
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
	for name, val := range e.emotionState {
		active = append(active, fmt.Sprintf("%s (%.2f)", name, val))
	}
	return active
}

func (e *EmotionEngine) PrintEmotions() {
	fmt.Println("🧠 Active Emotions:")
	for name, val := range e.emotionState {
		bar := strings.Repeat("█", int(val*10))
		fmt.Printf("• %-12s %5.2f  %s\n", name, val, bar)
	}
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
