
"C:\Documents\ARA-NODE_mvp\core\instincts.go":
package core

import (
	"strings"
	"time"
)

type InstinctEngine struct {
	LastInputTime time.Time
}

func NewInstinctEngine() *InstinctEngine {
	return &InstinctEngine{
		LastInputTime: time.Now(),
	}
}

func (ie *InstinctEngine) Tick(currentTime time.Time, signal string) (instincts []string) {
	var results []string

	// 1. Silence trigger (instinct_think)
	if currentTime.Sub(ie.LastInputTime) > 10*time.Second {
		results = append(results, "instinct_think")
		ie.LastInputTime = currentTime
	}

	// 2. Repeat detection (instinct_repeat)
	if isRepeat(signal) {
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

	return results
}

// isRepeat — проверка на повтор сигнала (заглушка)
func isRepeat(signal string) bool {
	// TODO: в будущем реализовать анализ дубликатов
	return false
}

"C:\Documents\ARA-NODE_mvp\core\memory_engine.go":
package core

import (
	"fmt"
	"sync"
	"time"
)

// MemoryEngine — сигнальная память агента
type MemoryEngine struct {
	QBits map[string]QBit
	Mu    sync.Mutex
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

"C:\Documents\ARA-NODE_mvp\core\prediction.go":
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

"C:\Documents\ARA-NODE_mvp\core\QBitEvolutionEngine.go":
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

"C:\Documents\ARA-NODE_mvp\core\reflex_engine.go":
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

"C:\Documents\ARA-NODE_mvp\core\signal_dictionary.go":
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


"C:\Documents\ARA-NODE_mvp\core\signal_engine.go":
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

"C:\Documents\ARA-NODE_mvp\core\standards.go":
package core

type StandardBlock struct {
	ID       string
	Keywords []string
	Priority float64
}


// Статические эталонные блоки миссий ARA
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

"C:\Documents\ARA-NODE_mvp\core\types.go":
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


"C:\Documents\ARA-NODE_mvp\core\will_engine.go":
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