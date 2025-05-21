# Module: QBitEvolutionEngine.go

---

## ✅ Purpose

The `QBitEvolutionEngine` manages the development or degradation of QBits in ARA-NODE. It promotes QBits to new functional types based on accumulated weight and demotes them when they lose importance.

---

## ⚙️ Functions

### `NewQBitEvolutionEngine(mem)`

* Initializes the engine with a pointer to `MemoryEngine`

### `EvolveAll()`

Iterates through all QBits and applies evolution rules:

* If `Weight > 2.5` and `Type == ""` → promoted to `reflex`
* If `Weight > 3.0` and `Type == "reflex"` → promoted to `generator`
* If `Weight < 0.1` and `Type != "standard"` → archived

---

## 🔄 Evolution Logic

```text
(weight > 2.5 && type == "")      → reflex
(weight > 3.0 && type == "reflex") → generator
(weight < 0.1 && type != standard) → archived
```

---

## 💬 Log Output

```text
[Evolve] 🌱 Promoted to reflex: qbit_x
[Evolve] 🔁 Reflex → generator: qbit_y
[Evolve] 💤 Archived: qbit_z
```

---

## 📈 Planned Improvements

* Add score tracking to guide evolution
* Allow custom evolution chains (e.g., learner → teacher)
* Block automatic evolution for locked QBits

---

## 📂 Dependencies

* Uses `MemoryEngine`, `QBit`
* Called from: `main.go` and `phantom.go`

---

## 🧪 Related Tests

| File         | Description                                       |
| ------------ | ------------------------------------------------- |
| `test_11.md` | Reflex/generator transitions logged and validated |
