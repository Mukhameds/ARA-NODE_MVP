# Module: types.go

---

## ✅ Purpose

Defines the core data structures for ARA-NODE: signals, QBits, reactions, and phantom tracking. These types form the foundational schema used across all modules.

---

## 📦 Core Types

### `Signal`

Represents any input or system-generated signal.

```go
type Signal struct {
  ID        string
  Content   string
  Tags      []string
  Timestamp time.Time
  Phase     float64
  Weight    float64
  Origin    string
  Type      string // user, instinct, background, prediction
}
```

* `Phase`: importance or excitation threshold
* `Weight`: persistence potential in memory

---

### `QBit`

QBit is the fundamental memory unit in ARA.

```go
type QBit struct {
  ID        string
  Content   string
  Tags      []string
  CreatedAt time.Time
  Weight    float64
  Phase     float64
  Type      string  // reflex, generator, standard, etc.
  Origin    string  // user, system, network
  Archived  bool
}
```

* Evolves into reflex/generator based on learning

---

### `Reaction`

Encapsulates a response to a signal.

```go
type Reaction struct {
  TriggeredBy string
  Response    string
  Tags        []string
  Confidence  float64
}
```

* Used by signal processors and user-facing layers

---

### `FanthomInterface`

Generic interface used by any phantom-capable module.

```go
type FanthomInterface interface {
  TriggerFromMatch(sig Signal)
}
```

---

### `PhantomLog`

Tracks the origin of a phantom node.

```go
type PhantomLog struct {
  PhantomID string
  SourceIDs []string
}
```

* Used to visualize signal ancestry chains

---

## 📂 Dependencies

* Used across: memory, phantom, signal, attention, bootstrap, CLI
* Self-contained: depends only on `time`

---

## 📈 Planned Improvements

* Add `Confidence` to QBits
* Add `DecayRate` per QBit
* Expand `Reaction` with `EmotionState`

---

## 🧪 Related Tests

| File         | Description                     |
| ------------ | ------------------------------- |
| `test_10.md` | Signal → QBit lifecycle tested  |
| `test_11.md` | Reaction confidence and tagging |
