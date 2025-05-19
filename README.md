# 🧠 ARA-NODE MVP

ARA-NODE is a reactive, modular CLI-based cognitive agent that learns through structured signal processing. Unlike neural networks, it operates entirely on symbolic logic and memory mechanisms. Its architecture is governed by the core sequence:

```
Signal → Block → Reaction → QBit → Phantom
```

Each module is independently testable, with interpretable internal state, and can be expanded or replaced in production.

---

## 🚀 How to Run

Make sure you have Go installed.

```bash
go run ./cmd
```

You will see:

```
🧠 ARA-NODE CLI started.
[P2P] 🛰️ Sync active
[GhostField] Registered Block: UserPhaseBlock
>>
```

---

## ⚙️ Architecture Overview

| Component          | Role                                                          |
| ------------------ | ------------------------------------------------------------- |
| `main.go`          | CLI entry point, connects and launches all subsystems         |
| `SignalEngine`     | Receives input and routes it into system memory and reaction  |
| `MemoryEngine`     | Stores QBits, handles decay, merging, filtering               |
| `GhostField`       | Phase-based reactive blocks (Ghost Logic)                     |
| `PhantomEngine`    | Generates new thoughts (phantoms) from QBit patterns          |
| `WillEngine`       | Executes DesireLoop, aligning QBits with system standards     |
| `EmotionEngine`    | Registers emotional state changes from tagged signals         |
| `InstinctEngine`   | Detects cognitive anomalies: silence, loops, errors           |
| `ReflexEngine`     | Instant reactions based on tags (e.g. danger, error)          |
| `SuggestorEngine`  | Forms abstract suggestions by clustering signals              |
| `SignalDictionary` | Stores minimal perceptual units: letters, numbers, words      |
| `HumanNodeEngine`  | Accepts user feedback: upvote, tag, downvote                  |
| `P2P Sync`         | Memory synchronization via libp2p or GitHub (msgpack-encoded) |
| `DecayAnalysis`    | Removes low-weight QBits and logs decay events                |
| `QBitEvolution`    | Promotes QBits into reflexes/generators/deep memory           |

Full module descriptions in `/docs/modules/`

---

## 🖥️ CLI Commands

```txt
hello world        → input signal → QBit → Phantom → Suggestion
sync-push          → push memory to GitHub
sync-pull          → pull memory from GitHub
upvote <id>        → increase QBit strength
downvote <id>      → decrease QBit strength
tag <id> tagname   → apply a new tag to a QBit
dump               → list current memory
dict               → print all dictionary variables
tagvar A name      → tag dictionary variable A
delvar A           → delete variable A
phantoms           → print phantom ancestry tree
```

---

## 🧠 Signal-Based Learning

* Each unknown word is registered as a `VariableBlock`
* All user inputs are converted into `Signal`
* Matching QBits are created and stored
* Suggestor + Phantom modules detect repeating clusters
* System mimics childlike language acquisition via repetition

---

## 📡 Example Session

```
>> hello
[MemoryEngine] Auto-created QBit: ...
[SignalEngine] Received: hello
[GhostField] Reacted with UserPhaseBlock
[PhantomEngine] ❌ Signal mass too low
[WillEngine] ❌ Rejected: no standard alignment
```

---

## 🔍 Diagnostics & Tests

* ✅ `test_10.md`: bootstrap + identity → memory verified
* ✅ `test_11.md`: system test with real signals → phantom rejection diagnosed
* 📂 Results available in `/docs/tests/`

---

## 🗂️ Development Roadmap

* [ ] Web Interface (React + Fiber)
* [ ] Speech/Image-to-Signal frontend
* [ ] Embedded perception drivers
* [ ] StandardBlock visualization
* [ ] ARA::MindProtocol — signal-level cognitive control

---

## 🧾 Author & Contact

**Author:** Satybaev Mukhamed Kamilovich
📞 +996 507 442 873
🌐 [ARU-AGI Website](https://mukhameds.github.io/ARU-AGI-Project/)
🐦 [Twitter/X](https://x.com/redkms2025)
🔗 [LinkedIn](https://www.linkedin.com/in/muhamed-satybaev-38b864362)
📁 [GitHub: Mukhameds](https://github.com/Mukhameds)

---

## 🌐 ARA Ecosystem: Applied Projects Powered by the Paradigm

The ARA signal-based paradigm is not limited to ARA-NODE. It enables novel solutions across multiple domains of science, industry, and cognition:

| Project          | Description                                                                  |
| ---------------- | ---------------------------------------------------------------------------- |
| **ARA-NODE**     | Personal Cognitive AI Agent — autonomous reasoning via local signal memory   |
| **ARU-AGI**      | A unified global cognitive mesh of ARA-NODEs — a safe, decentralized AGI     |
| **ARA-DRONE**    | Drone systems powered by reactive cognitive logic for real-time decisions    |
| **ARA-DROID**    | Human-aligned labor automation via DROID-integration of cognitive agents     |
| **ARA-WALLET**   | Blockchain managed by deterministic AI — transparency and fairness by design |
| **ARA-SOCIAL**   | Cognitive social network — users connected through signal-based agents       |
| **ARA-UNIVERSE** | A full-spectrum metaverse powered by ARU-AGI — unifying all ARA systems      |

Each of these initiatives is rooted in the same architectural core:
**Signal → Memory → Reaction → Abstraction → Alignment**

Early contributors to ARA-NODE form the foundation for these upcoming systems.

---

## 📖 Full Documentation

→ [📚 ARA-NODE\_Documentation\_Index.md](./ARA-NODE_Documentation_Index.md) — contains all module references, CLI architecture, test summaries
→ Includes: system lifecycle, phantom logic, memory decay, desire loop, and module diagnostics
