package main

import (
	"proxypin-go/internal/config"
	"proxypin-go/internal/gui"
)

func init() {
	config.InitConfig()
}

func main() {
	//core.StartServer()
	gui.Gui()
}
