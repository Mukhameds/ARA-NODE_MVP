package field

import (
	"fmt"
	"ara-node/core"
)

// GhostRocket — управляет множественными реакционными полями
type GhostRocket struct {
	Fields []*Matrix
	Name   string
}

// NewGhostRocket — создаёт новую "ракету" из полей
func NewGhostRocket(name string) *GhostRocket {
	return &GhostRocket{
		Name:   name,
		Fields: []*Matrix{},
	}
}

// AddField — подключает новое реакционное поле
func (r *GhostRocket) AddField(matrix *Matrix) {
	r.Fields = append(r.Fields, matrix)
	fmt.Printf("[Rocket:%s] 🚀 Field '%s' added.\n", r.Name, matrix.Name)
}

// Propagate — одновременно распространяет сигнал по всем полям
func (r *GhostRocket) Propagate(sig core.Signal) {
	fmt.Printf("[Rocket:%s] 🚀 Propagating signal: %s (%v)\n", r.Name, sig.ID, sig.Tags)
	for _, f := range r.Fields {
		go f.Propagate(sig)
	}
}

// ListFields — отладочный вывод подключённых полей
func (r *GhostRocket) ListFields() {
	fmt.Printf("[Rocket:%s] 🌌 Connected Fields:\n", r.Name)
	for _, f := range r.Fields {
		fmt.Printf("- %s (%d blocks)\n", f.Name, len(f.Blocks))
	}
}
