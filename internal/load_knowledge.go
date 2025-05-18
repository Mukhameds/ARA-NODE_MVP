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
