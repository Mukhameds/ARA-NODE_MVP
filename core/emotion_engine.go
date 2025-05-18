package core

import (
	"fmt"
	"sync"
)

// EmotionTrigger — условие и реакция на эмоциональный сигнал
type EmotionTrigger struct {
	Tag     string
	PhaseGT float64
	Action  func(sig Signal)
}

// EmotionEngine — реагирует на эмоциональные возбуждения и хранит текущие эмоции
type EmotionEngine struct {
	Rules   []EmotionTrigger
	current []string
	mu      sync.Mutex
}

func NewEmotionEngine() *EmotionEngine {
	return &EmotionEngine{
		Rules:   []EmotionTrigger{},
		current: []string{"neutral"},
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
		if Contains(sig.Tags, rule.Tag) && sig.Phase > rule.PhaseGT {

			fmt.Println("[Emotion] 💓 Triggered by:", sig.Content)
			rule.Action(sig)
			// Обновляем текущие эмоции при срабатывании триггера
			ee.UpdateEmotion(rule.Tag)
		}
	}
}

// CurrentEmotions возвращает текущие активные эмоции
func (ee *EmotionEngine) CurrentEmotions() []string {
	ee.mu.Lock()
	defer ee.mu.Unlock()
	return append([]string{}, ee.current...) // копия среза для безопасности
}

// UpdateEmotion добавляет новую эмоцию в текущие, предотвращая дубликаты
func (ee *EmotionEngine) UpdateEmotion(emotion string) {
	ee.mu.Lock()
	defer ee.mu.Unlock()

	for _, e := range ee.current {
		if e == emotion {
			return
		}
	}
	ee.current = append(ee.current, emotion)
	fmt.Println("[EmotionEngine] Updated emotions:", ee.current)
}

// ClearEmotions очищает текущие эмоции, оставляя только нейтральное состояние
func (ee *EmotionEngine) ClearEmotions() {
	ee.mu.Lock()
	defer ee.mu.Unlock()
	ee.current = []string{"neutral"}
	fmt.Println("[EmotionEngine] Emotions cleared")
}

// Базовые эмоциональные реакции
func DefaultEmotionSet(ee *EmotionEngine) {
	ee.AddTrigger("joy", 0.7, func(sig Signal) {
		fmt.Println("[Emotion] 😊 Joyful signal received.")
		ee.UpdateEmotion("joy")
	})
	ee.AddTrigger("frustration", 0.6, func(sig Signal) {
		fmt.Println("[Emotion] 😣 Frustration building up.")
		ee.UpdateEmotion("frustration")
	})
	ee.AddTrigger("fear", 0.6, func(sig Signal) {
		fmt.Println("[Emotion] 😨 Fear detected.")
		ee.UpdateEmotion("fear")
	})
	ee.AddTrigger("anger", 0.6, func(sig Signal) {
		fmt.Println("[Emotion] 😠 Anger detected.")
		ee.UpdateEmotion("anger")
	})
}

