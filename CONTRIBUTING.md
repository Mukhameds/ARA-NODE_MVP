# 🤝 CONTRIBUTING to ARA-NODE

ARA-NODE is a deterministic cognitive system. All contributions must respect its reactive paradigm: **Signal → Block → Reaction → Memory → Phantom → Intention**.

---

## 📜 Contribution Guidelines

### ✅ You Can Contribute:

* 🔧 New modules (memory logic, reaction types, phantom algorithms)
* 🧠 Improvements to signal flow or memory representation
* 🛠️ CLI utilities or visual diagnostic tools
* 📄 Documentation in `/docs/modules/`
* 🧪 Tests and reproducible signal loops

### ⚠️ Do Not Submit:

* ❌ LLM integration or statistical prediction code
* ❌ Neural network weights or black-box APIs
* ❌ Irreversible side-effects in signal handlers
* ❌ Global state hacks — all logic must be localized to modules

---

## 🗂️ Code Standards

* Language: Go 1.20+
* Format: `gofmt` + consistent indentation
* Architecture: modules must obey **Signal → QBit → Reaction** flow
* Logging: use standard `fmt.Println()` or `[Engine]` style prefixed logs
* Tests: prefer reproducible CLI interactions (see `test_10.md`, `test_11.md`)

---

## 🔄 Branching

* `main`: stable release
* `dev`: active development
* `feature/*`: use for large isolated modules

---

## 🧪 Submitting Tests

Each new module should be accompanied by a reproducible `.md` test case in `docs/tests/`:

```markdown
## Test Case: Reflex Reaction on Instinct Error
Signal: "unknown"
Expected:
- QBit stored
- Instinct triggers "instinct_error"
- ReflexEngine fires: "❗ error signal"
```

---

## ✉️ How to Contribute

1. Fork the repository
2. Create a new branch: `feature/my_module`
3. Implement changes
4. Add tests and documentation
5. Create a Pull Request into `dev`

We review all contributions for architectural integrity, memory safety, and reactive compliance.

---

## 👤 Maintainer

**Satybaev Mukhamed Kamilovich**
🔗 [GitHub: Mukhameds](https://github.com/Mukhameds)

---

Thank you for contributing to a project focused on verifiable, interpretable, and deterministic artificial cognition.
