package util

import (
	"os"
	"path/filepath"
)

func FileExist(filename string) bool {
	abs, _ := filepath.Abs(filename)

	_, err := os.Stat(abs)
	if err == nil {
		return true // 文件存在
	}
	if os.IsNotExist(err) {
		return false // 文件不存在
	}
	return false // 其他错误（如权限不足），通常也视为不存在
}
