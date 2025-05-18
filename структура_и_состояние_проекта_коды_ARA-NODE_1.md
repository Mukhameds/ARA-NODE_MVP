"C:\Documents\ARA-NODE_mvp\cmd\main.go":
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"ara-node/config"
	"ara-node/core"
	"ara-node/internal"

	_ "github.com/libp2p/go-libp2p/core/peer"
)

func main() {
	fmt.Println("🧠 ARA-NODE CLI started.")

	kernel := config.InitSelfKernel()
	fmt.Println("Agent ID:", kernel.AgentID)

	mem := core.NewMemoryEngine()
	signalDict := core.NewSignalDictionary(mem)

	p2p, err := internal.StartP2P(mem)
	if err != nil {
		fmt.Println("[P2P ❌]", err)
	} else {
		fmt.Println("[P2P] 🛰️ Sync active")
	}
	_ = p2p


	missionQ := mem.CreateQBit(kernel.CoreMission)
missionQ.Tags = []string{"mission", "identity"}
missionQ.Type = "system"
mem.StoreQBit(*missionQ)


	signalEngine := core.NewSignalEngine(mem)
	phantom := internal.NewFanthomEngine(mem)
	suggestor := internal.NewSuggestorEngine(mem)
	human := internal.NewHumanNodeEngine(mem)
	ghostField := core.NewGhostField()
	will := core.NewWillEngine(mem)
	will.DesireLoop()
	core.RunBootstrap(mem, signalDict)


	attention := core.NewAttentionEngine(mem, ghostField, phantom, signalEngine)
	attention.StartBackgroundThinking()

	reflex := core.NewReflexEngine()
	core.DefaultReflexSet(reflex)

	emotion := core.NewEmotionEngine()
	core.DefaultEmotionSet(emotion)

	decay := core.NewDecayAnalysisEngine(mem)

	evolution := core.NewQBitEvolutionEngine(mem)





	// Пример реактивного блока
	block := &core.Block{
		ID: "UserPhaseBlock",
		Rules: []core.ReactionRule{
			{
				MatchTags: []string{"user"},

				MinPhase: 0.8,
				Action: func(sig core.Signal) {
					fmt.Println("[Block Action] 🔁 Reacted to user signal:", sig.Content)
				},
			},
		},
	}
	ghostField.RegisterBlock(block)

	reader := bufio.NewReader(os.Stdin)

	for {
	fmt.Print("\n>> ")
	input, _ := reader.ReadString('\n')
	content := strings.TrimSpace(input)
	if content == "" {
		continue
	}
	attention.Suppress(3 * time.Second)

	// команды словаря переменных
	if content == "dict" {
		for _, vb := range signalDict.All() {
			fmt.Println("🔹", vb.ID, "|", vb.Tags, "| Reacts:", vb.Reacts)
		}
		continue
	}

	if content == "whoami" {
		fmt.Println("🤖 Agent ID:", kernel.AgentID)
		fmt.Println("🎯 Mission:", kernel.CoreMission)
		fmt.Println("🧬 Architect:", kernel.ArchitectID)
		fmt.Println("⏱ Born:", kernel.Inception.Format(time.RFC822))
		continue
	}

	if strings.HasPrefix(content, "tagvar ") {
		parts := strings.Split(content, " ")
		if len(parts) >= 3 {
			signalDict.Tag(parts[1], parts[2])
			fmt.Println("✅ Тег добавлен:", parts[1], "→", parts[2])
		}
		continue
	}

	if strings.HasPrefix(content, "delvar ") {
		parts := strings.Split(content, " ")
		if len(parts) >= 2 && signalDict.Delete(parts[1]) {
			fmt.Println("❌ Удалено:", parts[1])
		} else {
			fmt.Println("⚠️ Не найдено:", parts[1])
		}
		continue
	}

	// команды памяти и синхронизации
	if content == "dump" {
		mem.ListQBits()
		continue
	}

	if human.HandleCommand(content) {
		continue
	}

	if content == "sync-push" {
		err := internal.PushMemory(mem)
		if err != nil {
			fmt.Println("[GitSync ❌]", err)
		}
		continue
	}

	if content == "sync-pull" {
		err := internal.PullMemory(mem)
		if err != nil {
			fmt.Println("[GitSync ❌]", err)
		}
		continue
	}

	if strings.HasPrefix(content, "load_knowledge ") {
		path := strings.TrimPrefix(content, "load_knowledge ")
		err := internal.LoadKnowledge(path, mem)
		if err != nil {
			fmt.Println("⚠️ Ошибка загрузки:", err)
		}
		continue
	}

	if strings.HasPrefix(content, "load_profile ") {
		parts := strings.Split(content, " ")
		if len(parts) >= 3 {
			path := parts[1]
			license := parts[2]
			err := internal.LoadKnowledgeLicensed(path, license, mem)
			if err != nil {
				fmt.Println("⚠️ Ошибка профиля:", err)
			}
		}
		continue
	}

	if content == "phantoms" || content == "phantom-tree" {
		internal.PrintPhantomTree(mem)
		continue
	}

	if content == "decay-log" {
		decay.PrintDecayLog()
		continue
	}

	if content == "evolve" {
		evolution.EvolveAll()
		continue
	}


	// 🚀 Генерация основного сигнала
	signal := core.Signal{
		ID:        fmt.Sprintf("sig_%d", time.Now().UnixNano()),
		Content:   content,
		Tags:      []string{"user"},
		Timestamp: time.Now(),
		Phase:     0.85,
		Weight:    1.0,
		Origin:    "user",
	}

	reflex.React(signal)
	emotion.React(signal)
	reaction := signalEngine.ProcessSignal(signal)

	ghostField.Propagate(signal)
	phantom.TriggerFromMatch(signal)
	suggestor.SuggestFromQBits()
	mem.DecayQBits()
	decay.RunDecayCheck()
	evolution.EvolveAll()


	// 📡 сенсорная фильтрация через SignalDictionary
	tokens := strings.Fields(strings.ToLower(content))
	for _, token := range tokens {
		vb := signalDict.FindMatch(token)
		if vb != nil {
			allowed := false
			for _, tag := range vb.Tags {
				if strings.HasPrefix(tag, "letter") || strings.HasPrefix(tag, "number") || strings.HasPrefix(tag, "image") {
					allowed = true
					break
				}
			}
			if allowed {
				mem.Broadcast(vb.QBit)
			}
		}
	}

	instinctEngine := core.NewInstinctEngine()
	predictor := core.NewPredictionEngine()

	instincts := instinctEngine.Tick(time.Now(), content)
	for _, inst := range instincts {
		fmt.Println("[Instinct] Triggered:", inst)
		sig := core.Signal{
			ID:        fmt.Sprintf("inst_%d", time.Now().UnixNano()),
			Content:   inst,
			Tags:      []string{"instinct"},
			Timestamp: time.Now(),
			Phase:     0.9,
			Weight:    1.2,
			Origin:    "system",
		}
		signalEngine.ProcessSignal(sig)
		ghostField.Propagate(sig)
		phantom.TriggerFromMatch(sig)
	}

	if predicted, ok := predictor.Predict(content); ok {
		fmt.Println("[⏳ Prediction] Phantom may be:", predicted)
		signalEngine.ProcessSignal(core.Signal{
			ID:        fmt.Sprintf("pred_%d", time.Now().UnixNano()),
			Content:   predicted,
			Type:      "prediction",
			Tags:      []string{"predict"},
			Timestamp: time.Now(),
			Phase:     0.92,
			Weight:    1.1,
			Origin:    "system",
		})
	}

	fmt.Println("←", reaction.Response)
	fmt.Println("↪ Tags:", reaction.Tags)
}
}


"C:\Documents\ARA-NODE_mvp\config\manifest.go":
package config

import (
	"time"
	"fmt"
)

// SelfKernel — неизменяемая основа идентичности агента
type SelfKernel struct {
	AgentID     string
	ArchitectID string
	CoreMission string
	Inception   time.Time
}

// InitSelfKernel — инициализация ядра
func InitSelfKernel() *SelfKernel {
	kernel := &SelfKernel{
		AgentID:     "ARA::node::001",
		ArchitectID: "User::Architect",
		CoreMission: "Amplify and assist user cognition through signal logic.",
		Inception:   time.Now(),
	}
	fmt.Println("[SelfKernel] Initialized:", kernel.AgentID)
	return kernel
}

"C:\Documents\ARA-NODE_mvp\core\attention_engine.go":
package core

import (
	"fmt"
	"time"
	"math"
)

// AttentionEngine — генератор внутренних возбуждений
type AttentionEngine struct {
	Memory   *MemoryEngine
	Ghost    *GhostField
	Fanthom  FanthomInterface
	Engine   *SignalEngine
	SuppressedUntil time.Time
}

func NewAttentionEngine(mem *MemoryEngine, ghost *GhostField, fant FanthomInterface, eng *SignalEngine) *AttentionEngine {
	return &AttentionEngine{
		Memory:  mem,
		Ghost:   ghost,
		Fanthom: fant,
		Engine:  eng,
	}
}

// Suppress временно приостанавливает внутреннее мышление (например, при вводе пользователя)
func (ae *AttentionEngine) Suppress(d time.Duration) {
	ae.SuppressedUntil = time.Now().Add(d)
}

// StartBackgroundThinking запускает постоянное самовозбуждение
func (ae *AttentionEngine) StartBackgroundThinking() {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			if time.Now().Before(ae.SuppressedUntil) {
				continue
			}

			active := ae.Memory.FindAll(func(q QBit) bool {
				return q.Weight*q.Phase > 0.6 && !q.Archived && q.Type != "standard"
			})

			for _, q := range active {
				sig := Signal{
					ID:        fmt.Sprintf("bg_%d", time.Now().UnixNano()),
					Content:   q.Content,
					Tags:      q.Tags,
					Type:      "background",
					Origin:    "internal",
					Phase:     math.Min(q.Phase+0.05, 1.0),
					Weight:    q.Weight * 0.9,
					Timestamp: time.Now(),
				}

				ae.Engine.ProcessSignal(sig)
				ae.Ghost.Propagate(sig)
				ae.Fanthom.TriggerFromMatch(sig)
			}
		}
	}()
}


"C:\Documents\ARA-NODE_mvp\core\bootstrap.go":
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

"C:\Documents\ARA-NODE_mvp\core\decay_analysis_engine.go":
package core

import (
	"fmt"
	"time"
)

// DecayEvent — лог обнуления или вымирания узла
type DecayEvent struct {
	ID        string
	Reason    string
	Timestamp time.Time
}

// DecayAnalysisEngine — отслеживает процессы старения памяти
type DecayAnalysisEngine struct {
	Log []DecayEvent
	Mem *MemoryEngine
}

func NewDecayAnalysisEngine(mem *MemoryEngine) *DecayAnalysisEngine {
	return &DecayAnalysisEngine{
		Log: []DecayEvent{},
		Mem: mem,
	}
}

func (de *DecayAnalysisEngine) RunDecayCheck() {
	for id, q := range de.Mem.QBits {
		if q.Archived && q.Weight < 0.05 {
			delete(de.Mem.QBits, id)
			de.Log = append(de.Log, DecayEvent{
				ID:        id,
				Reason:    "fully decayed",
				Timestamp: time.Now(),
			})
			fmt.Println("[Decay] ☠️ Removed:", id)
		}
	}
}

func (de *DecayAnalysisEngine) PrintDecayLog() {
	if len(de.Log) == 0 {
		fmt.Println("[DecayLog] 🔹 Память ещё не подвергалась очистке.")
		return
	}
	fmt.Println("[DecayLog] 🧩 Deleted QBits:")
	for _, entry := range de.Log {
		fmt.Printf("⏱ %s | %s | %s\n", entry.Timestamp.Format(time.RFC822), entry.ID, entry.Reason)
	}
}


"C:\Documents\ARA-NODE_mvp\core\emotion_engine.go":

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

"C:\Documents\ARA-NODE_mvp\core\ghost_engine.go":
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

