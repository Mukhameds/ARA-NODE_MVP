package core

import "time"

// Signal — входящий или внутренний сигнал системы
type Signal struct {
	ID        string
	Content   string
	Tags      []string
	Timestamp time.Time
	Phase     float64  // фазовое значение сигнала
	Weight    float64  // значимость / энергия сигнала
	Origin    string   // источник: "user", "memory", "system", и т.п.
}

// QBit — единица сигнальной памяти
type QBit struct {
	ID        string
	Content   string
	Tags      []string
	CreatedAt time.Time
	Weight    float64
	Phase     float64
	Archived  bool
}

// Reaction — результат реакции на сигнал
type Reaction struct {
	TriggeredBy string   // ID сигнала
	Response    string   // текст или команда
	Tags        []string
	Confidence  float64  // степень уверенности
}

// PhaseMatch — фазовое совпадение для генерации фантома
type PhaseMatch struct {
	Delta     float64  // разница фаз
	Threshold float64  // порог совпадения
}
