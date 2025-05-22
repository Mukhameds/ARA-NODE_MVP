package internal

import (
	"fmt"
	"strings"
	"time"

	"ara-node/core"
)

// PhantomEngine — генератор фантомов
// PhantomEngine — генератор фантомов
type PhantomEngine struct {
	Memory     *core.MemoryEngine
	Instincts  *InstinctEngine
	Emotions   *EmotionEngine
	TimeEngine *TimeEngine // 🕒 биочасы
	Ghost      core.GhostLike
}

func NewPhantomEngine(mem *core.MemoryEngine, inst *InstinctEngine, emo *EmotionEngine, te *TimeEngine, ghost core.GhostLike) *PhantomEngine {
	return &PhantomEngine{
		Memory:     mem,
		Instincts:  inst,
		Emotions:   emo,
		TimeEngine: te,
		Ghost:      ghost,
	}
}





func (pe *PhantomEngine) TriggerFromMatch(sig core.Signal) {
	if sig.Weight < 0.5 {
		fmt.Println("[PhantomEngine] ❌ Signal weight too low, skip phantom generation")
		return
	}
	qbits := pe.Memory.FindByPhase(sig.Phase, 0.05)
	if len(qbits) < 2 {
		return
	}

	if uniqueSignalMass(qbits) < 1.5 {
		fmt.Println("[PhantomEngine] ❌ Unique signal mass too low — skip phantom")
		return
	}

	pe.GeneratePhantomChain(qbits)
}


func (pe *PhantomEngine) GeneratePhantomChain(chain []core.QBit) {
	var summary string
	var sources []string
	var signalMass float64
	seen := map[string]bool{}
	allPhantom := true
	phantomCount := 0

	for _, q := range chain {
		age := q.AgeFrame()
		if age == "emergent" || age == "legacy" {
			continue
		}

		if seen[q.ID] {
			fmt.Println("[PhantomEngine] ❌ Cycle detected, abort phantom generation")
			return
		}
		seen[q.ID] = true

		if strings.Contains(q.Content, "[phantom]") {
			phantomCount++
			if phantomCount > 1 {
				fmt.Println("[PhantomEngine] ❌ Too many phantom references, abort")
				return
			}
			continue
		}

		allPhantom = false

		inf := 1.0
		if q.Type == "standard" {
			inf += 1.5
		}
		if core.Contains(q.Tags, "instinct") {
			inf += 1.2
		}
		if core.Contains(q.Tags, "emotion") {
			inf += 1.1
		}

		boost := pe.Emotions.GetPhaseBoost(q.Tags)
		instinctBoost := pe.Instincts.GetInstinctBoost(q.Tags)
		timeFactor := pe.TimeEngine.TimeFactor()
		signalMass += (q.Phase + boost + instinctBoost) * q.Weight * inf * timeFactor

		summary += q.Content + " + "
		sources = append(sources, q.ID)
	}

	summary = strings.TrimSuffix(summary, " + ")


	// 🔍 Проверка конфликта перед генерацией фантома
	conflict := ConflictDetector{Memory: pe.Memory}
	phantomCandidate := core.QBit{
	Content: "[phantom] " + summary,
	Phase:   chain[0].Phase,
	Weight:  signalMass,
	}
if conflict.CheckConflict(phantomCandidate) {
	fmt.Println("[PhantomEngine] ❌ Phantom rejected due to internal contradiction.")
	return
}



	// Защита от повторной генерации фантома с тем же содержанием
if pe.Memory.ExistsQBit("[phantom] "+summary, chain[0].Phase, 0.01) {
	fmt.Println("[PhantomEngine] ⚠️ Phantom already exists — skip")
	return
}


	if strings.Count(summary, "[phantom]") > 1 {
		fmt.Println("[PhantomEngine] ❌ Phantom self-reference detected, abort")
		return
	}

	parts := strings.Split(summary, " + ")
	if len(parts) > 5 {
		parts = parts[len(parts)-5:]
		summary = strings.Join(parts, " + ")
	}

	var stdTags []string
	var stdWeightBonus float64
	if id, priority, score := core.MatchWithStandards(summary); id != "" {
		stdTags = []string{"standard_candidate", id}
		stdWeightBonus = priority * float64(score)
	}



	if allPhantom {
		fmt.Println("[PhantomEngine] ❌ All QBits are phantom, abort generation")
		return
	}
	if signalMass < 5.0 {
		fmt.Println("[PhantomEngine] ❌ Signal mass too low:", signalMass)
		return
	}

	for _, existing := range pe.Memory.FindByTag("phantom") {
		if existing.Content == "[phantom] "+summary {
			fmt.Println("[PhantomEngine] ❌ Duplicate phantom exists, skip")
			return
		}
	}

	if !pe.CheckInstinctEmotionAlignment(signalMass, summary) {
		fmt.Println("[PhantomEngine] ⚠️ Phantom temporarily rejected — tagged wait_for_merge")
		newQ := pe.Memory.CreateQBit("[phantom] " + summary)
		newQ.Tags = append([]string{"phantom", "wait_for_merge"}, stdTags...)
		newQ.Type = "phantom"
		newQ.Phase = chain[0].Phase
		newQ.Weight = (signalMass + stdWeightBonus) / float64(len(chain))
		pe.Memory.StoreQBit(*newQ)
		return
	}

	fmt.Println("[PhantomChain] 🧩 Related QBits:")
	for _, q := range chain {
		fmt.Printf("• %s | %.2f | %s\n", q.ID, q.Phase, q.Content)
	}
	fmt.Println("[PhantomChain] → Hypothesis: something meaningful links these signals.")

	newQ := pe.Memory.CreateQBit("[phantom] " + summary)
	newQ.Tags = append([]string{"phantom"}, stdTags...)
	newQ.Type = "phantom"
	newQ.Phase = chain[0].Phase
	newQ.Weight = (signalMass + stdWeightBonus) / float64(len(chain))
	pe.Memory.StoreQBit(*newQ)

	if pe.Ghost != nil {
	signal := core.SignalFromQBit(*newQ)
	pe.Ghost.Propagate(signal)
}


	go pe.DecayPhantom(newQ.ID, newQ.Weight)

	pe.Memory.PhantomTree = append(pe.Memory.PhantomTree, core.PhantomLog{
		PhantomID: newQ.ID,
		SourceIDs: sources,
	})

	fmt.Println("[PhantomEngine] 🔮 Phantom QBit:", newQ.ID)
	fmt.Println("[PhantomEngine] ↪ Sources:", strings.Join(sources, ","))
}





func (pe *PhantomEngine) CheckInstinctEmotionAlignment(signalMass float64, summary string) bool {
	instincts := pe.Instincts.Tick(time.Now(), summary)
	emotions := pe.Emotions.CurrentEmotions()

	allowedInstincts := []string{"instinct_think", "instinct_repeat"}
	blockedEmotions := []string{"fear", "anger", "disgust"}

	allow := false

	for _, inst := range instincts {
		for _, ai := range allowedInstincts {
			if inst.ID == ai {
				allow = true
				break
			}
		}
		if allow {
			break
		}
	}

	for _, emo := range emotions {
		for _, be := range blockedEmotions {
			if emo == be {
				allow = false
				break
			}
		}
		if !allow {
			break
		}
	}

	if signalMass < 5.0 {
		allow = false
	}

	return allow
}

func (pe *PhantomEngine) DecayPhantom(id string, weight float64) {
	if weight < 0.1 {
		pe.Memory.DeleteQBit(id)
		fmt.Println("[PhantomEngine] ⬇️ Phantom deleted due to low mass:", id)
	}
}


// ✅ Новая функция — вне тела предыдущей
func uniqueSignalMass(qbits []core.QBit) float64 {
	seen := make(map[string]bool)
	mass := 0.0
	for _, q := range qbits {
		if !seen[q.Content] {
			seen[q.Content] = true
			mass += q.Weight
		}
	}
	return mass
}


func (pe *PhantomEngine) TickUpdatePhantoms() {
	for _, q := range pe.Memory.FindByTag("wait_for_merge") {

if strings.Count(q.Content, "[phantom]") > 1 {
	fmt.Println("[PhantomEngine] ⚠️ Skip overloaded phantom:", q.ID)
	continue
}



		// 🔻 Перевод в глубокую память
		if q.Weight < 0.2 {
			q.Tags = append(q.Tags, "deep_memory")
			q.Tags = core.RemoveTag(q.Tags, "wait_for_merge")
			q.Weight = 0.05
			pe.Memory.UpdateQBit(q)
			fmt.Println("[PhantomEngine] 🧩 Moved to deep_memory:", q.ID)
			continue
		}


		// ✅ Проверка на эволюцию в стандартный блок
if core.Contains(q.Tags, "standard_candidate") && q.Weight > 2.0 {
	for _, tag := range q.Tags {
		if strings.HasPrefix(tag, "mission_") {
			stdID := tag
			std := core.GetStandardByID(stdID)
			if std != nil {
				q.Type = "standard"
				q.Tags = []string{"standard", stdID}
				q.Weight = 10.0
				pe.Memory.UpdateQBit(q)
				fmt.Println("[PhantomEngine] 🌐 Promoted to StandardBlock:", stdID)
				return // ⬅️ чтобы не сливался снова
			}
		}
	}
}



		// 🔁 Попытка слияния
		candidates := pe.Memory.FindByTag("wait_for_merge")
		var mergePool []string
		contentSet := make(map[string]bool)

		for _, c := range candidates {
			if c.ID == q.ID || !core.PhaseClose(q.Phase, c.Phase, 0.05) {
				continue
			}
			parts := strings.Split(c.Content, " + ")
			for _, p := range parts {
				contentSet[p] = true
			}
			mergePool = append(mergePool, c.ID)
		}

		// 🔘 Нет с кем слиться → затухание
		if len(mergePool) < 2 {
			q.Weight *= 0.95
			pe.Memory.UpdateQBit(q)
			continue
		}

		// 🧬 Объединение
		var merged []string
		for k := range contentSet {
			merged = append(merged, k)
		}
		summary := "[phantom] " + strings.Join(merged, " + ")
		if len(summary) > 128 {
			fmt.Println("[PhantomEngine] ⚠️ Merged phantom too long, skip")
			continue
		}

		newQ := pe.Memory.CreateQBit(summary)
		newQ.Type = "phantom"
		newQ.Tags = []string{"phantom"}
		newQ.Weight = q.Weight * 1.2 // частичное усиление
		newQ.Phase = q.Phase
		pe.Memory.StoreQBit(*newQ)

		for _, id := range mergePool {
			pe.Memory.DeleteQBit(id)
		}

		fmt.Println("[PhantomEngine] 🔄 Merged phantom created:", newQ.Content)
	}
}

func (pe *PhantomEngine) ReviveFromDeepMemory(sig core.Signal) {
	candidates := pe.Memory.FindByTag("deep_memory")
	for _, q := range candidates {

if strings.Contains(q.Content, "[phantom]") {
	continue // ⚠️ Не возбуждаем фантомы из глубокой памяти
}

		
		if core.PhaseClose(q.Phase, sig.Phase, 0.03) && strings.Contains(q.Content, sig.Content) {
			q.Weight += sig.Weight * 0.8
			if !core.Contains(q.Tags, "revived") {
				q.Tags = append(q.Tags, "revived")
			}
			pe.Memory.UpdateQBit(q)
			fmt.Println("[PhantomEngine] 🔁 Revived from deep_memory:", q.ID)
		}
	}
}
