# Module: main.go

---

## ✅ Purpose

`main.go` is the central CLI entry point of ARA-NODE. It initializes all core components, launches background processes, and handles real-time user input. It simulates a terminal-based thinking agent.

---

## ⚙️ Functions

### 🟢 Initialization Sequence:

* `InitSelfKernel()` — Sets fixed agent identity, mission, architect.
* `NewMemoryEngine()` — Creates QBit memory.
* `NewSignalDictionary()` — Loads perception dictionary.
* `StartP2P()` — Starts libp2p peer discovery and sync.
* `CreateQBit(CoreMission)` — Registers system mission.
* `NewSignalEngine()` — Signal reaction engine.
* `NewInstinctEngine()`, `NewEmotionEngine()` — Instinct and emotion logic.
* `NewPhantomEngine()` — Fanthom generator with alignment filters.
* `NewSuggestorEngine()` — Idea generation.
* `NewHumanNodeEngine()` — Accepts human feedback.
* `NewGhostField()` — Network of reactive blocks.
* `NewWillEngine()` — Desire loop, aligns thoughts to mission.
* `RunBootstrap()` — Asks user for goals, interests, role.
* `NewAttentionEngine()` — Launches background thinking.
* `NewReflexEngine()` — Instant tag-based reactions.
* `NewDecayAnalysisEngine()` — Logs memory degradation.
* `NewQBitEvolutionEngine()` — Evolves QBits into reflex/generator.

### 🔁 Main Loop:

```go
for {
  user_input → Signal → reflex + emotion + signalEngine
            → ghostField → phantom → suggestor
            → decay + evolve → prediction
            → console output
}
```

### 🧪 CLI Commands Supported:

| Command            | Description                        |
| ------------------ | ---------------------------------- |
| `dict`             | View variable blocks in dictionary |
| `whoami`           | Print agent identity and mission   |
| `tagvar A tag1`    | Add tag to variable A              |
| `delvar A`         | Delete variable A                  |
| `dump`             | Print all QBits                    |
| `phantoms`         | Print phantom tree                 |
| `sync-push`        | Push memory to GitHub              |
| `sync-pull`        | Pull memory from GitHub            |
| `load_knowledge p` | Load external knowledge JSON       |
| `load_profile p l` | Load licensed profile knowledge    |
| `decay-log`        | View memory decay history          |
| `evolve`           | Force evolution of QBits           |
| `upvote id`        | HumanNode vote boost               |
| `downvote id`      | HumanNode vote reduce              |
| `tag id label`     | HumanNode add tag                  |

---

## 🧠 Signal Flow

```
User input
   ↓
Signal (type=user)
   → ReflexEngine
   → EmotionEngine
   → SignalEngine.Process
   → GhostField.Propagate
   → PhantomEngine.TriggerFromMatch
   → Suggestor.SuggestFromQBits
   → Memory.DecayQBits + Evolution
   → AttentionEngine background activity
   → If matched → PredictionEngine
```

---

## 📈 Planned Improvements

* Command: `sync-peers` — show connected peers
* Modular input parser for CLI DSL
* Log-driven analytics of user interaction
* CLI history + replay

---

## 📂 Dependencies

* `core/*` (memory, signals, emotions, attention, phantom, ghost)
* `internal/*` (phantom tree, suggestor, sync, knowledge)
* `config/manifest.go`

---

## 🛠 Related Tests

| File         | Description                                      |
| ------------ | ------------------------------------------------ |
| `test_10.md` | Bootstrap QBit generation, sync activation       |
| `test_11.md` | Reactions to input, fanthom filtering, CLI logic |
