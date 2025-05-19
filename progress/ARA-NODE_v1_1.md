# 📍 ARA-NODE — Дорожная карта улучшений (v1.1+)

## 🛠 Цель

Создать полноценного когнитивного агента ARA-NODE, способного:

* интерпретировать сенсорные сигналы (буквы, слова, эмоции);
* формировать память и фантомы без внешнего контроля;
* самостоятельно формулировать гипотезы, цели, реакции;
* развивать структуру сознания через загрузку знаний и опыта.

---

## 🔧 Блоки доработки с детализацией

### 1. 🌀 FanthomEngine (phantom.go)

**Цель:** фантом — не просто реакция, а активный элемент памяти.

* [ ] ✅ Генерация нового QBit:

```go
newQ := fe.Memory.CreateQBit("[phantom] " + summary)
newQ.Tags = []string{"phantom"}
newQ.Type = "phantom"
```

* [ ] 📎 Привязка к исходным QBits через отдельную `PhantomLog`
* [ ] 💬 Лог: ID фантома, ID источников, фаза, вес

---

### 2. 🔁 GhostField (ghost\_engine.go)

**Цель:** блоки реагируют как нейроны с адаптацией и памятью.

* [ ] Множественные теги: `MatchTags []string`
* [ ] Реактивная память:

```go
LastTriggered time.Time
ReactionCount int
```

* [ ] Подавление фантома, если `sig.Phase < MinPhase`

---

### 3. 💡 Suggestor (suggestor.go)

**Цель:** делать гипотезы умнее.

* [ ] Выбор QBits:

```go
if q.Phase > 0.7 && q.Weight > 0.5 && contains(q.Tags, "user")
```

* [ ] Кэш последней идеи → избегать повторов
* [ ] Автоматическое сохранение гипотез как `QBit` типа `phantom`

---

### 4. 🔥 WillEngine (will\_engine.go)

**Цель:** постоянное сравнение целей и памяти

* [ ] При несоответствии:

```go
mem.AdjustWeight(q.ID, -0.2)
if q.Weight < 0.1 { q.Archived = true }
```

* [ ] Расширить `StandardLibrary`:

```go
mission_reflection: ["переосмысление", "ошибка", "анализ"]
```

---

### 5. 🧠 MemoryEngine

**Цель:** живущая память с самоочисткой и типами

* [ ] `Decay` → нелинейный:

```go
decayFactor := 0.3 * math.Log1p(age)
```

* [ ] Удаление QBits:

```go
if q.Archived && age > 3600 {
    delete(m.QBits, q.ID)
}
```

* [ ] Добавить `Type string` в `QBit{}`: `goal`, `phantom`, `fact`, `emotion`

---

### 6. ⏳ PredictionEngine

**Цель:** обучение повторяющимся паттернам

* [ ] Автообучение:

```go
Chains = append(Chains, []string{"q1", "q2", "q3"})
```

* [ ] Отладка:

```go
if predicted == userInput { fmt.Println("[Predict] 🎯 Confirmed") }
```

---

### 7. 👁 SignalDictionary

**Цель:** обучаемое сенсорное ядро

* [ ] Экспорт в `.json`:

```json
{"A": {"type": "letter", "tags": ["vowel"]}}
```

* [ ] Поддержка образов, emoji: `type:image`, `type:emoji`
* [ ] Коллизия:

```go
if _, ok := Variables[id]; ok { return nil }
```

---

### 8. 📊 Логика фантомных связей

**Цель:** граф мышления

* [ ] Хранилище `PhantomLog`: список `{phantomID, sources}`
* [ ] CLI-команда `phantoms`, `phantom-tree`
* [ ] Визуализация: `→ qbit_A + qbit_B → qbit_F`

---

### 9. 🌐 WebCLI/UI

**Цель:** наблюдение и управление мыслями

* [ ] WebTerminal (React)
* [ ] Граф QBits (Recharts + Graphlib)
* [ ] Кнопки: `📌 fix`, `⭐ upvote`, `🧠 focus`

---

### 10. 🧩 Структуризация знаний

* [ ] Расширение `QBit`: `Class`, `Layer`

```go
Class: "knowledge" / "emotion"
Layer: "conscious" / "subconscious"
```

* [ ] `QBitLink`: структура `from → to`, `relation`

---

### 11. 🔄 AttentionEngine

* [ ] Самовыбор QBit с max(weight \* phase)
* [ ] Вызов Suggestor при затишье
* [ ] Переключение фокуса при бездействии 30с

---

### 12. 📥 PerceptionStream

* [ ] Делить любой ввод на токены
* [ ] Хранить `EventChain`
* [ ] Ввод: "я иду быстро" → токены + связь

---

### 13. 🧠 GeneralizationEngine

* [ ] Сравнение цепей → создание обобщающего фантома
* [ ] Merge QBits с одинаковыми тегами/похожим содержанием

---

### 14. 🔧 CLI

* [ ] `mem-stats`, `phantoms`, `phantom-tree`
* [ ] `load_topic <file>`, `load_knowledge <file>`
* [ ] `focus <id>`, `highlight`, `archive <id>`

---

### 15. 🚀 Bootstrap Interview

* [ ] Модульные `BootstrapBlock{}`
* [ ] `Prompt()` → `Run(input)` сохраняет в память
* [ ] Опрос: цель, интерес, профиль, тип помощи

---

### 16. 📚 Импорт знаний

* [ ] `load_knowledge book.json`

```json
[
  {"content": "Сила = масса * ускорение", "tags": ["formula", "physics"]}
]
```

* [ ] Преобразование в QBits, добавление тегов, веса, фазы
* [ ] Определение источника: `learned_from: "book_physics"`

---

## 🧠 Версия: `ARA-NODE v1.1 — Полноценный когнитивный агент`

## 📌 Приоритет:

1. `fanthom.go`: фантом = память
2. `will_engine.go`: сравнение с миссией
3. `ghost_engine.go`: реактивные блоки
4. `bootstrap.go`: начало личности
5. `load_knowledge.go`: обучение книгами
6. `phantom_tree`: дерево мыслей
