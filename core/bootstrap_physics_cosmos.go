// core/bootstrap_physics_cosmos.go
package core

import (
	"fmt"
	"time"
)

// BootstrapPhysicsCosmos загружает сигнальные QBits по астрофизике и структуре Вселенной
func BootstrapPhysicsCosmos(mem *MemoryEngine) {
	cosmos := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Вселенная и её структура
		{"cosmos_universe", "the universe contains all of space, time, matter, and energy", []string{"physics", "cosmos", "universe", "core", "bootstrap"}},
		{"cosmos_galaxy", "a galaxy is a massive system of stars, gas, and dark matter", []string{"physics", "cosmos", "galaxy", "core", "bootstrap"}},
		{"cosmos_star", "a star is a massive glowing sphere of plasma", []string{"physics", "cosmos", "star", "core", "bootstrap"}},
		{"cosmos_planet", "a planet orbits a star and may have moons", []string{"physics", "cosmos", "planet", "core", "bootstrap"}},
		{"cosmos_solar_system", "our solar system includes the sun and all objects orbiting it", []string{"physics", "cosmos", "solar", "core", "bootstrap"}},

		// Гравитация и движение
		{"cosmos_gravity_scale", "gravity governs motion of planets, stars, and galaxies", []string{"physics", "cosmos", "gravity", "core", "bootstrap"}},
		{"cosmos_orbit", "an orbit is a stable curved path under gravity", []string{"physics", "cosmos", "orbit", "core", "bootstrap"}},
		{"cosmos_escape_velocity", "escape velocity is the speed needed to overcome gravity", []string{"physics", "cosmos", "velocity", "escape", "core", "bootstrap"}},

		// Расширение и происхождение
		{"cosmos_big_bang", "the universe began from a dense hot state — the Big Bang", []string{"physics", "cosmos", "origin", "bigbang", "core", "bootstrap"}},
		{"cosmos_expansion", "the universe is expanding — galaxies move away over time", []string{"physics", "cosmos", "expansion", "core", "bootstrap"}},
		{"cosmos_redshift", "redshift measures how much light stretches as galaxies recede", []string{"physics", "cosmos", "redshift", "core", "bootstrap"}},

		// Наблюдаемая и скрытая материя
		{"cosmos_dark_matter", "dark matter does not emit light but affects gravity", []string{"physics", "cosmos", "darkmatter", "core", "bootstrap"}},
		{"cosmos_dark_energy", "dark energy causes accelerated expansion of the universe", []string{"physics", "cosmos", "darkenergy", "core", "bootstrap"}},
	}

	for _, c := range cosmos {
		qbit := QBit{
			ID:        c.ID,
			Content:   c.Content,
			Tags:      c.Tags,
			Phase:     0.92,
			Weight:    0.98,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(qbit)
	}

	fmt.Println("🌌 [Bootstrap] Cosmos and astrophysics concepts loaded.")
}
