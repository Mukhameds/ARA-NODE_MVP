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

// Strength — сила сигнала (для WillEngine, FanthomEngine и др.)
func (s *Signal) Strength() float64 {
	return s.Phase * s.Weight
}

// HasTag — проверка наличия тега
func (s *Signal) HasTag(tag string) bool {
	for _, t := range s.Tags {
		if t == tag {
			return true
		}
	}
	return false
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

// Strength — сила узла в памяти
func (q *QBit) Strength() float64 {
	return q.Phase * q.Weight
}

// HasTag — проверка наличия тега
func (q *QBit) HasTag(tag string) bool {
	for _, t := range q.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// AgeFrame — возраст узла (семантический)
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

// DecayFactor — коэффициент затухания по возрасту
func (q *QBit) DecayFactor() float64 {
	age := time.Since(q.CreatedAt).Seconds()
	switch {
	case age < 60:
		return 1.0
	case age < 600:
		return 0.9
	case age < 3600:
		return 0.7
	default:
		return 0.4
	}
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


// GhostLike — интерфейс реактивного поля
type GhostLike interface {
	Propagate(sig Signal)
}


// PhantomLog — для построения дерева фантомов
type PhantomLog struct {
	PhantomID string
	SourceIDs []string
}

func SignalFromQBit(q QBit) Signal {
	return Signal{
		ID:        "sig_" + q.ID,
		Content:   q.Content,
		Tags:      q.Tags,
		Timestamp: time.Now(),
		Phase:     q.Phase,
		Weight:    q.Weight,
		Type:      "phantom",
		Origin:    "phantom_engine",
	}
}
