# Module: load\_knowledge.go

---

## ✅ Purpose

This module allows ARA-NODE to import structured knowledge from a JSON file and transform it into QBits. It is the main knowledge ingestion point.

---

## 📦 Structures

### `KnowledgeEntry`

```go
type KnowledgeEntry struct {
  Content string   `json:"content"`
  Tags    []string `json:"tags"`
  Source  string   `json:"source,omitempty"`
}
```

* Describes a unit of knowledge to be converted into a QBit
* Optional `Source` gets turned into a tag

---

## 🔧 Function

### `LoadKnowledge(path, mem)`

* Opens JSON file
* Parses an array of `KnowledgeEntry`
* For each entry:

  * Creates a QBit using `MemoryEngine`
  * Assigns tags
  * Appends `learned_from:<source>` if provided
  * Stores to memory

---

## 💬 Output Example

```text
[Knowledge] ✅ Photosynthesis is essential [qbit_171611...]
[Knowledge] 📚 Loaded 58 entries from data/biology.json
```

---

## 📈 Planned Improvements

* Support CSV, YAML
* Batch tagging and source classification
* Automatic duplicate detection

---

## 📂 Dependencies

* Uses: `core.MemoryEngine`
* Standard Go: `os`, `json`, `fmt`
* Used by: `main.go`, `knowledge_profile_loader.go`

---

## 🧪 Related Tests

| File         | Description                             |
| ------------ | --------------------------------------- |
| `test_11.md` | Loads multiple JSON entries into memory |
