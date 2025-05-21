# Module: signal\_dictionary.go

---

## ✅ Purpose

The `SignalDictionary` defines the perceptual vocabulary of ARA-NODE. It manages VariableBlocks — atomic units of recognition such as letters, words, images, and numeric symbols. Each block can link to memory and evolve over time.

---

## 🧩 Core Types

### `VariableBlock`

```go
type VariableBlock struct {
  ID     string
  Signal string
  Tags   []string
  Reacts []string
  QBit   *QBit
  Auto   bool
}
```

* Represents a minimal perceptual signal (e.g., "A", "42", "image1")
* Tags define category (`letter`, `number`, etc.)
* `Reacts` holds synonyms or token match patterns

### `SignalDictionary`

```go
type SignalDictionary struct {
  Variables map[string]*VariableBlock
  Memory    *MemoryEngine
}
```

---

## ⚙️ Functions

### `NewSignalDictionary(mem)`

* Initializes the dictionary and links to memory

### `FindMatch(token)`

* Finds a VariableBlock where `token` matches any `Reacts`

### `AutoLearn(token)`

* Creates a new VariableBlock with default tag `type:unknown`
* Generates a QBit for this token
* Sets `Auto = true`

### `Add(id, signal, tags, reacts)`

* Manually creates a VariableBlock with linked QBit

### `Delete(id)`

* Removes a VariableBlock

### `Tag(id, tag)`

* Adds tag to an existing block

### `All()`

* Returns all registered VariableBlocks

---

## 🧠 Example Flow

```text
"apple" → not in dictionary
→ AutoLearn("apple") → VariableBlock created + QBit stored
→ Next time → FindMatch("apple") → returns block
```

---

## 📈 Planned Improvements

* Frequency-based weighting
* Visual embedding (image hashes, OCR)
* Emotion binding for sensory signals
* Tag hierarchy (e.g., letter → character → symbol)

---

## 📂 Dependencies

* MemoryEngine, QBit
* Used in: `main.go`, `bootstrap.go`, `phantom.go`, `ghost`, etc.

---

## 🧪 Related Tests

| File         | Description                             |
| ------------ | --------------------------------------- |
| `test_10.md` | Bootstrap signals create VariableBlocks |
| `test_11.md` | AutoLearn and tag propagation observed  |
