package core

import (
	"fmt"
	"time"
)


// SignalEngine — обрабатывает входящие сигналы и вызывает реакцию
type SignalEngine struct {
	Memory *MemoryEngine
}

func NewSignalEngine(mem *MemoryEngine) *SignalEngine {
	return &SignalEngine{Memory: mem}
}

// ProcessSignal — основной метод приёма и реакции
func (se *SignalEngine) ProcessSignal(sig Signal) Reaction {
	fmt.Println("[SignalEngine] Received:", sig.Content)

	// Сохраняем сигнал как QBit
	qbit := QBit{
		ID:        "qbit_" + sig.ID,
		Content:   sig.Content,
		Tags:      sig.Tags,
		CreatedAt: time.Now(),
		Weight:    sig.Weight,
		Phase:     sig.Phase,
		Type:      sig.Type,
		Origin:    sig.Origin,
	}
	se.Memory.StoreQBit(qbit)

	// Проверка совпадения по фазе (заглушка)
	if sig.Phase > 0.8 {
		return Reaction{
			TriggeredBy: sig.ID,
			Response:    "Phantom triggered by phase match",
			Tags:        []string{"phantom"},
			Confidence:  0.95,
		}
	}

	// Обычная реакция
	return Reaction{
		TriggeredBy: sig.ID,
		Response:    "Signal processed and stored",
		Tags:        []string{"ack"},
		Confidence:  0.8,
	}
}
