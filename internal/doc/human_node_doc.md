# Module: human\_node.go

---

## ✅ Purpose

The `HumanNodeEngine` module allows a human user to directly provide feedback on memory QBits via CLI commands. It supports upvotes, downvotes, and tagging, and records a full feedback log.

---

## 📦 Structures

### `HumanFeedback`

```go
type HumanFeedback struct {
  QBitID    string
  Action    string // upvote / downvote / tag
  Value     string // tag name (if Action == tag)
  Timestamp time.Time
}
```

* Captures one feedback event from a user.

### `HumanNodeEngine`

```go
type HumanNodeEngine struct {
  Memory      *core.MemoryEngine
  FeedbackLog []HumanFeedback
}
```

* Connected to memory
* Stores all past feedback events

---

## 🔧 Functions

### `NewHumanNodeEngine(mem)`

* Creates an instance bound to a MemoryEngine

### `HandleCommand(input string) bool`

* Parses CLI input and performs one of:

  * `upvote <id>`: increases QBit weight by `+0.5`
  * `downvote <id>`: decreases QBit weight by `-0.5`
  * `tag <id> <tag>`: adds tag to QBit
* Appends to `FeedbackLog`
* Logs successful command

---

## 🧠 Example Input

```text
upvote qbit_12345
→ [HumanNode] ✅ upvote qbit_12345

tag qbit_67890 important
→ [HumanNode] ✅ tag qbit_67890 important
```

---

## 📈 Planned Improvements

* Add feedback weighting per user
* Add feedback expiration or decay
* CLI auto-completion for known QBit IDs and tags

---

## 📂 Dependencies

* `MemoryEngine`
* Used in: `main.go`

---

## 🧪 Related Tests

| File         | Description                          |
| ------------ | ------------------------------------ |
| `test_11.md` | Human feedback modifies memory state |
