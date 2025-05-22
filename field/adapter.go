package field

import "ara-node/core"

// ghostAdapter — адаптер между GhostRocket и интерфейсом GhostLike
type ghostAdapter struct {
	rocket *GhostRocket
}

// Propagate реализует core.GhostLike
func (g *ghostAdapter) Propagate(sig core.Signal) {
	g.rocket.Propagate(sig)
}

// RocketAdapter возвращает адаптер, соответствующий core.GhostLike
func RocketAdapter(r *GhostRocket) core.GhostLike {
	return &ghostAdapter{rocket: r}
}
