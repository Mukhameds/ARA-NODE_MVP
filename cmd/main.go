package main

import (
	"fmt"
	"time"

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

	// === SIGNAL ENGINE ===
	engine := core.NewSignalEngine(mem, ghost)
	prediction.Engine = engine
	prediction.Ghost = ghost

	// === BACKGROUND LOOPS ===
	internal.DefaultEmotionSet(emotion)
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

	// === CLI LOOP ===
	for {
		var input string
		fmt.Print("> ")
		fmt.Scanln(&input)

		if input == "exit" || input == "quit" {
			fmt.Println("👋 Завершение работы.")
			break
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
