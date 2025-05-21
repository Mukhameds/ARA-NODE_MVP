package core

import (
	"fmt"
	"strings"
	"time"
)

// ReactionRule — правило реакции блока на входной сигнал
type ReactionRule struct {
	MatchTags []string
	MinPhase  float64
	Action    func(sig Signal)
}

// Block — один реактивный модуль: emotion, reflex, suggestor и т.д.
type Block struct {
	Type     string
	Rules    []ReactionRule
	Cooldown time.Duration
	lastFire time.Time
}

func (b *Block) React(sig Signal) bool {
	if time.Since(b.lastFire) < b.Cooldown {
		return false
	}
	for _, r := range b.Rules {
		if tagsMatch(sig.Tags, r.MatchTags) && sig.Phase >= r.MinPhase {
			fmt.Printf("[Ghost] [%s] rule fired on signal: %s\n", b.Type, sig.ID)
			r.Action(sig)
			b.lastFire = time.Now()
			return true
		}
	}
	return false
}

func tagsMatch(signalTags, matchTags []string) bool {
	for _, mt := range matchTags {
		for _, st := range signalTags {
			if strings.Contains(st, mt) {
				return true
			}
		}
	}
	return false
}

// GhostField — содержит все реактивные блоки
type GhostField struct {
	Blocks []*Block
}

func NewGhostField() *GhostField {
	return &GhostField{Blocks: []*Block{}}
}

// Register — регистрирует внешний блок в реактивную сеть
func (g *GhostField) Register(blockType string, rule ReactionRule, cooldown time.Duration) {
	block := &Block{
		Type:     blockType,
		Cooldown: cooldown,
		Rules:    []ReactionRule{rule},
	}
	g.Blocks = append(g.Blocks, block)
}

// UnregisterBlock — удаляет блок по типу
func (g *GhostField) UnregisterBlock(blockType string) {
	var filtered []*Block
	for _, b := range g.Blocks {
		if b.Type != blockType {
			filtered = append(filtered, b)
		} else {
			fmt.Println("[GhostField] 🗑️ Removed Block:", b.Type)
		}
	}
	g.Blocks = filtered
}

// Propagate — распространяет сигнал по всем блокам
func (g *GhostField) Propagate(sig Signal) {
	// специальные сигналы
	for _, tag := range sig.Tags {
		switch tag {
		case "silent":
			fmt.Println("[GhostField] 🔕 Silent signal ignored.")
			return
		case "interrupt":
			fmt.Println("[GhostField] 🚫 Interrupt signal, halting propagation.")
			return
		case "ethereal":
			fmt.Println("[GhostField] 👻 Ethereal signal, observing only.")
			// не прерываем, но можно использовать в будущем
		}
	}

	reacted := false
	for _, b := range g.Blocks {
		if b.React(sig) {
			reacted = true
		}
	}
	if !reacted {
		fmt.Printf("[GhostField] ⚠️ No reaction for signal: %s (%v)\n", sig.Content, sig.Tags)
	}
}
