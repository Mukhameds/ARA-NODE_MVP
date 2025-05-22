// core/bootstrap_semantics.go
package core

import (
	"fmt"
	"time"
)

// BootstrapSemanticLinks загружает основные семантические связи между понятиями
func BootstrapSemanticLinks(mem *MemoryEngine) {
	relations := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// cause-effect
		{"sem_fire_burn", "fire causes burn", []string{"semantic", "cause-effect", "core", "bootstrap"}},
		{"sem_eat_full", "eating leads to fullness", []string{"semantic", "cause-effect", "core", "bootstrap"}},

		// is-a
		{"sem_cat_animal", "cat is an animal", []string{"semantic", "is-a", "category", "core", "bootstrap"}},
		{"sem_rose_flower", "rose is a flower", []string{"semantic", "is-a", "category", "core", "bootstrap"}},

		// part-of
		{"sem_leaf_tree", "leaf is part of tree", []string{"semantic", "part-of", "core", "bootstrap"}},
		{"sem_wheel_car", "wheel is part of car", []string{"semantic", "part-of", "core", "bootstrap"}},

		// synonym
		{"sem_big_large", "big means large", []string{"semantic", "synonym", "core", "bootstrap"}},
		{"sem_small_tiny", "small means tiny", []string{"semantic", "synonym", "core", "bootstrap"}},

		// antonym
		{"sem_hot_cold", "hot is opposite of cold", []string{"semantic", "antonym", "core", "bootstrap"}},
		{"sem_light_dark", "light is opposite of dark", []string{"semantic", "antonym", "core", "bootstrap"}},
	}

	for _, r := range relations {
		q := QBit{
			ID:        r.ID,
			Content:   r.Content,
			Tags:      r.Tags,
			Phase:     0.82,
			Weight:    0.95,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("🔗 [Bootstrap] Semantic relationships loaded.")
}
