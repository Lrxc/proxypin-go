package main

import (
	"proxypin-go/internal/config"
	"proxypin-go/internal/core"
	"proxypin-go/internal/gui"
	"proxypin-go/internal/system"
)

func init() {
	system.IsAlreadyRunning()
	config.InitConfig()
}

func main() {
	go core.StartServer()
	gui.Gui()
}
