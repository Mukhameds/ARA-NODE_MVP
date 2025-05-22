// core/bootstrap_physics.go
package core

import (
	"fmt"
	"time"
)

// BootstrapPhysicsConcepts загружает фундаментальные физические понятия в виде сигнальных QBits
func BootstrapPhysicsConcepts(mem *MemoryEngine) {
	concepts := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Материя и свойства
		{"phys_matter", "matter is anything that has mass and volume", []string{"physics", "matter", "definition", "core", "bootstrap"}},
		{"phys_mass", "mass is the amount of matter in an object", []string{"physics", "mass", "property", "core", "bootstrap"}},
		{"phys_volume", "volume is the amount of space an object occupies", []string{"physics", "volume", "property", "core", "bootstrap"}},

		// Сила и движение
		{"phys_force", "force causes change in motion", []string{"physics", "force", "core", "bootstrap"}},
		{"phys_gravity", "gravity pulls objects toward each other", []string{"physics", "gravity", "core", "bootstrap"}},
		{"phys_motion", "motion is change of position over time", []string{"physics", "motion", "core", "bootstrap"}},
		{"phys_speed", "speed is distance divided by time", []string{"physics", "speed", "definition", "core", "bootstrap"}},

		// Энергия и тепло
		{"phys_energy", "energy is the ability to do work", []string{"physics", "energy", "core", "bootstrap"}},
		{"phys_kinetic", "kinetic energy is energy of motion", []string{"physics", "energy", "kinetic", "core", "bootstrap"}},
		{"phys_potential", "potential energy is stored energy", []string{"physics", "energy", "potential", "core", "bootstrap"}},
		{"phys_heat", "heat is transfer of thermal energy", []string{"physics", "heat", "core", "bootstrap"}},
		{"phys_temperature", "temperature measures average kinetic energy", []string{"physics", "temperature", "core", "bootstrap"}},

		// Свет и волны
		{"phys_light", "light is electromagnetic radiation visible to the eye", []string{"physics", "light", "wave", "core", "bootstrap"}},
		{"phys_sound", "sound is a wave caused by vibration", []string{"physics", "sound", "wave", "core", "bootstrap"}},
		{"phys_wave", "a wave carries energy through space", []string{"physics", "wave", "core", "bootstrap"}},

		// Состояния вещества
		{"phys_solid", "a solid has definite shape and volume", []string{"physics", "state", "solid", "core", "bootstrap"}},
		{"phys_liquid", "a liquid has definite volume but not shape", []string{"physics", "state", "liquid", "core", "bootstrap"}},
		{"phys_gas", "a gas has no fixed shape or volume", []string{"physics", "state", "gas", "core", "bootstrap"}},
	}

	for _, c := range concepts {
		q := QBit{
			ID:        c.ID,
			Content:   c.Content,
			Tags:      c.Tags,
			Phase:     0.87,
			Weight:    0.95,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("🌌 [Bootstrap] Physical concepts loaded.")
}
