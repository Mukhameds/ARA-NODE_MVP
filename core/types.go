package core

import "time"

// Signal — входной сигнал, возбуждающий реакцию
type Signal struct {
	ID        string
	Content   string
	Tags      []string
	Timestamp time.Time
	Phase     float64
	Weight    float64
	Origin    string  // источник: user, instinct, prediction, background
	Type      string  // тип сигнала: user, instinct, prediction, background, etc.
}

// QBit — единица памяти (узел в памяти ARA)
type QBit struct {
	ID           string
	Content      string
	Tags         []string
	Type         string  // тип узла: standard, reflex, emotion, etc.
	Phase        float64
	Weight       float64
	Archived     bool
	Origin       string
	CreatedAt    time.Time
	LastAccessed time.Time
}

// Reaction — результат обработки сигнала
type Reaction struct {
	TriggeredBy string
	Response    string
	Tags        []string
	Confidence  float64
}

// FanthomInterface — интерфейс для фантомных генераторов
type FanthomInterface interface {
	TriggerFromMatch(sig Signal)
}

// PhantomLog — для построения дерева фантомов
type PhantomLog struct {
	PhantomID string
	SourceIDs []string
}


func (q *QBit) AgeFrame() string {
	age := time.Since(q.CreatedAt).Seconds()
	switch {
	case age < 60:
		return "emergent"
	case age < 600:
		return "forming"
	case age < 3600:
		return "mature"
	default:
		return "legacy"
	}
}
