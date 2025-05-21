# Module: phantom.go

---

## ✅ Purpose

The `PhantomEngine` is the heart of emergent cognition in ARA-NODE. It generates higher-order QBits (phantoms) by detecting meaningful clusters of signals and evolving them through memory, instinct, emotion, and standard alignment.

---

## 📦 Structure

```go
type PhantomEngine struct {
  Memory    *MemoryEngine
  Instincts *InstinctEngine
  Emotions  *EmotionEngine
}
```

---

## 🔧 Key Functions

### `TriggerFromMatch(sig)`

* Entry point: called by SignalEngine, Attention, Reflex
* Filters signals:

  * `sig.Weight < 0.5` → skip
  * `len(qbits) < 2` or `signalMass < 1.5` → skip
* Calls `GeneratePhantomChain()` if valid

### `GeneratePhantomChain(qbits)`

* Merges related QBits into a `[phantom]` node
* Applies:

  * Cycle/duplicate checks
  * Content filtering
  * Signal mass accumulation
  * Emotion/instinct gating
  * StandardBlock promotion (if `standard_candidate`)

### `DecayPhantom(id, weight)`

* Deletes phantom if weight < 0.1

### `TickUpdatePhantoms()`

* Background process
* Handles phantoms tagged as `wait_for_merge`
* Logic:

  * Promote to deep memory if weak
  * Merge similar phantoms into one
  * Promote to `standard` if eligible

### `ReviveFromDeepMemory(sig)`

* Revives relevant non-phantom deep memory blocks that match signal phase/content

### `uniqueSignalMass(qbits)`

* Returns cumulative weight of distinct QBits (by content)

### `CheckInstinctEmotionAlignment(mass, content)`

* Allows generation only if instincts like `instinct_think` or `instinct_repeat` are active
* Rejects if emotional state includes `fear`, `anger`, or `disgust`

---

## 🧠 Phantom Lifecycle

```text
Signal → QBits → TriggerFromMatch → GeneratePhantomChain
→ wait_for_merge → Merge or Promote → deep_memory / standard
```

---

## 💬 Output Examples

```text
[PhantomEngine] ❌ Signal weight too low, skip phantom generation
[PhantomChain] → Hypothesis: something meaningful links these signals.
[PhantomEngine] 🔮 Phantom QBit: qbit_...
```

---

## 📈 Planned Improvements

* Self-reflection analysis
* Weighted recurrence tracking
* Emotion embedding in phantom tags

---

## 📂 Dependencies

* `MemoryEngine`, `Signal`, `InstinctEngine`, `EmotionEngine`, `Standards`, `QBit`, `PhantomLog`

---

## 🧪 Related Tests

| File         | Description                                |
| ------------ | ------------------------------------------ |
| `test_11.md` | Phantom created, rejected, merged, revived |
