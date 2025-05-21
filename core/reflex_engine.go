package core

import (
	"fmt"
	"time"
)

// ReflexRule — мгновенная реакция на сигнал
type ReflexRule struct {
	MatchTag string
	Action   func(sig Signal)
}

// ReflexEngine — хранит и вызывает рефлексы
type ReflexEngine struct {
	Rules          []ReflexRule
	CooldownPerTag map[string]time.Time
	MinInterval    time.Duration
}

func NewReflexEngine() *ReflexEngine {
	return &ReflexEngine{
		Rules:          []ReflexRule{},
		CooldownPerTag: make(map[string]time.Time),
		MinInterval:    3 * time.Second,
	}
}

func (re *ReflexEngine) AddRule(tag string, action func(sig Signal)) {
	re.Rules = append(re.Rules, ReflexRule{
		MatchTag: tag,
		Action:   action,
	})
}

func (re *ReflexEngine) React(sig Signal) {
	now := time.Now()
	for _, rule := range re.Rules {
		if containsTag(sig.Tags, rule.MatchTag) {
			last, exists := re.CooldownPerTag[rule.MatchTag]
			if exists && now.Sub(last) < re.MinInterval {
				continue // подавлено из-за частоты
			}
			fmt.Println("[Reflex] ⚡ Instant reaction to:", sig.Content)
			re.CooldownPerTag[rule.MatchTag] = now
			rule.Action(sig)
		}
	}
}

func containsTag(tags []string, key string) bool {
	for _, t := range tags {
		if t == key {
			return true
		}
	}
	return false
}

// Пример предустановленных рефлексов
func DefaultReflexSet(re *ReflexEngine) {
	re.AddRule("instinct_error", func(sig Signal) {
		fmt.Println("[Reflex] ❗ System error reflex triggered.")
	})
	re.AddRule("danger", func(sig Signal) {
		fmt.Println("[Reflex] 🚨 Danger signal! Executing safety protocol...")
	})
	re.AddRule("fail", func(sig Signal) {
		fmt.Println("[Reflex] 😤 Fail detected. Reacting emotionally.")
	})
}
