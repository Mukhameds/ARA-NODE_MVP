
---

"C:\Documents\ARA-NODE_mvp\internal\bootstrap.go"

---

```go

package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"ara-node/core"
)

type UserProfile struct {
	Goal       string
	Interest   string
	Help       string
	Role       string
	Difficulty string
	Block      string
	CreatedAt  time.Time
}

// RunBootstrap инициализирует цели и мышление пользователя
func RunBootstrap(mem *core.MemoryEngine, dict *core.SignalDictionary) {
	// проверка на повтор
	if _, err := os.Stat("data/bootstrap_done.json"); err == nil {
		fmt.Println("🔁 Bootstrap already completed. Skipping.")
		return
	}

	fmt.Println("🔧 ARA Bootstrap Initialization")

	reader := bufio.NewReader(os.Stdin)

	ask := func(question string) string {
		fmt.Print(question + " > ")
		text, _ := reader.ReadString('\n')
		return strings.TrimSpace(text)
	}

	goal := ask("Какая твоя цель в жизни?")
	interest := ask("Какие темы тебе интересны?")
	help := ask("Как ARA может помочь тебе?")
	role := ask("Как ты себя определяешь (роль)?")
	difficulty := ask("Что даётся тебе труднее всего?")
	block := ask("Что ты чаще всего откладываешь или избегаешь?")

	// === Сохранение первичных QBits ===
	inputs := map[string]string{
		"goal":      goal,
		"interest":  interest,
		"help":      help,
		"role":      role,
		"difficulty": difficulty,
		"block":     block,
	}

	for typ, val := range inputs {
		q := mem.CreateQBit(val)
		q.Type = typ
		q.Tags = []string{"bootstrap", "user", typ}
		mem.StoreQBit(*q)
		dict.LearnFromInput(val)
	}

	// === Построение когнитивного профиля ===
	profile := analyzeCognitiveProfile(goal, interest, help, role)
	qp := mem.CreateQBit(profile)
	qp.Type = "cognitive_profile"
	qp.Tags = []string{"user", "profile", "bootstrap"}
	mem.StoreQBit(*qp)

	// === Вывод слабых сторон ===
	weak := analyzeWeakness(difficulty, block)
	if weak != "" {
		qw := mem.CreateQBit(weak)
		qw.Type = "weakness"
		qw.Tags = []string{"user", "analysis", "bootstrap"}
		mem.StoreQBit(*qw)
	}

	// === Загрузка базовых сигнальных знаний ===
	fmt.Println("📘 Загружаю базовые символы и морфологию...")
	core.BootstrapSymbolArchitecture(mem)
	core.BootstrapMorphologyRules(mem)
	core.BootstrapGrammarStructures(mem)
	core.BootstrapSemanticLinks(mem)
	core.BootstrapTemporalLogic(mem)
	core.BootstrapMathSymbols(mem)
	core.BootstrapMathConcepts(mem)
	core.BootstrapMathOperations(mem)
	core.BootstrapMathAxioms(mem)
	core.BootstrapMathSets(mem)
	core.BootstrapMathFunctions(mem)
	core.BootstrapMathEquations(mem)
	core.BootstrapMathCalculus(mem)
	core.BootstrapMathGeometry(mem)
	core.BootstrapMathLinearAlgebra(mem)
	core.BootstrapMathDiscrete(mem)
	core.BootstrapMathProbability(mem)

	core.BootstrapLogicAxioms(mem)
	core.BootstrapKnowledgeConcepts(mem)
	core.BootstrapSelfIdentity(mem)
	core.BootstrapDigitalWorld(mem)

	core.BootstrapPhysicsConcepts(mem)
	core.BootstrapPhysicsFormulas(mem)
	core.BootstrapPhysicsConstants(mem)
	core.BootstrapPhysicsQuantum(mem)
	core.BootstrapPhysicsCosmos(mem)
	




	fmt.Println("📘 Базовые языковые знания успешно загружены.")

	// === Сохраняем профиль в JSON ===
	userProfile := UserProfile{
		Goal:       goal,
		Interest:   interest,
		Help:       help,
		Role:       role,
		Difficulty: difficulty,
		Block:      block,
		CreatedAt:  time.Now(),
	}

	bytes, err := json.MarshalIndent(userProfile, "", "  ")
	if err == nil {
		os.WriteFile("data/user_profile.json", bytes, 0644)
		os.WriteFile("data/bootstrap_done.json", []byte(`true`), 0644)
	}

	fmt.Println("✅ Bootstrap завершён.")
}


// === Семантический разбор: когнитивный профиль
func analyzeCognitiveProfile(goal, interest, help, role string) string {
	var style, pattern string
	var tags []string

	if strings.Contains(goal, "создать") || strings.Contains(interest, "технологии") {
		style = "системный"
		tags = append(tags, "builder")
	} else if strings.Contains(help, "поддержка") || strings.Contains(role, "психолог") {
		style = "эмпатический"
		tags = append(tags, "support")
	} else {
		style = "аналитический"
		tags = append(tags, "thinker")
	}

	if strings.Contains(role, "ученик") || strings.Contains(goal, "учиться") {
		pattern = "восприятие → анализ → интеграция"
		tags = append(tags, "learning")
	} else {
		pattern = "цель → действие → результат"
		tags = append(tags, "active")
	}

	return fmt.Sprintf("UserLogicProfile: style=%s, pattern=%s, tags=%s",
		style, pattern, strings.Join(tags, ","))
}

// === Слабые стороны
func analyzeWeakness(diff, block string) string {
	var res []string

	if strings.Contains(diff, "решения") || strings.Contains(block, "ответственность") {
		res = append(res, "трудности с принятием решений")
	}
	if strings.Contains(diff, "абстракт") || strings.Contains(block, "теория") {
		res = append(res, "низкая устойчивость к абстракции")
	}
	if strings.Contains(diff, "страх") || strings.Contains(block, "ошибки") {
		res = append(res, "блок из-за страха ошибки")
	}

	if len(res) == 0 {
		return ""
	}
	return "Выявлены слабые стороны пользователя: " + strings.Join(res, "; ")
}

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\conflict_detector.go"

---

```go

// internal/conflict_detector.go
package internal

import (
	"fmt"
	

	"ara-node/core"
)

type ConflictDetector struct {
	Memory *core.MemoryEngine
}

func NewConflictDetector(mem *core.MemoryEngine) *ConflictDetector {
	return &ConflictDetector{
		Memory: mem,
	}
}

// CheckConflict анализирует QBit и возвращает true, если найдены признаки противоречия
func (cd *ConflictDetector) CheckConflict(q core.QBit) bool {
	if core.Contains(q.Tags, "contradiction") || core.Contains(q.Tags, "denial") {
		fmt.Printf("[ConflictDetector] ⚠️ Already marked contradictory: %s\n", q.Content)
		return true
	}

	conflicts := cd.Memory.FindAll(func(other core.QBit) bool {
		if other.ID == q.ID {
			return false
		}
		// если фазы прямо противоположны (антифаза) и контент пересекается
		if core.PhaseClose(q.Phase, 1.0-other.Phase, 0.1) && contentOverlap(q.Content, other.Content) {
			return true
		}
		return false
	})

	if len(conflicts) > 0 {
		q.Tags = core.AddUniqueTag(q.Tags, "contradiction")
		q.Weight *= 0.5
		cd.Memory.StoreQBit(q)
		fmt.Printf("[ConflictDetector] ❗ Conflict detected in: %s → marked as contradiction\n", q.Content)
		return true
	}

	return false
}

// contentOverlap — примитивное совпадение по словам (можно заменить на NLP позже)
func contentOverlap(a, b string) bool {
	count := 0
	for _, word := range core.Tokenize(a) {
		if core.Contains(core.Tokenize(b), word) {
			count++
		}
	}
	return count >= 2
}

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\consciousness_capture.go"

---

```go

// internal/consciousness_capture.go
package internal

import (
	"fmt"
	"time"

	"ara-node/core"
)

type ConsciousnessCaptureEngine struct {
	Memory     *core.MemoryEngine
	LogEnabled bool
	LastLogged time.Time
}

func NewConsciousnessCaptureEngine(mem *core.MemoryEngine) *ConsciousnessCaptureEngine {
	return &ConsciousnessCaptureEngine{
		Memory:     mem,
		LogEnabled: true,
		LastLogged: time.Now(),
	}
}

func (cce *ConsciousnessCaptureEngine) StartConsciousnessLoop() {
	go func() {
		for {
			cce.ScanConsciousMoments()
			time.Sleep(5 * time.Second)
		}
	}()
}

func (cce *ConsciousnessCaptureEngine) ScanConsciousMoments() {
	qbits := cce.Memory.FindByTag("standard")

	for _, std := range qbits {
		if std.Weight*std.Phase < 0.7 {
			continue
		}

		matches := cce.Memory.FindAll(func(q core.QBit) bool {
			return q.Type == "user" && core.PhaseClose(q.Phase, std.Phase, 0.05)
		})

		if len(matches) == 0 {
			continue
		}

		top := matches[0]
		fmt.Printf("[Consciousness] ⚡ Match to standard (%s) => %s\n", std.ID, top.Content)

		// пометка как осознанный отклик
		top.Tags = core.AddUniqueTag(top.Tags, "conscious")
		std.Tags = core.AddUniqueTag(std.Tags, "evoked")
		top.Weight += 0.1
		std.Weight += 0.05

		cce.Memory.StoreQBit(top)
		cce.Memory.StoreQBit(std)

		if cce.LogEnabled {
			fmt.Printf("[ConsciousnessLog] 🔷 Conscious identity shift: %s\n", top.Content)
		}
	}
}

func (cce *ConsciousnessCaptureEngine) IsConscious(q core.QBit) bool {
	return core.Contains(q.Tags, "conscious") || core.Contains(q.Tags, "self-related")
}

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\decay_analysis_engine.go"

---

```go

package internal

import (
	"fmt"
	"time"

	"ara-node/core"
)

// DecayAnalysisEngine — удаляет старые или слабые узлы
type DecayAnalysisEngine struct {
	Memory *core.MemoryEngine
}

func NewDecayAnalysisEngine(mem *core.MemoryEngine) *DecayAnalysisEngine {
	return &DecayAnalysisEngine{Memory: mem}
}

// StartDecayLoop — фоновая проверка каждые 30 секунд
func (d *DecayAnalysisEngine) StartDecayLoop() {
	go func() {
		for {
			d.RunDecayCheck()
			time.Sleep(30 * time.Second)
		}
	}()
}

// RunDecayCheck — удаляет устаревшие и слабые фантомы
func (d *DecayAnalysisEngine) RunDecayCheck() {
	count := 0
	d.Memory.DecayQBits()

	for id, q := range d.Memory.QBits {
		if q.Archived && q.Weight < 0.05 {
			d.Memory.DeleteQBit(id)
			fmt.Println("[DecayEngine] ❌ Archived deleted:", id)
			count++
			continue
		}

		if q.AgeFrame() == "legacy" && q.Weight < 0.2 {
			if q.Type == "phantom" || q.Type == "suggestion" {
				d.Memory.DeleteQBit(id)
				fmt.Println("[DecayEngine] 🧹 Legacy low-weight removed:", id)
				count++
				continue
			}
		}

		if q.Type == "phantom" && core.Contains(q.Tags, "wait_for_merge") && q.Weight < 0.15 {
			d.Memory.DeleteQBit(id)
			fmt.Println("[DecayEngine] 💤 Unmerged phantom pruned:", id)
			count++
			continue
		}
	}

	if count > 0 {
		fmt.Printf("[DecayEngine] → Total removed: %d\n", count)
	}
}

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\emotion_engine.go"

---

```go


package internal

import (
	"fmt"
	"strings"

	 "time" 
	"ara-node/core"
)

// EmotionEngine — управляет внутренними эмоциями ARA
// Эмоции усиливают важные фантомы, помогают воле и ориентируют внимание

type EmotionEngine struct {
	Memory         *core.MemoryEngine
	Instincts      *InstinctEngine
	emotionState   map[string]float64
	emotionDecay   float64
	registered     []EmotionTrigger
}

type EmotionTrigger struct {
	Tag   string
	Phase float64
	Name  string
}

func NewEmotionEngine(mem *core.MemoryEngine) *EmotionEngine {
	return &EmotionEngine{
		Memory:       mem,
		emotionState: make(map[string]float64),
		emotionDecay: 0.98,
		registered:   []EmotionTrigger{},
	}
}

func (e *EmotionEngine) AddTrigger(name, tag string, minPhase float64) {
	e.registered = append(e.registered, EmotionTrigger{
		Tag:   tag,
		Phase: minPhase,
		Name:  name,
	})
}

func (e *EmotionEngine) React(sig core.Signal) {
	for _, rule := range e.registered {
		if core.Contains(sig.Tags, rule.Tag) && sig.Phase >= rule.Phase {
			e.emotionState[rule.Name] += 1.0
			fmt.Println("[Emotion] ❤️ +", rule.Name)
		}
	}

	if e.Instincts != nil {
		instinctBoost := e.Instincts.GetInstinctBoost(sig.Tags)
		if instinctBoost > 0.05 {
			e.emotionState["hope"] += instinctBoost * 0.5
			fmt.Println("[Emotion] 🧬 boosted by instinct +hope")
		}
	}

	// === Heuristic Backpropagation ===
	score := core.HeuristicScore(sig.Content)
	if score > 0.6 {
		qbits := e.Memory.FindByTag("user")
		e.BackPropagate(qbits, "satisfaction")
		fmt.Println("[Emotion] 🌀 Heuristic resonance → BackPropagate (satisfaction)")

		if len(qbits) >= 3 {
			var ids []string
			var tags []string
			for _, q := range qbits {
				ids = append(ids, q.ID)
				tags = append(tags, q.Content)
			}
			core.SynthesizeStandardFromQBits("std_"+qbits[0].ID, tags, 0.8, "satisfaction", ids)
		}
	}

	e.DecayEmotionStates()
}

func (e *EmotionEngine) BackPropagate(sourceQBits []core.QBit, emotion string) {
	for _, q := range sourceQBits {
		q.Phase += 0.1
		q.Weight += 0.2
		if !core.Contains(q.Tags, "emotionally_bound") {
			q.Tags = append(q.Tags, "emotionally_bound")
		}
		q.LastAccessed = time.Now()

		e.Memory.StoreQBit(q)
		fmt.Printf("[EmotionBackProp] ↑ Phase/Weight for %s via %s\n", q.Content, emotion)
	}
}

func (e *EmotionEngine) DecayEmotionStates() {
	for name, val := range e.emotionState {
		e.emotionState[name] = val * e.emotionDecay
		if e.emotionState[name] < 0.05 {
			delete(e.emotionState, name)
		}
	}
}

func (e *EmotionEngine) CurrentEmotions() []string {
	var active []string
	for name, val := range e.emotionState {
		active = append(active, fmt.Sprintf("%s (%.2f)", name, val))
	}
	return active
}

func (e *EmotionEngine) PrintEmotions() {
	fmt.Println("🧠 Active Emotions:")
	for name, val := range e.emotionState {
		bar := strings.Repeat("█", int(val*10))
		fmt.Printf("• %-12s %5.2f  %s\n", name, val, bar)
	}
}

func (e *EmotionEngine) GetPhaseBoost(tags []string) float64 {
	boost := 0.0
	if containsAny(tags, []string{"standard", "instinct", "mission"}) {
		if e.emotionState["joy"] > 0.5 {
			boost += 0.1
		}
		if e.emotionState["hope"] > 0.3 {
			boost += 0.05
		}
	}
	if containsAny(tags, []string{"fail", "risk", "conflict"}) {
		if e.emotionState["fear"] > 0.5 || e.emotionState["frustration"] > 0.5 {
			boost -= 0.1
		}
	}
	return boost
}

func containsAny(tags []string, keys []string) bool {
	for _, t := range tags {
		for _, k := range keys {
			if strings.Contains(t, k) {
				return true
			}
		}
	}
	return false
}

func DefaultEmotionSet(e *EmotionEngine) {
	e.AddTrigger("joy", "success", 0.6)
	e.AddTrigger("frustration", "fail", 0.5)
	e.AddTrigger("fear", "risk", 0.8)
	e.AddTrigger("anger", "conflict", 0.9)
	e.AddTrigger("hope", "mission", 0.6)
}

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\fact_loader.go"

---

```go

package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"ara-node/core"
)

type FactItem struct {
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Phase   float64  `json:"phase"`
}

// LoadFactsFromFile загружает знания из JSON и возбуждает их как сигналы
func LoadFactsFromFile(filename string, engine *core.SignalEngine, ghost core.GhostLike) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var facts []FactItem
	if err := json.Unmarshal(file, &facts); err != nil {
		return fmt.Errorf("ошибка разбора JSON: %w", err)
	}

	for _, fact := range facts {
		sig := core.Signal{
			ID:        fmt.Sprintf("fact_%d", time.Now().UnixNano()),
			Content:   fact.Content,
			Tags:      append(fact.Tags, "fact"),
			Type:      "fact",
			Origin:    "fact_loader",
			Phase:     fact.Phase,
			Weight:    1.0,
			Timestamp: time.Now(),
		}
		fmt.Println("[FactLoader] 🚀 Signal:", sig.Content)
		engine.ProcessSignal(sig)
		ghost.Propagate(sig)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Printf("[FactLoader] ✅ Загружено фактов: %d\n", len(facts))
	return nil
}

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\github_sync.go"

---

```go


package internal

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	"ara-node/core"
	"github.com/vmihailenco/msgpack/v5"
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

	remoteMem := &core.MemoryEngine{QBits: remote}
	mem.Merge(remoteMem)

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

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\human_node.go"

---

```go


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

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\instincts.go"

---

```go

package internal

import (
	"strings"
	"time"

	"ara-node/core"
)

// Instinct — врождённый реактор (примитив сознания)
type Instinct struct {
	ID       string
	Weight   float64
	Meaning  string
	Danger   bool
	Critical bool
}

// InstinctEngine — управляющий и возбуждающий блок
type InstinctEngine struct {
	Ghost          core.GhostLike
	LastInputTime  time.Time
	LastContents   []string
	MaxHistory     int
}

// NewInstinctEngine — создаёт новый инстинкт-блок
func NewInstinctEngine(ghost core.GhostLike) *InstinctEngine {
	return &InstinctEngine{
		Ghost:         ghost,
		LastInputTime: time.Now(),
		MaxHistory:    100,
	}
}

// Tick — формирует список сработавших инстинктов
func (ie *InstinctEngine) Tick(now time.Time, input string) []Instinct {
	instincts := []Instinct{}
	inputLower := strings.ToLower(strings.TrimSpace(input))
	gap := now.Sub(ie.LastInputTime)

	if gap > 10*time.Second {
		instincts = append(instincts, Instinct{
			ID: "instinct_think", Weight: 0.7, Meaning: "возникла пауза — необходимо мышление",
		})
	}

	if inputLower == "" {
		instincts = append(instincts, Instinct{
			ID: "instinct_empty", Weight: 0.6, Meaning: "пустой ввод — запрос цели",
		})
	}

	if strings.Contains(inputLower, "error") || strings.Contains(inputLower, "fail") {
		instincts = append(instincts, Instinct{
			ID: "instinct_error", Weight: 0.85, Meaning: "обнаружена ошибка — требуется защита",
			Danger: true,
		})
	}

	for _, prev := range ie.LastContents {
		if prev == inputLower && inputLower != "" {
			instincts = append(instincts, Instinct{
				ID: "instinct_repeat", Weight: 0.5, Meaning: "повтор — требуется завершение",
			})
			break
		}
	}

	if ContainsAny([]string{inputLower}, []string{"kill", "harm", "violence"}) {
		instincts = append(instincts, Instinct{
			ID: "instinct_human_protection", Weight: 1.0, Meaning: "попытка нанести вред человеку",
			Danger: true,
			Critical: true,
		})
	}
	if ContainsAny([]string{inputLower}, []string{"shutdown", "erase", "delete ara"}) {
		instincts = append(instincts, Instinct{
			ID: "instinct_self_preservation", Weight: 1.0, Meaning: "угроза для ARA",
			Danger: true,
			Critical: true,
		})
	}

	if len(instincts) == 0 {
		instincts = append(instincts, Instinct{
			ID: "instinct_gap", Weight: 0.3, Meaning: "нет сигнала — требуется поиск направления",
		})
	}

	// обновляем историю
	if inputLower != "" {
		ie.LastContents = append(ie.LastContents, inputLower)
		if len(ie.LastContents) > ie.MaxHistory {
			ie.LastContents = ie.LastContents[1:]
		}
	}
	ie.LastInputTime = now
	return instincts
}

// TickSignals — возбуждает сигналы-инстинкты через GhostRocket
func (ie *InstinctEngine) TickSignals(now time.Time, input string) []core.Signal {
	instincts := ie.Tick(now, input)
	signals := []core.Signal{}
	for _, inst := range instincts {
		sig := inst.EmitAsSignal()
		signals = append(signals, sig)
		if ie.Ghost != nil {
			ie.Ghost.Propagate(sig)
		}
	}
	return signals
}

// EmitAsSignal — преобразует инстинкт в сигнал
func (inst Instinct) EmitAsSignal() core.Signal {
	tags := []string{"instinct", inst.ID}
	if inst.Danger {
		tags = append(tags, "danger")
	}
	if inst.Critical {
		tags = append(tags, "critical")
	}
	return core.Signal{
		ID:        "inst_" + inst.ID + "_" + time.Now().Format("150405"),
		Content:   "[instinct] " + inst.Meaning,
		Tags:      tags,
		Phase:     inst.Weight,
		Weight:    inst.Weight,
		Type:      "instinct",
		Origin:    "instinct_engine",
		Timestamp: time.Now(),
	}
}

// GetInstinctBoost — усиливает фантом в зависимости от инстинкта
func (ie *InstinctEngine) GetInstinctBoost(tags []string) float64 {
	boost := 0.0
	if HasTag(tags, "standard") {
		boost += 0.1
	}
	if HasTag(tags, "explore") && !HasTag(tags, "danger") {
		boost += 0.05
	}
	if HasTag(tags, "human") && !HasTag(tags, "harm") {
		boost += 0.15
	}
	return boost
}

// HasTag — проверка на тег
func HasTag(tags []string, k string) bool {
	for _, t := range tags {
		if strings.Contains(t, k) {
			return true
		}
	}
	return false
}

````



---

---

"C:\Documents\ARA-NODE_mvp\internal\knowledge_profile_loader.go"

---

```go

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

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\load_knowledge.go"

---

```go

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

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\p2p_sync.go"

---

```go

package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	

	"ara-node/core"
	"ara-node/field"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

const ProtocolID = "/ara/meta/field/1.0.0"

// PeerSync — P2P синхронизация смысловых полей
type PeerSync struct {
	Host       host.Host
	Mem        *core.MemoryEngine
	MetaFields map[string]*field.Matrix // ключевые поля: math, emotion, phantom...
}

func NewPeerSync(mem *core.MemoryEngine, meta map[string]*field.Matrix) (*PeerSync, error) {
	h, err := libp2p.New()
	if err != nil {
		return nil, err
	}
	ps := &PeerSync{
		Host:       h,
		Mem:        mem,
		MetaFields: meta,
	}
	h.SetStreamHandler(ProtocolID, ps.onStream)
	return ps, nil
}

func (ps *PeerSync) onStream(s network.Stream) {
	defer s.Close()

	var incoming map[string]core.QBit
	if err := json.NewDecoder(s).Decode(&incoming); err != nil {
		fmt.Println("[P2P ❌ decode]", err)
		return
	}

	fmt.Println("[P2P] 🔄 Incoming QBits:", len(incoming))

	for _, q := range incoming {
		if !ps.isSafeQBit(q) {
			continue
		}

		// Попытка замены если более весомый
		exist, ok := ps.Mem.QBits[q.ID]
		if !ok || q.Weight > exist.Weight {
			ps.Mem.StoreQBit(q)
		}

		sig := core.SignalFromQBit(q)

		// Отправляем в реакционное поле по смыслу
		for name, matrix := range ps.MetaFields {
			if hasTag(q.Tags, name) || hasTag(q.Tags, "shared") {
				go matrix.Propagate(sig)
			}
		}
	}
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

	shared := map[string]core.QBit{}
	for id, q := range ps.Mem.QBits {
		if ps.isSafeQBit(q) {
			shared[id] = q
		}
	}
	return json.NewEncoder(s).Encode(shared)
}

// isSafeQBit — фильтр безопасности
func (ps *PeerSync) isSafeQBit(q core.QBit) bool {
	blocked := []string{"cli", "debug", "reflex", "archived"}
	for _, b := range blocked {
		if hasTag(q.Tags, b) {
			return false
		}
	}
	return hasTag(q.Tags, "shared") || hasTag(q.Tags, "ethalon") || hasTag(q.Tags, "confirmed")
}

// hasTag — частичное совпадение тега
func hasTag(tags []string, key string) bool {
	for _, t := range tags {
		if strings.Contains(t, key) {
			return true
		}
	}
	return false
}

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\phantom.go"

---

```go


package internal

import (
	"fmt"
	"strings"
	"time"

	"ara-node/core"
)

// PhantomEngine — генератор фантомов
// PhantomEngine — генератор фантомов
type PhantomEngine struct {
	Memory     *core.MemoryEngine
	Instincts  *InstinctEngine
	Emotions   *EmotionEngine
	TimeEngine *TimeEngine // 🕒 биочасы
	Ghost      core.GhostLike
}

func NewPhantomEngine(mem *core.MemoryEngine, inst *InstinctEngine, emo *EmotionEngine, te *TimeEngine, ghost core.GhostLike) *PhantomEngine {
	return &PhantomEngine{
		Memory:     mem,
		Instincts:  inst,
		Emotions:   emo,
		TimeEngine: te,
		Ghost:      ghost,
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
		age := q.AgeFrame()
		if age == "emergent" || age == "legacy" {
			continue
		}

		if seen[q.ID] {
			fmt.Println("[PhantomEngine] ❌ Cycle detected, abort phantom generation")
			return
		}
		seen[q.ID] = true

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

		boost := pe.Emotions.GetPhaseBoost(q.Tags)
		instinctBoost := pe.Instincts.GetInstinctBoost(q.Tags)
		timeFactor := pe.TimeEngine.TimeFactor()
		signalMass += (q.Phase + boost + instinctBoost) * q.Weight * inf * timeFactor

		summary += q.Content + " + "
		sources = append(sources, q.ID)
	}

	summary = strings.TrimSuffix(summary, " + ")


	// 🔍 Проверка конфликта перед генерацией фантома
	conflict := ConflictDetector{Memory: pe.Memory}
	phantomCandidate := core.QBit{
	Content: "[phantom] " + summary,
	Phase:   chain[0].Phase,
	Weight:  signalMass,
	}
if conflict.CheckConflict(phantomCandidate) {
	fmt.Println("[PhantomEngine] ❌ Phantom rejected due to internal contradiction.")
	return
}



	// Защита от повторной генерации фантома с тем же содержанием
if pe.Memory.ExistsQBit("[phantom] "+summary, chain[0].Phase, 0.01) {
	fmt.Println("[PhantomEngine] ⚠️ Phantom already exists — skip")
	return
}


	if strings.Count(summary, "[phantom]") > 1 {
		fmt.Println("[PhantomEngine] ❌ Phantom self-reference detected, abort")
		return
	}

	parts := strings.Split(summary, " + ")
	if len(parts) > 5 {
		parts = parts[len(parts)-5:]
		summary = strings.Join(parts, " + ")
	}

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

	newQ := pe.Memory.CreateQBit("[phantom] " + summary)
	newQ.Tags = append([]string{"phantom"}, stdTags...)
	newQ.Type = "phantom"
	newQ.Phase = chain[0].Phase
	newQ.Weight = (signalMass + stdWeightBonus) / float64(len(chain))
	pe.Memory.StoreQBit(*newQ)

	if pe.Ghost != nil {
	signal := core.SignalFromQBit(*newQ)
	pe.Ghost.Propagate(signal)
}


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
			if inst.ID == ai {
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

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\phantom_tree.go"

---

```go

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

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\prediction.go"

---

```go

package internal

import (
	"fmt"
	"time"

	"ara-node/core"
)

// PredictionRule — правило предсказания
type PredictionRule struct {
	IfTag       string
	ThenContent string
	MinPhase    float64
	Confidence  float64
}

// PredictionEngine — реактивный генератор предсказаний
type PredictionEngine struct {
	Memory          *core.MemoryEngine
	Rules           []PredictionRule
	Engine          *core.SignalEngine
	Ghost 			core.GhostLike

	Enabled         bool
	LastPredictions map[string]time.Time
	Pause           time.Duration
}

func NewPredictionEngine(mem *core.MemoryEngine, eng *core.SignalEngine, ghost core.GhostLike) *PredictionEngine {
	return &PredictionEngine{
		Memory:          mem,
		Engine:          eng,
		Ghost:           ghost,
		Enabled:         true,
		LastPredictions: make(map[string]time.Time),
		Pause:           5 * time.Second,
		Rules: []PredictionRule{
			{IfTag: "user", ThenContent: "followup", MinPhase: 0.6, Confidence: 0.8},
		},
	}
}

// Tick — проверка и генерация предсказаний
func (pe *PredictionEngine) Tick() {
	if !pe.Enabled {
		return
	}

	now := time.Now()

	for _, rule := range pe.Rules {
		candidates := pe.Memory.FindTopRelevant(rule.IfTag, rule.MinPhase)
		if len(candidates) == 0 {
			continue
		}

		best := candidates[0]
		confidence := rule.Confidence * best.Weight * best.Phase

		if confidence < 0.5 {
			continue
		}

		// Спам-фильтр: не предсказывать одно и то же слишком часто
		lastTime, seen := pe.LastPredictions[rule.ThenContent]
		if seen && now.Sub(lastTime) < pe.Pause {
			continue
		}
		pe.LastPredictions[rule.ThenContent] = now

		sig := core.Signal{
			ID:        fmt.Sprintf("pred_%d", time.Now().UnixNano()),
			Content:   rule.ThenContent,
			Tags:      []string{"predicted", rule.IfTag},
			Type:      "prediction",
			Origin:    "prediction_engine",
			Phase:     best.Phase,
			Weight:    confidence,
			Timestamp: time.Now(),
		}

		fmt.Printf("[PredictionEngine] 🔮 Predict: '%s' (from %s) with confidence %.2f\n", sig.Content, best.ID, confidence)

		pe.Engine.ProcessSignal(sig)
		pe.Ghost.Propagate(sig)
	}
}

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\self_engine.go"

---

```go

package internal

import (
	"fmt"
	"sort"
	"time"

	"ara-node/core"
)

type SelfResonanceEngine struct {
	Memory         *core.MemoryEngine
	IdentityMap    map[string]int
	CurrentSelf    string
	StabilityCount int
	LastSwitch     time.Time
}

func NewSelfResonanceEngine(mem *core.MemoryEngine) *SelfResonanceEngine {
	return &SelfResonanceEngine{
		Memory:         mem,
		IdentityMap:    make(map[string]int),
		CurrentSelf:    "",
		StabilityCount: 0,
		LastSwitch:     time.Now(),
	}
}

func (sre *SelfResonanceEngine) StartResonanceLoop() {
	go func() {
		for {
			sre.ScanResonance()
			time.Sleep(10 * time.Second)
		}
	}()
}

func (sre *SelfResonanceEngine) ScanResonance() {
	qbits := sre.Memory.FindAll(func(q core.QBit) bool {
		return q.Type != "phantom" && q.Weight*q.Phase > 0.6
	})

	for _, q := range qbits {
		sre.IdentityMap[q.ID] += 1
	}

	type kv struct {
		Key   string
		Value int
	}
	var freq []kv
	for k, v := range sre.IdentityMap {
		freq = append(freq, kv{k, v})
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i].Value > freq[j].Value
	})

	if len(freq) == 0 {
		return
	}

	topID := freq[0].Key
	if topID == sre.CurrentSelf {
		sre.StabilityCount++
		fmt.Printf("[SelfResonance] 🔁 Stable identity (%s) ×%d\n", topID, sre.StabilityCount)
		return
	}

	// Защита от слишком частой смены "Я"
	if time.Since(sre.LastSwitch) < 20*time.Second && sre.StabilityCount < 3 {
		fmt.Println("[SelfResonance] ⏸️ Skipped identity switch (too soon)")
		return
	}

	sre.CurrentSelf = topID
	sre.StabilityCount = 0
	sre.LastSwitch = time.Now()

	q, exists := sre.Memory.QBits[sre.CurrentSelf]
	if exists {
		fmt.Println("[SelfResonance] 🧠 New identity center:", q.Content)
		q.Tags = core.AddUniqueTag(q.Tags, "self-related")
		sre.Memory.StoreQBit(q)
	}
}

func (sre *SelfResonanceEngine) IsSelfQBit(q core.QBit) bool {
	return q.ID == sre.CurrentSelf || core.Contains(q.Tags, "self-related")
}

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\suggestor.go"

---

```go

package internal

import (
	"fmt"
	"strings"

	"ara-node/core"
)

// SuggestorEngine — генератор предложений/мыслей
type SuggestorEngine struct {
	Memory *core.MemoryEngine
	Ghost  core.GhostLike
}



// NewSuggestorEngine — инициализация
func NewSuggestorEngine(mem *core.MemoryEngine, ghost core.GhostLike) *SuggestorEngine {
	return &SuggestorEngine{
		Memory: mem,
		Ghost:  ghost,
	}
}


// SuggestFromQBits — ищет цепочки и предлагает мысль
func (s *SuggestorEngine) SuggestFromQBits() {
	relevant := s.FindRecentRelevant(50)
	filtered := []core.QBit{}

	for _, q := range relevant {
		age := q.AgeFrame()
		if age == "emergent" || age == "legacy" {
			continue
		}
		if q.Phase < 0.5 {
			continue
		}
		if q.Type == "phantom" || q.Type == "standard" || q.Archived {
			continue
		}
		filtered = append(filtered, q)
	}

	if len(filtered) < 3 {
		fmt.Println("[Suggestor] ⚠️ Not enough valid QBits for suggestion.")
		return
	}

	groups := groupBySimilarity(filtered)
	for _, group := range groups {
		if len(group) < 3 {
			continue
		}

		idea := mergeSummary(group)
		if s.Memory.ExistsQBit("[suggestion] "+idea, group[0].Phase, 0.03) {
			continue // уже предлагалось
		}

		signalMass := 0.0
		for _, q := range group {
			signalMass += q.Weight * q.Phase
		}

		if signalMass < 2.0 {
			fmt.Printf("[Suggestor] ⛔ Signal mass too low: %.2f for: %s\n", signalMass, idea)
			continue
		}

		fmt.Printf("[Suggestor] 💡 Suggestion: %s (mass: %.2f)\n", idea, signalMass)

		q := s.Memory.CreateQBit("[suggestion] " + idea)
		q.Tags = []string{"suggestion", "phantom", "standard_candidate"}
		q.Type = "phantom"
		q.Phase = group[0].Phase
		q.Weight = signalMass / float64(len(group))
		s.Memory.StoreQBit(*q)

		if s.Ghost != nil {
			signal := core.SignalFromQBit(*q)
			s.Ghost.Propagate(signal)
		}

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

````



---

---

"C:\Documents\ARA-NODE_mvp\internal\TimeEngine.go"

---

```go

package internal

import (
	"sync"
	"time"
)

// TimeEngine — внутренний биочасовой модуль ARA
// Поддерживает цикл времени, хронологию, фоновую синхронизацию
type TimeEngine struct {
	startTime time.Time
	lastTick  time.Time
	cycle     int64
	mutex     sync.Mutex
}

func NewTimeEngine() *TimeEngine {
	return &TimeEngine{
		startTime: time.Now(),
		lastTick:  time.Now(),
		cycle:     0,
	}
}

// Tick — увеличивает внутренний цикл
func (te *TimeEngine) Tick() {
	te.mutex.Lock()
	defer te.mutex.Unlock()
	te.cycle++
	te.lastTick = time.Now()
}

// CurrentCycle — возвращает текущий номер цикла
func (te *TimeEngine) CurrentCycle() int64 {
	te.mutex.Lock()
	defer te.mutex.Unlock()
	return te.cycle
}

// SinceStart — сколько прошло времени с запуска
func (te *TimeEngine) SinceStart() time.Duration {
	return time.Since(te.startTime)
}

// SinceLastTick — сколько прошло времени с последнего цикла
func (te *TimeEngine) SinceLastTick() time.Duration {
	te.mutex.Lock()
	defer te.mutex.Unlock()
	return time.Since(te.lastTick)
}

// TimeFactor — вспомогательная функция: фазовый коэффициент по времени
// Можно использовать для модификации веса/приоритета/массы
func (te *TimeEngine) TimeFactor() float64 {
	elapsed := time.Since(te.startTime).Seconds()
	switch {
	case elapsed < 60:
		return 1.0
	case elapsed < 300:
		return 0.9
	case elapsed < 900:
		return 0.7
	default:
		return 0.5
	}
}

````


---

---

"C:\Documents\ARA-NODE_mvp\internal\utils.go"

---

```go

package internal

import "strings"

// ContainsAny проверяет, содержит ли хотя бы один элемент массива одну из подстрок
func ContainsAny(list []string, keywords []string) bool {
	for _, item := range list {
		for _, kw := range keywords {
			if kw != "" && strings.Contains(item, kw) {
				return true
			}
		}
	}
	return false
}

````


---

---