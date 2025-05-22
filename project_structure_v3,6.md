# 📁 ARA-NODE v3.6 — System Structure (Updated)

## 🧠 Overview

ARA-NODE v3.6 is an advanced reactive cognitive agent based on the paradigm:

```
Signal → Field → Reaction → Memory → Phantom → Ethalon
```

This version introduces a complete replacement of the legacy `GhostField` architecture with the modular `GhostRocket` engine, multi-field signal propagation (`Matrix` blocks), and a reactive system supporting autonomous cognition, signal merging, and self-evolving memory.

Fully modularized in Go, the system offers a CLI entry point, reactive memory architecture, and distributed knowledge flow.

---

## 📁 Project Directory Structure

### 🔧 CLI Entry Point (`cmd/`)

* **`main.go`**: Main orchestrator; handles user inputs, bootstraps agents, launches background thinking, controls propagation through `GhostRocket`.

### ⚙️ System Configuration (`config/`)

* **`manifest.go`**: Defines the immutable mission and core identity (`SelfKernel`).

### 🧠 Core Modules (`core/`)

| Module                     | Functionality                                                 |
| -------------------------- | ------------------------------------------------------------- |
| `signal_engine.go`         | Converts input to signals; emits QBits into memory.           |
| `memory_engine.go`         | Manages QBit lifecycle: store, decay, retrieve.               |
| `signal_dictionary.go`     | Basic symbol units and signal-token mapping.                  |
| `attention_engine.go`      | Handles background thought and focus switching.               |
| `emotion_engine.go`        | Affects QBits via emotional response and backpropagation.     |
| `instincts.go`             | Generates instinctive signals and regulatory reactions.       |
| `reflex_engine.go`         | Immediate low-latency reaction to dangerous input.            |
| `will_engine.go`           | Filters based on core mission alignment and ethalons.         |
| `decay_analysis_engine.go` | Prunes obsolete or inactive memory traces.                    |
| `QBitEvolutionEngine.go`   | Promotes persistent QBits into higher roles (reflex/ethalon). |
| `prediction.go`            | Suggests future paths based on signal patterns.               |
| `standards.go`             | Contains ethalons and mission-aligned QBit verification.      |
| `types.go`                 | Core data types: `QBit`, `Signal`, `Reaction`.                |
| `helpers.go`               | Shared utilities: tag search, phase checks, etc.              |
| `shutdown_engine.go`       | Terminates the system safely under critical state.            |
| `resonance_matrix.go`      | Tracks QBit associative strengthening/weakening.              |
| `dual_processor.go`        | Dual hemisphere parallel processing engine.                   |

### 🔬 Internal Modules (`internal/`)

| Module                        | Functionality                                                  |
| ----------------------------- | -------------------------------------------------------------- |
| `phantom.go`                  | Creates, evolves, and filters phantom thoughts.                |
| `phantom_tree.go`             | Visualizes phantom structure and ancestry.                     |
| `suggestor.go`                | Generates suggestions from QBit clusters and semantic cues.    |
| `prediction.go`               | Signal prediction and expectancy computation.                  |
| `fact_loader.go`              | Loads structured facts from JSON into memory.                  |
| `bootstrap.go`                | Cognitive user profiling and initial memory seeds.             |
| `conflict_detector.go`        | Detects contradictions between signals.                        |
| `consciousness_capture.go`    | Detects moments of self-aware resonance.                       |
| `self_engine.go`              | Maintains stable identity and traces dominant patterns.        |
| `p2p_sync.go`                 | Signal-level P2P exchange of meta-fields (ethalons, phantoms). |
| `github_sync.go`              | Synchronizes memory with GitHub for distributed persistence.   |
| `human_node.go`               | Accepts external reinforcement and correction from humans.     |
| `knowledge_profile_loader.go` | Loads knowledge sets under a license or context.               |

### 🌐 Signal Field Architecture (`field/`)

| File         | Functionality                                                    |
| ------------ | ---------------------------------------------------------------- |
| `matrix.go`  | Defines field structure and propagation rules.                   |
| `rocket.go`  | Central hub for signal routing across matrices.                  |
| `adapter.go` | Connects `GhostRocket` to core system via `GhostLike` interface. |

### 🧬 Data Storage (`data/`)

* **`memory.msgpack`**: Serialized memory file (QBits, signals).
* **`*.json`**: Static bootstrap or knowledge datasets.

---

## ⟳ Signal Flow in v3.6

```
User Input or Instinct
   ↓
Signal → SignalEngine → Memory (QBit)
   ↓
GhostRocket → Fields (math/emotion/phantom/...)
   ↓
Field Rule Match → Reaction Modules Triggered
   ↓
PhantomEngine / Suggestor / Emotion / Will
   ↓
SelfEngine + ResonanceMatrix + ConflictDetector
   ↓
AttentionEngine + PredictionEngine
   ↓
Memory Update or Phantom Promotion
```

---

## 🔒 Safety & Ethics Layer

* Instincts prevent harmful decisions (e.g., self-destruction or user harm).
* EmotionEngine adjusts phase/mass for dangerous patterns.
* Standards block propagation of signals that don't align with core mission.
* ShutdownEngine exits system if mass drops critically.

---

## 📈 Transition from v3.5 to v3.6

| Feature            | v3.5                      | v3.6+GhostRocket                     |
| ------------------ | ------------------------- | ------------------------------------ |
| Signal Propagation | `GhostField` (monolithic) | `GhostRocket` + `Matrix` (modular)   |
| Thinking Engine    | `SignalEngine` only       | `DualProcessor` + `AttentionEngine`  |
| Phantom Thinking   | Basic                     | Merged with Suggestor + FieldTrigger |
| Ethalon Filtering  | Static                    | Fully dynamic and recursive          |
| Memory Growth      | Manual                    | Reflex & Phase-triggered Evolution   |
| Self-Awareness     | N/A                       | `SelfEngine` + `ConsciousCapture`    |
| Network Sharing    | GitHub only               | +P2P MetaField Sync (`p2p_sync.go`)  |

---

## 📖 Documentation & Contact

* **Docs Index**: [ARA-NODE Documentation Index](./ARA-NODE_Documentation_Index.md)
* **Author**: Mukhamed Kamilovich Satybaev
* 🌐 [ARU-AGI Project](https://mukhameds.github.io/ARU-AGI-Project/)
* 🕞 [Twitter/X](https://x.com/redkms2025)
* 🔗 [LinkedIn](https://www.linkedin.com/in/muhamed-satybaev-38b864362)
* 📁 [GitHub: Mukhameds](https://github.com/Mukhameds)
