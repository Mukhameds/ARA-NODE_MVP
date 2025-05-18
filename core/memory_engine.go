package core

import (
	"fmt"
	"sync"
	"time"
)

// MemoryEngine — сигнальная память агента
type MemoryEngine struct {
	QBits       map[string]QBit
	Mu          sync.Mutex
	PhantomTree []PhantomLog
}

// NewMemoryEngine — инициализация памяти
func NewMemoryEngine() *MemoryEngine {
	return &MemoryEngine{
		QBits: make(map[string]QBit),
	}
}

// StoreQBit — сохранить QBit в память
func (m *MemoryEngine) StoreQBit(q QBit) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	m.QBits[q.ID] = q
	fmt.Println("[MemoryEngine] Stored QBit:", q.ID)
}

// GetQBit — получить QBit по ID
func (m *MemoryEngine) GetQBit(id string) (QBit, bool) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	q, exists := m.QBits[id]
	return q, exists
}

// DeleteQBit — удалить QBit по ID
func (m *MemoryEngine) DeleteQBit(id string) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	if _, exists := m.QBits[id]; exists {
		delete(m.QBits, id)
		fmt.Println("[MemoryEngine] ❌ QBit deleted:", id)
	}
}

// DecayQBits — уменьшает вес старых или слабых QBit
func (m *MemoryEngine) DecayQBits() {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	now := time.Now()
	for id, q := range m.QBits {
		age := now.Sub(q.CreatedAt).Seconds()
		decayFactor := 0.5 * age
		q.Weight -= decayFactor
		if q.Weight < 0 {
			q.Weight = 0
		}
		if q.Weight < 0.2 {
			q.Archived = true
		}
		fmt.Printf("[Decay] %s → age: %.1fs, decay: %.2f, new weight: %.2f\n", id, age, decayFactor, q.Weight)
		m.QBits[id] = q
	}
}

// FindByTag — вернуть QBit по тегу
func (m *MemoryEngine) FindByTag(tag string) []QBit {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	var result []QBit
	for _, q := range m.QBits {
		if q.Archived {
			continue
		}
		for _, t := range q.Tags {
			if t == tag {
				result = append(result, q)
				break
			}
		}
	}
	return result
}

// FindByPhase — вернуть QBit с близкой фазой
func (m *MemoryEngine) FindByPhase(target float64, tolerance float64) []QBit {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	var result []QBit
	for _, q := range m.QBits {
		if q.Archived {
			continue
		}
		if abs(q.Phase-target) <= tolerance {
			result = append(result, q)
		}
	}
	return result
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

// ListQBits — выводит все неархивированные QBit
func (m *MemoryEngine) ListQBits() {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	fmt.Println("\n[Memory Dump]")
	for _, q := range m.QBits {
		if q.Archived {
			continue
		}
		fmt.Printf("• ID: %s | Phase: %.2f | Weight: %.2f | Tags: %v\n", q.ID, q.Phase, q.Weight, q.Tags)
	}
}

func (m *MemoryEngine) AdjustWeight(id string, delta float64) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	q, exists := m.QBits[id]
	if !exists || q.Archived {
		return
	}
	q.Weight += delta
	if q.Weight < 0 {
		q.Weight = 0
	}
	m.QBits[id] = q
}

func (m *MemoryEngine) AddTag(id string, tag string) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	q, exists := m.QBits[id]
	if !exists || q.Archived {
		return
	}
	for _, t := range q.Tags {
		if t == tag {
			return
		}
	}
	q.Tags = append(q.Tags, tag)
	m.QBits[id] = q
}

func (m *MemoryEngine) Merge(other map[string]QBit) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	for id, q := range other {
		if _, exists := m.QBits[id]; !exists {
			m.QBits[id] = q
		}
	}
}

// CreateQBit — создать и сохранить новый QBit
func (m *MemoryEngine) CreateQBit(content string) *QBit {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	q := QBit{
		ID:        fmt.Sprintf("qbit_%d", time.Now().UnixNano()),
		Content:   content,
		Tags:      []string{"auto"},
		CreatedAt: time.Now(),
		Phase:     0.75,
		Weight:    1.0,
	}

	m.QBits[q.ID] = q
	fmt.Println("[MemoryEngine] Auto-created QBit:", q.ID)
	return &q
}

// Broadcast — возбуждает QBit во всей памяти (прямая трансляция)
func (m *MemoryEngine) Broadcast(q *QBit) {
	if q == nil {
		return
	}
	m.StoreQBit(*q)
	fmt.Println("[MemoryEngine] 📡 Broadcast QBit:", q.ID)
}

// FindAll — вернуть все QBits, удовлетворяющие фильтру
func (m *MemoryEngine) FindAll(filter func(QBit) bool) []QBit {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	var result []QBit
	for _, q := range m.QBits {
		if filter(q) {
			result = append(result, q)
		}
	}
	return result
}

func (m *MemoryEngine) UpdateQBit(qbit QBit) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	if _, exists := m.QBits[qbit.ID]; exists {
		m.QBits[qbit.ID] = qbit
	}
}
