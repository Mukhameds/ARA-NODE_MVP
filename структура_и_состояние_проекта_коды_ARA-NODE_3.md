
"C:\Documents\ARA-NODE_mvp\internal\github_sync.go":
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


"C:\Documents\ARA-NODE_mvp\internal\human_node.go":
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

"C:\Documents\ARA-NODE_mvp\internal\knowledge_profile_loader.go":
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


"C:\Documents\ARA-NODE_mvp\internal\load_knowledge.go":
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

"C:\Documents\ARA-NODE_mvp\internal\p2p_sync.go":
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

"C:\Documents\ARA-NODE_mvp\internal\phantom.go":
package internal

import (
	"fmt"
	"strings"
	"ara-node/core"
)

// FanthomEngine — генератор фантомов
type FanthomEngine struct {
	Memory *core.MemoryEngine
}

// NewFanthomEngine — инициализация
func NewFanthomEngine(mem *core.MemoryEngine) *FanthomEngine {
	return &FanthomEngine{
		Memory: mem,
	}
}

// TriggerFromMatch — ищет совпадения и запускает фантом
func (fe *FanthomEngine) TriggerFromMatch(sig core.Signal) {
	qbits := fe.Memory.FindByPhase(sig.Phase, 0.05)
	if len(qbits) >= 2 {
		fmt.Println("[FanthomEngine] ⚡ Phase-match found. Generating phantom...")
		fe.GeneratePhantomChain(qbits)
	}
}

// GeneratePhantomChain — строит фантом из цепочки QBit
func (fe *FanthomEngine) GeneratePhantomChain(chain []core.QBit) {
	fmt.Println("[FanthomChain] 🧩 Related QBits:")
	var summary string
	var sources []string

	for _, q := range chain {
		fmt.Printf("• %s | %.2f | %s\n", q.ID, q.Phase, q.Content)
		summary += q.Content + " + "
		sources = append(sources, q.ID)
	}

	summary = strings.TrimSuffix(summary, " + ")
	fmt.Println("[FanthomChain] → Hypothesis: something meaningful links these signals.")

	// Создать и сохранить QBit-фантом
	newQ := fe.Memory.CreateQBit("[phantom] " + summary)
	newQ.Tags = []string{"phantom"}
	newQ.Type = "phantom"
	newQ.Phase = chain[0].Phase
	newQ.Weight = 0.8
	fe.Memory.StoreQBit(*newQ)

	// Лог
	fmt.Println("[FanthomEngine] 🔮 Phantom QBit:", newQ.ID)
	fmt.Println("[FanthomEngine] ↪ Sources:", strings.Join(sources, ","))
}


"C:\Documents\ARA-NODE_mvp\internal\phantom_tree.go":
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


"C:\Documents\ARA-NODE_mvp\internal\suggestor.go":
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
	return &SuggestorEngine{
		Memory: mem,
	}
}

// SuggestFromQBits — ищет цепочки и предлагает мысль
func (s *SuggestorEngine) SuggestFromQBits() {
	qbits := s.Memory.FindByTag("user")
	if len(qbits) < 2 {
		return
	}

	// Объединение смыслов
	var ideas []string
	for _, q := range qbits {
		ideas = append(ideas, q.Content)
		if len(ideas) >= 3 {
			break
		}
	}

	// Генерация фразы
	suggestion := s.GenerateSuggestion(ideas)
	fmt.Println("[Suggestor] 💡", suggestion)
}

// GenerateSuggestion — строит предложение на основе смыслов
func (s *SuggestorEngine) GenerateSuggestion(ideas []string) string {
	if len(ideas) == 0 {
		return "No suggestion available."
	}
	return fmt.Sprintf("Would you like to explore the idea: \"%s\" + ...?", strings.Join(ideas, " + "))
}


"C:\Documents\ARA-NODE_mvp\tests\test_1.md":
Microsoft Windows [Version 10.0.26100.4061]
(c) Microsoft Corporation. All rights reserved.

C:\Users\99650>cd C:\Documents\ARA-NODE_mvp

C:\Documents\ARA-NODE_mvp>go run ./cmd
🧠 ARA-NODE CLI started.
[SelfKernel] Initialized: ARA::node::001
Agent ID: ARA::node::001
[P2P] Started with ID: 12D3KooWGQdzbdXUbyKcWwXW81z61asYHVrGknyLqVCujMpESkpw
[P2P] 🛰️ Sync active
[MemoryEngine] Auto-created QBit: qbit_1747584347075192700
[MemoryEngine] Stored QBit: qbit_1747584347075192700
🧬 [ARA Bootstrap] Начало инициализации личности...
🧠 Какая твоя главная цель в жизни?
→ q1
[MemoryEngine] Auto-created QBit: qbit_1747584354375436200
[MemoryEngine] Auto-created QBit: qbit_1747584354376264200
🧠 Какие темы тебе наиболее интересны?
→ q2
[MemoryEngine] Auto-created QBit: qbit_1747584356028223500
[MemoryEngine] Auto-created QBit: qbit_1747584356029729200
🧠 Как ты хочешь, чтобы ARA помогала тебе?
→ q3
[MemoryEngine] Auto-created QBit: qbit_1747584357372099200
[MemoryEngine] Auto-created QBit: qbit_1747584357372099200
🧠 Кто ты по жизни? (учёный, инженер, философ...)
→ q4
[MemoryEngine] Auto-created QBit: qbit_1747584358926437500
[MemoryEngine] Auto-created QBit: qbit_1747584358926437500
✅ [Bootstrap] Базовые цели и профиль сохранены.
[GhostField] Registered Block: UserPhaseBlock

>> [SignalEngine] Received: q2
[MemoryEngine] Stored QBit: qbit_bg_1747584363928309900
[SignalEngine] Received: q2
[MemoryEngine] Stored QBit: qbit_bg_1747584363930160400
[FanthomEngine] ⚡ Phase-match found. Generating phantom...
[FanthomChain] 🧩 Related QBits:
• qbit_bg_1747584363928309900 | 0.80 | q2
• qbit_bg_1747584363930160400 | 0.80 | q2
[FanthomChain] → Hypothesis: something meaningful links these signals.
[MemoryEngine] Auto-created QBit: qbit_1747584363930864800
[MemoryEngine] Stored QBit: qbit_1747584363930864800
[FanthomEngine] 🔮 Phantom QBit: qbit_1747584363930864800
[FanthomEngine] ↪ Sources: qbit_bg_1747584363928309900,qbit_bg_1747584363930160400
[SignalEngine] Received: q3
[MemoryEngine] Stored QBit: qbit_bg_1747584363930864800
[FanthomEngine] ⚡ Phase-match found. Generating phantom...
[FanthomChain] 🧩 Related QBits:
• qbit_bg_1747584363928309900 | 0.80 | q2
• qbit_bg_1747584363930160400 | 0.80 | q2
• qbit_1747584363930864800 | 0.80 | [phantom] q2 + q2
• qbit_bg_1747584363930864800 | 0.80 | q3
[FanthomChain] → Hypothesis: something meaningful links these signals.
[MemoryEngine] Auto-created QBit: qbit_1747584363932425700
[MemoryEngine] Stored QBit: qbit_1747584363932425700
[FanthomEngine] 🔮 Phantom QBit: qbit_1747584363932425700
[FanthomEngine] ↪ Sources: qbit_bg_1747584363928309900,qbit_bg_1747584363930160400,qbit_1747584363930864800,qbit_bg_1747584363930864800
[SignalEngine] Received: q4
[MemoryEngine] Stored QBit: qbit_bg_1747584363932425700
[FanthomEngine] ⚡ Phase-match found. Generating phantom...
[FanthomChain] 🧩 Related QBits:
• qbit_bg_1747584363928309900 | 0.80 | q2
• qbit_bg_1747584363930160400 | 0.80 | q2
• qbit_1747584363930864800 | 0.80 | [phantom] q2 + q2
• qbit_bg_1747584363930864800 | 0.80 | q3
• qbit_1747584363932425700 | 0.80 | [phantom] q2 + q2 + [phantom] q2 + q2 + q3
• qbit_bg_1747584363932425700 | 0.80 | q4
[FanthomChain] → Hypothesis: something meaningful links these signals.
[MemoryEngine] Auto-created QBit: qbit_1747584363933600700
[MemoryEngine] Stored QBit: qbit_1747584363933600700
[FanthomEngine] 🔮 Phantom QBit: qbit_1747584363933600700
[FanthomEngine] ↪ Sources: qbit_bg_1747584363928309900,qbit_bg_1747584363930160400,qbit_1747584363930864800,qbit_bg_1747584363930864800,qbit_1747584363932425700,qbit_bg_1747584363932425700
[SignalEngine] Received: Amplify and assist user cognition through signal logic.
[MemoryEngine] Stored QBit: qbit_bg_1747584363933600700
[FanthomEngine] ⚡ Phase-match found. Generating phantom...
[FanthomChain] 🧩 Related QBits:
• qbit_1747584363933600700 | 0.80 | [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + q4
• qbit_bg_1747584363932425700 | 0.80 | q4
• qbit_bg_1747584363933600700 | 0.80 | Amplify and assist user cognition through signal logic.
• qbit_1747584363930864800 | 0.80 | [phantom] q2 + q2
• qbit_bg_1747584363930864800 | 0.80 | q3
• qbit_1747584363932425700 | 0.80 | [phantom] q2 + q2 + [phantom] q2 + q2 + q3
• qbit_bg_1747584363928309900 | 0.80 | q2
• qbit_bg_1747584363930160400 | 0.80 | q2
[FanthomChain] → Hypothesis: something meaningful links these signals.
[MemoryEngine] Auto-created QBit: qbit_1747584363934667900
[MemoryEngine] Stored QBit: qbit_1747584363934667900
[FanthomEngine] 🔮 Phantom QBit: qbit_1747584363934667900
[FanthomEngine] ↪ Sources: qbit_1747584363933600700,qbit_bg_1747584363932425700,qbit_bg_1747584363933600700,qbit_1747584363930864800,qbit_bg_1747584363930864800,qbit_1747584363932425700,qbit_bg_1747584363928309900,qbit_bg_1747584363930160400
[SignalEngine] Received: q1
[MemoryEngine] Stored QBit: qbit_bg_1747584363936279600
[FanthomEngine] ⚡ Phase-match found. Generating phantom...
[FanthomChain] 🧩 Related QBits:
• qbit_1747584363930864800 | 0.80 | [phantom] q2 + q2
• qbit_bg_1747584363930864800 | 0.80 | q3
• qbit_1747584363932425700 | 0.80 | [phantom] q2 + q2 + [phantom] q2 + q2 + q3
• qbit_1747584363934667900 | 0.80 | [phantom] [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + q4 + q4 + Amplify and assist user cognition through signal logic. + [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + q2 + q2
• qbit_bg_1747584363928309900 | 0.80 | q2
• qbit_bg_1747584363930160400 | 0.80 | q2
• qbit_bg_1747584363936279600 | 0.80 | q1
• qbit_1747584363933600700 | 0.80 | [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + q4
• qbit_bg_1747584363932425700 | 0.80 | q4
• qbit_bg_1747584363933600700 | 0.80 | Amplify and assist user cognition through signal logic.
[FanthomChain] → Hypothesis: something meaningful links these signals.
[MemoryEngine] Auto-created QBit: qbit_1747584363936872900
[MemoryEngine] Stored QBit: qbit_1747584363936872900
[FanthomEngine] 🔮 Phantom QBit: qbit_1747584363936872900
[FanthomEngine] ↪ Sources: qbit_1747584363930864800,qbit_bg_1747584363930864800,qbit_1747584363932425700,qbit_1747584363934667900,qbit_bg_1747584363928309900,qbit_bg_1747584363930160400,qbit_bg_1747584363936279600,qbit_1747584363933600700,qbit_bg_1747584363932425700,qbit_bg_1747584363933600700
[SignalEngine] Received: q1
[MemoryEngine] Stored QBit: qbit_bg_1747584363936872900
[FanthomEngine] ⚡ Phase-match found. Generating phantom...
[FanthomChain] 🧩 Related QBits:
• qbit_bg_1747584363936872900 | 0.80 | q1
• qbit_bg_1747584363928309900 | 0.80 | q2
• qbit_bg_1747584363930160400 | 0.80 | q2
• qbit_bg_1747584363936279600 | 0.80 | q1
• qbit_1747584363933600700 | 0.80 | [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + q4
• qbit_bg_1747584363932425700 | 0.80 | q4
• qbit_bg_1747584363933600700 | 0.80 | Amplify and assist user cognition through signal logic.
• qbit_1747584363936872900 | 0.80 | [phantom] [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + [phantom] [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + q4 + q4 + Amplify and assist user cognition through signal logic. + [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + q2 + q2 + q2 + q2 + q1 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + q4 + q4 + Amplify and assist user cognition through signal logic.
• qbit_1747584363930864800 | 0.80 | [phantom] q2 + q2
• qbit_bg_1747584363930864800 | 0.80 | q3
• qbit_1747584363932425700 | 0.80 | [phantom] q2 + q2 + [phantom] q2 + q2 + q3
• qbit_1747584363934667900 | 0.80 | [phantom] [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + q4 + q4 + Amplify and assist user cognition through signal logic. + [phantom] q2 + q2 + q3 + [phantom] q2 + q2 + [phantom] q2 + q2 + q3 + q2 + q2
[FanthomChain] → Hypothesis: something meaningful links these signals.
[MemoryEngine] Auto-created QBit: qbit_1747584363937945200
[MemoryEngine] Stored QBit: qbit_1747584363937945200
[FanthomEngine] 🔮 Phantom QBit: qbit_1747584363937945200
[FanthomEngine] ↪ Sources: qbit_bg_1747584363936872900,qbit_bg_1747584363928309900,qbit_bg_1747584363930160400,qbit_bg_1747584363936279600,qbit_1747584363933600700,qbit_bg_1747584363932425700,qbit_bg_1747584363933600700,qbit_1747584363936872900,qbit_1747584363930864800,qbit_bg_1747584363930864800,qbit_1747584363932425700,qbit_1747584363934667900


