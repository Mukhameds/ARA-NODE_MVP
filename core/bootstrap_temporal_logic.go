// core/bootstrap_temporal_logic.go
package core

import (
	"fmt"
	"time"
)

// BootstrapTemporalLogic инициализирует временные отношения и причинно-следственные QBits
func BootstrapTemporalLogic(mem *MemoryEngine) {
	temporalConcepts := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Time concepts
		{"time_now", "now is the current moment", []string{"time", "concept", "core", "bootstrap"}},
		{"time_yesterday", "yesterday was before today", []string{"time", "sequence", "core", "bootstrap"}},
		{"time_tomorrow", "tomorrow comes after today", []string{"time", "sequence", "core", "bootstrap"}},
		{"time_past", "past is before present", []string{"time", "relation", "core", "bootstrap"}},
		{"time_future", "future is after present", []string{"time", "relation", "core", "bootstrap"}},
		{"time_difference", "yesterday is not today", []string{"time", "comparison", "core", "bootstrap"}},

		// Temporal logic / events
		{"event_before", "event A happens before event B", []string{"time", "event", "order", "core", "bootstrap"}},
		{"event_after", "event B happens after event A", []string{"time", "event", "order", "core", "bootstrap"}},
		{"event_if_then", "if it rains then it gets wet", []string{"causality", "if-then", "core", "bootstrap"}},
		{"event_sequence", "wake up → brush teeth → eat", []string{"event", "sequence", "routine", "core", "bootstrap"}},
		{"event_repeat", "sunrise repeats every day", []string{"event", "cycle", "core", "bootstrap"}},
	}

	for _, c := range temporalConcepts {
		q := QBit{
			ID:        c.ID,
			Content:   c.Content,
			Tags:      c.Tags,
			Phase:     0.82,
			Weight:    0.93,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("⏳ [Bootstrap] Temporal logic and causality loaded.")
}
