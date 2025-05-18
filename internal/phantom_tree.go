package internal

import (
	"fmt"
	"ara-node/core"
)

// PrintPhantomTree — выводит дерево фантомов
func PrintPhantomTree(mem *core.MemoryEngine) {
	if len(mem.PhantomTree) == 0 {
		fmt.Println("[PhantomTree] ⚠️ Нет фантомов в журнале.")
		return
	}

	fmt.Println("🌱 [PhantomTree] Дерево фантомов:")
	for _, p := range mem.PhantomTree {
		fmt.Printf("🔮 %s\n", p.PhantomID)
		for _, src := range p.SourceIDs {
			if q, ok := mem.QBits[src]; ok {
				fmt.Printf("   ↪ %s | %s\n", src, q.Content)
			} else {
				fmt.Printf("   ↪ %s | [not found]\n", src)
			}
		}
	}
}
