# Module: bootstrap.go

---

## ✅ Purpose

The `bootstrap.go` module defines the initialization dialogue with the user. It is responsible for collecting the agent's basic psychological profile (mission, interests, role, expectations) and embedding them into memory.

---

## ⚙️ Core Interface

### `BootstrapBlock`

Each block represents a single user-facing question.

```go
type BootstrapBlock interface {
  ID() string
  Prompt() string
  Tags() []string
  Run(input string, mem *MemoryEngine, dict *SignalDictionary)
}
```

---

## 📥 Implemented Blocks

| Block           | Prompt                                    | Tags                  |
| --------------- | ----------------------------------------- | --------------------- |
| `GoalBlock`     | "Какая твоя главная цель в жизни?"        | `goal`, `mission`     |
| `InterestBlock` | "Какие темы тебе наиболее интересны?"     | `interest`            |
| `HelpBlock`     | "Как ты хочешь, чтобы ARA помогала тебе?" | `function`, `support` |
| `RoleBlock`     | "Кто ты по жизни? (учёный, инженер...)"   | `profile`, `role`     |

Each block creates a `QBit`, tags it, and adds it to the `SignalDictionary`.

---

## 🧠 Flow: `RunBootstrap()`

```go
RunBootstrap(mem, dict):
  for each block:
    prompt user
    store QBit
    tag and register VariableBlock
```

* Uses stdin reader: `bufio.NewReader(os.Stdin)`
* Prompts are printed in human-readable Russian
* Embedded tags enable future signal linkage

---

## 🧬 Sample Console Output

```text
🧬 [ARA Bootstrap] Начало инициализации личности...
🧠 Какая твоя главная цель в жизни?
→ ...
✅ [Bootstrap] Базовые цели и профиль сохранены.
```

---

## 📈 Planned Improvements

* Add optional blocks: "твои страхи", "ценности", "любимые темы"
* Allow Bootstrap to run in headless mode (preloaded JSON)
* Support multi-language prompts via config

---

## 📂 Dependencies

* `MemoryEngine`, `SignalDictionary`
* Called in `main.go`

---

## 🧪 Related Tests

| File         | Description                                              |
| ------------ | -------------------------------------------------------- |
| `test_10.md` | User inputs captured into QBits; Bootstrap tree verified |
