package internal

import (
	"sync"
	"time"
)

// TimeEngine — внутренний биочасовой модуль ARA
// Поддерживает цикл времени, хронологию, фоновую синхронизацию
type TimeEngine struct {
	startTime time.Time
	lastTick  time.Time
	cycle     int64
	mutex     sync.Mutex
}

func NewTimeEngine() *TimeEngine {
	return &TimeEngine{
		startTime: time.Now(),
		lastTick:  time.Now(),
		cycle:     0,
	}
}

// Tick — увеличивает внутренний цикл
func (te *TimeEngine) Tick() {
	te.mutex.Lock()
	defer te.mutex.Unlock()
	te.cycle++
	te.lastTick = time.Now()
}

// CurrentCycle — возвращает текущий номер цикла
func (te *TimeEngine) CurrentCycle() int64 {
	te.mutex.Lock()
	defer te.mutex.Unlock()
	return te.cycle
}

// SinceStart — сколько прошло времени с запуска
func (te *TimeEngine) SinceStart() time.Duration {
	return time.Since(te.startTime)
}

// SinceLastTick — сколько прошло времени с последнего цикла
func (te *TimeEngine) SinceLastTick() time.Duration {
	te.mutex.Lock()
	defer te.mutex.Unlock()
	return time.Since(te.lastTick)
}

// TimeFactor — вспомогательная функция: фазовый коэффициент по времени
// Можно использовать для модификации веса/приоритета/массы
func (te *TimeEngine) TimeFactor() float64 {
	elapsed := time.Since(te.startTime).Seconds()
	switch {
	case elapsed < 60:
		return 1.0
	case elapsed < 300:
		return 0.9
	case elapsed < 900:
		return 0.7
	default:
		return 0.5
	}
}
