# Module: emotion\_engine.go

---

## ✅ Purpose

The `EmotionEngine` in ARA-NODE manages the emotional state of the agent. It detects emotionally tagged signals, triggers associated responses, and maintains an internal list of current active emotions.

---

## ⚙️ Functions

### `NewEmotionEngine()`

* Initializes with default emotional state: `neutral`

### `AddTrigger(tag, phaseGT, action)`

* Adds a rule for triggering on signals with a given tag and phase

### `React(sig Signal)`

* Checks signal tags/phase against rules
* Executes `Action` on match
* Updates emotional state

### `CurrentEmotions()`

* Returns current emotions (thread-safe copy)

### `UpdateEmotion(emotion string)`

* Adds a new emotion to the active list (if not duplicate)

### `ClearEmotions()`

* Resets emotions to `neutral`

### `DefaultEmotionSet(engine)`

* Registers base emotions:

  * `joy` (😊)
  * `frustration` (😣)
  * `fear` (😨)
  * `anger` (😠)

---

## 📦 Struct: `EmotionTrigger`

```go
type EmotionTrigger struct {
  Tag     string
  PhaseGT float64
  Action  func(sig Signal)
}
```

---

## 🧠 Example Flow

```text
Signal with tag "joy" + Phase=0.75 →
→ React() matches rule → executes Action →
→ Updates internal state to include "joy"
```

---

## 💬 Log Output

```text
[Emotion] 😊 Joyful signal received.
[EmotionEngine] Updated emotions: [neutral joy]
```

---

## 📈 Planned Improvements

* Emotion decay over time
* Intensity scaling with repeated triggers
* Emotional conflict detection (e.g. joy + anger)

---

## 📂 Dependencies

* `Signal`, `Contains` util
* Used by: `main.go`, `phantom.go`, `attention_engine.go`

---

## 🧪 Related Tests

| File         | Description                                        |
| ------------ | -------------------------------------------------- |
| `test_11.md` | Emotion triggered by signal, state update verified |
