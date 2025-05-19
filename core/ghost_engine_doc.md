# Module: ghost\_engine.go

---

## ✅ Purpose

The `GhostEngine` module implements decentralized reactivity in ARA-NODE using signal-block logic. It defines `Block` nodes that react to signals based on phase and tag-matching rules. The `GhostField` holds a list of all blocks and propagates signals across them.

---

## ⚙️ Core Types

### `ReactionRule`

```go
type ReactionRule struct {
  MatchTags []string
  MinPhase  float64
  Action    func(sig Signal)
}
```

* Defines how a block reacts to specific tags and signal strength.

### `Block`

```go
type Block struct {
  ID            string
  Rules         []ReactionRule
  LastTriggered time.Time
  ReactionCount int
}
```

* A reactive node that holds multiple rules and tracks reaction history.

### `GhostField`

```go
type GhostField struct {
  Blocks []*Block
}
```

* A registry of all reactive blocks in the system.

---

## 🧠 Functions

### `Block.React(sig Signal)`

* For each rule:

  * Match by tag and `Phase >= MinPhase`
  * If matched, execute `Action(sig)`
  * Log trigger to console

### `NewGhostField()`

* Returns a new empty field

### `RegisterBlock(b *Block)`

* Adds a block to the field
* Logs registration

### `Propagate(sig Signal)`

* Sends signal to all registered blocks
* Calls `React(sig)` on each

---

## 📡 Example

```text
Signal: {Tags: [user], Phase: 0.85} →
→ Propagate()
→ Block "UserPhaseBlock" triggers action:
   [Block UserPhaseBlock] Reacting to signal: Hello
```

---

## 📈 Planned Improvements

* Block memory: last 10 signals per block
* Adaptive rules: learn new triggers based on history
* Phase-tuned feedback loops

---

## 📂 Dependencies

* `Signal` type from `types.go`
* Used by: `main.go`, `attention_engine.go`, `phantom.go`

---

## 🧪 Related Tests

| File         | Description                       |
| ------------ | --------------------------------- |
| `test_11.md` | Ghost block reacts to user signal |
