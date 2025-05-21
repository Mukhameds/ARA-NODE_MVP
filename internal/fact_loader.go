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
func LoadFactsFromFile(filename string, engine *core.SignalEngine, ghost *core.GhostField) error {
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
