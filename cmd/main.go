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

	signalEngine := core.NewSignalEngine(mem)
	phantom := internal.NewFanthomEngine(mem)
	suggestor := internal.NewSuggestorEngine(mem)
	human := internal.NewHumanNodeEngine(mem)
	ghostField := core.NewGhostField()
	will := core.NewWillEngine(mem)
	will.DesireLoop()

	// Пример реактивного блока
	block := &core.Block{
		ID: "UserPhaseBlock",
		Rules: []core.ReactionRule{
			{
				MatchTag: "user",
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

		// команды словаря переменных
		if content == "dict" {
			for _, vb := range signalDict.All() {
				fmt.Println("🔹", vb.ID, "|", vb.Tags, "| Reacts:", vb.Reacts)
			}
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

		// 📡 обработка сигнального словаря (переменные)
		tokens := strings.Fields(strings.ToLower(content))
		for _, token := range tokens {
			vb := signalDict.FindMatch(token)
			if vb == nil {
				vb = signalDict.AutoLearn(token)
			}
			mem.Broadcast(vb.QBit)
		}

		// 🚀 основная реакция
		signal := core.Signal{
			ID:        fmt.Sprintf("sig_%d", time.Now().UnixNano()),
			Content:   content,
			Tags:      []string{"user"},
			Timestamp: time.Now(),
			Phase:     0.85,
			Weight:    1.0,
			Origin:    "user",
		}

		reaction := signalEngine.ProcessSignal(signal)

		ghostField.Propagate(signal)
		phantom.TriggerFromMatch(signal)
		suggestor.SuggestFromQBits()
		mem.DecayQBits()
		fmt.Println("[Main] Decaying QBits...")

		fmt.Println("←", reaction.Response)
		fmt.Println("↪ Tags:", reaction.Tags)
	}
}
