# Module: knowledge\_profile\_loader.go

---

## ✅ Purpose

This module adds license-based access control to knowledge loading in ARA-NODE. It wraps the general-purpose `LoadKnowledge()` function with a license check.

---

## 🔧 Functions

### `verifyLicense(license, path) → bool`

* Placeholder function
* Returns `true` for license keys `"dev"` or `"free"`
* Future hook for license validation logic

### `LoadKnowledgeLicensed(path, license, mem)`

* Calls `verifyLicense()`
* If valid:

  * Calls `LoadKnowledge(path, mem)`
* If invalid:

  * Returns error: `"❌ Invalid license key"`

---

## 💡 Use Case

```text
User wants to load proprietary data file
→ Calls: load_profile data.json dev
→ If license is invalid, aborts
→ Else, data is imported
```

---

## 📈 Planned Improvements

* Real license file validation (hash, token, or public key)
* License tier differentiation (dev / pro / enterprise)
* Expiration and license metadata parsing

---

## 📂 Dependencies

* `core.MemoryEngine`
* Delegates to `LoadKnowledge()` defined elsewhere

---

## 🧪 Related Tests

| File         | Description                         |
| ------------ | ----------------------------------- |
| `test_11.md` | Valid/invalid license cases checked |
