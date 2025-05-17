package internal

import (
	"fmt"
	"strings"
	"time"
	"ara-node/core"
)

type HumanFeedback struct {
	QBitID    string
	Action    string // upvote / downvote / tag
	Value     string // tag name (если Action == tag)
	Timestamp time.Time
}

type HumanNodeEngine struct {
	Memory       *core.MemoryEngine
	FeedbackLog  []HumanFeedback
}

func NewHumanNodeEngine(mem *core.MemoryEngine) *HumanNodeEngine {
	return &HumanNodeEngine{
		Memory: mem,
	}
}

func (h *HumanNodeEngine) HandleCommand(input string) bool {
	parts := strings.Fields(input)
	if len(parts) < 2 {
		return false
	}

	action := parts[0]
	id := parts[1]
	var tag string
	if action == "tag" && len(parts) > 2 {
		tag = parts[2]
	}

	switch action {
	case "upvote":
		h.Memory.AdjustWeight(id, +0.5)
	case "downvote":
		h.Memory.AdjustWeight(id, -0.5)
	case "tag":
		h.Memory.AddTag(id, tag)
	default:
		return false
	}

	h.FeedbackLog = append(h.FeedbackLog, HumanFeedback{
		QBitID:    id,
		Action:    action,
		Value:     tag,
		Timestamp: time.Now(),
	})

	fmt.Println("[HumanNode] ✅", action, id, tag)
	return true
}
