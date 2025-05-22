// core/bootstrap_physics_formulas.go
package core

import (
	"fmt"
	"time"
)

// BootstrapPhysicsFormulas загружает базовые физические формулы как сигнальные QBits
func BootstrapPhysicsFormulas(mem *MemoryEngine) {
	formulas := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Классическая механика
		{"phys_force", "F = m · a", []string{"physics", "formula", "force", "newton", "core", "bootstrap"}},
		{"phys_acceleration", "a = Δv / Δt", []string{"physics", "formula", "acceleration", "core", "bootstrap"}},
		{"phys_velocity", "v = d / t", []string{"physics", "formula", "velocity", "core", "bootstrap"}},
		{"phys_momentum", "p = m · v", []string{"physics", "formula", "momentum", "core", "bootstrap"}},
		{"phys_work", "W = F · d", []string{"physics", "formula", "work", "core", "bootstrap"}},
		{"phys_power", "P = W / t", []string{"physics", "formula", "power", "core", "bootstrap"}},

		// Энергия
		{"phys_kinetic_energy", "KE = ½ · m · v²", []string{"physics", "formula", "energy", "kinetic", "core", "bootstrap"}},
		{"phys_potential_energy", "PE = m · g · h", []string{"physics", "formula", "energy", "potential", "core", "bootstrap"}},
		{"phys_einstein", "E = m · c²", []string{"physics", "formula", "energy", "relativity", "core", "bootstrap"}},

		// Материя
		{"phys_density", "ρ = m / V", []string{"physics", "formula", "density", "core", "bootstrap"}},

		// Электричество
		{"phys_ohm_law", "V = I · R", []string{"physics", "formula", "electricity", "ohm", "core", "bootstrap"}},
		{"phys_power_electric", "P = V · I", []string{"physics", "formula", "electricity", "power", "core", "bootstrap"}},
	}

	for _, f := range formulas {
		q := QBit{
			ID:        f.ID,
			Content:   f.Content,
			Tags:      f.Tags,
			Phase:     0.91,
			Weight:    0.98,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("⚙️ [Bootstrap] Physics formulas and physical laws loaded.")
}
