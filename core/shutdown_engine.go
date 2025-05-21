package core

import (
	"fmt"
	"sync"
	"time"
)

// ShutdownEngine управляет реактивным отключением модулей в случае сигнального коллапса

type ShutdownEngine struct {
	Active           bool
	SignalMass       float64
	Threshold        float64
	CriticalModules  []string
	ShutdownLog      []string
	Mutex            sync.Mutex
	StepDelay        time.Duration
	OnModuleShutdown func(string)
}

func NewShutdownEngine(threshold float64, stepDelay time.Duration) *ShutdownEngine {
	return &ShutdownEngine{
		Active:          false,
		SignalMass:      1.0,
		Threshold:       threshold,
		CriticalModules: []string{"suggestor", "reflex", "phantom", "attention", "emotion", "will", "ghost", "signal"},
		ShutdownLog:     []string{},
		StepDelay:       stepDelay,
	}
}

// UpdateMass обновляет текущую фазовую массу системы (сумма активных резонансов / смыслов)
func (se *ShutdownEngine) UpdateMass(current float64) {
	se.Mutex.Lock()
	defer se.Mutex.Unlock()
	se.SignalMass = current
	if current < se.Threshold && !se.Active {
		fmt.Println("[ShutdownEngine] ⚠️ Critical signal mass detected! Initiating collapse...")
		se.Active = true
		go se.StartCollapse()
	}
}

// StartCollapse — последовательное отключение модулей
func (se *ShutdownEngine) StartCollapse() {
	for _, module := range se.CriticalModules {
		se.ShutdownLog = append(se.ShutdownLog, fmt.Sprintf("Module %s: disabled", module))
		fmt.Printf("[ShutdownEngine] ❌ Module %s is shutting down\n", module)
		if se.OnModuleShutdown != nil {
			se.OnModuleShutdown(module)
		}
		time.Sleep(se.StepDelay)
	}
	fmt.Println("[ShutdownEngine] 💀 ARA-NODE has ceased functioning.")
}

// IsShuttingDown возвращает статус критического отключения
func (se *ShutdownEngine) IsShuttingDown() bool {
	se.Mutex.Lock()
	defer se.Mutex.Unlock()
	return se.Active
}

// Log возвращает историю отключения
func (se *ShutdownEngine) Log() []string {
	return se.ShutdownLog
}
