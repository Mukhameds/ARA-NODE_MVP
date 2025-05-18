package core

type PredictionEngine struct {
	Chains [][]string // Примитивный шаблон: [q1, q2] → q3
}

func NewPredictionEngine() *PredictionEngine {
	return &PredictionEngine{
		Chains: [][]string{
			{"q1", "q2", "q3"},
			{"q5", "q1", "q3"},
		},
	}
}

func (pe *PredictionEngine) Predict(input string) (predicted string, ok bool) {
	for _, chain := range pe.Chains {
		if len(chain) < 3 {
			continue
		}
		if chain[0] == input || chain[1] == input {
			return chain[2], true
		}
	}
	return "", false
}
