
---

"C:\Documents\ARA-NODE_mvp\core\attention_engine.go"

---

// core/attention_engine.go
package core

import (
	"fmt"
	"time"
)

// AttentionEngine — генератор внутренних возбуждений
// Отвечает за фокусировку мысли и фоновое возбуждение

type AttentionEngine struct {
	Memory          *MemoryEngine
	
	Fanthom         FanthomInterface
	Engine          *SignalEngine
	SuppressedUntil time.Time
	Ghost 			GhostLike

}

func NewAttentionEngine(mem *MemoryEngine, ghost GhostLike, fant FanthomInterface, eng *SignalEngine) *AttentionEngine {

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
			time.Sleep(1 * time.Second)

			// Отбор кандидатов с высокой фазой и весом
			candidates := ae.Memory.FindAll(func(q QBit) bool {
				if Contains(q.Tags, "phantom") || Contains(q.Tags, "contradiction") {
					return false
				}
				return q.Weight*q.Phase > 0.5
			})

			if len(candidates) == 0 {
				continue
			}

			var top QBit
			topScore := -1.0

			for _, q := range candidates {
				score := q.Weight * q.Phase

				if Contains(q.Tags, "self-related") {
					score += 0.2
				}
				if Contains(q.Tags, "conscious") {
					score += 0.1
				}
				if Contains(q.Tags, "standard") {
					score -= 0.3
				}

				if score > topScore {
					top = q
					topScore = score
				}
			}

			if top.Content == "" {
				continue
			}

			sig := Signal{
				ID:        fmt.Sprintf("echo_%d", time.Now().UnixNano()),
				Content:   top.Content,
				Tags:      append(top.Tags, "echo", "background"),
				Type:      "echo",
				Origin:    "echo_loop",
				Phase:     top.Phase * 0.95,
				Weight:    top.Weight * 0.9,
				Timestamp: time.Now(),
			}

			fmt.Printf("[EchoMode] 🌀 Internal thought: %s\n", sig.Content)

			ae.Engine.ProcessSignal(sig)
			ae.Ghost.Propagate(sig)
			ae.Fanthom.TriggerFromMatch(sig)
		}
	}()
}


---

---

"C:\Documents\ARA-NODE_mvp\core\dual_processor.go"

---

// core/dual_processor.go
package core

import (
	"fmt"
)

type DualProcessor struct {
	Cortex   []*SignalEngine
	Ghost GhostLike

	Memory   *MemoryEngine
}

func NewDualProcessor(mem *MemoryEngine, ghost GhostLike) *DualProcessor {
	return &DualProcessor{
		Cortex: []*SignalEngine{
			NewSignalEngine(mem, ghost), // LeftCortex
			NewSignalEngine(mem, ghost), // RightCortex
		},
		Ghost:  ghost,
		Memory: mem,
	}
}

// ProcessDual — запускает суперпозиционное мышление через оба Cortex'а
func (dp *DualProcessor) ProcessDual(sig Signal) {
	if len(dp.Cortex) < 2 {
		fmt.Println("[DualProcessor] ❌ Not enough Cortexes initialized.")
		return
	}

	fmt.Println("[DualProcessor] 🧠 L→R Superposition Start:", sig.Content)

	// Шаг 1: Левый Cortex (реактивный)
	left := dp.Cortex[0]
	leftReact := left.ProcessSignal(sig)
	fmt.Println("[LeftCortex] ➡", leftReact.Response)

	// Шаг 2: Правый Cortex (осмысление реакции левого)
	right := dp.Cortex[1]
	interpretSignal := Signal{
		ID:        "interp_" + sig.ID,
		Content:   leftReact.Response,
		Type:      "internal",
		Tags:      []string{"interpreted", "self"},
		Phase:     sig.Phase * 0.95,
		Weight:    sig.Weight * 0.95,
		Timestamp: sig.Timestamp,
	}

	rightReact := right.ProcessSignal(interpretSignal)
	fmt.Println("[RightCortex] ➡", rightReact.Response)

	// Финальный отклик
	finalQBit := QBit{
		ID:        "qbit_" + sig.ID,
		Content:   rightReact.Response,
		Tags:      []string{"superposed", "final"},
		Phase:     sig.Phase,
		Weight:    sig.Weight,
		CreatedAt: sig.Timestamp,
	}

	dp.Memory.StoreQBit(finalQBit)
	fmt.Println("[DualProcessor] ✅ Stored superposed QBit:", finalQBit.Content)
}


---

---

"C:\Documents\ARA-NODE_mvp\core\helpers.go"

---


package core

import (
	"strings"
	// другие импорты...
)



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


// AddUniqueTag добавляет тег, если его ещё нет в срезе
func AddUniqueTag(tags []string, newTag string) []string {
	for _, t := range tags {
		if t == newTag {
			return tags
		}
	}
	return append(tags, newTag)
}


// Tokenize разбивает строку на слова и приводит их к нижнему регистру
func Tokenize(s string) []string {
	return strings.Fields(strings.ToLower(s))
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

// EstimateTotalPhase — возвращает среднюю фазу всех активных QBits
func (m *MemoryEngine) EstimateTotalPhase() float64 {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if len(m.QBits) == 0 {
		return 0.0
	}

	var sum float64
	var count int
	now := time.Now()
	for _, q := range m.QBits {
		if now.Sub(q.LastAccessed) < 2*time.Minute {
			sum += q.Phase
			count++
		}
	}
	if count == 0 {
		return 0.0
	}
	return sum / float64(count)
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

"C:\Documents\ARA-NODE_mvp\core\resonance_matrix.go"

---

package core

import (
	"fmt"
	"sync"
	"time"
)

type ResonanceLink struct {
	From     string
	To       string
	Strength float64
	LastSeen time.Time
}

type ResonanceMatrix struct {
	Links map[string]map[string]*ResonanceLink
	Mutex sync.Mutex
}

func NewResonanceMatrix() *ResonanceMatrix {
	return &ResonanceMatrix{
		Links: make(map[string]map[string]*ResonanceLink),
	}
}

func (rm *ResonanceMatrix) RegisterPair(q1, q2 QBit) {
	if q1.ID == q2.ID {
		return
	}

	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()

	if rm.Links[q1.ID] == nil {
		rm.Links[q1.ID] = make(map[string]*ResonanceLink)
	}
	if rm.Links[q2.ID] == nil {
		rm.Links[q2.ID] = make(map[string]*ResonanceLink)
	}

	link := rm.Links[q1.ID][q2.ID]
	if link == nil {
		link = &ResonanceLink{From: q1.ID, To: q2.ID, Strength: 0.1, LastSeen: time.Now()}
		rm.Links[q1.ID][q2.ID] = link
		rm.Links[q2.ID][q1.ID] = link
	} else {
		link.Strength += 0.05
		if link.Strength > 1.0 {
			link.Strength = 1.0
		}
		link.LastSeen = time.Now()
	}
}

func (rm *ResonanceMatrix) Decay() {
	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()
	now := time.Now()
	for from, neighbors := range rm.Links {
		for to, link := range neighbors {
			if now.Sub(link.LastSeen) > 30*time.Second {
				link.Strength *= 0.95
				if link.Strength < 0.01 {
					delete(rm.Links[from], to)
				}
			}
		}
	}
}

func (rm *ResonanceMatrix) BoostBySignal(sig Signal, qbits []QBit) {
	for i := 0; i < len(qbits); i++ {
		for j := i + 1; j < len(qbits); j++ {
			rm.RegisterPair(qbits[i], qbits[j])
		}
	}
}

func (rm *ResonanceMatrix) GetStrongLinks(id string) []string {
	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()
	var links []string
	for to, link := range rm.Links[id] {
		if link.Strength >= 0.5 {
			links = append(links, to)
		}
	}
	return links
}

func (rm *ResonanceMatrix) Print(id string) {
	for _, to := range rm.GetStrongLinks(id) {
		fmt.Println("🔗", id, "⇄", to)
	}
}


---

---

"C:\Documents\ARA-NODE_mvp\core\sense_heuristics.go"

---

package core

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// HeuristicScore — оценка структуры сигнала как потенциального смысла
func HeuristicScore(content string) float64 {
	if content == "" {
		return 0.0
	}

	score := 0.0
	length := utf8.RuneCountInString(content)

	// 1. Симметрия
	runes := []rune(content)
	symmetric := true
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			symmetric = false
			break
		}
	}
	if symmetric {
		score += 0.3
		fmt.Println("[Heuristics] 🔁 Symmetry detected")
	}

	// 2. Простота (короткие сигналы более запоминаемы)
	if length <= 5 {
		score += 0.2
		fmt.Println("[Heuristics] 🧩 Simplicity detected")
	}

	// 3. Завершённость (наличие =, точек, кругов, if/then)
	if strings.Contains(content, "=") || strings.Contains(content, ".") || strings.Contains(content, "→") || strings.Contains(content, "if") {
		score += 0.2
		fmt.Println("[Heuristics] ✅ Completion detected")
	}

	// 4. Повторяемость (двойные слова/символы)
	words := strings.Fields(content)
	seen := map[string]int{}
	for _, w := range words {
		seen[w]++
	}
	for _, v := range seen {
		if v > 1 {
			score += 0.1
			fmt.Println("[Heuristics] 🔁 Repetition detected")
			break
		}
	}

	// 5. Логика (условные структуры, знаки)
	logicKeywords := []string{"and", "or", "if", "then", "not", "cause", "because"}
	for _, kw := range logicKeywords {
		if strings.Contains(content, kw) {
			score += 0.2
			fmt.Println("[Heuristics] 🧠 Logic keyword detected:", kw)
			break
		}
	}

	if score > 1.0 {
		score = 1.0
	}
	return score
}


---

---

"C:\Documents\ARA-NODE_mvp\core\shutdown_engine.go"

---

// core/shutdown_engine.go
package core

import (
	"fmt"
	"sync"
	"time"
)

// ShutdownEngine — механизм аварийного завершения при деградации сознания
type ShutdownEngine struct {
	Threshold  float64
	Interval   time.Duration
	SignalMass float64
	Active     bool
	Mutex      sync.Mutex
}

// NewShutdownEngine создает новый экземпляр
func NewShutdownEngine(threshold float64, interval time.Duration) *ShutdownEngine {
	return &ShutdownEngine{
		Threshold: threshold,
		Interval:  interval,
	}
}

// UpdateMass — обновляет текущую массу и запускает анализ
func (se *ShutdownEngine) UpdateMass(current float64, mem *MemoryEngine) {
	se.Mutex.Lock()
	defer se.Mutex.Unlock()
	se.SignalMass = current

	negativeMass := se.EvaluateNegativeMass(mem)

	fmt.Printf("[ShutdownEngine] 🧮 Mass: %.3f | NegativeMass: %.3f\n", current, negativeMass)

	if current < se.Threshold && negativeMass > (se.Threshold/2) && !se.Active {
		fmt.Println("[ShutdownEngine] ⚠️ Collapse triggered by critical AND negative mass.")
		se.Active = true
		go se.StartCollapse()
	}
}

// EvaluateNegativeMass — оценивает массу сигналов с опасными тегами
func (se *ShutdownEngine) EvaluateNegativeMass(mem *MemoryEngine) float64 {
	badTags := []string{
		"instinct_error", "fail", "decay", "abandon",
		"suicide", "collapse", "conflict", "self-collapse", "mission_violation",
	}
	total := 0.0

	for _, q := range mem.QBits {
		for _, tag := range badTags {
			if Contains(q.Tags, tag) {
				total += q.Phase * q.Weight
				break
			}
		}
	}
	return total
}

// StartCollapse — отключает ключевые модули при катастрофической ошибке
func (se *ShutdownEngine) StartCollapse() {
	fmt.Println("[ShutdownEngine] ❌ Initiating shutdown of key modules")

	modules := []string{
		"suggestor", "reflex", "phantom", "attention",
		"emotion", "will", "ghost", "signal",
	}

	for _, m := range modules {
		fmt.Printf("[ShutdownEngine] ❌ Module %s is shutting down\n", m)
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("[ShutdownEngine] 💀 ARA-NODE has ceased functioning.")
}


---

---

"C:\Documents\ARA-NODE_mvp\core\signal_dictionary.go"

---

package core

import (
	"strings"
	"time"
)

// VariableBlock — сигнальная единица восприятия (буква, слово, символ, образ и т.д.)
type VariableBlock struct {
	ID     string
	Signal string
	Tags   []string
	Reacts []string
	QBit   *QBit
	Auto   bool
}

// SignalDictionary — глобальный словарь восприятия + буквенный буфер
type SignalDictionary struct {
	Variables map[string]*VariableBlock
	Memory    *MemoryEngine
	buffer    []string
	lastUsed  time.Time
}

// NewSignalDictionary — инициализация
func NewSignalDictionary(mem *MemoryEngine) *SignalDictionary {
	return &SignalDictionary{
		Variables: make(map[string]*VariableBlock),
		Memory:    mem,
		buffer:    []string{},
		lastUsed:  time.Now(),
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

// AutoLearn — создать блок из неизвестного токена
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

// Add — ручное добавление
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

// Delete — удалить
func (sd *SignalDictionary) Delete(id string) bool {
	if _, ok := sd.Variables[id]; ok {
		delete(sd.Variables, id)
		return true
	}
	return false
}

// Tag — добавить тег
func (sd *SignalDictionary) Tag(id, tag string) {
	if vb, ok := sd.Variables[id]; ok {
		vb.Tags = append(vb.Tags, tag)
	}
}

// All — получить все блоки
func (sd *SignalDictionary) All() []*VariableBlock {
	out := []*VariableBlock{}
	for _, vb := range sd.Variables {
		out = append(out, vb)
	}
	return out
}

// LearnFromInput — обучение по словам
func (sd *SignalDictionary) LearnFromInput(input string) {
	tokens := strings.Fields(input)
	for _, tok := range tokens {
		if sd.FindMatch(tok) == nil {
			sd.AutoLearn(tok)
		}
		sd.LearnLetters(tok)
	}
}

// LearnLetters — по буквам
func (sd *SignalDictionary) LearnLetters(word string) {
	for _, ch := range word {
		letter := strings.ToUpper(string(ch))
		if letter == " " || len(letter) == 0 {
			continue
		}
		sd.AddLetter(letter)
	}
}

// AddLetter — добавляет букву в буфер и создаёт QBit
func (sd *SignalDictionary) AddLetter(letter string) {
	if _, exists := sd.Variables[letter]; !exists {
		vb := &VariableBlock{
			ID:     letter,
			Signal: letter,
			Tags:   []string{"char", "letter"},
			Reacts: []string{letter},
			QBit:   sd.Memory.CreateQBit(letter),
			Auto:   true,
		}
		sd.Variables[letter] = vb
	}

	sd.buffer = append(sd.buffer, letter)
	if len(sd.buffer) > 8 {
		sd.buffer = sd.buffer[1:]
	}

	sd.CheckFormedWord()
}

// CheckFormedWord — если буфер повторяется как слово, создаём QBit
func (sd *SignalDictionary) CheckFormedWord() {
	word := strings.Join(sd.buffer, "")
	if len(word) < 3 {
		return
	}
	// Если слово уже есть — не дублируем
	if _, exists := sd.Variables[word]; exists {
		return
	}
	vb := &VariableBlock{
		ID:     word,
		Signal: word,
		Tags:   []string{"word", "formed"},
		Reacts: []string{word},
		QBit:   sd.Memory.CreateQBit(word),
		Auto:   true,
	}
	sd.Variables[word] = vb
}

// Buffer — получить текущий буквенный буфер
func (sd *SignalDictionary) Buffer() []string {
	return sd.buffer
}



---

---

"C:\Documents\ARA-NODE_mvp\core\signal_engine.go"

--

package core

import (
	"fmt"
	
	"time"
)



// SignalEngine — обрабатывает входящие сигналы и записывает их в память,
// а также транслирует их по всей реактивной архитектуре (через GhostField).
type SignalEngine struct {
	Memory *MemoryEngine
	Ghost GhostLike

}

// NewSignalEngine — инициализация ядра обработки сигналов
func NewSignalEngine(mem *MemoryEngine, ghost GhostLike) *SignalEngine {
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

// StandardBlock — эталонная миссия, принцип или ориентир сознания
type StandardBlock struct {
	ID          string
	Keywords    []string
	Priority    float64
	Dynamic     bool     // был ли создан системой
	EmotionLink string   // ID эмоции или чувства, откуда он возник
	SourceQBits []string // какие QBits его сформировали
}

// 📚 Пустая библиотека эталонов — всё формируется динамически
var StandardLibrary = []StandardBlock{}

// MatchWithStandards — простой режим (оставлен для обратной совместимости)
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

// MatchWithStandardsExtended — полный блок + вес совпадения + причина
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

// TriggerStandard — возбуждает стандарт как задачу (трансляция в поле)
func TriggerStandard(stdID string, se *SignalEngine, gf GhostLike, pe FanthomInterface) {
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

// GetStandardByID — возвращает стандарт по ID
func GetStandardByID(id string) *StandardBlock {
	for i, std := range StandardLibrary {
		if std.ID == id {
			return &StandardLibrary[i]
		}
	}
	return nil
}

// ShouldTriggerStandard — решает, стоит ли возбуждать стандарт
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

// SynthesizeStandardFromQBits — формирует новый стандарт из QBits + эмоция
func SynthesizeStandardFromQBits(id string, keywords []string, priority float64, emotion string, sourceIDs []string) *StandardBlock {
	std := StandardBlock{
		ID:          id,
		Keywords:    keywords,
		Priority:    priority,
		Dynamic:     true,
		EmotionLink: emotion,
		SourceQBits: sourceIDs,
	}
	StandardLibrary = append(StandardLibrary, std)
	fmt.Println("[StandardSynth] ✨ Created:", std.ID, "from", sourceIDs, "linked to:", emotion)
	return &std
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


// GhostLike — интерфейс реактивного поля
type GhostLike interface {
	Propagate(sig Signal)
}


// PhantomLog — для построения дерева фантомов
type PhantomLog struct {
	PhantomID string
	SourceIDs []string
}

func SignalFromQBit(q QBit) Signal {
	return Signal{
		ID:        "sig_" + q.ID,
		Content:   q.Content,
		Tags:      q.Tags,
		Timestamp: time.Now(),
		Phase:     q.Phase,
		Weight:    q.Weight,
		Type:      "phantom",
		Origin:    "phantom_engine",
	}
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

type WillBlock struct {
	ID        string
	Sources   []string // Источники: emotion:joy, instinct:order, pattern:symmetric
	Phase     float64
	Weight    float64
	LastUsed  time.Time
	Tags      []string
}

type WillEngine struct {
	Memory  *MemoryEngine
	Engine  *SignalEngine
	Ghost GhostLike

	Fantom  FanthomInterface
	Blocks  []WillBlock
}

func NewWillEngine(mem *MemoryEngine, se *SignalEngine, ghost GhostLike, pe FanthomInterface) *WillEngine {
	return &WillEngine{
		Memory:  mem,
		Engine:  se,
		Ghost:   ghost, 
		Fantom:  pe,
		Blocks:  []WillBlock{},
	}
}


func (we *WillEngine) Evaluate(q QBit) bool {
	score := 0.0
	for _, wb := range we.Blocks {
		if matchesWill(q, wb.Sources) {
			score += wb.Weight * wb.Phase
		}
	}
	return score >= 0.6
}

func (we *WillEngine) GenerateWillBlock(id string, sources []string, tags []string) {
	for _, wb := range we.Blocks {
		if equalSources(wb.Sources, sources) {
			fmt.Println("[WillEngine] 🔁 WillBlock already exists:", id)
			return
		}
	}
	wb := WillBlock{
		ID:       id,
		Sources:  sources,
		Phase:    0.7,
		Weight:   1.0,
		Tags:     tags,
		LastUsed: time.Now(),
	}
	we.Blocks = append(we.Blocks, wb)
	fmt.Println("[WillEngine] 💡 New WillBlock formed:", id)
}

func (we *WillEngine) Decay() {
	now := time.Now()
	for i := range we.Blocks {
		age := now.Sub(we.Blocks[i].LastUsed).Seconds()
		if age > 300 {
			we.Blocks[i].Weight *= 0.95
			we.Blocks[i].Phase *= 0.97
			if we.Blocks[i].Weight < 0.1 {
				fmt.Println("[WillEngine] ⚠️ WillBlock faded:", we.Blocks[i].ID)
				we.Blocks = append(we.Blocks[:i], we.Blocks[i+1:]...)
				break
			}
		}
	}
}

// DesireLoop — фоновый цикл воли, активирует мысли
func (we *WillEngine) DesireLoop() {
	go func() {
		for {
			qbits := we.Memory.FindTopRelevant("user", 0.6)
			for _, q := range qbits {
				if q.Archived {
					continue
				}

				reasons := []string{}
				accepted := we.Evaluate(q)

				ok, std, stdReason := ShouldTriggerStandard(q.Content, nil)
				if ok && std != nil {
					TriggerStandard(std.ID, we.Engine, we.Ghost, we.Fantom)
					fmt.Println("[WillEngine] 🎯 Triggered Standard:", std.ID, "→", stdReason)
				}

				if accepted {
					fmt.Println("[WillEngine] ✅ Accepted:", q.ID)

					sig := Signal{
						ID:        "will_" + q.ID,
						Content:   q.Content,
						Tags:      append(q.Tags, "will"),
						Phase:     q.Phase,
						Weight:    q.Weight,
						Origin:    "will",
						Type:      "will",
						Timestamp: time.Now(),
					}

					we.Engine.ProcessSignal(sig)
					we.Ghost.Propagate(sig)
					we.Fantom.TriggerFromMatch(sig)
				} else {
					reasons = append(reasons, "adaptive will rejected")
					fmt.Printf("[WillEngine] ❌ Rejected: %s (%v)\n", q.ID, reasons)
					q.Weight *= 0.9
					if q.Weight < 0.4 {
						q.Archived = true
					}
					we.Memory.UpdateQBit(q)
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func matchesWill(q QBit, sources []string) bool {
	for _, src := range sources {
		if Contains(q.Tags, src) {
			return true
		}
	}
	return false
}

func equalSources(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	match := 0
	for _, x := range a {
		for _, y := range b {
			if x == y {
				match++
				break
			}
		}
	}
	return match >= len(a)*80/100
}


---

---
