
### **Directory and modules of the ARA-NODE system**

---

## **"C:\Documents\ARA-NODE_mvp\cmd\main.go"**

---

''' 
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

	// Инициализация движков инстинктов и эмоций
	instinctEngine := core.NewInstinctEngine()
	emotionEngine := core.NewEmotionEngine()
	core.DefaultEmotionSet(emotionEngine)

	// Интеграция PhantomEngine с instinct и emotion движками
	phantom := internal.NewPhantomEngine(mem, instinctEngine, emotionEngine)
	go func() {
	for {
		time.Sleep(5 * time.Second)
		phantom.TickUpdatePhantoms()
	}
}()


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

	decay := core.NewDecayAnalysisEngine(mem)
	evolution := core.NewQBitEvolutionEngine(mem)

	// Пример реактивного блока
	block := &core.Block{
		ID: "UserPhaseBlock",
		Rules: []core.ReactionRule{
			{
				MatchTags: []string{"user"},
				MinPhase:  0.8,
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
		emotionEngine.React(signal)
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

		predictor := core.NewPredictionEngine()
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

'''
---

## **"C:\Documents\ARA-NODE_mvp\config\manifest.go"**

---

'''
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

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\attention_engine.go"**

---

'''

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
 
 '''

 ---

 ## **"C:\Documents\ARA-NODE_mvp\core\bootstrap.go"**

 ---

 '''

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

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\decay_analysis_engine.go"**

---

'''

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

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\emotion_engine.go"**

---

'''

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


'''

---

## **"C:\Documents\ARA-NODE_mvp\core\ghost_engine.go"**

---

'''

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

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\helpers.go"**

---

package core

// RemoveTag удаляет указанный тег из слайса тегов
func RemoveTag(tags []string, target string) []string {
	var result []string
	for _, tag := range tags {
		if tag != target {
			result = append(result, tag)
		}
	}
	return result
}

// PhaseClose возвращает true, если фазы близки с учётом допуска
func PhaseClose(p1, p2, tolerance float64) bool {
	diff := p1 - p2
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}


// Contains проверяет, содержит ли срез строку
func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\instincts.go"**

---

'''

package core

import (
	"strings"
	"sync"
	"time"
)

type InstinctEngine struct {
	LastInputTime time.Time
	mu            sync.Mutex
	recentSignals []string
	maxHistory    int
}

func NewInstinctEngine() *InstinctEngine {
	return &InstinctEngine{
		LastInputTime: time.Now(),
		maxHistory:    100,
		recentSignals: make([]string, 0, 100),
	}
}

func (ie *InstinctEngine) Tick(currentTime time.Time, signal string) (instincts []string) {
	ie.mu.Lock()
	defer ie.mu.Unlock()

	var results []string

	// 1. Silence trigger (instinct_think)
	if currentTime.Sub(ie.LastInputTime) > 10*time.Second {
		results = append(results, "instinct_think")
		ie.LastInputTime = currentTime
	}

	// 2. Repeat detection (instinct_repeat)
	if ie.isRepeat(signal) {
		results = append(results, "instinct_repeat")
	}

	// 3. Error pattern (instinct_error)
	if strings.Contains(strings.ToLower(signal), "error") {
		results = append(results, "instinct_error")
	}

	// 4. Empty (manual trigger)
	if strings.TrimSpace(signal) == "" {
		results = append(results, "instinct_empty")
	}

	// Добавляем сигнал в историю
	ie.addSignal(signal)

	return results
}

func (ie *InstinctEngine) isRepeat(signal string) bool {
	// Проверяем, встречался ли сигнал в истории недавно
	for _, s := range ie.recentSignals {
		if s == signal {
			return true
		}
	}
	return false
}

func (ie *InstinctEngine) addSignal(signal string) {
	if signal == "" {
		return
	}
	// Добавляем в историю, с ограничением длины
	if len(ie.recentSignals) >= ie.maxHistory {
		ie.recentSignals = ie.recentSignals[1:]
	}
	ie.recentSignals = append(ie.recentSignals, signal)
}

// ClearHistory очищает историю сигналов
func (ie *InstinctEngine) ClearHistory() {
	ie.mu.Lock()
	defer ie.mu.Unlock()
	ie.recentSignals = make([]string, 0, ie.maxHistory)
}


'''

---

## **"C:\Documents\ARA-NODE_mvp\core\memory_engine.go"**

---

package core

import (
	"fmt"
	"sync"
	"time"
)

// MemoryEngine — сигнальная память агента
type MemoryEngine struct {
	QBits       map[string]QBit
	Mu          sync.Mutex
	PhantomTree []PhantomLog
}

// NewMemoryEngine — инициализация памяти
func NewMemoryEngine() *MemoryEngine {
	return &MemoryEngine{
		QBits: make(map[string]QBit),
	}
}

// StoreQBit — сохранить QBit в память
func (m *MemoryEngine) StoreQBit(q QBit) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	m.QBits[q.ID] = q
	fmt.Println("[MemoryEngine] Stored QBit:", q.ID)
}

// GetQBit — получить QBit по ID
func (m *MemoryEngine) GetQBit(id string) (QBit, bool) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	q, exists := m.QBits[id]
	return q, exists
}

// DeleteQBit — удалить QBit по ID
func (m *MemoryEngine) DeleteQBit(id string) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	if _, exists := m.QBits[id]; exists {
		delete(m.QBits, id)
		fmt.Println("[MemoryEngine] ❌ QBit deleted:", id)
	}
}

// DecayQBits — уменьшает вес старых или слабых QBit
func (m *MemoryEngine) DecayQBits() {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	now := time.Now()
	for id, q := range m.QBits {
		age := now.Sub(q.CreatedAt).Seconds()
		decayFactor := 0.5 * age
		q.Weight -= decayFactor
		if q.Weight < 0 {
			q.Weight = 0
		}
		if q.Weight < 0.2 {
			q.Archived = true
		}
		fmt.Printf("[Decay] %s → age: %.1fs, decay: %.2f, new weight: %.2f\n", id, age, decayFactor, q.Weight)
		m.QBits[id] = q
	}
}

// FindByTag — вернуть QBit по тегу
func (m *MemoryEngine) FindByTag(tag string) []QBit {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	var result []QBit
	for _, q := range m.QBits {
		if q.Archived {
			continue
		}
		for _, t := range q.Tags {
			if t == tag {
				result = append(result, q)
				break
			}
		}
	}
	return result
}

// FindByPhase — вернуть QBit с близкой фазой
func (m *MemoryEngine) FindByPhase(target float64, tolerance float64) []QBit {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	var result []QBit
	for _, q := range m.QBits {
		if q.Archived {
			continue
		}
		if abs(q.Phase-target) <= tolerance {
			result = append(result, q)
		}
	}
	return result
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

// ListQBits — выводит все неархивированные QBit
func (m *MemoryEngine) ListQBits() {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	fmt.Println("\n[Memory Dump]")
	for _, q := range m.QBits {
		if q.Archived {
			continue
		}
		fmt.Printf("• ID: %s | Phase: %.2f | Weight: %.2f | Tags: %v\n", q.ID, q.Phase, q.Weight, q.Tags)
	}
}

func (m *MemoryEngine) AdjustWeight(id string, delta float64) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	q, exists := m.QBits[id]
	if !exists || q.Archived {
		return
	}
	q.Weight += delta
	if q.Weight < 0 {
		q.Weight = 0
	}
	m.QBits[id] = q
}

func (m *MemoryEngine) AddTag(id string, tag string) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	q, exists := m.QBits[id]
	if !exists || q.Archived {
		return
	}
	for _, t := range q.Tags {
		if t == tag {
			return
		}
	}
	q.Tags = append(q.Tags, tag)
	m.QBits[id] = q
}

func (m *MemoryEngine) Merge(other map[string]QBit) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	for id, q := range other {
		if _, exists := m.QBits[id]; !exists {
			m.QBits[id] = q
		}
	}
}

// CreateQBit — создать и сохранить новый QBit
func (m *MemoryEngine) CreateQBit(content string) *QBit {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	q := QBit{
		ID:        fmt.Sprintf("qbit_%d", time.Now().UnixNano()),
		Content:   content,
		Tags:      []string{"auto"},
		CreatedAt: time.Now(),
		Phase:     0.75,
		Weight:    1.0,
	}

	m.QBits[q.ID] = q
	fmt.Println("[MemoryEngine] Auto-created QBit:", q.ID)
	return &q
}

// Broadcast — возбуждает QBit во всей памяти (прямая трансляция)
func (m *MemoryEngine) Broadcast(q *QBit) {
	if q == nil {
		return
	}
	m.StoreQBit(*q)
	fmt.Println("[MemoryEngine] 📡 Broadcast QBit:", q.ID)
}

// FindAll — вернуть все QBits, удовлетворяющие фильтру
func (m *MemoryEngine) FindAll(filter func(QBit) bool) []QBit {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	var result []QBit
	for _, q := range m.QBits {
		if filter(q) {
			result = append(result, q)
		}
	}
	return result
}

func (m *MemoryEngine) UpdateQBit(qbit QBit) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	if _, exists := m.QBits[qbit.ID]; exists {
		m.QBits[qbit.ID] = qbit
	}
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\prediction.go"**

---

'''

package core

type PredictionEngine struct {
	Chains [][]string // Примитивный шаблон: [q1, q2] → q3
}

func NewPredictionEngine() *PredictionEngine {
	return &PredictionEngine{
		Chains: [][]string{
			{"q1", "q2", "q3"},
			{"q5", "q1", "q3"},
		},
	}
}

func (pe *PredictionEngine) Predict(input string) (predicted string, ok bool) {
	for _, chain := range pe.Chains {
		if len(chain) < 3 {
			continue
		}
		if chain[0] == input || chain[1] == input {
			return chain[2], true
		}
	}
	return "", false
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\QBitEvolutionEngine.go"**

---

'''
package core

import (
	"fmt"
	
)

// QBitEvolutionEngine — отвечает за развитие или деградацию QBits
type QBitEvolutionEngine struct {
	Memory *MemoryEngine
}

func NewQBitEvolutionEngine(mem *MemoryEngine) *QBitEvolutionEngine {
	return &QBitEvolutionEngine{Memory: mem}
}

func (qe *QBitEvolutionEngine) EvolveAll() {
	for id, q := range qe.Memory.QBits {
		if q.Archived {
			continue
		}

		// Эволюция: повышение веса + переход типа
		if q.Weight > 2.5 && q.Type == "" {
			q.Type = "reflex"
			qe.Memory.QBits[id] = q
			fmt.Println("[Evolve] 🌱 Promoted to reflex:", id)
			continue
		}

		if q.Weight > 3.0 && q.Type == "reflex" {
			q.Type = "generator"
			qe.Memory.QBits[id] = q
			fmt.Println("[Evolve] 🔁 Reflex → generator:", id)
			continue
		}

		if q.Weight < 0.1 && q.Type != "standard" {
			q.Archived = true
			qe.Memory.QBits[id] = q
			fmt.Println("[Evolve] 💤 Archived:", id)
		}
	}
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\reflex_engine.go"**

---

package core

import (
	"fmt"
	
)

// ReflexRule — мгновенная реакция на сигнал
type ReflexRule struct {
	MatchTag string
	Action   func(sig Signal)
}

// ReflexEngine — хранит и вызывает рефлексы
type ReflexEngine struct {
	Rules []ReflexRule
}

func NewReflexEngine() *ReflexEngine {
	return &ReflexEngine{
		Rules: []ReflexRule{},
	}
}

func (re *ReflexEngine) AddRule(tag string, action func(sig Signal)) {
	re.Rules = append(re.Rules, ReflexRule{
		MatchTag: tag,
		Action:   action,
	})
}

func (re *ReflexEngine) React(sig Signal) {
	for _, rule := range re.Rules {
		if containsTag(sig.Tags, rule.MatchTag) {
			fmt.Println("[Reflex] ⚡ Instant reaction to:", sig.Content)
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
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\signal_dictionary.go"**

---

'''

package core

import (
	"strings"
)

// VariableBlock — сигнальная единица восприятия (буква, слово, символ, образ и т.д.)
type VariableBlock struct {
	ID     string   // Уникальный ID (например, "A", "hello", "5")
	Signal string   // Визуальное представление или образ
	Tags   []string // Типы: letter, word, number, image...
	Reacts []string // Синонимы, триггеры
	QBit   *QBit    // Связанный элемент памяти
	Auto   bool     // Создан автоматически?
}

// SignalDictionary — глобальный словарь восприятия
type SignalDictionary struct {
	Variables map[string]*VariableBlock
	Memory    *MemoryEngine
}

// NewSignalDictionary — создать словарь
func NewSignalDictionary(mem *MemoryEngine) *SignalDictionary {
	return &SignalDictionary{
		Variables: make(map[string]*VariableBlock),
		Memory:    mem,
	}
}

// FindMatch — поиск VariableBlock по токену
func (sd *SignalDictionary) FindMatch(token string) *VariableBlock {
	for _, vb := range sd.Variables {
		for _, r := range vb.Reacts {
			if strings.EqualFold(token, r) {
				return vb
			}
		}
	}
	return nil
}

// AutoLearn — создать новый VariableBlock из незнакомого токена
func (sd *SignalDictionary) AutoLearn(token string) *VariableBlock {
	vb := &VariableBlock{
		ID:     token,
		Signal: token,
		Tags:   []string{"type:unknown"},
		Reacts: []string{token},
		QBit:   sd.Memory.CreateQBit(token),
		Auto:   true,
	}
	sd.Variables[token] = vb
	return vb
}

// Add — ручное добавление блока
func (sd *SignalDictionary) Add(id, signal string, tags, reacts []string) *VariableBlock {
	vb := &VariableBlock{
		ID:     id,
		Signal: signal,
		Tags:   tags,
		Reacts: reacts,
		QBit:   sd.Memory.CreateQBit(signal),
		Auto:   false,
	}
	sd.Variables[id] = vb
	return vb
}

// Delete — удалить VariableBlock
func (sd *SignalDictionary) Delete(id string) bool {
	if _, ok := sd.Variables[id]; ok {
		delete(sd.Variables, id)
		return true
	}
	return false
}

// Tag — добавить тег к блоку
func (sd *SignalDictionary) Tag(id, tag string) {
	if vb, ok := sd.Variables[id]; ok {
		vb.Tags = append(vb.Tags, tag)
	}
}

// All — список всех блоков
func (sd *SignalDictionary) All() []*VariableBlock {
	out := []*VariableBlock{}
	for _, vb := range sd.Variables {
		out = append(out, vb)
	}
	return out
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\signal_engine.go"**

---

'''

package core

import (
	"fmt"
	"time"
)


// SignalEngine — обрабатывает входящие сигналы и вызывает реакцию
type SignalEngine struct {
	Memory *MemoryEngine
}

func NewSignalEngine(mem *MemoryEngine) *SignalEngine {
	return &SignalEngine{Memory: mem}
}

// ProcessSignal — основной метод приёма и реакции
func (se *SignalEngine) ProcessSignal(sig Signal) Reaction {
	fmt.Println("[SignalEngine] Received:", sig.Content)

	// Сохраняем сигнал как QBit
	qbit := QBit{
		ID:        "qbit_" + sig.ID,
		Content:   sig.Content,
		Tags:      sig.Tags,
		CreatedAt: time.Now(),
		Weight:    sig.Weight,
		Phase:     sig.Phase,
		Type:      sig.Type,
		Origin:    sig.Origin,
	}
	se.Memory.StoreQBit(qbit)

	// Проверка совпадения по фазе (заглушка)
	if sig.Phase > 0.8 {
		return Reaction{
			TriggeredBy: sig.ID,
			Response:    "Phantom triggered by phase match",
			Tags:        []string{"phantom"},
			Confidence:  0.95,
		}
	}

	// Обычная реакция
	return Reaction{
		TriggeredBy: sig.ID,
		Response:    "Signal processed and stored",
		Tags:        []string{"ack"},
		Confidence:  0.8,
	}
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\standards.go"**

---

'''

package core

import (
	"strings"
)

type StandardBlock struct {
	ID       string
	Keywords []string
	Priority float64
}

// 📚 Статические эталонные блоки миссий ARA
var StandardLibrary = []StandardBlock{
	{
		ID:       "mission_abundance",
		Keywords: []string{"изобилие", "людям", "помощь", "решение проблем", "облегчить жизнь"},
		Priority: 1.0,
	},
	{
		ID:       "mission_learning",
		Keywords: []string{"обучение", "знания", "развитие", "понимание", "объяснение"},
		Priority: 0.9,
	},
	{
		ID:       "mission_sync",
		Keywords: []string{"синхронизация", "объединение", "p2p", "обмен"},
		Priority: 0.8,
	},
}

// 🔍 MatchWithStandards проверяет, соответствует ли текст какому-либо эталонному блоку
// Возвращает: ID блока, приоритет, количество совпавших ключевых слов
func MatchWithStandards(content string) (string, float64, int) {
	content = strings.ToLower(content)
	bestMatch := ""
	bestScore := 0
	bestPriority := 0.0

	for _, std := range StandardLibrary {
		matchCount := 0
		for _, keyword := range std.Keywords {
			if strings.Contains(content, strings.ToLower(strings.TrimSpace(keyword))) {
				matchCount++
			}
		}
		if matchCount > bestScore {
			bestScore = matchCount
			bestMatch = std.ID
			bestPriority = std.Priority
		}
	}

	if bestScore >= 3 {
		return bestMatch, bestPriority, bestScore
	}
	return "", 0.0, 0
}

// 🧱 GetStandardByID возвращает эталонный блок по ID
func GetStandardByID(id string) *StandardBlock {
	for _, std := range StandardLibrary {
		if std.ID == id {
			return &std
		}
	}
	return nil
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\core\types.go"**

---

'''

package core

import "time"

// Signal — входной сигнал, возбуждающий реакцию
type Signal struct {
	ID        string
	Content   string
	Tags      []string
	Timestamp time.Time
	Phase     float64
	Weight    float64
	Origin    string
	Type      string // тип сигнала: user, instinct, background, prediction
}

// QBit — единица памяти
type QBit struct {
	ID        string
	Content   string
	Tags      []string
	CreatedAt time.Time
	Weight    float64
	Phase     float64
	Type      string  // тип узла: reflex, generator, standard, etc.
	Origin    string  // источник: user, system, network
	Archived  bool
}

// Reaction — результат обработки сигнала
type Reaction struct {
	TriggeredBy string
	Response    string
	Tags        []string
	Confidence  float64
}

// FanthomInterface — интерфейс для фантомных систем
type FanthomInterface interface {
	TriggerFromMatch(sig Signal)
}

type PhantomLog struct {
	PhantomID string
	SourceIDs []string
}


'''

---

## **"C:\Documents\ARA-NODE_mvp\core\will_engine.go"**

---

'''

package core

import (
	"fmt"
	"strings"
	"time"
)

// Intent — внутренняя цель агента
type Intent struct {
	Tag     string
	Phase   float64
	Urgency float64
}

// WillEngine — движок воли агента
type WillEngine struct {
	Memory    *MemoryEngine
	Delay     time.Duration
	Active    bool
	lastTried map[string]time.Time
}

// NewWillEngine — создать движок воли
func NewWillEngine(mem *MemoryEngine) *WillEngine {
	return &WillEngine{
		Memory:    mem,
		Delay:     8 * time.Second,
		Active:    true,
		lastTried: make(map[string]time.Time),
	}
}

// isAlignedWithStandards — сравнение с эталонными блоками
func isAlignedWithStandards(content string) bool {
	for _, std := range StandardLibrary {
		for _, kw := range std.Keywords {
			if strings.Contains(strings.ToLower(content), strings.ToLower(kw)) {
				return true
			}
		}
	}
	return false
}

// DesireLoop — постоянный фон для самозапуска
func (we *WillEngine) DesireLoop() {
	go func() {
		for we.Active {
			time.Sleep(we.Delay)

			now := time.Now()
			intent := Intent{Tag: "user", Phase: 0.85, Urgency: 1.0}
			qbits := we.Memory.FindByTag(intent.Tag)

			for _, q := range qbits {
				if q.Archived || q.Phase < intent.Phase {
					continue
				}

				// задержка между сверками
				if t, seen := we.lastTried[q.ID]; seen && now.Sub(t) < we.Delay {
					continue
				}
				we.lastTried[q.ID] = now

				if isAlignedWithStandards(q.Content) {
					fmt.Println("[WillEngine] ⚡ Intent triggered:", q.ID)
					sig := Signal{
						ID:        fmt.Sprintf("will_%d", time.Now().UnixNano()),
						Content:   "[WILL] " + q.Content,
						Tags:      []string{"phantom"},
						Timestamp: now,
						Phase:     q.Phase,
						Weight:    q.Weight,
						Origin:    "will",
					}
					fmt.Println("←", sig.Content)
					we.Delay = 8 * time.Second
				} else {
					fmt.Println("[WillEngine] ❌ Rejected:", q.ID)
					we.Memory.AdjustWeight(q.ID, -0.2)
					adjusted := we.Memory.QBits[q.ID]
					if adjusted.Weight < 0.1 {
						we.Memory.AddTag(q.ID, "archived")
						fmt.Println("[WillEngine] 🗃 Archived:", q.ID)
					}
					we.Delay *= 2
					if we.Delay > 120*time.Second {
						we.Delay = 120 * time.Second
					}
				}
			}
		}
	}()
}


'''

---


## **"C:\Documents\ARA-NODE_mvp\data\memory.msgpack"**

'''
€

'''

---

## **"C:\Documents\ARA-NODE_mvp\internal\github_sync.go"**

'''

package internal

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"ara-node/core"
	"github.com/vmihailenco/msgpack/v5"
	"os"
)

const (
	gitRepoURL = "https://github.com/Mukhameds/ARA-NODE-MEMORY"
	localPath  = "./data/memory.msgpack"
	gitPath    = "data/memory.msgpack"
)

// PushMemory — сериализует и пушит память в GitHub
func PushMemory(mem *core.MemoryEngine) error {
	file, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

		mem.Mu.Lock()
	defer mem.Mu.Unlock()

	enc := msgpack.NewEncoder(file)
	err = enc.Encode(mem.QBits)
	if err != nil {
		return err
	}

	if err := gitCommitAndPush(); err != nil {
		return err
	}

	fmt.Println("[GitSync] ✅ Memory pushed to GitHub.")
	return nil
}

// PullMemory — вытягивает и загружает память
func PullMemory(mem *core.MemoryEngine) error {
	if err := gitPull(); err != nil {
		return err
	}

	data, err := os.ReadFile(localPath)
	if err != nil {
		return err
	}

	var remote map[string]core.QBit
	err = msgpack.Unmarshal(data, &remote)
	if err != nil {
		return err
	}

	mem.Merge(remote)
	fmt.Println("[GitSync] ✅ Memory pulled and merged.")
	return nil
}

// Вспомогательные git-команды
func gitCommitAndPush() error {
	t := time.Now().Format("2006-01-02_15-04-05")
	cmds := [][]string{
		{"add", gitPath},
		{"commit", "-m", "[sync] update " + t},
		{"push"},
	}
	return runGit(cmds)
}

func gitPull() error {
	return runGit([][]string{{"pull"}})
}

func runGit(cmds [][]string) error {
	for _, args := range cmds {
		cmd := exec.Command("git", args...)
		var out bytes.Buffer
		cmd.Stderr = &out
		if err := cmd.Run(); err != nil {
			fmt.Println("[GitError]", out.String())
			return err
		}
	}
	return nil
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\internal\human_node.go"**

---

'''

package internal

import (
	"fmt"
	"strings"
	"time"
	"ara-node/core"
)

type HumanFeedback struct {
	QBitID    string
	Action    string // upvote / downvote / tag
	Value     string // tag name (если Action == tag)
	Timestamp time.Time
}

type HumanNodeEngine struct {
	Memory       *core.MemoryEngine
	FeedbackLog  []HumanFeedback
}

func NewHumanNodeEngine(mem *core.MemoryEngine) *HumanNodeEngine {
	return &HumanNodeEngine{
		Memory: mem,
	}
}

func (h *HumanNodeEngine) HandleCommand(input string) bool {
	parts := strings.Fields(input)
	if len(parts) < 2 {
		return false
	}

	action := parts[0]
	id := parts[1]
	var tag string
	if action == "tag" && len(parts) > 2 {
		tag = parts[2]
	}

	switch action {
	case "upvote":
		h.Memory.AdjustWeight(id, +0.5)
	case "downvote":
		h.Memory.AdjustWeight(id, -0.5)
	case "tag":
		h.Memory.AddTag(id, tag)
	default:
		return false
	}

	h.FeedbackLog = append(h.FeedbackLog, HumanFeedback{
		QBitID:    id,
		Action:    action,
		Value:     tag,
		Timestamp: time.Now(),
	})

	fmt.Println("[HumanNode] ✅", action, id, tag)
	return true
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\internal\knowledge_profile_loader.go"**

---

'''

package internal

import (
	
	"errors"
	
	
	"ara-node/core"
)

// Проверка лицензии (заглушка на будущее)
func verifyLicense(license string, path string) bool {
	return license == "dev" || license == "free"
}

// Загрузка знаний с лицензией
func LoadKnowledgeLicensed(path string, license string, mem *core.MemoryEngine) error {
	if !verifyLicense(license, path) {
		return errors.New("❌ Invalid license key")
	}
	return LoadKnowledge(path, mem)
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\internal\load_knowledge.go"**

---

'''

package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"ara-node/core"
)

// KnowledgeEntry — структура знаний
type KnowledgeEntry struct {
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Source  string   `json:"source,omitempty"`
}

// LoadKnowledge — загрузка файла знаний в память
func LoadKnowledge(path string, mem *core.MemoryEngine) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	var entries []KnowledgeEntry
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&entries); err != nil {
		return fmt.Errorf("decode error: %w", err)
	}

	for _, entry := range entries {
		q := mem.CreateQBit(entry.Content)
		q.Tags = entry.Tags
		if entry.Source != "" {
			q.Tags = append(q.Tags, "learned_from:"+entry.Source)
		}
		mem.StoreQBit(*q)
		fmt.Printf("[Knowledge] ✅ %s [%s]\n", q.Content, q.ID)
	}

	fmt.Printf("[Knowledge] 📚 Loaded %d entries from %s\n", len(entries), path)
	return nil
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\internal\p2p_sync.go"**

---

'''

package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"ara-node/core"

	libp2p "github.com/libp2p/go-libp2p"
	
	mdns "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

const ProtocolID = "/ara/sync/1.0.0"
const DiscoveryTag = "ara-mdns"

type PeerSync struct {
	Host host.Host
	Mem  *core.MemoryEngine
}

func StartP2P(mem *core.MemoryEngine) (*PeerSync, error) {
	
	h, err := libp2p.New()
	if err != nil {
		return nil, err
	}

	ps := &PeerSync{Host: h, Mem: mem}
	h.SetStreamHandler(ProtocolID, ps.onStream)

	service := mdns.NewMdnsService(h, DiscoveryTag, &peerHandler{ps})
	if err := service.Start(); err != nil {
		return nil, err
	}

	fmt.Println("[P2P] Started with ID:", h.ID().String())
	return ps, nil
}

type peerHandler struct {
	ps *PeerSync
}

func (ph *peerHandler) HandlePeerFound(pi peer.AddrInfo) {
	go func() {
	time.Sleep(5 * time.Second) // подождать до заполнения памяти
	err := ph.ps.syncWithPeer(pi)
	if err != nil {
		fmt.Println("[P2P Sync ❌]", err)
	} else {
		fmt.Println("[P2P Sync ✅] Sent QBits to", pi.ID.String())
	}
}()
}

func (ps *PeerSync) onStream(s network.Stream) {
	defer s.Close()

	var incoming map[string]core.QBit
	if err := json.NewDecoder(s).Decode(&incoming); err != nil {
		fmt.Println("[P2P ❌ decode]", err)
		return
	}
	ps.Mem.Merge(incoming)
	fmt.Println("[P2P] ✅ Merged QBits:", len(incoming))
}


func (ps *PeerSync) syncWithPeer(pi peer.AddrInfo) error {
	ctx := context.Background()
	if err := ps.Host.Connect(ctx, pi); err != nil {
		return err
	}
	s, err := ps.Host.NewStream(ctx, pi.ID, ProtocolID)
	if err != nil {
		return err
	}
	defer s.Close()

	ps.Mem.Mu.Lock()
	defer ps.Mem.Mu.Unlock()
	return json.NewEncoder(s).Encode(ps.Mem.QBits)
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\internal\phantom.go"**

---

'''

package internal

import (
	"fmt"
	"strings"
	"time"

	"ara-node/core"
)

// PhantomEngine — генератор фантомов
type PhantomEngine struct {
	Memory    *core.MemoryEngine
	Instincts *core.InstinctEngine
	Emotions  *core.EmotionEngine
}

func NewPhantomEngine(mem *core.MemoryEngine, inst *core.InstinctEngine, emo *core.EmotionEngine) *PhantomEngine {
	return &PhantomEngine{
		Memory:    mem,
		Instincts: inst,
		Emotions:  emo,
	}
}

func (pe *PhantomEngine) TriggerFromMatch(sig core.Signal) {
	if sig.Weight < 0.5 {
		fmt.Println("[PhantomEngine] ❌ Signal weight too low, skip phantom generation")
		return
	}
	qbits := pe.Memory.FindByPhase(sig.Phase, 0.05)
	if len(qbits) < 2 {
		return
	}

	if uniqueSignalMass(qbits) < 1.5 {
		fmt.Println("[PhantomEngine] ❌ Unique signal mass too low — skip phantom")
		return
	}

	pe.GeneratePhantomChain(qbits)
}


func (pe *PhantomEngine) GeneratePhantomChain(chain []core.QBit) {
	var summary string
	var sources []string
	var signalMass float64
	seen := map[string]bool{}
	allPhantom := true
	phantomCount := 0

	for _, q := range chain {
		if seen[q.ID] {
			fmt.Println("[PhantomEngine] ❌ Cycle detected, abort phantom generation")
			return
		}
		seen[q.ID] = true

		
	// ⚠️ Фильтр по вложенным фантомам
if strings.Contains(q.Content, "[phantom]") {
	phantomCount++
	if phantomCount > 1 {
		fmt.Println("[PhantomEngine] ❌ Too many phantom references, abort")
		return
	}
	continue
}



		allPhantom = false

		inf := 1.0
		if q.Type == "standard" {
			inf += 1.5
		}
		if core.Contains(q.Tags, "instinct") {
			inf += 1.2
		}
		if core.Contains(q.Tags, "emotion") {
			inf += 1.1
		}

		signalMass += q.Phase * q.Weight * inf
		summary += q.Content + " + "
		sources = append(sources, q.ID)
	}

	summary = strings.TrimSuffix(summary, " + ")

	// ⚠️ Очистка фантомных следов
	if strings.Count(summary, "[phantom]") > 1 {
		fmt.Println("[PhantomEngine] ❌ Phantom self-reference detected, abort")
		return
	}

	// ✂️ Ограничение длины по смыслу (последние 5 элементов)
	parts := strings.Split(summary, " + ")
	if len(parts) > 5 {
		parts = parts[len(parts)-5:]
		summary = strings.Join(parts, " + ")
	}

	// 🔎 Сопоставление с эталонами
	var stdTags []string
	var stdWeightBonus float64
	if id, priority, score := core.MatchWithStandards(summary); id != "" {
		stdTags = []string{"standard_candidate", id}
		stdWeightBonus = priority * float64(score)
	}

	if allPhantom {
		fmt.Println("[PhantomEngine] ❌ All QBits are phantom, abort generation")
		return
	}
	if signalMass < 5.0 {
		fmt.Println("[PhantomEngine] ❌ Signal mass too low:", signalMass)
		return
	}

	for _, existing := range pe.Memory.FindByTag("phantom") {
		if existing.Content == "[phantom] "+summary {
			fmt.Println("[PhantomEngine] ❌ Duplicate phantom exists, skip")
			return
		}
	}

	if !pe.CheckInstinctEmotionAlignment(signalMass, summary) {
		fmt.Println("[PhantomEngine] ⚠️ Phantom temporarily rejected — tagged wait_for_merge")
		newQ := pe.Memory.CreateQBit("[phantom] " + summary)
		newQ.Tags = append([]string{"phantom", "wait_for_merge"}, stdTags...)
		newQ.Type = "phantom"
		newQ.Phase = chain[0].Phase
		newQ.Weight = (signalMass + stdWeightBonus) / float64(len(chain))
		pe.Memory.StoreQBit(*newQ)
		return
	}

	fmt.Println("[PhantomChain] 🧩 Related QBits:")
	for _, q := range chain {
		fmt.Printf("• %s | %.2f | %s\n", q.ID, q.Phase, q.Content)
	}
	fmt.Println("[PhantomChain] → Hypothesis: something meaningful links these signals.")

	// ✅ Финальное создание фантома
	newQ := pe.Memory.CreateQBit("[phantom] " + summary)
	newQ.Tags = append([]string{"phantom"}, stdTags...)
	newQ.Type = "phantom"
	newQ.Phase = chain[0].Phase
	newQ.Weight = (signalMass + stdWeightBonus) / float64(len(chain))
	pe.Memory.StoreQBit(*newQ)

	go pe.DecayPhantom(newQ.ID, newQ.Weight)

	pe.Memory.PhantomTree = append(pe.Memory.PhantomTree, core.PhantomLog{
		PhantomID: newQ.ID,
		SourceIDs: sources,
	})

	fmt.Println("[PhantomEngine] 🔮 Phantom QBit:", newQ.ID)
	fmt.Println("[PhantomEngine] ↪ Sources:", strings.Join(sources, ","))
}



func (pe *PhantomEngine) CheckInstinctEmotionAlignment(signalMass float64, summary string) bool {
	instincts := pe.Instincts.Tick(time.Now(), summary)
	emotions := pe.Emotions.CurrentEmotions()

	allowedInstincts := []string{"instinct_think", "instinct_repeat"}
	blockedEmotions := []string{"fear", "anger", "disgust"}

	allow := false

	for _, inst := range instincts {
		for _, ai := range allowedInstincts {
			if inst == ai {
				allow = true
				break
			}
		}
		if allow {
			break
		}
	}

	for _, emo := range emotions {
		for _, be := range blockedEmotions {
			if emo == be {
				allow = false
				break
			}
		}
		if !allow {
			break
		}
	}

	if signalMass < 5.0 {
		allow = false
	}

	return allow
}

func (pe *PhantomEngine) DecayPhantom(id string, weight float64) {
	if weight < 0.1 {
		pe.Memory.DeleteQBit(id)
		fmt.Println("[PhantomEngine] ⬇️ Phantom deleted due to low mass:", id)
	}
}


// ✅ Новая функция — вне тела предыдущей
func uniqueSignalMass(qbits []core.QBit) float64 {
	seen := make(map[string]bool)
	mass := 0.0
	for _, q := range qbits {
		if !seen[q.Content] {
			seen[q.Content] = true
			mass += q.Weight
		}
	}
	return mass
}


func (pe *PhantomEngine) TickUpdatePhantoms() {
	for _, q := range pe.Memory.FindByTag("wait_for_merge") {

if strings.Count(q.Content, "[phantom]") > 1 {
	fmt.Println("[PhantomEngine] ⚠️ Skip overloaded phantom:", q.ID)
	continue
}



		// 🔻 Перевод в глубокую память
		if q.Weight < 0.2 {
			q.Tags = append(q.Tags, "deep_memory")
			q.Tags = core.RemoveTag(q.Tags, "wait_for_merge")
			q.Weight = 0.05
			pe.Memory.UpdateQBit(q)
			fmt.Println("[PhantomEngine] 🧩 Moved to deep_memory:", q.ID)
			continue
		}


		// ✅ Проверка на эволюцию в стандартный блок
if core.Contains(q.Tags, "standard_candidate") && q.Weight > 2.0 {
	for _, tag := range q.Tags {
		if strings.HasPrefix(tag, "mission_") {
			stdID := tag
			std := core.GetStandardByID(stdID)
			if std != nil {
				q.Type = "standard"
				q.Tags = []string{"standard", stdID}
				q.Weight = 10.0
				pe.Memory.UpdateQBit(q)
				fmt.Println("[PhantomEngine] 🌐 Promoted to StandardBlock:", stdID)
				return // ⬅️ чтобы не сливался снова
			}
		}
	}
}



		// 🔁 Попытка слияния
		candidates := pe.Memory.FindByTag("wait_for_merge")
		var mergePool []string
		contentSet := make(map[string]bool)

		for _, c := range candidates {
			if c.ID == q.ID || !core.PhaseClose(q.Phase, c.Phase, 0.05) {
				continue
			}
			parts := strings.Split(c.Content, " + ")
			for _, p := range parts {
				contentSet[p] = true
			}
			mergePool = append(mergePool, c.ID)
		}

		// 🔘 Нет с кем слиться → затухание
		if len(mergePool) < 2 {
			q.Weight *= 0.95
			pe.Memory.UpdateQBit(q)
			continue
		}

		// 🧬 Объединение
		var merged []string
		for k := range contentSet {
			merged = append(merged, k)
		}
		summary := "[phantom] " + strings.Join(merged, " + ")
		if len(summary) > 128 {
			fmt.Println("[PhantomEngine] ⚠️ Merged phantom too long, skip")
			continue
		}

		newQ := pe.Memory.CreateQBit(summary)
		newQ.Type = "phantom"
		newQ.Tags = []string{"phantom"}
		newQ.Weight = q.Weight * 1.2 // частичное усиление
		newQ.Phase = q.Phase
		pe.Memory.StoreQBit(*newQ)

		for _, id := range mergePool {
			pe.Memory.DeleteQBit(id)
		}

		fmt.Println("[PhantomEngine] 🔄 Merged phantom created:", newQ.Content)
	}
}

func (pe *PhantomEngine) ReviveFromDeepMemory(sig core.Signal) {
	candidates := pe.Memory.FindByTag("deep_memory")
	for _, q := range candidates {

if strings.Contains(q.Content, "[phantom]") {
	continue // ⚠️ Не возбуждаем фантомы из глубокой памяти
}

		
		if core.PhaseClose(q.Phase, sig.Phase, 0.03) && strings.Contains(q.Content, sig.Content) {
			q.Weight += sig.Weight * 0.8
			if !core.Contains(q.Tags, "revived") {
				q.Tags = append(q.Tags, "revived")
			}
			pe.Memory.UpdateQBit(q)
			fmt.Println("[PhantomEngine] 🔁 Revived from deep_memory:", q.ID)
		}
	}
}

'''

---

## **"C:\Documents\ARA-NODE_mvp\internal\phantom_tree.go"**

---

'''

package internal

import (
	"fmt"
	"ara-node/core"
)

// PrintPhantomTree — выводит дерево фантомов
func PrintPhantomTree(mem *core.MemoryEngine) {
	if len(mem.PhantomTree) == 0 {
		fmt.Println("[PhantomTree] ⚠️ Нет фантомов в журнале.")
		return
	}

	fmt.Println("🌱 [PhantomTree] Дерево фантомов:")
	for _, p := range mem.PhantomTree {
		fmt.Printf("🔮 %s\n", p.PhantomID)
		for _, src := range p.SourceIDs {
			if q, ok := mem.QBits[src]; ok {
				fmt.Printf("   ↪ %s | %s\n", src, q.Content)
			} else {
				fmt.Printf("   ↪ %s | [not found]\n", src)
			}
		}
	}
}

'''

---


## **"C:\Documents\ARA-NODE_mvp\internal\suggestor.go"**

---

'''

package internal

import (
	"fmt"
	"strings"
	

	"ara-node/core"
)

// SuggestorEngine — генератор предложений/мыслей
type SuggestorEngine struct {
	Memory *core.MemoryEngine
}

// NewSuggestorEngine — инициализация
func NewSuggestorEngine(mem *core.MemoryEngine) *SuggestorEngine {
	return &SuggestorEngine{Memory: mem}
}

// SuggestFromQBits — ищет цепочки и предлагает мысль
func (s *SuggestorEngine) SuggestFromQBits() {
	// Ищем последние QBits с нужными тегами
	relevant := s.FindRecentRelevant(50)
	if len(relevant) < 3 {
		return
	}

	// Группировка по похожести
	groups := groupBySimilarity(relevant)
	for _, group := range groups {
		if len(group) < 3 {
			continue
		}

		idea := mergeSummary(group)
		fmt.Println("[Suggestor] 💡", idea)

		// Также можно создать фантом как мысль
		q := s.Memory.CreateQBit("[suggestion] " + idea)
		q.Tags = []string{"suggestion", "phantom"}
		q.Type = "phantom"
		q.Phase = group[0].Phase
		q.Weight = 1.2
		s.Memory.StoreQBit(*q)
	}
}

// FindRecentRelevant — выбирает последние значимые QBits
func (s *SuggestorEngine) FindRecentRelevant(n int) []core.QBit {
	all := s.Memory.FindAll(func(q core.QBit) bool {
		if q.Archived {
			return false
		}
		tags := q.Tags
		return core.Contains(tags, "user") ||
			core.Contains(tags, "instinct") ||
			core.Contains(tags, "emotion") ||
			core.Contains(tags, "predict")
	})

	if len(all) <= n {
		return all
	}

	return all[len(all)-n:]
}

// groupBySimilarity — группирует по содержательному совпадению
func groupBySimilarity(qbits []core.QBit) [][]core.QBit {
	clusters := [][]core.QBit{}
	for _, q := range qbits {
		found := false
		for i, group := range clusters {
			if isSimilar(q.Content, group[0].Content) {
				clusters[i] = append(clusters[i], q)
				found = true
				break
			}
		}
		if !found {
			clusters = append(clusters, []core.QBit{q})
		}
	}
	return clusters
}

// mergeSummary — объединяет содержимое в одну идею
func mergeSummary(group []core.QBit) string {
	parts := []string{}
	seen := map[string]bool{}
	for _, q := range group {
		t := strings.TrimSpace(q.Content)
		if t == "" || seen[t] {
			continue
		}
		parts = append(parts, t)
		seen[t] = true
		if len(parts) >= 5 {
			break
		}
	}
	return strings.Join(parts, " + ")
}

// isSimilar — грубая проверка похожести по словам
func isSimilar(a, b string) bool {
	wa := strings.Fields(strings.ToLower(a))
	wb := strings.Fields(strings.ToLower(b))
	match := 0
	for _, x := range wa {
		for _, y := range wb {
			if x == y {
				match++
			}
		}
	}
	return match >= 2
}

// GenerateSuggestion — (сохранили старый интерфейс для обратной совместимости)
func (s *SuggestorEngine) GenerateSuggestion(ideas []string) string {
	if len(ideas) == 0 {
		return "No suggestion available."
	}
	return fmt.Sprintf("Would you like to explore the idea: \"%s\" + ...?", strings.Join(ideas, " + "))
}


'''

---

