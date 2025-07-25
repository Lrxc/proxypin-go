package main

import (
	"proxypin-go/internal/config"
	"proxypin-go/internal/gui"
	"proxypin-go/internal/system"
)

func init() {
	system.IsAlreadyRunning()
	config.InitConfig()
	config.InitLog()
}

func main() {
	gui.Gui()
}
