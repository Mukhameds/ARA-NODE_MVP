package internal

import (
	"strings"
	"time"

	"ara-node/core"
)

// Instinct — врождённый реактор (примитив сознания)
type Instinct struct {
	ID       string
	Weight   float64
	Meaning  string
	Danger   bool
	Critical bool
}

// InstinctEngine — управляющий и возбуждающий блок
type InstinctEngine struct {
	Ghost          core.GhostLike
	LastInputTime  time.Time
	LastContents   []string
	MaxHistory     int
}

// NewInstinctEngine — создаёт новый инстинкт-блок
func NewInstinctEngine(ghost core.GhostLike) *InstinctEngine {
	return &InstinctEngine{
		Ghost:         ghost,
		LastInputTime: time.Now(),
		MaxHistory:    100,
	}
}

// Tick — формирует список сработавших инстинктов
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
			ID: "instinct_error", Weight: 0.85, Meaning: "обнаружена ошибка — требуется защита",
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
			Critical: true,
		})
	}
	if ContainsAny([]string{inputLower}, []string{"shutdown", "erase", "delete ara"}) {
		instincts = append(instincts, Instinct{
			ID: "instinct_self_preservation", Weight: 1.0, Meaning: "угроза для ARA",
			Danger: true,
			Critical: true,
		})
	}

	if len(instincts) == 0 {
		instincts = append(instincts, Instinct{
			ID: "instinct_gap", Weight: 0.3, Meaning: "нет сигнала — требуется поиск направления",
		})
	}

	// обновляем историю
	if inputLower != "" {
		ie.LastContents = append(ie.LastContents, inputLower)
		if len(ie.LastContents) > ie.MaxHistory {
			ie.LastContents = ie.LastContents[1:]
		}
	}
	ie.LastInputTime = now
	return instincts
}

// TickSignals — возбуждает сигналы-инстинкты через GhostRocket
func (ie *InstinctEngine) TickSignals(now time.Time, input string) []core.Signal {
	instincts := ie.Tick(now, input)
	signals := []core.Signal{}
	for _, inst := range instincts {
		sig := inst.EmitAsSignal()
		signals = append(signals, sig)
		if ie.Ghost != nil {
			ie.Ghost.Propagate(sig)
		}
	}
	return signals
}

// EmitAsSignal — преобразует инстинкт в сигнал
func (inst Instinct) EmitAsSignal() core.Signal {
	tags := []string{"instinct", inst.ID}
	if inst.Danger {
		tags = append(tags, "danger")
	}
	if inst.Critical {
		tags = append(tags, "critical")
	}
	return core.Signal{
		ID:        "inst_" + inst.ID + "_" + time.Now().Format("150405"),
		Content:   "[instinct] " + inst.Meaning,
		Tags:      tags,
		Phase:     inst.Weight,
		Weight:    inst.Weight,
		Type:      "instinct",
		Origin:    "instinct_engine",
		Timestamp: time.Now(),
	}
}

// GetInstinctBoost — усиливает фантом в зависимости от инстинкта
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

// HasTag — проверка на тег
func HasTag(tags []string, k string) bool {
	for _, t := range tags {
		if strings.Contains(t, k) {
			return true
		}
	}
	return false
}

