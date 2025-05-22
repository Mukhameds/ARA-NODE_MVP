package field

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"ara-node/core"
)

// ReactionRule — правило, по которому блок реагирует на сигнал
type ReactionRule struct {
	MatchTags []string              // нужные теги
	MinPhase  float64               // минимальная фаза
	Action    func(sig core.Signal) // реакция
}

// Block — реактивный модуль (мозговой узел)
type Block struct {
	Type     string
	Rules    []ReactionRule
	Cooldown time.Duration
	lastFire time.Time
	lock     sync.Mutex
}

// CanReact — проверка, может ли блок сейчас среагировать (по cooldown)
func (b *Block) CanReact() bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	return time.Since(b.lastFire) >= b.Cooldown
}

// MarkReacted — обновляет момент последнего срабатывания
func (b *Block) MarkReacted() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.lastFire = time.Now()
}

// React — запускает все подходящие правила блока параллельно
func (b *Block) React(sig core.Signal) {
	if !b.CanReact() {
		return
	}

	triggered := false
	for _, r := range b.Rules {
		if tagsMatch(sig.Tags, r.MatchTags) && sig.Phase >= r.MinPhase {
			go r.Action(sig)
			fmt.Printf("[Field] [%s] rule fired on signal: %s\n", b.Type, sig.ID)
			triggered = true
		}
	}
	if triggered {
		b.MarkReacted()
	}
}

// Matrix — реакционное поле (эквивалент GhostField)
type Matrix struct {
	Name   string
	Blocks []*Block
}

// NewMatrix — создаёт новое поле
func NewMatrix(name string) *Matrix {
	return &Matrix{
		Name:   name,
		Blocks: []*Block{},
	}
}

// RegisterBlock — регистрирует правило в блоке или создаёт блок
func (m *Matrix) RegisterBlock(blockType string, rule ReactionRule, cooldown time.Duration) {
	for _, b := range m.Blocks {
		if b.Type == blockType {
			b.Rules = append(b.Rules, rule)
			return
		}
	}
	m.Blocks = append(m.Blocks, &Block{
		Type:     blockType,
		Cooldown: cooldown,
		Rules:    []ReactionRule{rule},
	})
}

// Propagate — распространяет сигнал по всей матрице
func (m *Matrix) Propagate(sig core.Signal) {
	for _, tag := range sig.Tags {
		switch tag {
		case "silent":
			fmt.Printf("[Field:%s] 🔕 Silent signal skipped.\n", m.Name)
			return
		case "interrupt":
			fmt.Printf("[Field:%s] 🚫 Interrupt signal received, skipping.\n", m.Name)
			return
		}
	}

	for _, b := range m.Blocks {
		go b.React(sig)
	}
}

// tagsMatch — частичное сравнение тегов (суперпозиционно)
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
