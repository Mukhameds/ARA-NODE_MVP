package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// BootstrapBlock — интерфейс опросного блока
type BootstrapBlock interface {
	ID() string
	Prompt() string
	Tags() []string
	Run(input string, mem *MemoryEngine, dict *SignalDictionary)
}

// GoalBlock — цель жизни
type GoalBlock struct{}
func (b GoalBlock) ID() string         { return "q_mission" }
func (b GoalBlock) Prompt() string     { return "Какая твоя главная цель в жизни?" }
func (b GoalBlock) Tags() []string     { return []string{"goal", "mission"} }
func (b GoalBlock) Run(input string, mem *MemoryEngine, dict *SignalDictionary) {
	q := mem.CreateQBit(input)
	q.Tags = b.Tags()
	dict.Add(b.ID(), input, b.Tags(), []string{input})
}

// InterestBlock — интересы
type InterestBlock struct{}
func (b InterestBlock) ID() string     { return "q_interest" }
func (b InterestBlock) Prompt() string { return "Какие темы тебе наиболее интересны?" }
func (b InterestBlock) Tags() []string { return []string{"interest"} }
func (b InterestBlock) Run(input string, mem *MemoryEngine, dict *SignalDictionary) {
	q := mem.CreateQBit(input)
	q.Tags = b.Tags()
	dict.Add(b.ID(), input, b.Tags(), []string{input})
}

// HelpBlock — как помочь
type HelpBlock struct{}
func (b HelpBlock) ID() string         { return "q_help" }
func (b HelpBlock) Prompt() string     { return "Как ты хочешь, чтобы ARA помогала тебе?" }
func (b HelpBlock) Tags() []string     { return []string{"function", "support"} }
func (b HelpBlock) Run(input string, mem *MemoryEngine, dict *SignalDictionary) {
	q := mem.CreateQBit(input)
	q.Tags = b.Tags()
	dict.Add(b.ID(), input, b.Tags(), []string{input})
}

// RoleBlock — кто ты
type RoleBlock struct{}
func (b RoleBlock) ID() string         { return "q_role" }
func (b RoleBlock) Prompt() string     { return "Кто ты по жизни? (учёный, инженер, философ...)" }
func (b RoleBlock) Tags() []string     { return []string{"profile", "role"} }
func (b RoleBlock) Run(input string, mem *MemoryEngine, dict *SignalDictionary) {
	q := mem.CreateQBit(input)
	q.Tags = b.Tags()
	dict.Add(b.ID(), input, b.Tags(), []string{input})
}

// RunBootstrap — запуск всех Bootstrap-блоков
func RunBootstrap(mem *MemoryEngine, dict *SignalDictionary) {
	blocks := []BootstrapBlock{
		GoalBlock{}, InterestBlock{}, HelpBlock{}, RoleBlock{},
	}
	fmt.Println("🧬 [ARA Bootstrap] Начало инициализации личности...")

	reader := bufio.NewReader(os.Stdin)

	for _, b := range blocks {
		fmt.Println("🧠", b.Prompt())
		fmt.Print("→ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			b.Run(input, mem, dict)
		}
	}

	fmt.Println("✅ [Bootstrap] Базовые цели и профиль сохранены.")
}
