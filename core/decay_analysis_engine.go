package core

import (
	"fmt"
	"time"
)

// DecayEvent — лог обнуления или вымирания узла
type DecayEvent struct {
	ID        string
	Reason    string
	Timestamp time.Time
}

// DecayAnalysisEngine — отслеживает процессы старения памяти
type DecayAnalysisEngine struct {
	Log []DecayEvent
	Mem *MemoryEngine
}

func NewDecayAnalysisEngine(mem *MemoryEngine) *DecayAnalysisEngine {
	return &DecayAnalysisEngine{
		Log: []DecayEvent{},
		Mem: mem,
	}
}

func (de *DecayAnalysisEngine) RunDecayCheck() {
	for id, q := range de.Mem.QBits {
		if q.Archived && q.Weight < 0.05 {
			delete(de.Mem.QBits, id)
			de.Log = append(de.Log, DecayEvent{
				ID:        id,
				Reason:    "fully decayed",
				Timestamp: time.Now(),
			})
			fmt.Println("[Decay] ☠️ Removed:", id)
		}
	}
}

func (de *DecayAnalysisEngine) PrintDecayLog() {
	if len(de.Log) == 0 {
		fmt.Println("[DecayLog] 🔹 Память ещё не подвергалась очистке.")
		return
	}
	fmt.Println("[DecayLog] 🧩 Deleted QBits:")
	for _, entry := range de.Log {
		fmt.Printf("⏱ %s | %s | %s\n", entry.Timestamp.Format(time.RFC822), entry.ID, entry.Reason)
	}
}
