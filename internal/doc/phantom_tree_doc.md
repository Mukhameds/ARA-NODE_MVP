# Module: phantom\_tree.go

---

## ✅ Purpose

The `phantom_tree.go` module provides a diagnostic function to print the source structure of all generated phantoms in ARA-NODE. It reads from the `PhantomTree` and visualizes origin QBits.

---

## 🔧 Function

### `PrintPhantomTree(mem)`

* Takes `*MemoryEngine`
* Iterates over `mem.PhantomTree` entries
* For each phantom:

  * Prints its ID
  * Lists source QBits with their IDs and content
  * Marks `[not found]` if the QBit no longer exists

---

## 💬 Output Example

```text
🌱 [PhantomTree] Дерево фантомов:
🔮 qbit_17823...
   ↪ qbit_1001 | water is life
   ↪ qbit_1040 | plant absorbs light
```

If no phantoms exist:

```text
[PhantomTree] ⚠️ Нет фантомов в журнале.
```

---

## 📈 Planned Improvements

* Add tree depth metrics
* Export to DOT/GraphViz format
* Time-stamp phantom creation for sorting

---

## 📂 Dependencies

* Depends on `MemoryEngine`, `PhantomLog`, `QBit`
* Used in CLI command `phantoms`

---

## 🧪 Related Tests

| File         | Description                      |
| ------------ | -------------------------------- |
| `test_11.md` | Phantom tree printed after merge |
