package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"ara-node/core"
)

type UserProfile struct {
	Goal       string
	Interest   string
	Help       string
	Role       string
	Difficulty string
	Block      string
	CreatedAt  time.Time
}

// RunBootstrap инициализирует цели и мышление пользователя
func RunBootstrap(mem *core.MemoryEngine, dict *core.SignalDictionary) {
	// проверка на повтор
	if _, err := os.Stat("data/bootstrap_done.json"); err == nil {
		fmt.Println("🔁 Bootstrap already completed. Skipping.")
		return
	}

	fmt.Println("🔧 ARA Bootstrap Initialization")

	reader := bufio.NewReader(os.Stdin)

	ask := func(question string) string {
		fmt.Print(question + " > ")
		text, _ := reader.ReadString('\n')
		return strings.TrimSpace(text)
	}

	goal := ask("Какая твоя цель в жизни?")
	interest := ask("Какие темы тебе интересны?")
	help := ask("Как ARA может помочь тебе?")
	role := ask("Как ты себя определяешь (роль)?")
	difficulty := ask("Что даётся тебе труднее всего?")
	block := ask("Что ты чаще всего откладываешь или избегаешь?")

	// === Сохранение первичных QBits ===
	inputs := map[string]string{
		"goal":      goal,
		"interest":  interest,
		"help":      help,
		"role":      role,
		"difficulty": difficulty,
		"block":     block,
	}

	for typ, val := range inputs {
		q := mem.CreateQBit(val)
		q.Type = typ
		q.Tags = []string{"bootstrap", "user", typ}
		mem.StoreQBit(*q)
		dict.LearnFromInput(val)
	}

	// === Построение когнитивного профиля ===
	profile := analyzeCognitiveProfile(goal, interest, help, role)
	qp := mem.CreateQBit(profile)
	qp.Type = "cognitive_profile"
	qp.Tags = []string{"user", "profile", "bootstrap"}
	mem.StoreQBit(*qp)

	// === Вывод слабых сторон ===
	weak := analyzeWeakness(difficulty, block)
	if weak != "" {
		qw := mem.CreateQBit(weak)
		qw.Type = "weakness"
		qw.Tags = []string{"user", "analysis", "bootstrap"}
		mem.StoreQBit(*qw)
	}

	// === Сохраняем профиль в JSON ===
	userProfile := UserProfile{
		Goal:       goal,
		Interest:   interest,
		Help:       help,
		Role:       role,
		Difficulty: difficulty,
		Block:      block,
		CreatedAt:  time.Now(),
	}

	bytes, err := json.MarshalIndent(userProfile, "", "  ")
	if err == nil {
		os.WriteFile("data/user_profile.json", bytes, 0644)
		os.WriteFile("data/bootstrap_done.json", []byte(`true`), 0644)
	}

	fmt.Println("✅ Bootstrap завершён.")
}

// === Семантический разбор: когнитивный профиль
func analyzeCognitiveProfile(goal, interest, help, role string) string {
	var style, pattern string
	var tags []string

	if strings.Contains(goal, "создать") || strings.Contains(interest, "технологии") {
		style = "системный"
		tags = append(tags, "builder")
	} else if strings.Contains(help, "поддержка") || strings.Contains(role, "психолог") {
		style = "эмпатический"
		tags = append(tags, "support")
	} else {
		style = "аналитический"
		tags = append(tags, "thinker")
	}

	if strings.Contains(role, "ученик") || strings.Contains(goal, "учиться") {
		pattern = "восприятие → анализ → интеграция"
		tags = append(tags, "learning")
	} else {
		pattern = "цель → действие → результат"
		tags = append(tags, "active")
	}

	return fmt.Sprintf("UserLogicProfile: style=%s, pattern=%s, tags=%s",
		style, pattern, strings.Join(tags, ","))
}

// === Слабые стороны
func analyzeWeakness(diff, block string) string {
	var res []string

	if strings.Contains(diff, "решения") || strings.Contains(block, "ответственность") {
		res = append(res, "трудности с принятием решений")
	}
	if strings.Contains(diff, "абстракт") || strings.Contains(block, "теория") {
		res = append(res, "низкая устойчивость к абстракции")
	}
	if strings.Contains(diff, "страх") || strings.Contains(block, "ошибки") {
		res = append(res, "блок из-за страха ошибки")
	}

	if len(res) == 0 {
		return ""
	}
	return "Выявлены слабые стороны пользователя: " + strings.Join(res, "; ")
}
