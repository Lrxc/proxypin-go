package system

import (
	"fmt"
	"os/exec"
	"runtime"
)

// 安装证书到系统
func InstallCert(certPath string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("certutil", "-addstore", "Root", certPath)
	case "darwin": // macOS
		cmd = exec.Command("security", "add-trusted-cert", "-d", "-r", "trustRoot", "-k", "/Library/Keychains/System.keychain", certPath)
	case "linux":
		// Linux 有多种发行版，这里以Ubuntu为例
		cmd = exec.Command("sudo", "cp", certPath, "/usr/local/share/ca-certificates/")
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("sudo", "update-ca-certificates")
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	output, err := cmd.CombinedOutput()
	fmt.Println("install cert: ", string(output))
	return err
}
