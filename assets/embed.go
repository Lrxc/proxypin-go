package assets

import (
	"embed"
)

//go:embed *
var Resources embed.FS

func ReadByte(scriptPath string) ([]byte, error) {
	return Resources.ReadFile(scriptPath)
}

func Read(scriptPath string) []byte {
	file, err := Resources.ReadFile(scriptPath)
	if err != nil {
		return nil
	}
	return file
}
