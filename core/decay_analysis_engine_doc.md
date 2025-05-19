# Module: decay\_analysis\_engine.go

---

## ✅ Purpose

The `DecayAnalysisEngine` is responsible for cleaning up the memory in ARA-NODE by removing QBits that have fully decayed. It logs these decay events for future introspection or analysis.

---

## ⚙️ Functions

### `NewDecayAnalysisEngine(mem *MemoryEngine)`

* Creates a new decay engine instance with a memory reference and empty event log.

### `RunDecayCheck()`

* Iterates over all QBits
* Removes any QBit with:

  * `Archived == true`
  * `Weight < 0.05`
* Logs the deletion event in `DecayEvent` log
* Console output: `☠️ Removed: {id}`

### `PrintDecayLog()`

* Outputs a formatted list of decay events
* If none occurred, prints a notice

---

## 📦 Struct: `DecayEvent`

```go
type DecayEvent struct {
  ID        string
  Reason    string
  Timestamp time.Time
}
```

---

## 🧠 Decay Policy

```text
If QBit.Archived && QBit.Weight < 0.05:
  → delete from memory
  → log decay event
```

---

## 📈 Planned Improvements

* Add retention policy: delay deletion N seconds after archive
* Group decay stats by tag/type
* Track decay trends over time for optimization

---

## 📂 Dependencies

* `MemoryEngine`
* `QBit` structure and `Weight`/`Archived` fields

---

## 🧪 Related Tests

| File         | Description                                        |
| ------------ | -------------------------------------------------- |
| `test_11.md` | Auto-decay logic, decay log correctness validation |
