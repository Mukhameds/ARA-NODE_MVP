// core/bootstrap_physics_constants.go
package core

import (
	"fmt"
	"time"
)

// BootstrapPhysicsConstants загружает физические константы как QBits
func BootstrapPhysicsConstants(mem *MemoryEngine) {
	constants := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Основные фундаментальные константы
		{"const_g", "gravitational acceleration g ≈ 9.81 m/s²", []string{"physics", "constant", "gravity", "core", "bootstrap"}},
		{"const_c", "speed of light c ≈ 299,792,458 m/s", []string{"physics", "constant", "light", "core", "bootstrap"}},
		{"const_h", "Planck constant h ≈ 6.626×10⁻³⁴ Js", []string{"physics", "constant", "quantum", "planck", "core", "bootstrap"}},
		{"const_e", "elementary charge e ≈ 1.602×10⁻¹⁹ C", []string{"physics", "constant", "charge", "electron", "core", "bootstrap"}},
		{"const_k", "Boltzmann constant k ≈ 1.381×10⁻²³ J/K", []string{"physics", "constant", "thermo", "boltzmann", "core", "bootstrap"}},
		{"const_na", "Avogadro constant Nₐ ≈ 6.022×10²³ mol⁻¹", []string{"physics", "constant", "avogadro", "core", "bootstrap"}},

		// Массы и заряды
		{"const_mass_electron", "mass of electron ≈ 9.109×10⁻³¹ kg", []string{"physics", "constant", "mass", "electron", "core", "bootstrap"}},
		{"const_mass_proton", "mass of proton ≈ 1.673×10⁻²⁷ kg", []string{"physics", "constant", "mass", "proton", "core", "bootstrap"}},

		// Температурные пределы
		{"const_absolute_zero", "absolute zero = 0 K = −273.15 °C", []string{"physics", "constant", "temperature", "core", "bootstrap"}},

		// Прочие полезные
		{"const_gas", "ideal gas constant R ≈ 8.314 J/(mol·K)", []string{"physics", "constant", "gas", "core", "bootstrap"}},
		{"const_mu0", "magnetic constant μ₀ ≈ 4π×10⁻⁷ N/A²", []string{"physics", "constant", "magnetism", "core", "bootstrap"}},
		{"const_epsilon0", "electric constant ε₀ ≈ 8.854×10⁻¹² F/m", []string{"physics", "constant", "electric", "core", "bootstrap"}},
	}

	for _, c := range constants {
		q := QBit{
			ID:        c.ID,
			Content:   c.Content,
			Tags:      c.Tags,
			Phase:     0.92,
			Weight:    0.99,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("🔬 [Bootstrap] Physical constants loaded.")
}
