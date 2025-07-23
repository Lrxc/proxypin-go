package main

import (
	"proxypin-go/internal"
)

func init() {
	internal.InitConfig()
}

func main() {
	internal.SysProxy()

	internal.StartServer()
}
