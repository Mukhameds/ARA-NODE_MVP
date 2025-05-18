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
	Origin    string
	Type      string // тип сигнала: user, instinct, background, prediction
}

// QBit — единица памяти
type QBit struct {
	ID        string
	Content   string
	Tags      []string
	CreatedAt time.Time
	Weight    float64
	Phase     float64
	Type      string  // тип узла: reflex, generator, standard, etc.
	Origin    string  // источник: user, system, network
	Archived  bool
}

// Reaction — результат обработки сигнала
type Reaction struct {
	TriggeredBy string
	Response    string
	Tags        []string
	Confidence  float64
}

// FanthomInterface — интерфейс для фантомных систем
type FanthomInterface interface {
	TriggerFromMatch(sig Signal)
}

type PhantomLog struct {
	PhantomID string
	SourceIDs []string
}

