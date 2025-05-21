package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"sync"
	"time"
)

// MemoryEngine — сигнальное хранилище ARA
type MemoryEngine struct {
	QBits       map[string]QBit
	PhantomTree []PhantomLog
	mutex       sync.Mutex
}

func NewMemoryEngine() *MemoryEngine {
	return &MemoryEngine{
		QBits:       make(map[string]QBit),
		PhantomTree: []PhantomLog{},
	}
}

func (m *MemoryEngine) CreateQBit(content string) *QBit {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	q := QBit{
		ID:           fmt.Sprintf("qbit_%d", time.Now().UnixNano()),
		Content:      content,
		CreatedAt:    time.Now(),
		LastAccessed: time.Now(),
		Phase:        0.7,
		Weight:       1.0,
		Tags:         []string{},
		Type:         "",
	}
	m.QBits[q.ID] = q
	return &q
}

func (m *MemoryEngine) StoreQBit(q QBit) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	q.LastAccessed = time.Now()
	m.QBits[q.ID] = q
}

func (m *MemoryEngine) UpdateQBit(q QBit) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	q.LastAccessed = time.Now()
	m.QBits[q.ID] = q
}

func (m *MemoryEngine) FindByTag(tag string) []QBit {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	results := []QBit{}
	for _, q := range m.QBits {
		if Contains(q.Tags, tag) {
			q.LastAccessed = time.Now()
			results = append(results, q)
		}
	}
	return results
}

func (m *MemoryEngine) FindByPhase(phase float64, tolerance float64) []QBit {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	results := []QBit{}
	for _, q := range m.QBits {
		if PhaseClose(q.Phase, phase, tolerance) {
			q.LastAccessed = time.Now()
			results = append(results, q)
		}
	}
	return results
}

func (m *MemoryEngine) FindTopRelevant(tag string, minPhase float64) []QBit {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var filtered []QBit
	for _, q := range m.QBits {
		if Contains(q.Tags, tag) && q.Phase >= minPhase {
			filtered = append(filtered, q)
		}
	}
	sort.SliceStable(filtered, func(i, j int) bool {
		return filtered[i].Weight*filtered[i].Phase > filtered[j].Weight*filtered[j].Phase
	})
	return filtered
}

func (m *MemoryEngine) FindAll(filter func(q QBit) bool) []QBit {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var all []QBit
	for _, q := range m.QBits {
		if filter(q) {
			q.LastAccessed = time.Now()
			all = append(all, q)
		}
	}
	return all
}

func (m *MemoryEngine) AdjustWeight(id string, delta float64) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	q := m.QBits[id]
	q.Weight += delta
	if q.Weight < 0 {
		q.Weight = 0
	}
	q.LastAccessed = time.Now()
	m.QBits[id] = q
}

func (m *MemoryEngine) AddTag(id, tag string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	q := m.QBits[id]
	if !Contains(q.Tags, tag) {
		q.Tags = append(q.Tags, tag)
	}
	m.QBits[id] = q
}

func (m *MemoryEngine) DeleteQBit(id string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.QBits, id)
	fmt.Println("[MemoryEngine] ❌ QBit deleted:", id)
}

func (m *MemoryEngine) DeleteByTag(tag string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	count := 0
	for id, q := range m.QBits {
		if Contains(q.Tags, tag) {
			delete(m.QBits, id)
			count++
		}
	}
	fmt.Println("[MemoryEngine] 🧹 Deleted QBits by tag:", tag, "→", count)
}

func (m *MemoryEngine) ExistsQBit(content string, phase float64, tol float64) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, q := range m.QBits {
		if q.Content == content && PhaseClose(q.Phase, phase, tol) {
			return true
		}
	}
	return false
}

func (m *MemoryEngine) DecayQBits() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	now := time.Now()
	for id, q := range m.QBits {
		if q.Type == "standard" || Contains(q.Tags, "instinct") {
			continue
		}
		age := now.Sub(q.LastAccessed).Seconds()
		decay := 1.0 - (age / 1800.0) // затухает за 30 минут
		if decay < 0.5 {
			q.Weight *= decay
			if q.Weight < 0.05 {
				delete(m.QBits, id)
				continue
			}
			m.QBits[id] = q
		}
	}
}

func (m *MemoryEngine) ListQBits() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	fmt.Println("\n🧠 Current Memory:")
	for _, q := range m.QBits {
		fmt.Printf("%s | %.2f | %s | Tags: %v\n", q.ID, q.Phase, q.Content, q.Tags)
	}
}

func (m *MemoryEngine) Merge(other *MemoryEngine) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for id, q := range other.QBits {
		if _, exists := m.QBits[id]; !exists {
			m.QBits[id] = q
		}
	}
	fmt.Println("[MemoryEngine] ✅ Merged external memory:", len(other.QBits), "entries")
}

func (m *MemoryEngine) ExportJSON(filename string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	data, err := json.MarshalIndent(m.QBits, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (m *MemoryEngine) ImportJSON(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	var imported map[string]QBit
	if err := json.Unmarshal(bytes, &imported); err != nil {
		return err
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for id, q := range imported {
		if _, exists := m.QBits[id]; !exists {
			m.QBits[id] = q
		}
	}
	fmt.Println("[MemoryEngine] 📥 Imported QBits:", len(imported))
	return nil
}

// EstimateTotalPhase — возвращает среднюю фазу всех активных QBits
func (m *MemoryEngine) EstimateTotalPhase() float64 {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if len(m.QBits) == 0 {
		return 0.0
	}

	var sum float64
	var count int
	now := time.Now()
	for _, q := range m.QBits {
		if now.Sub(q.LastAccessed) < 2*time.Minute {
			sum += q.Phase
			count++
		}
	}
	if count == 0 {
		return 0.0
	}
	return sum / float64(count)
}
