# Module: reflex\_engine.go

---

## ✅ Purpose

The `ReflexEngine` provides instant, tag-triggered reactions in ARA-NODE. It is the fastest layer of the system’s reactive hierarchy, responding immediately to critical or pre-defined signals.

---

## ⚙️ Structure

### `ReflexRule`

```go
type ReflexRule struct {
  MatchTag string
  Action   func(sig Signal)
}
```

* Associates a specific signal tag with a function that will be executed.

### `ReflexEngine`

```go
type ReflexEngine struct {
  Rules []ReflexRule
}
```

* Stores all registered reflex rules.

---

## 🔧 Functions

### `NewReflexEngine()`

* Returns a new empty reflex engine

### `AddRule(tag, action)`

* Adds a new reflex rule

### `React(sig)`

* Checks if `sig.Tags` contains any `MatchTag`
* Executes corresponding `Action`
* Prints debug log on match

### `DefaultReflexSet(engine)`

* Adds pre-configured reflexes:

  * `instinct_error`: triggers system error alert
  * `danger`: triggers safety protocol

---

## 💬 Log Output Example

```text
[Reflex] ⚡ Instant reaction to: overload
[Reflex] ❗ System error reflex triggered.
```

---

## 📈 Planned Improvements

* Add reflex suppression window to avoid overload
* Reflex confidence/priority levels
* Reflex chaining: cascade one into another

---

## 📂 Dependencies

* `Signal` type, `containsTag` utility
* Used in: `main.go`, `instinct`, `phantom`, `background`

---

## 🧪 Related Tests

| File         | Description                                             |
| ------------ | ------------------------------------------------------- |
| `test_11.md` | Reflex triggered by `instinct_error` or `danger` signal |
