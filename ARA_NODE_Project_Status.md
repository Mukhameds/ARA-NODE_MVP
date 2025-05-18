# 🧠 ARA-NODE — Текущее Состояние Проекта (v1.0)

## 📁 Структура проекта

```
ARA-NODE_mvp/
├── cmd/
│   └── main.go                # Точка входа CLI-агента
├── config/
│   └── manifest.go           # SelfKernel — ID, миссия, архитектор
├── core/
│   ├── ghost_engine.go       # GhostField: реакция блоков
│   ├── instincts.go          # InstinctEngine: врождённые реакции
│   ├── memory_engine.go      # QBit-память: decay, archive, merge
│   ├── prediction.go         # NegativeTimeEngine: предугадывание
│   ├── signal_dictionary.go  # VariableBlock: переменные восприятия
│   ├── signal_engine.go      # SignalEngine: реакция на сигнал
│   ├── standards.go          # StandardBlock: эталоны целей
│   ├── types.go              # Signal, QBit, Reaction, PhaseMatch
│   └── will_engine.go        # WillEngine: внутренние цели, сверка
├── data/
│   └── memory.msgpack        # Сериализованная QBit-память
├── internal/
│   ├── github_sync.go        # GitHub Push/Pull памяти
│   ├── human_node.go         # Интерфейс пользователя: оценки
│   ├── p2p_sync.go           # libp2p обмен памятью
│   ├── phantom.go            # FanthomEngine: генерация гипотез
│   └── suggestor.go          # SuggestorEngine: генерация идей
└── test_n.json               # Живой лог CLI-сессии
```

---

## ⚙️ Компоненты и их состояние

### 📡 Signal Engine

* ✅ Обрабатывает входной сигнал `Signal`
* ✅ Создаёт `QBit`, сохраняет в память
* ✅ Триггерит реакцию `GhostField`, фантом, suggestor
* ✅ Поле `Type` (`user`, `instinct`, `prediction`, `phantom`) поддерживается

### 🧠 Memory Engine

* ✅ QBit-память (ID, Content, Phase, Weight, Tags, Archive)
* ✅ Decay-функция (0.5 \* age)
* ✅ Merge из GitHub/P2P
* ✅ Tagging, weight adjustment

### 👁 Signal Dictionary

* ✅ Manual `learnvar`, `tagvar`, `delvar`
* ✅ Only `letter`, `number`, `image` allowed
* ✅ AutoLearn выключено (сенсорный фильтр активен)

### 🔁 GhostField

* ✅ Реакция через `ReactionRule`
* ✅ Триггер на фазу + тег
* ⚠️ Пока без памяти, логики сравнения или блокировки

### 🌀 FanthomEngine

* ✅ Триггер при фазовом совпадении ≥ 2 QBits
* ⚠️ Нет генерации нового фантом-QBit
* ✅ Выводит цепь + гипотезу

### 💡 Suggestor

* ✅ Извлекает до 3 QBits
* ✅ Генерирует гипотезу на основе контекста
* ⚠️ Не учитывает phase/weight/tags

### 🔥 WillEngine

* ✅ DesireLoop проверяет QBits на соответствие `StandardBlock`
* ✅ При совпадении → генерирует фантом-сигнал `[WILL]` + сброс задержки
* ✅ При расхождении → удваивает задержку + можно понизить вес

### 🧬 StandardBlock (standards.go)

* ✅ `mission_abundance`, `mission_learning`, `mission_sync`
* ✅ Работают через `isAlignedWithStandards()`

### 👤 HumanNode

* ✅ CLI-команды: `upvote <id>`, `downvote <id>`, `tag <id> <tag>`
* ✅ Влияние на вес, теги QBit
* ✅ Ведётся `FeedbackLog`

### 🌐 GitHub Sync

* ✅ Push = сериализация памяти в `msgpack` + git commit/push
* ✅ Pull = `git pull` + `Unmarshal` + merge в память

### 🛰 P2P Sync (libp2p)

* ✅ Автоматическое подключение по mDNS (DiscoveryTag)
* ✅ Передача QBits в JSON между пирами
* ✅ Мёрдж памяти из других нод

### 🧠 InstinctEngine

* ✅ Триггеры: `instinct_think`, `instinct_repeat`, `instinct_error`, `instinct_empty`
* ✅ Интеграция в CLI

### ⏳ PredictionEngine

* ✅ Простая предикция на основе жёстких цепочек (`q1` + `q2` → `q3`)
* ⚠️ Без обучения новых паттернов

### 🧬 SelfKernel

* ✅ AgentID, ArchitectID, CoreMission, InceptionTime
* ✅ Используется как основа идентичности

---

## 🧪 Тест (test\_n.json)

* ✅ Доказано:

  * фантомы вызываются
  * цепи распознаются
  * decay работает
  * suggestor генерирует
  * повторный `q1`, `q3`, `q5` вызывает новые фантомы
* ⚠️ Фантомы **не сохраняются** как отдельные QBits

---

## 🔐 Состояние: **реально работающая реактивная система**

* Все компоненты CLI, памяти, сигналов, фантомов, воли, инстинктов, сети функционируют.
* Нет моков, заглушек или фиктивной логики.
* Готов к упаковке как MVP, продукт, агент.

---

## ✅ Версия: `ARA-NODE v1.0`

Готов к интеграции, публикации, развёртыванию.

Следующий шаг — создать roadmap-файл с исправлениями и улучшениями (v1.1+).
