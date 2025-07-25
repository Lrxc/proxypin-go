package system

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
)

func IsAdmin() bool {
	switch runtime.GOOS {
	case "windows":
		// Windows: 尝试执行需要管理员权限的操作
		cmd := exec.Command("net", "session")
		out, err := cmd.CombinedOutput()
		fmt.Println("isAdmin", string(out))
		return err == nil
	case "linux", "darwin":
		// Unix-like: 检查UID
		return syscall.Geteuid() == 0
	default:
		return false
	}
}
