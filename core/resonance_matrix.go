package core

import (
	"fmt"
	"sync"
	"time"
)

type ResonanceLink struct {
	From     string
	To       string
	Strength float64
	LastSeen time.Time
}

type ResonanceMatrix struct {
	Links map[string]map[string]*ResonanceLink
	Mutex sync.Mutex
}

func NewResonanceMatrix() *ResonanceMatrix {
	return &ResonanceMatrix{
		Links: make(map[string]map[string]*ResonanceLink),
	}
}

func (rm *ResonanceMatrix) RegisterPair(q1, q2 QBit) {
	if q1.ID == q2.ID {
		return
	}

	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()

	if rm.Links[q1.ID] == nil {
		rm.Links[q1.ID] = make(map[string]*ResonanceLink)
	}
	if rm.Links[q2.ID] == nil {
		rm.Links[q2.ID] = make(map[string]*ResonanceLink)
	}

	link := rm.Links[q1.ID][q2.ID]
	if link == nil {
		link = &ResonanceLink{From: q1.ID, To: q2.ID, Strength: 0.1, LastSeen: time.Now()}
		rm.Links[q1.ID][q2.ID] = link
		rm.Links[q2.ID][q1.ID] = link
	} else {
		link.Strength += 0.05
		if link.Strength > 1.0 {
			link.Strength = 1.0
		}
		link.LastSeen = time.Now()
	}
}

func (rm *ResonanceMatrix) Decay() {
	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()
	now := time.Now()
	for from, neighbors := range rm.Links {
		for to, link := range neighbors {
			if now.Sub(link.LastSeen) > 30*time.Second {
				link.Strength *= 0.95
				if link.Strength < 0.01 {
					delete(rm.Links[from], to)
				}
			}
		}
	}
}

func (rm *ResonanceMatrix) BoostBySignal(sig Signal, qbits []QBit) {
	for i := 0; i < len(qbits); i++ {
		for j := i + 1; j < len(qbits); j++ {
			rm.RegisterPair(qbits[i], qbits[j])
		}
	}
}

func (rm *ResonanceMatrix) GetStrongLinks(id string) []string {
	rm.Mutex.Lock()
	defer rm.Mutex.Unlock()
	var links []string
	for to, link := range rm.Links[id] {
		if link.Strength >= 0.5 {
			links = append(links, to)
		}
	}
	return links
}

func (rm *ResonanceMatrix) Print(id string) {
	for _, to := range rm.GetStrongLinks(id) {
		fmt.Println("🔗", id, "⇄", to)
	}
}
