# Module: helpers.go

---

## ✅ Purpose

The `helpers.go` module defines small utility functions used across the ARA-NODE system. These assist with tag manipulation, phase comparison, and list checks, supporting core modules like memory, phantom, and ghost logic.

---

## ⚙️ Functions

### `RemoveTag(tags []string, target string) []string`

* Returns a copy of `tags` without `target`

```go
[]string{"a", "b", "c"}, "b" → ["a", "c"]
```

### `PhaseClose(p1, p2, tolerance float64) bool`

* Compares two phase values within a defined tolerance

```go
PhaseClose(0.82, 0.85, 0.05) → true
```

### `Contains(slice []string, item string) bool`

* Returns true if `item` is in `slice`

```go
Contains([]string{"user", "phantom"}, "phantom") → true
```

---

## 📈 Planned Improvements

* Add `Distinct` helper for deduplicating tag lists
* Add `TagPrefixMatch(slice, prefix string)`
* Optimize with generics if Go version >= 1.18

---

## 📂 Dependencies

* Pure helper module — no external or internal dependencies
* Used in: `phantom.go`, `emotion_engine.go`, `ghost_engine.go`, `will_engine.go`

---

## 🧪 Related Tests

| File         | Description                       |
| ------------ | --------------------------------- |
| `test_11.md` | Phase tolerance and tag filtering |
