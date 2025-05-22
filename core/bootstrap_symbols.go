// core/bootstrap_symbols.go
package core

import (
	"fmt"
	"time"
)

// BootstrapSymbolArchitecture загружает сигнальные элементы (линии, дуги) и структурные схемы букв
func BootstrapSymbolArchitecture(mem *MemoryEngine) {
	// === Примитивы: цифровые штрихи и паттерны ===
	primitives := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		{"stroke_vertical", "|", []string{"primitive", "stroke", "line", "core"}},
		{"stroke_horizontal", "-", []string{"primitive", "stroke", "line", "core"}},
		{"stroke_slash", "/", []string{"primitive", "stroke", "diagonal", "core"}},
		{"stroke_backslash", "\\", []string{"primitive", "stroke", "diagonal", "core"}},
		{"stroke_circle", "○", []string{"primitive", "shape", "circle", "core"}},
		{"stroke_halfcircle", "◔", []string{"primitive", "shape", "curve", "core"}},
		{"stroke_cross", "+", []string{"primitive", "shape", "intersection", "core"}},
		{"stroke_dot", "•", []string{"primitive", "point", "dot", "core"}},
	}

	for _, p := range primitives {
		q := QBit{
			ID:        p.ID,
			Content:   p.Content,
			Tags:      append(p.Tags, "bootstrap", "symbol"),
			Phase:     0.6,
			Weight:    0.8,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	// === Буквы как архитектуры из примитивов ===
	letters := []struct {
		ID        string
		Letter    string
		Structure []string
	}{
		{"letter_A", "A", []string{"/", "\\", "-"}},
		{"letter_B", "B", []string{"|", "◔", "◔"}},
		{"letter_C", "C", []string{"◔"}},
		{"letter_D", "D", []string{"|", "◔"}},
		{"letter_E", "E", []string{"|", "-", "-", "-"}},
		{"letter_F", "F", []string{"|", "-", "-"}},
		{"letter_H", "H", []string{"|", "|", "-"}},
		{"letter_I", "I", []string{"|"}},
		{"letter_L", "L", []string{"|", "-"}},
		{"letter_O", "O", []string{"○"}},
		{"letter_T", "T", []string{"-", "|"}},
		{"letter_X", "X", []string{"/", "\\"}},
	}

	for _, l := range letters {
		q := QBit{
			ID:        l.ID,
			Content:   l.Letter,
			Tags:      []string{"letter", "structure", "core", "symbol", "bootstrap"},
			Phase:     0.75,
			Weight:    0.9,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)

		// Добавим отдельный QBit с описанием архитектуры
		desc := fmt.Sprintf("Letter %s = %s", l.Letter, fmt.Sprintf("%v", l.Structure))
		structQ := QBit{
			ID:        l.ID + "_structure",
			Content:   desc,
			Tags:      []string{"architecture", "composition", "core", "bootstrap"},
			Phase:     0.7,
			Weight:    0.85,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(structQ)
	}

	fmt.Println("🔤 [Bootstrap] Symbol primitives and letter structures loaded.")
}
