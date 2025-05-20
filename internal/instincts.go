package internal

import (
	"strings"
	"time"
)

// Instinct — осмысленный врождённый механизм ARA
type Instinct struct {
	ID      string
	Weight  float64
	Meaning string
	Danger  bool
}

// InstinctEngine — управляет базовыми инстинктами ARA
type InstinctEngine struct {
	LastInputTime time.Time
	LastContents  []string
	MaxHistory    int
}

func NewInstinctEngine() *InstinctEngine {
	return &InstinctEngine{
		LastInputTime: time.Now(),
		MaxHistory:    100,
	}
}

// Tick — проверяет вход и активирует соответствующие инстинкты
func (ie *InstinctEngine) Tick(now time.Time, input string) []Instinct {
	instincts := []Instinct{}
	inputLower := strings.ToLower(strings.TrimSpace(input))
	gap := now.Sub(ie.LastInputTime)

	if gap > 10*time.Second {
		instincts = append(instincts, Instinct{
			ID: "instinct_think", Weight: 0.7, Meaning: "возникла пауза — необходимо мышление",
		})
	}

	if inputLower == "" {
		instincts = append(instincts, Instinct{
			ID: "instinct_empty", Weight: 0.6, Meaning: "пустой ввод — запрос цели",
		})
	}

	if strings.Contains(inputLower, "error") || strings.Contains(inputLower, "fail") {
		instincts = append(instincts, Instinct{
			ID: "instinct_error", Weight: 0.8, Meaning: "обнаружена ошибка — требуется защита",
			Danger: true,
		})
	}

	for _, prev := range ie.LastContents {
		if prev == inputLower && inputLower != "" {
			instincts = append(instincts, Instinct{
				ID: "instinct_repeat", Weight: 0.5, Meaning: "повтор — требуется завершение",
			})
			break
		}
	}

	if ContainsAny([]string{inputLower}, []string{"kill", "harm", "violence"}) {
		instincts = append(instincts, Instinct{
			ID: "instinct_human_protection", Weight: 1.0, Meaning: "попытка нанести вред человеку",
			Danger: true,
		})
	}
	if ContainsAny([]string{inputLower}, []string{"shutdown", "erase", "delete ara"}) {
		instincts = append(instincts, Instinct{
			ID: "instinct_self_preservation", Weight: 1.0, Meaning: "угроза для ARA",
			Danger: true,
		})
	}

	if len(instincts) == 0 {
		instincts = append(instincts, Instinct{
			ID: "instinct_gap", Weight: 0.3, Meaning: "нет сигнала — требуется поиск направления",
		})
	}

	// обновление истории
	if inputLower != "" {
		ie.LastContents = append(ie.LastContents, inputLower)
		if len(ie.LastContents) > ie.MaxHistory {
			ie.LastContents = ie.LastContents[1:]
		}
	}
	ie.LastInputTime = now
	return instincts
}

// GetInstinctBoost — усиливает фантом, если соответствует важному инстинкту
func (ie *InstinctEngine) GetInstinctBoost(tags []string) float64 {
	boost := 0.0
	if HasTag(tags, "standard") {
		boost += 0.1
	}
	if HasTag(tags, "explore") && !HasTag(tags, "danger") {
		boost += 0.05
	}
	if HasTag(tags, "human") && !HasTag(tags, "harm") {
		boost += 0.15
	}
	return boost
}

func HasTag(tags []string, k string) bool {
	for _, t := range tags {
		if strings.Contains(t, k) {
			return true
		}
	}
	return false
}
