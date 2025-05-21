# Module: memory\_engine.go

---

## ✅ Purpose

The `MemoryEngine` is the core memory component of ARA-NODE. It stores, updates, filters, and decays QBits. This module is central to all cognitive processes and is accessed by nearly every subsystem.

---

## 📦 Core Structure

### `MemoryEngine`

```go
type MemoryEngine struct {
  QBits       map[string]QBit
  Mu          sync.Mutex
  PhantomTree []PhantomLog
}
```

---

## ⚙️ Functions

### Creation & Storage

* `NewMemoryEngine()` → initializes empty QBit map
* `CreateQBit(content)` → generates and stores a new QBit
* `StoreQBit(q)` → saves QBit in map
* `Broadcast(q)` → stores and logs QBit as a broadcast

### Retrieval

* `GetQBit(id)` → fetch QBit by ID
* `FindByTag(tag)` → filter QBits by tag
* `FindByPhase(target, tolerance)` → select by phase proximity
* `FindAll(filter)` → generic predicate-based search

### Update & Maintenance

* `AdjustWeight(id, delta)` → changes weight of a QBit
* `AddTag(id, tag)` → adds a tag to a QBit
* `UpdateQBit(qbit)` → full QBit replacement

### Decay & Cleanup

* `DecayQBits()` → applies age-based weight reduction
* `DeleteQBit(id)` → hard removal

### Output

* `ListQBits()` → prints all active (non-archived) QBits

### Merging

* `Merge(other)` → adds new QBits from another memory store

---

## 🧠 QBit Lifecycle

```text
CreateQBit → Store → Adjust/AddTag → Decay → Archive → Delete
```

---

## 🧬 Example Console Output

```text
[MemoryEngine] Auto-created QBit: qbit_1716112311223344
[Decay] qbit_xxx → age: 2.1s, decay: 1.05, new weight: 0.3
[MemoryEngine] 📡 Broadcast QBit: qbit_abc
```

---

## 📈 Planned Improvements

* Index by tags and phase ranges for faster lookup
* Soft tagging (decay-based tag suppression)
* Add QBit confidence score

---

## 📂 Dependencies

* QBit, PhantomLog from `types.go`
* Used by: signal engine, phantom, will, suggestor, p2p sync, bootstrap

---

## 🧪 Related Tests

| File         | Description                                    |
| ------------ | ---------------------------------------------- |
| `test_11.md` | Decay behavior, merge verification, tag update |
