package internal

import (
	
	"errors"
	
	
	"ara-node/core"
)

// Проверка лицензии (заглушка на будущее)
func verifyLicense(license string, path string) bool {
	return license == "dev" || license == "free"
}

// Загрузка знаний с лицензией
func LoadKnowledgeLicensed(path string, license string, mem *core.MemoryEngine) error {
	if !verifyLicense(license, path) {
		return errors.New("❌ Invalid license key")
	}
	return LoadKnowledge(path, mem)
}
