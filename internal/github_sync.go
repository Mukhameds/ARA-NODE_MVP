package internal

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"ara-node/core"
	"github.com/vmihailenco/msgpack/v5"
	"os"
)

const (
	gitRepoURL = "https://github.com/Mukhameds/ARA-NODE-MEMORY"
	localPath  = "./data/memory.msgpack"
	gitPath    = "data/memory.msgpack"
)

// PushMemory — сериализует и пушит память в GitHub
func PushMemory(mem *core.MemoryEngine) error {
	file, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

		mem.Mu.Lock()
	defer mem.Mu.Unlock()

	enc := msgpack.NewEncoder(file)
	err = enc.Encode(mem.QBits)
	if err != nil {
		return err
	}

	if err := gitCommitAndPush(); err != nil {
		return err
	}

	fmt.Println("[GitSync] ✅ Memory pushed to GitHub.")
	return nil
}

// PullMemory — вытягивает и загружает память
func PullMemory(mem *core.MemoryEngine) error {
	if err := gitPull(); err != nil {
		return err
	}

	data, err := os.ReadFile(localPath)
	if err != nil {
		return err
	}

	var remote map[string]core.QBit
	err = msgpack.Unmarshal(data, &remote)
	if err != nil {
		return err
	}

	mem.Merge(remote)
	fmt.Println("[GitSync] ✅ Memory pulled and merged.")
	return nil
}

// Вспомогательные git-команды
func gitCommitAndPush() error {
	t := time.Now().Format("2006-01-02_15-04-05")
	cmds := [][]string{
		{"add", gitPath},
		{"commit", "-m", "[sync] update " + t},
		{"push"},
	}
	return runGit(cmds)
}

func gitPull() error {
	return runGit([][]string{{"pull"}})
}

func runGit(cmds [][]string) error {
	for _, args := range cmds {
		cmd := exec.Command("git", args...)
		var out bytes.Buffer
		cmd.Stderr = &out
		if err := cmd.Run(); err != nil {
			fmt.Println("[GitError]", out.String())
			return err
		}
	}
	return nil
}
