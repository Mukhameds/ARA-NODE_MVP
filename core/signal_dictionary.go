package core

import (
	"strings"
	"time"
)

// VariableBlock — сигнальная единица восприятия (буква, слово, символ, образ и т.д.)
type VariableBlock struct {
	ID     string
	Signal string
	Tags   []string
	Reacts []string
	QBit   *QBit
	Auto   bool
}

// SignalDictionary — глобальный словарь восприятия + буквенный буфер
type SignalDictionary struct {
	Variables map[string]*VariableBlock
	Memory    *MemoryEngine
	buffer    []string
	lastUsed  time.Time
}

// NewSignalDictionary — инициализация
func NewSignalDictionary(mem *MemoryEngine) *SignalDictionary {
	return &SignalDictionary{
		Variables: make(map[string]*VariableBlock),
		Memory:    mem,
		buffer:    []string{},
		lastUsed:  time.Now(),
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

// AutoLearn — создать блок из неизвестного токена
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

// Add — ручное добавление
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

// Delete — удалить
func (sd *SignalDictionary) Delete(id string) bool {
	if _, ok := sd.Variables[id]; ok {
		delete(sd.Variables, id)
		return true
	}
	return false
}

// Tag — добавить тег
func (sd *SignalDictionary) Tag(id, tag string) {
	if vb, ok := sd.Variables[id]; ok {
		vb.Tags = append(vb.Tags, tag)
	}
}

// All — получить все блоки
func (sd *SignalDictionary) All() []*VariableBlock {
	out := []*VariableBlock{}
	for _, vb := range sd.Variables {
		out = append(out, vb)
	}
	return out
}

// LearnFromInput — обучение по словам
func (sd *SignalDictionary) LearnFromInput(input string) {
	tokens := strings.Fields(input)
	for _, tok := range tokens {
		if sd.FindMatch(tok) == nil {
			sd.AutoLearn(tok)
		}
		sd.LearnLetters(tok)
	}
}

// LearnLetters — по буквам
func (sd *SignalDictionary) LearnLetters(word string) {
	for _, ch := range word {
		letter := strings.ToUpper(string(ch))
		if letter == " " || len(letter) == 0 {
			continue
		}
		sd.AddLetter(letter)
	}
}

// AddLetter — добавляет букву в буфер и создаёт QBit
func (sd *SignalDictionary) AddLetter(letter string) {
	if _, exists := sd.Variables[letter]; !exists {
		vb := &VariableBlock{
			ID:     letter,
			Signal: letter,
			Tags:   []string{"char", "letter"},
			Reacts: []string{letter},
			QBit:   sd.Memory.CreateQBit(letter),
			Auto:   true,
		}
		sd.Variables[letter] = vb
	}

	sd.buffer = append(sd.buffer, letter)
	if len(sd.buffer) > 8 {
		sd.buffer = sd.buffer[1:]
	}

	sd.CheckFormedWord()
}

// CheckFormedWord — если буфер повторяется как слово, создаём QBit
func (sd *SignalDictionary) CheckFormedWord() {
	word := strings.Join(sd.buffer, "")
	if len(word) < 3 {
		return
	}
	// Если слово уже есть — не дублируем
	if _, exists := sd.Variables[word]; exists {
		return
	}
	vb := &VariableBlock{
		ID:     word,
		Signal: word,
		Tags:   []string{"word", "formed"},
		Reacts: []string{word},
		QBit:   sd.Memory.CreateQBit(word),
		Auto:   true,
	}
	sd.Variables[word] = vb
}

// Buffer — получить текущий буквенный буфер
func (sd *SignalDictionary) Buffer() []string {
	return sd.buffer
}

