package core

import (
	"fmt"
	
)

// EmotionTrigger — условие и реакция на эмоциональный сигнал
type EmotionTrigger struct {
	Tag     string
	PhaseGT float64
	Action  func(sig Signal)
}

// EmotionEngine — реагирует на эмоциональные возбуждения
type EmotionEngine struct {
	Rules []EmotionTrigger
}

func NewEmotionEngine() *EmotionEngine {
	return &EmotionEngine{
		Rules: []EmotionTrigger{},
	}
}

func (ee *EmotionEngine) AddTrigger(tag string, minPhase float64, action func(sig Signal)) {
	ee.Rules = append(ee.Rules, EmotionTrigger{
		Tag:     tag,
		PhaseGT: minPhase,
		Action:  action,
	})
}

func (ee *EmotionEngine) React(sig Signal) {
	for _, rule := range ee.Rules {
		if contains(sig.Tags, rule.Tag) && sig.Phase > rule.PhaseGT {
			fmt.Println("[Emotion] 💓 Triggered by:", sig.Content)
			rule.Action(sig)
		}
	}
}

// Базовые эмоциональные реакции
func DefaultEmotionSet(ee *EmotionEngine) {
	ee.AddTrigger("joy", 0.7, func(sig Signal) {
		fmt.Println("[Emotion] 😊 Joyful signal received.")
	})
	ee.AddTrigger("frustration", 0.6, func(sig Signal) {
		fmt.Println("[Emotion] 😣 Frustration building up.")
	})
}
