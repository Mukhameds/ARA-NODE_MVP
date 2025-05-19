# Module: github\_sync.go

---

## ✅ Purpose

This module enables Git-based synchronization of ARA-NODE’s memory. It serializes QBits using MsgPack and pushes/pulls from a remote GitHub repository.

---

## 🌐 Git Integration

### Constants

```go
const (
  gitRepoURL = "https://github.com/Mukhameds/ARA-NODE-MEMORY"
  localPath  = "./data/memory.msgpack"
  gitPath    = "data/memory.msgpack"
)
```

* Defines local and remote paths for memory sync.

---

## ⚙️ Functions

### `PushMemory(mem)`

* Serializes `MemoryEngine.QBits` to `memory.msgpack`
* Executes:

  * `git add`
  * `git commit -m "[sync] update TIMESTAMP"`
  * `git push`

### `PullMemory(mem)`

* Runs `git pull`
* Loads updated `memory.msgpack`
* Decodes into `map[string]QBit`
* Merges with existing memory

---

## 🧪 Internal Git Wrappers

* `gitCommitAndPush()`
* `gitPull()`
* `runGit(cmds [][]string)` — executes `git` commands via `exec.Command`

---

## 💬 Output Examples

```text
[GitSync] ✅ Memory pushed to GitHub.
[GitSync] ✅ Memory pulled and merged.
[GitError] error: failed to push some refs
```

---

## 📈 Planned Improvements

* Add authentication for private repos
* Conflict detection and resolution
* Progress reporting for large memory states
* Remote health check before push

---

## 📂 Dependencies

* `core.MemoryEngine`, `msgpack`
* Uses Go stdlib: `os`, `exec`, `time`, `bytes`

---

## 🧪 Related Tests

| File         | Description                                  |
| ------------ | -------------------------------------------- |
| `test_11.md` | Memory serialized, committed, pushed, merged |
