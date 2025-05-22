// core/bootstrap_math_probability.go
package core

import (
	"fmt"
	"time"
)

// BootstrapMathProbability загружает сигнальные понятия вероятности и случайности
func BootstrapMathProbability(mem *MemoryEngine) {
	prob := []struct {
		ID      string
		Content string
		Tags    []string
	}{
		// Случайность и вероятность
		{"probability_def", "probability measures the likelihood of an event", []string{"math", "probability", "definition", "core", "bootstrap"}},
		{"prob_event", "an event is an outcome or set of outcomes", []string{"math", "probability", "event", "core", "bootstrap"}},
		{"prob_range", "probability is always between 0 and 1", []string{"math", "probability", "range", "core", "bootstrap"}},

		// Примеры и базовые правила
		{"prob_example", "the probability of a fair coin landing heads is 0.5", []string{"math", "probability", "example", "core", "bootstrap"}},
		{"prob_sum_rule", "P(A ∪ B) = P(A) + P(B) − P(A ∩ B)", []string{"math", "probability", "rule", "core", "bootstrap"}},
		{"prob_independent", "events A and B are independent if P(A ∩ B) = P(A)·P(B)", []string{"math", "probability", "independence", "core", "bootstrap"}},
		{"prob_conditional", "P(A|B) = P(A ∩ B) / P(B)", []string{"math", "probability", "conditional", "core", "bootstrap"}},

		// Ожидание и дисперсия
		{"prob_expectation", "expected value is the average outcome weighted by probability", []string{"math", "probability", "expectation", "core", "bootstrap"}},
		{"prob_variance", "variance measures spread of outcomes around the expected value", []string{"math", "probability", "variance", "core", "bootstrap"}},
		{"prob_stddev", "standard deviation is the square root of variance", []string{"math", "probability", "deviation", "core", "bootstrap"}},

		// Распределения
		{"prob_distribution", "a probability distribution assigns values to events", []string{"math", "probability", "distribution", "core", "bootstrap"}},
		{"prob_uniform", "uniform distribution assigns equal probability to all outcomes", []string{"math", "probability", "distribution", "uniform", "core", "bootstrap"}},
		{"prob_normal", "normal distribution is symmetric and bell-shaped", []string{"math", "probability", "distribution", "normal", "core", "bootstrap"}},
	}

	for _, p := range prob {
		q := QBit{
			ID:        p.ID,
			Content:   p.Content,
			Tags:      p.Tags,
			Phase:     0.89,
			Weight:    0.97,
			CreatedAt: time.Now(),
		}
		mem.StoreQBit(q)
	}

	fmt.Println("🎲 [Bootstrap] Probability and randomness concepts loaded.")
}
