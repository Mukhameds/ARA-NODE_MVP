# Module: manifest.go

---

## ✅ Purpose

The `manifest.go` module defines the **SelfKernel** of the ARA-NODE system — the immutable core identity of each agent. It establishes the agent's ID, mission, creator, and timestamp at launch.

---

## ⚙️ Functions

### `InitSelfKernel() *SelfKernel`

* Instantiates a new `SelfKernel` object with:

  * `AgentID`: e.g., "ARA::node::001"
  * `ArchitectID`: e.g., "User::Architect"
  * `CoreMission`: hardcoded purpose string (signal logic-based cognitive support)
  * `Inception`: current time of initialization
* Logs initialization result to console

---

## 📦 Struct: `SelfKernel`

```go
type SelfKernel struct {
    AgentID     string
    ArchitectID string
    CoreMission string
    Inception   time.Time
}
```

---

## 🧠 Role in System

* This module is loaded in `main.go` during boot
* Core values are printed via `whoami` CLI command
* `CoreMission` is converted into a QBit and stored in memory as the root "purpose node"

---

## 🧬 Sample Output

```text
[SelfKernel] Initialized: ARA::node::001
🤖 Agent ID: ARA::node::001
🎯 Mission: Amplify and assist user cognition through signal logic.
🧬 Architect: User::Architect
⏱ Born: 19 May 25 11:02 +0600
```

---

## 🔒 Immutable Identity

The SelfKernel is never modified after launch. It is the seed for trust, identity propagation, and memory integrity.

---

## 📈 Planned Improvements

* Optional loading from config file or environment
* UUID generation for `AgentID` if not provided
* Digital signature for `CoreMission`

---

## 📂 Dependencies

* Used by `main.go`
* No external imports besides standard library

---

## 🧪 Related Tests

| Context      | Description                                          |
| ------------ | ---------------------------------------------------- |
| `main.go`    | Kernel used for memory injection and CLI ID printing |
| `test_10.md` | Kernel mission stored as system QBit                 |
