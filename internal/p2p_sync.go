package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"ara-node/core"

	libp2p "github.com/libp2p/go-libp2p"
	
	mdns "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

const ProtocolID = "/ara/sync/1.0.0"
const DiscoveryTag = "ara-mdns"

type PeerSync struct {
	Host host.Host
	Mem  *core.MemoryEngine
}

func StartP2P(mem *core.MemoryEngine) (*PeerSync, error) {
	
	h, err := libp2p.New()
	if err != nil {
		return nil, err
	}

	ps := &PeerSync{Host: h, Mem: mem}
	h.SetStreamHandler(ProtocolID, ps.onStream)

	service := mdns.NewMdnsService(h, DiscoveryTag, &peerHandler{ps})
	if err := service.Start(); err != nil {
		return nil, err
	}

	fmt.Println("[P2P] Started with ID:", h.ID().String())
	return ps, nil
}

type peerHandler struct {
	ps *PeerSync
}

func (ph *peerHandler) HandlePeerFound(pi peer.AddrInfo) {
	go func() {
	time.Sleep(5 * time.Second) // подождать до заполнения памяти
	err := ph.ps.syncWithPeer(pi)
	if err != nil {
		fmt.Println("[P2P Sync ❌]", err)
	} else {
		fmt.Println("[P2P Sync ✅] Sent QBits to", pi.ID.String())
	}
}()
}

func (ps *PeerSync) onStream(s network.Stream) {
	defer s.Close()

	var incoming map[string]core.QBit
	if err := json.NewDecoder(s).Decode(&incoming); err != nil {
		fmt.Println("[P2P ❌ decode]", err)
		return
	}
	ps.Mem.Merge(incoming)
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

	ps.Mem.Mu.Lock()
	defer ps.Mem.Mu.Unlock()
	return json.NewEncoder(s).Encode(ps.Mem.QBits)
}
