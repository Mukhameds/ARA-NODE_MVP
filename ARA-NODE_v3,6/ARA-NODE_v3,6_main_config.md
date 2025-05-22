
---

"C:\Documents\ARA-NODE_mvp\cmd\main.go"

---

package main

import (
	"fmt"
	"strings"
	"time"

	"ara-node/core"
	"ara-node/field"
	"ara-node/internal"
)

func main() {
	fmt.Println("🧠 ARA-NODE CLI started.")

	// === CORE INIT ===
	mem := core.NewMemoryEngine()
	dict := core.NewSignalDictionary(mem)
	internal.RunBootstrap(mem, dict)
	shutdown := core.NewShutdownEngine(1.0, 2*time.Second)

	// === GHOST ROCKET SETUP ===
	rocket := field.NewGhostRocket("MainMind")

	mathField := field.NewMatrix("math")
	emotionField := field.NewMatrix("emotion")
	phantomField := field.NewMatrix("phantom")
	instinctField := field.NewMatrix("instinct")

	rocket.AddField(mathField)
	rocket.AddField(emotionField)
	rocket.AddField(phantomField)
	rocket.AddField(instinctField)

	// === ROCKET ADAPTER ===
	adapter := field.RocketAdapter(rocket)

	// === MODULES ===
	emotion := internal.NewEmotionEngine(mem)
	instinct := internal.NewInstinctEngine(adapter)
	emotion.Instincts = instinct

	timeEngine := internal.NewTimeEngine()
	phantom := internal.NewPhantomEngine(mem, instinct, emotion, timeEngine, adapter)
	suggestor := internal.NewSuggestorEngine(mem, adapter)
	prediction := internal.NewPredictionEngine(mem, nil, nil)
	reflex := core.NewReflexEngine()
	will := core.NewWillEngine(mem, nil, nil, phantom)
	resonance := core.NewResonanceMatrix()
	attention := core.NewAttentionEngine(mem, adapter, phantom, nil)
	dual := core.NewDualProcessor(mem, adapter)

	// === Register reaction blocks ===
	emotionField.RegisterBlock("emotion", field.ReactionRule{
		MatchTags: []string{"emotion", "instinct", "success", "fail"},
		MinPhase:  0.5,
		Action:    emotion.React,
	}, 200*time.Millisecond)

	phantomField.RegisterBlock("phantom", field.ReactionRule{
		MatchTags: []string{"phantom", "standard", "merge"},
		MinPhase:  0.7,
		Action:    phantom.TriggerFromMatch,
	}, 200*time.Millisecond)

	instinctField.RegisterBlock("reflex", field.ReactionRule{
		MatchTags: []string{"danger", "reflex", "instinct_error"},
		MinPhase:  0.5,
		Action:    reflex.React,
	}, 100*time.Millisecond)

	mathField.RegisterBlock("suggestor", field.ReactionRule{
		MatchTags: []string{"user", "phantom", "background", "core", "math", "physics", "symbol"},
		MinPhase:  0.7,
		Action: func(sig core.Signal) {
			go suggestor.SuggestFromQBits()
		},
	}, 2*time.Second)

	// === META FIELDS ===
	metaFieldMath := field.NewMatrix("meta_math")
	metaFieldEmotion := field.NewMatrix("meta_emotion")
	metaFieldPhantom := field.NewMatrix("meta_phantom")

	rocket.AddField(metaFieldMath)
	rocket.AddField(metaFieldEmotion)
	rocket.AddField(metaFieldPhantom)

	//meta := map[string]*field.Matrix{
	//	"math":    metaFieldMath,
	//	"emotion": metaFieldEmotion,
	//	"phantom": metaFieldPhantom,
	//}

	// === PeerSync (можно включить позже) ===
	/*
		peerSync, err := internal.NewPeerSync(mem, meta)
		if err != nil {
			fmt.Println("❌ PeerSync init error:", err)
		} else {
			fmt.Println("🌐 PeerSync initialized.")
		}
	*/

	// === ENGINE WIRING ===
	engine := core.NewSignalEngine(mem, adapter)
	prediction.Engine = engine
	prediction.Ghost = adapter
	will.Engine = engine
	will.Ghost = adapter
	attention.Engine = engine

	// === BACKGROUND ENGINES ===
	attention.StartBackgroundThinking()
	will.DesireLoop()

	go func() {
		for {
			time.Sleep(6 * time.Second)
			suggestor.SuggestFromQBits()
		}
	}()

	go func() {
		for {
			time.Sleep(8 * time.Second)
			top := mem.FindTopRelevant("core", 0.6)
			if len(top) >= 3 {
				phantom.GeneratePhantomChain(top[:5])
			}
		}
	}()

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

	go func() {
		for {
			time.Sleep(15 * time.Second)
			resonance.Decay()
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			currentMass := mem.EstimateTotalPhase()
			fmt.Printf("[MassCheck] 🧮 Estimated signal mass: %.3f\n", currentMass)
			shutdown.UpdateMass(currentMass, mem)
		}
	}()

	// === CLI LOOP ===
	for {
		var input string
		fmt.Print("> ")
		fmt.Scanln(&input)

		switch {
		case input == "exit" || input == "quit":
			fmt.Println("👋 Завершение работы.")
			return

		case input == "view":
			mem.ListQBits()

		case input == "help":
			fmt.Println("🆘 Команды:\n- help\n- view\n- view emotions\n- delete <qbit_id>\n- sync\n- loadfacts\n- exit")

		case input == "view emotions":
			for _, e := range emotion.CurrentEmotions() {
				fmt.Println("❤️", e)
			}

		case strings.HasPrefix(input, "delete "):
			id := strings.TrimPrefix(input, "delete ")
			mem.DeleteQBit(id)

		case input == "sync":
			fmt.Println("[Sync] 🔄 Запуск синхронизации (заглушка)...")

		case input == "loadfacts":
			err := internal.LoadFactsFromFile("data/core_knowledge.json", engine, adapter)
			if err != nil {
				fmt.Println("❌ Ошибка загрузки фактов:", err)
			} else {
				fmt.Println("📚 Факты успешно загружены.")
			}

		default:
			// === Инстинктивные сигналы до основного
			signals := instinct.TickSignals(time.Now(), input)
			for _, sig := range signals {
				dpSignal := sig
				dpSignal.Type = "user"
				dpSignal.Tags = append(dpSignal.Tags, "cli")
				dpSignal.Origin = "instinct"
				dpSignal.Phase = 0.7
				dpSignal.Weight = 0.8
				dpSignal.Timestamp = time.Now()
				dual.ProcessDual(dpSignal)
			}

			// === Основной пользовательский сигнал
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
			dual.ProcessDual(sig)

			matched := mem.FindByTag("user")
			resonance.BoostBySignal(sig, matched)
			resonance.Print(sig.ID)
		}
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