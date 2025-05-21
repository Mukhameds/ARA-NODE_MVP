# Module: signal\_engine.go

---

## ✅ Purpose

The `SignalEngine` is the entry point for all signals in ARA-NODE. It processes incoming `Signal` structs, stores them as `QBit`s in memory, and optionally triggers reactions or phantom logic based on phase.

---

## ⚙️ Functions

### `NewSignalEngine(mem *MemoryEngine)`

* Initializes a signal processor with reference to main memory.

### `ProcessSignal(sig Signal) Reaction`

* Logs the received signal
* Converts it to `QBit`
* Stores it in `MemoryEngine`
* Checks for phase-based phantom trigger (if `Phase > 0.8`)
* Returns corresponding `Reaction`

---

## 📥 Input

```go
Signal{
  ID:        string,
  Content:   string,
  Tags:      []string,
  Timestamp: time.Time,
  Phase:     float64,
  Weight:    float64,
  Origin:    string,
  Type:      string, // user, instinct, background, etc.
}
```

---

## 🧠 Internal Logic

```text
Signal received → wrapped as QBit → stored in Memory →
if Phase > 0.8:
    Reaction{phantom-triggered}
else:
    Reaction{acknowledged}
```

---

## 🧪 Reaction Output

```go
Reaction{
  TriggeredBy: sig.ID,
  Response:    string,
  Tags:        []string,
  Confidence:  float64,
}
```

* Tags include: `phantom` or `ack`
* Confidence: 0.8–0.95

---

## 🔄 Signal Flow

```text
Signal → SignalEngine.ProcessSignal → QBit → Memory →
       ↪ Reaction → WillEngine, EmotionEngine, Suggestor
```

---

## 💬 Log Output

```text
[SignalEngine] Received: {content}
[MemoryEngine] Stored QBit: {id}
```

---

## 📈 Planned Improvements

* Add phase tolerance-based classification
* Rate-limit duplicate signals
* Confidence scoring model for signal quality
* Signal type-specific routing

---

## 📂 Dependencies

* `core/memory_engine.go`
* `core/types.go`
* Used by: `main.go`, `attention_engine.go`, `phantom.go`

---

## 🧪 Related Tests

| File         | Description                       |
| ------------ | --------------------------------- |
| `test_10.md` | Signal stored and acknowledged    |
| `test_11.md` | Phantom-trigger reaction observed |
