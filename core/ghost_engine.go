package core

import (
	"fmt"
	"time"
)

type ReactionRule struct {
	MatchTags []string
	MinPhase  float64
	Action    func(sig Signal)
}


// Block — реактивный узел, срабатывающий на сигнал
type Block struct {
	ID            string
	Rules         []ReactionRule
	LastTriggered time.Time
	ReactionCount int
}

// React — проверка и реакция на сигнал
func (b *Block) React(sig Signal) {
	for _, rule := range b.Rules {
		if sig.Phase < rule.MinPhase {
			continue
		}
		for _, match := range rule.MatchTags {
			if contains(sig.Tags, match) {
				fmt.Printf("[Block %s] Reacting to signal: %s\n", b.ID, sig.Content)
				b.LastTriggered = time.Now()
				b.ReactionCount++
				rule.Action(sig)
				break
			}
		}
	}
}

// contains — проверка наличия тега
func contains(tags []string, key string) bool {
	for _, tag := range tags {
		if tag == key {
			return true
		}
	}
	return false
}

// GhostField — сеть блоков
type GhostField struct {
	Blocks []*Block
}

// NewGhostField — инициализация пустого поля
func NewGhostField() *GhostField {
	return &GhostField{
		Blocks: []*Block{},
	}
}

// RegisterBlock — добавление нового блока
func (gf *GhostField) RegisterBlock(b *Block) {
	gf.Blocks = append(gf.Blocks, b)
	fmt.Println("[GhostField] Registered Block:", b.ID)
}

// Propagate — передача сигнала по полю
func (gf *GhostField) Propagate(sig Signal) {
	for _, block := range gf.Blocks {
		block.React(sig)
	}
}
