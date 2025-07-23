package resources

import (
	"embed"
)

//go:embed *
var Resources embed.FS

func ReadByte(scriptPath string) ([]byte, error) {
	return Resources.ReadFile(scriptPath)
}
