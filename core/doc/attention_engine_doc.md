# Module: attention\_engine.go

---

## ✅ Purpose

The `AttentionEngine` drives autonomous cognitive activity in ARA-NODE. It generates periodic background signals by reactivating high-relevance QBits, mimicking subconscious or reflective thought.

---

## ⚙️ Functions

### `NewAttentionEngine(...)`

* Initializes the engine with references to:

  * `MemoryEngine`
  * `GhostField`
  * `PhantomEngine`
  * `SignalEngine`

### `Suppress(d time.Duration)`

* Temporarily blocks background activity (e.g., during active user input)
* Sets `SuppressedUntil = now + d`

### `StartBackgroundThinking()`

* Launches a `goroutine` that:

  * Every 5 seconds:

    * Skips if `SuppressedUntil` is active
    * Finds QBits with `Weight * Phase > 0.6`
    * Emits synthetic `Signal{Type: background, Origin: internal}`
    * Reactivates them via SignalEngine, GhostField, PhantomEngine

---

## 🧠 Internal Logic

```text
If not suppressed:
  Scan memory for relevant QBits
  → Form Signal
  → ProcessSignal
  → Propagate via GhostField
  → Trigger PhantomEngine
```

---

## 📥 Generated Signal

```go
Signal{
  ID:        "bg_...",
  Content:   QBit.Content,
  Tags:      QBit.Tags,
  Type:      "background",
  Origin:    "internal",
  Phase:     min(QBit.Phase + 0.05, 1.0),
  Weight:    QBit.Weight * 0.9,
  Timestamp: now,
}
```

---

## 📈 Planned Improvements

* Add context drift detector: when attention starts moving across themes
* Implement cooldown for specific QBits to prevent over-firing
* Track frequency of background activations for analysis

---

## 📂 Dependencies

* `MemoryEngine` (for QBit recall)
* `SignalEngine` (to reprocess)
* `GhostField` (for signal propagation)
* `FanthomInterface` (for phantom triggering)

---

## 🧪 Related Tests

| File         | Description                                                     |
| ------------ | --------------------------------------------------------------- |
| `test_11.md` | Background signals triggered from memory, signal flow validated |
