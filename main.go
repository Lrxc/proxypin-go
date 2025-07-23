package main

import (
	"proxypin-go/internal/config"
	"proxypin-go/internal/core"
	"proxypin-go/internal/system"
)

func init() {
	config.InitConfig()
	system.SysProxy()
}

func main() {
	core.StartServer()
}
