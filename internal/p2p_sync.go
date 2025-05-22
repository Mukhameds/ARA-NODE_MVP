package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	

	"ara-node/core"
	"ara-node/field"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

const ProtocolID = "/ara/meta/field/1.0.0"

// PeerSync — P2P синхронизация смысловых полей
type PeerSync struct {
	Host       host.Host
	Mem        *core.MemoryEngine
	MetaFields map[string]*field.Matrix // ключевые поля: math, emotion, phantom...
}

func NewPeerSync(mem *core.MemoryEngine, meta map[string]*field.Matrix) (*PeerSync, error) {
	h, err := libp2p.New()
	if err != nil {
		return nil, err
	}
	ps := &PeerSync{
		Host:       h,
		Mem:        mem,
		MetaFields: meta,
	}
	h.SetStreamHandler(ProtocolID, ps.onStream)
	return ps, nil
}

func (ps *PeerSync) onStream(s network.Stream) {
	defer s.Close()

	var incoming map[string]core.QBit
	if err := json.NewDecoder(s).Decode(&incoming); err != nil {
		fmt.Println("[P2P ❌ decode]", err)
		return
	}

	fmt.Println("[P2P] 🔄 Incoming QBits:", len(incoming))

	for _, q := range incoming {
		if !ps.isSafeQBit(q) {
			continue
		}

		// Попытка замены если более весомый
		exist, ok := ps.Mem.QBits[q.ID]
		if !ok || q.Weight > exist.Weight {
			ps.Mem.StoreQBit(q)
		}

		sig := core.SignalFromQBit(q)

		// Отправляем в реакционное поле по смыслу
		for name, matrix := range ps.MetaFields {
			if hasTag(q.Tags, name) || hasTag(q.Tags, "shared") {
				go matrix.Propagate(sig)
			}
		}
	}
}

func (ps *PeerSync) syncWithPeer(pi peer.AddrInfo) error {
	ctx := context.Background()
	if err := ps.Host.Connect(ctx, pi); err != nil {
		return err
	}
	s, err := ps.Host.NewStream(ctx, pi.ID, ProtocolID)
	if err != nil {
		return err
	}
	defer s.Close()

	shared := map[string]core.QBit{}
	for id, q := range ps.Mem.QBits {
		if ps.isSafeQBit(q) {
			shared[id] = q
		}
	}
	return json.NewEncoder(s).Encode(shared)
}

// isSafeQBit — фильтр безопасности
func (ps *PeerSync) isSafeQBit(q core.QBit) bool {
	blocked := []string{"cli", "debug", "reflex", "archived"}
	for _, b := range blocked {
		if hasTag(q.Tags, b) {
			return false
		}
	}
	return hasTag(q.Tags, "shared") || hasTag(q.Tags, "ethalon") || hasTag(q.Tags, "confirmed")
}

// hasTag — частичное совпадение тега
func hasTag(tags []string, key string) bool {
	for _, t := range tags {
		if strings.Contains(t, key) {
			return true
		}
	}
	return false
}
