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
