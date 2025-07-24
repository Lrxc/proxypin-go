package system

import (
	"fmt"
	"github.com/Trisia/gosysproxy"
	"os"
	"os/signal"
	"proxypin-go/internal/config"
)

func SysProxyOn() {
	if !config.Conf.System.AutoEnable {
		return
	}

	// 启动时设置系统代理
	addr := fmt.Sprintf("%s:%d", config.Conf.System.Host, config.Conf.System.Port)
	if err := gosysproxy.SetGlobalProxy(addr); err != nil {
		fmt.Errorf("system proxy err: %v", err)
	}
	fmt.Println("system proxy on: ", addr)

	go ExitFunc()
}

func SysProxyOff() {
	if !config.Conf.System.AutoEnable {
		return
	}

	gosysproxy.Off()
	fmt.Println("system proxy off")
}

func ExitFunc() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("exit: ", s)

	SysProxyOff()
	os.Exit(0)
}
