
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

"C:\Documents\ARA-NODE_mvp\config\manifest.go"

---

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


---

---
