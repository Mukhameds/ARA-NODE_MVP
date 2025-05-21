# Module: suggestor.go

---

## ✅ Purpose

The `SuggestorEngine` proposes new ideas by analyzing recent QBits. It clusters similar signals and generates new phantom thoughts labeled as `suggestion`. This is one of ARA-NODE’s internal idea generation mechanisms.

---

## ⚙️ Structure

```go
type SuggestorEngine struct {
  Memory *core.MemoryEngine
}
```

---

## 🔧 Key Functions

### `NewSuggestorEngine(mem)`

* Initializes Suggestor with reference to memory

### `SuggestFromQBits()`

* Retrieves relevant recent QBits using `FindRecentRelevant`
* Groups QBits using `groupBySimilarity`
* Merges content into summarized idea with `mergeSummary`
* Logs result
* Creates a new phantom QBit with tag `suggestion`

### `FindRecentRelevant(n)`

* Returns last `n` QBits tagged with `user`, `instinct`, `emotion`, or `predict`

### `groupBySimilarity(qbits)`

* Forms clusters where first items are textually similar (≥2 common words)

### `mergeSummary(group)`

* Joins up to 5 unique, non-empty contents into one line

### `isSimilar(a, b)`

* Returns true if two strings have ≥2 shared words

### `GenerateSuggestion(ideas)`

* Legacy interface returning a summary prompt

---

## 💬 Output Example

```text
[Suggestor] 💡 cats absorb warmth + sleep is vital + soft touch triggers memory
```

---

## 📈 Planned Improvements

* Semantic embedding instead of word matching
* Memory feedback loop from suggestion usage
* Emotion-aware filtering

---

## 📂 Dependencies

* `MemoryEngine`, `QBit`, basic utils
* Used by: `main.go`, `phantom.go`, `background`

---

## 🧪 Related Tests

| File         | Description                           |
| ------------ | ------------------------------------- |
| `test_11.md` | Suggestion formed and phantom created |
