# Module: will\_engine.go

---

## ✅ Purpose

The `WillEngine` manages the autonomous desire-driven behavior of ARA-NODE. It scans memory for signals aligned with the system's standards and triggers internal activity based on predefined urgency and phase criteria.

---

## ⚙️ Core Structures

### `Intent`

```go
type Intent struct {
  Tag     string
  Phase   float64
  Urgency float64
}
```

* Encapsulates the agent’s active cognitive focus.

### `WillEngine`

```go
type WillEngine struct {
  Memory    *MemoryEngine
  Delay     time.Duration
  Active    bool
  lastTried map[string]time.Time
}
```

* Drives the DesireLoop and manages retry/backoff mechanics.

---

## 🔧 Functions

### `NewWillEngine(mem)`

* Initializes the engine with memory reference, delay timer, and retry tracking map.

### `DesireLoop()`

* Executes in background via goroutine.
* Every `Delay` interval:

  * Retrieves all QBits tagged `user`
  * Skips archived or weak phase signals
  * If `isAlignedWithStandards(content)`:

    * Triggers signal → Phase/Weight-preserved
    * Logs match
    * Resets delay
  * Else:

    * Reduces QBit weight
    * Archives if weight drops below 0.1
    * Doubles delay (up to 120s max)

### `isAlignedWithStandards(content)`

* Checks signal content for presence of standard keywords
* Uses `StandardLibrary`

---

## 🧠 Example Flow

```text
QBit{Tag: user, Phase: 0.9} → DesireLoop → standards match
→ Emit signal: [WILL] Help people...
→ If no match → reduce weight → possibly archive
```

---

## 💬 Log Output

```text
[WillEngine] ⚡ Intent triggered: qbit_...
[WillEngine] ❌ Rejected: qbit_...
[WillEngine] 🗃 Archived: qbit_...
```

---

## 📈 Planned Improvements

* Intent prioritization by urgency
* Multi-phase interaction cycles
* Support for goal queues and branching

---

## 📂 Dependencies

* `MemoryEngine`, `Signal`, `StandardLibrary`
* Used in: `main.go`, `phantom`, `background`

---

## 🧪 Related Tests

| File         | Description                             |
| ------------ | --------------------------------------- |
| `test_11.md` | Intent loop runs, matches, and archives |
