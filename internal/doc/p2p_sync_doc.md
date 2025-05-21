# Module: p2p\_sync.go

---

## ✅ Purpose

The `p2p_sync.go` module enables decentralized memory synchronization using libp2p and mDNS. ARA-NODE instances can discover each other locally and exchange their QBit maps directly.

---

## 📦 Core Types

### `PeerSync`

```go
type PeerSync struct {
  Host host.Host
  Mem  *core.MemoryEngine
}
```

* Holds a libp2p host and pointer to memory.

### `peerHandler`

* Reacts to peer discovery events
* Calls `syncWithPeer()` after a short delay

---

## ⚙️ Functions

### `StartP2P(mem)`

* Creates libp2p host
* Registers stream handler for protocol `/ara/sync/1.0.0`
* Starts mDNS discovery with tag `ara-mdns`
* Returns `PeerSync`

### `syncWithPeer(peerInfo)`

* Connects to given peer
* Opens stream and sends `MemoryEngine.QBits` as JSON

### `onStream(stream)`

* Handles incoming libp2p stream
* Decodes JSON `map[string]QBit`
* Calls `MemoryEngine.Merge()`

---

## 🔄 Protocol Constants

```go
const ProtocolID = "/ara/sync/1.0.0"
const DiscoveryTag = "ara-mdns"
```

---

## 💬 Output Examples

```text
[P2P] Started with ID: 12D3KooW...
[P2P Sync ✅] Sent QBits to peer XYZ
[P2P] ✅ Merged QBits: 17
```

---

## 📈 Planned Improvements

* Replace JSON with MsgPack for performance
* Limit QBits by tag, phase, or freshness
* Add sync status reporting to CLI

---

## 📂 Dependencies

* `libp2p`, `mdns`, `core.MemoryEngine`
* Uses: `json`, `context`, `time`, `fmt`

---

## 🧪 Related Tests

| File         | Description                            |
| ------------ | -------------------------------------- |
| `test_11.md` | Peer discovery and memory merge tested |
