package core

import (
	"fmt"
	
	"time"
)



// SignalEngine — обрабатывает входящие сигналы и записывает их в память,
// а также транслирует их по всей реактивной архитектуре (через GhostField).
type SignalEngine struct {
	Memory *MemoryEngine
	Ghost  *GhostField
}

// NewSignalEngine — инициализация ядра обработки сигналов
func NewSignalEngine(mem *MemoryEngine, ghost *GhostField) *SignalEngine {
	return &SignalEngine{
		Memory: mem,
		Ghost:  ghost,
	}
}

// ProcessSignal — принимает сигнал, сохраняет как QBit, запускает реакцию
func (se *SignalEngine) ProcessSignal(sig Signal) Reaction {
	fmt.Println("[SignalEngine] Received:", sig.Content)

	qbit := QBit{
		ID:           "qbit_" + sig.ID,
		Content:      sig.Content,
		Tags:         sig.Tags,
		CreatedAt:    time.Now(),
		LastAccessed: time.Now(),
		Weight:       sig.Weight,
		Phase:        sig.Phase,
		Type:         sig.Type,
		Origin:       sig.Origin,
	}
	se.Memory.StoreQBit(qbit)

	// Транслируем сигнал во всё поле
	if se.Ghost != nil {
		se.Ghost.Propagate(sig)
	}

	// Формируем реакцию по локальной памяти
	conf := 0.5
	tags := []string{"ack"}

	if sig.Phase > 0.85 {
		conf += 0.2
		tags = append(tags, "high_phase")
	}
	if sig.Origin == "will" || sig.Type == "phantom" {
		conf += 0.1
		tags = append(tags, "internal")
	}
	if matches := se.Memory.FindByPhase(sig.Phase, 0.03); len(matches) >= 2 {
		conf += 0.1
		tags = append(tags, "resonance")
	}

	return Reaction{
		TriggeredBy: sig.ID,
		Response:    "Signal dispatched to memory and network",
		Tags:        tags,
		Confidence:  conf,
	}
}
