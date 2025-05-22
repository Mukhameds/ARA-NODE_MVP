package internal

import (
	"fmt"
	"sort"
	"time"

	"ara-node/core"
)

type SelfResonanceEngine struct {
	Memory         *core.MemoryEngine
	IdentityMap    map[string]int
	CurrentSelf    string
	StabilityCount int
	LastSwitch     time.Time
}

func NewSelfResonanceEngine(mem *core.MemoryEngine) *SelfResonanceEngine {
	return &SelfResonanceEngine{
		Memory:         mem,
		IdentityMap:    make(map[string]int),
		CurrentSelf:    "",
		StabilityCount: 0,
		LastSwitch:     time.Now(),
	}
}

func (sre *SelfResonanceEngine) StartResonanceLoop() {
	go func() {
		for {
			sre.ScanResonance()
			time.Sleep(10 * time.Second)
		}
	}()
}

func (sre *SelfResonanceEngine) ScanResonance() {
	qbits := sre.Memory.FindAll(func(q core.QBit) bool {
		return q.Type != "phantom" && q.Weight*q.Phase > 0.6
	})

	for _, q := range qbits {
		sre.IdentityMap[q.ID] += 1
	}

	type kv struct {
		Key   string
		Value int
	}
	var freq []kv
	for k, v := range sre.IdentityMap {
		freq = append(freq, kv{k, v})
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i].Value > freq[j].Value
	})

	if len(freq) == 0 {
		return
	}

	topID := freq[0].Key
	if topID == sre.CurrentSelf {
		sre.StabilityCount++
		fmt.Printf("[SelfResonance] 🔁 Stable identity (%s) ×%d\n", topID, sre.StabilityCount)
		return
	}

	// Защита от слишком частой смены "Я"
	if time.Since(sre.LastSwitch) < 20*time.Second && sre.StabilityCount < 3 {
		fmt.Println("[SelfResonance] ⏸️ Skipped identity switch (too soon)")
		return
	}

	sre.CurrentSelf = topID
	sre.StabilityCount = 0
	sre.LastSwitch = time.Now()

	q, exists := sre.Memory.QBits[sre.CurrentSelf]
	if exists {
		fmt.Println("[SelfResonance] 🧠 New identity center:", q.Content)
		q.Tags = core.AddUniqueTag(q.Tags, "self-related")
		sre.Memory.StoreQBit(q)
	}
}

func (sre *SelfResonanceEngine) IsSelfQBit(q core.QBit) bool {
	return q.ID == sre.CurrentSelf || core.Contains(q.Tags, "self-related")
}
