
---

"C:\Documents\ARA-NODE_mvp\cmd\main.go"

---

package main

import (
	"fmt"
	"time"
	"strings"

	"ara-node/core"
	"ara-node/internal"
)

func main() {
	fmt.Println("🧠 ARA-NODE CLI started.")

	// === CORE INITIALIZATION ===
	mem := core.NewMemoryEngine()
	dict := core.NewSignalDictionary(mem)
	internal.RunBootstrap(mem, dict)

	// === MODULES ===
	emotion := internal.NewEmotionEngine(mem)
	instinct := internal.NewInstinctEngine()
	emotion.Instincts = instinct
	timeEngine := internal.NewTimeEngine()
	phantom := internal.NewPhantomEngine(mem, instinct, emotion, timeEngine)
	suggestor := internal.NewSuggestorEngine(mem)
	decay := internal.NewDecayAnalysisEngine(mem)
	prediction := internal.NewPredictionEngine(mem, nil, nil)
	reflex := core.NewReflexEngine()
	will := core.NewWillEngine(mem, nil, nil, phantom)

	// === GHOST FIELD ===
	ghost := core.NewGhostField()

	ghost.Register("emotion", core.ReactionRule{
		MatchTags: []string{"emotion", "instinct", "success", "fail"},
		MinPhase:  0.5,
		Action:    emotion.React,
	}, 200*time.Millisecond)

	ghost.Register("suggestor", core.ReactionRule{
		MatchTags: []string{"user", "phantom"},
		MinPhase:  0.6,
		Action: func(sig core.Signal) {
			go suggestor.SuggestFromQBits()
		},
	}, 1*time.Second)

	ghost.Register("phantom", core.ReactionRule{
		MatchTags: []string{"phantom", "standard", "merge"},
		MinPhase:  0.7,
		Action:    phantom.TriggerFromMatch,
	}, 200*time.Millisecond)

	ghost.Register("reflex", core.ReactionRule{
		MatchTags: []string{"danger", "reflex", "instinct_error"},
		MinPhase:  0.5,
		Action:    reflex.React,
	}, 100*time.Millisecond)

	// === SIGNAL ENGINE ===
	engine := core.NewSignalEngine(mem, ghost)
	prediction.Engine = engine
	prediction.Ghost = ghost
	will.Engine = engine
	will.Ghost = ghost

	// === BACKGROUND LOOPS ===
	internal.DefaultEmotionSet(emotion)
	core.DefaultReflexSet(reflex)
	decay.StartDecayLoop()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			prediction.Tick()
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			timeEngine.Tick()
		}
	}()

	attention := core.NewAttentionEngine(mem, ghost, phantom, engine)
	attention.StartBackgroundThinking()
	will.DesireLoop()

	// === CLI LOOP ===
		// === CLI LOOP ===
	for {
		var input string
		fmt.Print("> ")
		fmt.Scanln(&input)

		if input == "exit" || input == "quit" {
			fmt.Println("👋 Завершение работы.")
			break
		}

		if input == "view" {
			mem.ListQBits()
			continue
		}

		if input == "help" {
			fmt.Println("🆘 Команды:\n- help\n- view\n- view emotions\n- delete <qbit_id>\n- sync\n- exit")
			continue
		}

		if input == "view emotions" {
			for _, e := range emotion.CurrentEmotions() {
				fmt.Println("❤️", e)
			}
			continue
		}

		if strings.HasPrefix(input, "delete ") {
			id := strings.TrimPrefix(input, "delete ")
			mem.DeleteQBit(id)
			continue
		}

		if input == "sync" {
			fmt.Println("[Sync] 🔄 Запуск синхронизации (заглушка)...")
			// здесь позже будет: go internal.GitHubSync(mem)
			continue
		}

		// === 🔁 Инстинктивная реакция до основного сигнала
		signals := instinct.TickSignals(time.Now(), input)
		for _, sig := range signals {
			engine.ProcessSignal(sig)
			ghost.Propagate(sig)
			phantom.TriggerFromMatch(sig)
		}

		sig := core.Signal{
			ID:        fmt.Sprintf("sig_%d", time.Now().UnixNano()),
			Content:   input,
			Tags:      []string{"user"},
			Type:      "user",
			Origin:    "cli",
			Phase:     0.75,
			Weight:    1.0,
			Timestamp: time.Now(),
		}

		engine.ProcessSignal(sig)
		ghost.Propagate(sig)
		phantom.TriggerFromMatch(sig)
	}
}


---

---

"C:\Documents\ARA-NODE_mvp\core\attention_engine.go"

---

package core

import (
	"fmt"
	"math"
	"time"
)

// AttentionEngine — генератор внутренних возбуждений
// Отвечает за фокусировку мысли и фоновое возбуждение

type AttentionEngine struct {
	Memory          *MemoryEngine
	Ghost           *GhostField
	Fanthom         FanthomInterface
	Engine          *SignalEngine
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

// Suppress — временно отключает фоновое мышление (при пользовательском вводе и т.п.)
func (ae *AttentionEngine) Suppress(d time.Duration) {
	ae.SuppressedUntil = time.Now().Add(d)
	fmt.Println("[Attention] ⏸️ Suppressed for", d)
}

// StartBackgroundThinking — фоновое мышление по резонансу
func (ae *AttentionEngine) StartBackgroundThinking() {
	go func() {
		for {
			if time.Now().Before(ae.SuppressedUntil) {
				time.Sleep(200 * time.Millisecond)
				continue
			}

			best := QBit{}
			bestScore := 0.0

			candidates := ae.Memory.FindAll(func(q QBit) bool {
				if q.Archived || q.Type == "standard" || q.Type == "phantom" {
					return false
				}
				age := q.AgeFrame()
				return age == "fresh" && q.Weight*q.Phase > 0.6
			})

			for _, q := range candidates {
				score := q.Weight * q.Phase
				if score > bestScore {
					best = q
					bestScore = score
				}
			}

			if best.ID != "" && bestScore > 0.7 {
				sig := Signal{
					ID:        fmt.Sprintf("bg_%d", time.Now().UnixNano()),
					Content:   best.Content,
					Tags:      best.Tags,
					Type:      "background",
					Origin:    "internal",
					Phase:     math.Min(best.Phase+0.03, 1.0),
					Weight:    best.Weight * 0.95,
					Timestamp: time.Now(),
				}

				fmt.Println("[Attention] 🧠 Background focus on:", best.Content)
				ae.Engine.ProcessSignal(sig)
				ae.Ghost.Propagate(sig)
				ae.Fanthom.TriggerFromMatch(sig)
			}

			time.Sleep(1 * time.Second)
		}
	}()
}

---

---

"C:\Documents\ARA-NODE_mvp\core\ghost_engine.go"

---

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

---

---

"C:\Documents\ARA-NODE_mvp\core\helpers.go"

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

---

---

"C:\Documents\ARA-NODE_mvp\core\memory_engine.go"

---

package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"sync"
	"time"
)

// MemoryEngine — сигнальное хранилище ARA
type MemoryEngine struct {
	QBits       map[string]QBit
	PhantomTree []PhantomLog
	mutex       sync.Mutex
}

func NewMemoryEngine() *MemoryEngine {
	return &MemoryEngine{
		QBits:       make(map[string]QBit),
		PhantomTree: []PhantomLog{},
	}
}

func (m *MemoryEngine) CreateQBit(content string) *QBit {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	q := QBit{
		ID:           fmt.Sprintf("qbit_%d", time.Now().UnixNano()),
		Content:      content,
		CreatedAt:    time.Now(),
		LastAccessed: time.Now(),
		Phase:        0.7,
		Weight:       1.0,
		Tags:         []string{},
		Type:         "",
	}
	m.QBits[q.ID] = q
	return &q
}

func (m *MemoryEngine) StoreQBit(q QBit) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	q.LastAccessed = time.Now()
	m.QBits[q.ID] = q
}

func (m *MemoryEngine) UpdateQBit(q QBit) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	q.LastAccessed = time.Now()
	m.QBits[q.ID] = q
}

func (m *MemoryEngine) FindByTag(tag string) []QBit {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	results := []QBit{}
	for _, q := range m.QBits {
		if Contains(q.Tags, tag) {
			q.LastAccessed = time.Now()
			results = append(results, q)
		}
	}
	return results
}

func (m *MemoryEngine) FindByPhase(phase float64, tolerance float64) []QBit {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	results := []QBit{}
	for _, q := range m.QBits {
		if PhaseClose(q.Phase, phase, tolerance) {
			q.LastAccessed = time.Now()
			results = append(results, q)
		}
	}
	return results
}

func (m *MemoryEngine) FindTopRelevant(tag string, minPhase float64) []QBit {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var filtered []QBit
	for _, q := range m.QBits {
		if Contains(q.Tags, tag) && q.Phase >= minPhase {
			filtered = append(filtered, q)
		}
	}
	sort.SliceStable(filtered, func(i, j int) bool {
		return filtered[i].Weight*filtered[i].Phase > filtered[j].Weight*filtered[j].Phase
	})
	return filtered
}

func (m *MemoryEngine) FindAll(filter func(q QBit) bool) []QBit {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var all []QBit
	for _, q := range m.QBits {
		if filter(q) {
			q.LastAccessed = time.Now()
			all = append(all, q)
		}
	}
	return all
}

func (m *MemoryEngine) AdjustWeight(id string, delta float64) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	q := m.QBits[id]
	q.Weight += delta
	if q.Weight < 0 {
		q.Weight = 0
	}
	q.LastAccessed = time.Now()
	m.QBits[id] = q
}

func (m *MemoryEngine) AddTag(id, tag string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	q := m.QBits[id]
	if !Contains(q.Tags, tag) {
		q.Tags = append(q.Tags, tag)
	}
	m.QBits[id] = q
}

func (m *MemoryEngine) DeleteQBit(id string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.QBits, id)
	fmt.Println("[MemoryEngine] ❌ QBit deleted:", id)
}

func (m *MemoryEngine) DeleteByTag(tag string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	count := 0
	for id, q := range m.QBits {
		if Contains(q.Tags, tag) {
			delete(m.QBits, id)
			count++
		}
	}
	fmt.Println("[MemoryEngine] 🧹 Deleted QBits by tag:", tag, "→", count)
}

func (m *MemoryEngine) ExistsQBit(content string, phase float64, tol float64) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, q := range m.QBits {
		if q.Content == content && PhaseClose(q.Phase, phase, tol) {
			return true
		}
	}
	return false
}

func (m *MemoryEngine) DecayQBits() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	now := time.Now()
	for id, q := range m.QBits {
		if q.Type == "standard" || Contains(q.Tags, "instinct") {
			continue
		}
		age := now.Sub(q.LastAccessed).Seconds()
		decay := 1.0 - (age / 1800.0) // затухает за 30 минут
		if decay < 0.5 {
			q.Weight *= decay
			if q.Weight < 0.05 {
				delete(m.QBits, id)
				continue
			}
			m.QBits[id] = q
		}
	}
}

func (m *MemoryEngine) ListQBits() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	fmt.Println("\n🧠 Current Memory:")
	for _, q := range m.QBits {
		fmt.Printf("%s | %.2f | %s | Tags: %v\n", q.ID, q.Phase, q.Content, q.Tags)
	}
}

func (m *MemoryEngine) Merge(other *MemoryEngine) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for id, q := range other.QBits {
		if _, exists := m.QBits[id]; !exists {
			m.QBits[id] = q
		}
	}
	fmt.Println("[MemoryEngine] ✅ Merged external memory:", len(other.QBits), "entries")
}

func (m *MemoryEngine) ExportJSON(filename string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	data, err := json.MarshalIndent(m.QBits, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (m *MemoryEngine) ImportJSON(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	var imported map[string]QBit
	if err := json.Unmarshal(bytes, &imported); err != nil {
		return err
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for id, q := range imported {
		if _, exists := m.QBits[id]; !exists {
			m.QBits[id] = q
		}
	}
	fmt.Println("[MemoryEngine] 📥 Imported QBits:", len(imported))
	return nil
}

---

---

"C:\Documents\ARA-NODE_mvp\core\QBitEvolutionEngine.go"

---

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

---

---

"C:\Documents\ARA-NODE_mvp\core\reflex_engine.go"

---

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

---

---

"C:\Documents\ARA-NODE_mvp\core\signal_dictionary.go"

---

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

// LearnFromInput — разбивает строку на токены и обучает словарь
func (sd *SignalDictionary) LearnFromInput(input string) {
	tokens := strings.Fields(input)
	for _, tok := range tokens {
		if sd.FindMatch(tok) == nil {
			sd.AutoLearn(tok)
		}
	}
}

---

---

"C:\Documents\ARA-NODE_mvp\core\signal_engine.go"

---

package core

import (
	"fmt"
	
	"time"
)



// SignalEngine — обрабатывает входящие сигналы и записывает их в память,
// а также транслирует их по всей реактивной архитектуре (через GhostField).
type SignalEngine struct {
	Memory *MemoryEngine
	Ghost  *GhostField
}

// NewSignalEngine — инициализация ядра обработки сигналов
func NewSignalEngine(mem *MemoryEngine, ghost *GhostField) *SignalEngine {
	return &SignalEngine{
		Memory: mem,
		Ghost:  ghost,
	}
}

// ProcessSignal — принимает сигнал, сохраняет как QBit, запускает реакцию
func (se *SignalEngine) ProcessSignal(sig Signal) Reaction {
	fmt.Println("[SignalEngine] Received:", sig.Content)

	qbit := QBit{
		ID:           "qbit_" + sig.ID,
		Content:      sig.Content,
		Tags:         sig.Tags,
		CreatedAt:    time.Now(),
		LastAccessed: time.Now(),
		Weight:       sig.Weight,
		Phase:        sig.Phase,
		Type:         sig.Type,
		Origin:       sig.Origin,
	}
	se.Memory.StoreQBit(qbit)

	// Транслируем сигнал во всё поле
	if se.Ghost != nil {
		se.Ghost.Propagate(sig)
	}

	// Формируем реакцию по локальной памяти
	conf := 0.5
	tags := []string{"ack"}

	if sig.Phase > 0.85 {
		conf += 0.2
		tags = append(tags, "high_phase")
	}
	if sig.Origin == "will" || sig.Type == "phantom" {
		conf += 0.1
		tags = append(tags, "internal")
	}
	if matches := se.Memory.FindByPhase(sig.Phase, 0.03); len(matches) >= 2 {
		conf += 0.1
		tags = append(tags, "resonance")
	}

	return Reaction{
		TriggeredBy: sig.ID,
		Response:    "Signal dispatched to memory and network",
		Tags:        tags,
		Confidence:  conf,
	}
}


---

---

"C:\Documents\ARA-NODE_mvp\core\standards.go"

---

package core

import (
	"strings"
	"fmt"
	"time"
)

// StandardBlock — эталонная миссия или принцип
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

// 🔍 MatchWithStandards — простой режим (оставлен для обратной совместимости)
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

// 🧠 MatchWithStandardsExtended — полный блок + вес совпадения + причина
func MatchWithStandardsExtended(content string) (*StandardBlock, float64, string) {
	content = strings.ToLower(content)
	var best *StandardBlock
	bestScore := 0.0
	reason := ""

	for _, std := range StandardLibrary {
		matchCount := 0
		for _, keyword := range std.Keywords {
			if strings.Contains(content, keyword) {
				matchCount++
			}
		}
		score := float64(matchCount) * std.Priority
		if score > bestScore {
			bestScore = score
			best = &std
			reason = fmt.Sprintf("Matched %d keywords × priority %.2f", matchCount, std.Priority)
		}
	}

	if bestScore >= 2.0 {
		return best, bestScore, reason
	}
	return nil, 0.0, "No significant match"
}

// 🚀 TriggerStandard — возбуждает стандарт как задачу (трансляция в поле)
func TriggerStandard(stdID string, se *SignalEngine, gf *GhostField, pe FanthomInterface) {
	std := GetStandardByID(stdID)
	if std == nil {
		fmt.Println("[StandardTrigger] ❌ Not found:", stdID)
		return
	}

	sig := Signal{
		ID:        "std_" + std.ID,
		Content:   strings.Join(std.Keywords, " "),
		Tags:      append([]string{"standard", std.ID}, std.Keywords...),
		Phase:     std.Priority,
		Weight:    std.Priority * 1.0,
		Origin:    "standard_trigger",
		Type:      "mission",
		Timestamp: time.Now(),
	}

	se.ProcessSignal(sig)
	gf.Propagate(sig)
	pe.TriggerFromMatch(sig)

	fmt.Println("[StandardTrigger] 🚩 Broadcasted:", std.ID)
}

// 📎 GetStandardByID — возвращает стандарт по ID
func GetStandardByID(id string) *StandardBlock {
	for _, std := range StandardLibrary {
		if std.ID == id {
			return &std
		}
	}
	return nil
}

// 🔍 ShouldTriggerStandard — решает, стоит ли возбуждать стандарт
func ShouldTriggerStandard(content string, alreadyActive map[string]bool) (bool, *StandardBlock, string) {
	std, score, reason := MatchWithStandardsExtended(content)
	if std == nil || score < 2.0 {
		return false, nil, "Not strong enough match"
	}
	if alreadyActive != nil && alreadyActive[std.ID] {
		return false, std, "Already active"
	}
	return true, std, reason
}


---

---

"C:\Documents\ARA-NODE_mvp\core\types.go"

---

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
	Origin    string  // источник: user, instinct, prediction, background
	Type      string  // тип сигнала: user, instinct, prediction, background, etc.
}

// Strength — сила сигнала (для WillEngine, FanthomEngine и др.)
func (s *Signal) Strength() float64 {
	return s.Phase * s.Weight
}

// HasTag — проверка наличия тега
func (s *Signal) HasTag(tag string) bool {
	for _, t := range s.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// QBit — единица памяти (узел в памяти ARA)
type QBit struct {
	ID           string
	Content      string
	Tags         []string
	Type         string  // тип узла: standard, reflex, emotion, etc.
	Phase        float64
	Weight       float64
	Archived     bool
	Origin       string
	CreatedAt    time.Time
	LastAccessed time.Time
}

// Strength — сила узла в памяти
func (q *QBit) Strength() float64 {
	return q.Phase * q.Weight
}

// HasTag — проверка наличия тега
func (q *QBit) HasTag(tag string) bool {
	for _, t := range q.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// AgeFrame — возраст узла (семантический)
func (q *QBit) AgeFrame() string {
	age := time.Since(q.CreatedAt).Seconds()
	switch {
	case age < 60:
		return "emergent"
	case age < 600:
		return "forming"
	case age < 3600:
		return "mature"
	default:
		return "legacy"
	}
}

// DecayFactor — коэффициент затухания по возрасту
func (q *QBit) DecayFactor() float64 {
	age := time.Since(q.CreatedAt).Seconds()
	switch {
	case age < 60:
		return 1.0
	case age < 600:
		return 0.9
	case age < 3600:
		return 0.7
	default:
		return 0.4
	}
}

// Reaction — результат обработки сигнала
type Reaction struct {
	TriggeredBy string
	Response    string
	Tags        []string
	Confidence  float64
}

// FanthomInterface — интерфейс для фантомных генераторов
type FanthomInterface interface {
	TriggerFromMatch(sig Signal)
}

// PhantomLog — для построения дерева фантомов
type PhantomLog struct {
	PhantomID string
	SourceIDs []string
}

---

---

"C:\Documents\ARA-NODE_mvp\core\will_engine.go"

---

package core

import (
	"fmt"
	"time"
)

// WillEngine — движок воли агента
// Отслеживает сигналы пользователя и проверяет соответствие миссии
// Генерирует волевые сигналы при совпадении с эталонами

type WillEngine struct {
	Memory *MemoryEngine
	Engine *SignalEngine
	Ghost  *GhostField
	Fantom FanthomInterface
}

func NewWillEngine(mem *MemoryEngine, eng *SignalEngine, gf *GhostField, f FanthomInterface) *WillEngine {
	return &WillEngine{
		Memory: mem,
		Engine: eng,
		Ghost:  gf,
		Fantom: f,
	}
}

func (we *WillEngine) DesireLoop() {
	go func() {
		delay := 5 * time.Second
		for {
			candidates := we.Memory.FindByTag("user")
			for _, q := range candidates {
				reasons := []string{}
				alignment := isAlignedWithStandards(q.Content)

				if q.Phase < 0.85 {
					reasons = append(reasons, "phase < 0.85")
				}
				if q.AgeFrame() == "emergent" || q.AgeFrame() == "legacy" {
					reasons = append(reasons, "age not stable")
				}
				if !alignment {
					reasons = append(reasons, "not aligned with standards")
				}

				// 🔁 Проверка: возбуждать ли стандарт
				ok, std, stdReason := ShouldTriggerStandard(q.Content, nil)
				if ok && std != nil {
					TriggerStandard(std.ID, we.Engine, we.Ghost, we.Fantom)
					fmt.Println("[WillEngine] 🎯 Triggered Standard:", std.ID, "→", stdReason)
				}

				// ✅ Если нет причин отказа — запускаем волю
				if len(reasons) == 0 {
					fmt.Println("[WillEngine] ✅ Accepted:", q.ID)

					sig := Signal{
						ID:        fmt.Sprintf("will_%d", time.Now().UnixNano()),
						Content:   q.Content,
						Tags:      append(q.Tags, "will", "intent"),
						Type:      "will",
						Origin:    "internal",
						Phase:     q.Phase,
						Weight:    q.Weight,
						Timestamp: time.Now(),
					}

					we.Engine.ProcessSignal(sig)
					we.Ghost.Propagate(sig)
					we.Fantom.TriggerFromMatch(sig)
				} else {
					fmt.Printf("[WillEngine] ❌ Rejected: %s (%s)\n", q.ID, reasons)
					q.Weight *= 0.9
					if q.Weight < 0.4 {
						q.Archived = true
					}
					we.Memory.UpdateQBit(q)
				}
			}
			time.Sleep(delay)
		}
	}()
}

func isAlignedWithStandards(content string) bool {
	id, _, score := MatchWithStandards(content)
	return id != "" && score >= 3
}

---

---
