package internal

import (
	"context"
	"encoding/json"
	"fmt"

	"ara-node/core"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

const ProtocolID = "/ara/memory/1.0.0"

// PeerSync — P2P синхронизация памяти
type PeerSync struct {
	Host host.Host
	Mem  *core.MemoryEngine
}

func NewPeerSync(mem *core.MemoryEngine) (*PeerSync, error) {
	h, err := libp2p.New()
	if err != nil {
		return nil, err
	}
	ps := &PeerSync{Host: h, Mem: mem}
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
	remote := &core.MemoryEngine{QBits: incoming}
	ps.Mem.Merge(remote)
	fmt.Println("[P2P] ✅ Merged QBits:", len(incoming))
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

	return json.NewEncoder(s).Encode(ps.Mem.QBits)
}
