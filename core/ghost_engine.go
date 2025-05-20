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

func (b *Block) React(sig Signal) {
	if time.Since(b.lastFire) < b.Cooldown {
		return
	}
	for _, r := range b.Rules {
		if tagsMatch(sig.Tags, r.MatchTags) && sig.Phase >= r.MinPhase {
			fmt.Printf("[Ghost] [%s] rule fired on signal: %s\n", b.Type, sig.ID)
			r.Action(sig)
			b.lastFire = time.Now()
			return
		}
	}
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

// Propagate — распространяет сигнал по всем блокам
func (g *GhostField) Propagate(sig Signal) {
	for _, b := range g.Blocks {
		b.React(sig)
	}
}
