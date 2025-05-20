# Module: instincts.go

---

## ✅ Purpose

The `InstinctEngine` models basic reflexive responses in ARA-NODE. It monitors signal input patterns and generates instinctive triggers like silence, repetition, or error detection. These instincts influence Phantom and Emotion engines.

---

## ⚙️ Functions

### `NewInstinctEngine()`

* Initializes with:

  * current timestamp
  * history buffer (max 100 signals)

### `Tick(currentTime, signal) → []string`

* Evaluates current signal and system time
* Possible triggers:

  * `instinct_think`: silence timeout >10s
  * `instinct_repeat`: signal repeated
  * `instinct_error`: contains "error"
  * `instinct_empty`: input is empty or whitespace
* Appends signal to history

### `isRepeat(signal string)`

* Checks if signal is in recent history

### `addSignal(signal string)`

* Adds signal to history, trimming oldest if full

### `ClearHistory()`

* Clears history buffer manually

---

## 🧠 Instinct Logic

```text
If:
  • silence > 10s        → instinct_think
  • same signal again    → instinct_repeat
  • contains 'error'     → instinct_error
  • empty input          → instinct_empty
```

---

## 🔄 Internal State

* `LastInputTime`: tracks silence gaps
* `recentSignals`: buffer of last inputs
* `maxHistory`: 100 items default

---

## 📈 Planned Improvements

* Add confidence scores per instinct
* Support time-weighted history for forgetting
* Separate instinct types (cognitive, danger, survival)

---

## 📂 Dependencies

* Used by: `phantom.go`, `main.go`, `background_thinking`
* Relies on `strings`, `sync`, `time`

---

## 🧪 Related Tests

| File         | Description                         |
| ------------ | ----------------------------------- |
| `test_11.md` | Triggered instincts from input flow |
