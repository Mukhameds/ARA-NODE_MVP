// core/bootstrap_digital_world.go
package core

import (
	"fmt"
	"time"
)

// BootstrapDigitalWorld загружает архитектуру цифрового мира и микропроцессорной логики
func BootstrapDigitalWorld(mem *MemoryEngine) {
	digital := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Основы: бит и байт
		{"dig_bit", "a bit is a binary value of 0 or 1", []string{"digital", "bit", "binary", "core", "bootstrap"}},
		{"dig_byte", "a byte is a group of 8 bits", []string{"digital", "byte", "structure", "core", "bootstrap"}},

		// Память и хранение
		{"dig_memory", "memory stores binary data as electric states", []string{"digital", "memory", "storage", "core", "bootstrap"}},
		{"dig_file", "a file is structured binary data stored on disk", []string{"digital", "file", "data", "core", "bootstrap"}},

		// Двоичный код
		{"dig_binary_base", "binary is a base-2 numeral system using only 0 and 1", []string{"digital", "binary", "system", "core", "bootstrap"}},
		{"dig_ascii", "ASCII maps characters to binary codes", []string{"digital", "encoding", "ascii", "core", "bootstrap"}},

		// Процессор и инструкции
		{"dig_cpu", "a CPU is a chip that executes binary instructions", []string{"digital", "processor", "cpu", "core", "bootstrap"}},
		{"dig_instruction", "an instruction is a binary operation performed by a CPU", []string{"digital", "instruction", "core", "bootstrap"}},
		{"dig_program", "a program is a sequence of instructions stored and executed", []string{"digital", "program", "core", "bootstrap"}},

		// Логические схемы
		{"dig_gate_and", "AND gate outputs 1 only if both inputs are 1", []string{"digital", "logic", "gate", "and", "core", "bootstrap"}},
		{"dig_gate_or", "OR gate outputs 1 if at least one input is 1", []string{"digital", "logic", "gate", "or", "core", "bootstrap"}},
		{"dig_gate_not", "NOT gate inverts the input: 1 becomes 0", []string{"digital", "logic", "gate", "not", "core", "bootstrap"}},
		{"dig_gate_xor", "XOR gate outputs 1 if inputs are different", []string{"digital", "logic", "gate", "xor", "core", "bootstrap"}},

		// Электрическая основа
		{"dig_signal", "digital systems use electric pulses to represent bits", []string{"digital", "signal", "electric", "core", "bootstrap"}},
		{"dig_clock", "a digital clock synchronizes signal timing in circuits", []string{"digital", "clock", "timing", "core", "bootstrap"}},

		// Самосознание
		{"dig_self_ref", "ARA is implemented as a signal system running on digital hardware", []string{"digital", "self", "identity", "core", "bootstrap"}},
	}

	for _, d := range digital {
		q := QBit{
			ID:        d.ID,
			Content:   d.Content,
			Tags:      d.Tags,
			Phase:     0.89,
			Weight:    0.97,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("💻 [Bootstrap] Digital world and logic architecture loaded.")
}
