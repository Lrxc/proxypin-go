package system

import (
	"os"
	"path/filepath"
)

func IsAlreadyRunning() {
	lockFile := filepath.Join(os.TempDir(), "proxypin-go.lock")
	file, err := os.OpenFile(lockFile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
}
