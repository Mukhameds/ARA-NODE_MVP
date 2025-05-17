package core

import (
	"strings"
)

// VariableBlock — сигнальная единица восприятия (буква, слово, символ, образ и т.д.)
type VariableBlock struct {
	ID     string   // Уникальный ID (например, "A", "hello", "5")
	Signal string   // Визуальное представление или образ
	Tags   []string // Типы: letter, word, number, image...
	Reacts []string // Синонимы, триггеры
	QBit   *QBit    // Связанный элемент памяти
	Auto   bool     // Создан автоматически?
}

// SignalDictionary — глобальный словарь восприятия
type SignalDictionary struct {
	Variables map[string]*VariableBlock
	Memory    *MemoryEngine
}

// NewSignalDictionary — создать словарь
func NewSignalDictionary(mem *MemoryEngine) *SignalDictionary {
	return &SignalDictionary{
		Variables: make(map[string]*VariableBlock),
		Memory:    mem,
	}
}

// FindMatch — поиск VariableBlock по токену
func (sd *SignalDictionary) FindMatch(token string) *VariableBlock {
	for _, vb := range sd.Variables {
		for _, r := range vb.Reacts {
			if strings.EqualFold(token, r) {
				return vb
			}
		}
	}
	return nil
}

// AutoLearn — создать новый VariableBlock из незнакомого токена
func (sd *SignalDictionary) AutoLearn(token string) *VariableBlock {
	vb := &VariableBlock{
		ID:     token,
		Signal: token,
		Tags:   []string{"type:unknown"},
		Reacts: []string{token},
		QBit:   sd.Memory.CreateQBit(token),
		Auto:   true,
	}
	sd.Variables[token] = vb
	return vb
}

// Add — ручное добавление блока
func (sd *SignalDictionary) Add(id, signal string, tags, reacts []string) *VariableBlock {
	vb := &VariableBlock{
		ID:     id,
		Signal: signal,
		Tags:   tags,
		Reacts: reacts,
		QBit:   sd.Memory.CreateQBit(signal),
		Auto:   false,
	}
	sd.Variables[id] = vb
	return vb
}

// Delete — удалить VariableBlock
func (sd *SignalDictionary) Delete(id string) bool {
	if _, ok := sd.Variables[id]; ok {
		delete(sd.Variables, id)
		return true
	}
	return false
}

// Tag — добавить тег к блоку
func (sd *SignalDictionary) Tag(id, tag string) {
	if vb, ok := sd.Variables[id]; ok {
		vb.Tags = append(vb.Tags, tag)
	}
}

// All — список всех блоков
func (sd *SignalDictionary) All() []*VariableBlock {
	out := []*VariableBlock{}
	for _, vb := range sd.Variables {
		out = append(out, vb)
	}
	return out
}
