// core/bootstrap_physics_quantum.go
package core

import (
	"fmt"
	"time"
)

// BootstrapPhysicsQuantum загружает сигнальные QBits квантовой физики
func BootstrapPhysicsQuantum(mem *MemoryEngine) {
	quantum := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Основные концепции
		{"quantum_def", "quantum physics studies behavior of matter and energy at small scales", []string{"physics", "quantum", "definition", "core", "bootstrap"}},
		{"quantum_superposition", "superposition means a particle can exist in multiple states at once", []string{"physics", "quantum", "superposition", "core", "bootstrap"}},
		{"quantum_entanglement", "entanglement links particles so that the state of one affects the other instantly", []string{"physics", "quantum", "entanglement", "core", "bootstrap"}},
		{"quantum_uncertainty", "Heisenberg uncertainty principle: Δx · Δp ≥ ħ / 2", []string{"physics", "quantum", "uncertainty", "core", "bootstrap"}},
		{"quantum_wave_dual", "wave-particle duality: particles behave like both waves and particles", []string{"physics", "quantum", "duality", "core", "bootstrap"}},

		// Частицы и энергия
		{"quantum_quantization", "energy comes in discrete packets called quanta", []string{"physics", "quantum", "energy", "core", "bootstrap"}},
		{"quantum_levels", "electrons occupy quantized energy levels", []string{"physics", "quantum", "energy", "core", "bootstrap"}},
		{"quantum_spin", "spin is an intrinsic form of angular momentum", []string{"physics", "quantum", "spin", "core", "bootstrap"}},

		// Измерение и интерпретации
		{"quantum_measurement", "measurement collapses a quantum system into a definite state", []string{"physics", "quantum", "measurement", "core", "bootstrap"}},
		{"quantum_probabilistic", "quantum predictions are probabilistic, not deterministic", []string{"physics", "quantum", "probability", "core", "bootstrap"}},
		{"quantum_planck_scale", "Planck scale defines the smallest meaningful units in physics", []string{"physics", "quantum", "planck", "core", "bootstrap"}},
	}

	for _, q := range quantum {
		qbit := QBit{
			ID:        q.ID,
			Content:   q.Content,
			Tags:      q.Tags,
			Phase:     0.93,
			Weight:    0.98,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(qbit)
	}

	fmt.Println("🧬 [Bootstrap] Quantum physics concepts loaded.")
}
