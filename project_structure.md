# 📁 ARA-NODE System Structure

## 🧠 Overview

ARA-NODE is a reactive cognitive agent based on the paradigm: **Signal → Block → Reaction**. The system is built in modular Go code with a clean CLI entry point, autonomous memory, emotion, instinct, will, and phantom logic.

---

## 📁 Project Directory

### `C:\Documents\ARA-NODE_mvp\`

Root folder of the project.

### 🔧 `cmd/`

* `main.go` — CLI entry point, connects all modules, handles user input, emits signals, processes loops.

### ⚙️ `config/`

* `manifest.go` — SelfKernel: immutable identity of the agent (ID, mission, architect, time of birth).

### 🧠 `core/` — Core signal architecture

| File                       | Purpose                                                    |
| -------------------------- | ---------------------------------------------------------- |
| `signal_engine.go`         | Processes all incoming signals into QBits and reactions.   |
| `memory_engine.go`         | Stores, decays, evolves QBits; core memory.                |
| `signal_dictionary.go`     | VariableBlock dictionary: letters, symbols, words, images. |
| `attention_engine.go`      | Internal excitation generator (background thinking).       |
| `emotion_engine.go`        | Reacts to emotional triggers; holds current emotions.      |
| `instincts.go`             | Detects silence, loops, errors (instinct logic).           |
| `reflex_engine.go`         | Immediate reflexes triggered by tags.                      |
| `ghost_engine.go`          | GhostField and Block logic: reaction rules.                |
| `will_engine.go`           | DesireLoop; filters signals against mission standards.     |
| `decay_analysis_engine.go` | Detects and logs decayed QBits.                            |
| `QBitEvolutionEngine.go`   | Evolves QBits to reflex or generator type.                 |
| `prediction.go`            | Basic prediction engine using primitive chains.            |
| `standards.go`             | Standard mission blocks and keyword matching.              |
| `bootstrap.go`             | Initial user interview: purpose, interests, profile.       |
| `types.go`                 | Core types: Signal, QBit, Reaction, PhantomLog, etc.       |
| `helpers.go`               | Utility functions (phase diff, tag removal, etc).          |

### 🔬 `internal/` — Interface, I/O, Extensions

| File                          | Purpose                                             |
| ----------------------------- | --------------------------------------------------- |
| `phantom.go`                  | Phantom generation, merging, decay, deep memory.    |
| `phantom_tree.go`             | Console tree display of phantom source chains.      |
| `suggestor.go`                | SuggestorEngine: idea proposal from QBit chains.    |
| `p2p_sync.go`                 | libp2p: P2P discovery, stream sync, memory merge.   |
| `github_sync.go`              | GitHub memory sync via msgpack and `git push/pull`. |
| `human_node.go`               | User feedback: upvote, downvote, tag QBits.         |
| `load_knowledge.go`           | Load external knowledge (JSON) into memory.         |
| `knowledge_profile_loader.go` | License-checked loader wrapper.                     |

### 🧬 `data/`

* `memory.msgpack` — Serialized QBit memory used in GitHub sync.

---

## 🔄 Runtime Flow (Simplified)

```
User Input
   ↓
Signal → SignalEngine
   ↓
Memory + GhostField
   ↓
PhantomEngine + SuggestorEngine
   ↓
Emotion + Will + Reflex + Instinct
   ↓
Background Thinking (Attention)
   ↓
Decay + Evolution
```

---


