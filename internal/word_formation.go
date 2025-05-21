package internal

import (
	"fmt"
	"strings"
	"time"

	"ara-node/core"
)

// WordFormationEngine — анализирует поток букв и формирует устойчивые слова
// Используется поверх SignalDictionary (буфер уже там)
type WordFormationEngine struct {
	Dict       *core.SignalDictionary
	Memory     *core.MemoryEngine
	WordCount  map[string]int
	LastSeen   map[string]time.Time
	ConfirmMin int
}

func NewWordFormationEngine(dict *core.SignalDictionary, mem *core.MemoryEngine) *WordFormationEngine {
	return &WordFormationEngine{
		Dict:       dict,
		Memory:     mem,
		WordCount:  make(map[string]int),
		LastSeen:   make(map[string]time.Time),
		ConfirmMin: 3, // минимум повторений до подтверждения
	}
}

// Tick — проверяет буфер словаря и формирует слова, если они устойчивые
func (w *WordFormationEngine) Tick() {
	word := strings.Join(w.Dict.Buffer(), "")
	if len(word) < 3 || len(word) > 12 {
		return
	}

	w.WordCount[word]++
	w.LastSeen[word] = time.Now()

	if w.WordCount[word] >= w.ConfirmMin {
		if vb := w.Dict.FindMatch(word); vb == nil {
			vb := w.Dict.Add(word, word, []string{"word", "confirmed"}, []string{word})
			vb.QBit.Tags = append(vb.QBit.Tags, "confirmed")
			vb.QBit.Phase += 0.1
			w.Memory.StoreQBit(*vb.QBit)
			fmt.Println("[WordFormationEngine] ✅ Confirmed word:", word)
		}
	}
}

// Decay — снижает счётчик редко встречающихся слов
func (w *WordFormationEngine) Decay() {
	now := time.Now()
	for word, t := range w.LastSeen {
		if now.Sub(t) > 60*time.Second {
			w.WordCount[word]--
			if w.WordCount[word] <= 0 {
				delete(w.WordCount, word)
				delete(w.LastSeen, word)
			}
		}
	}
}
