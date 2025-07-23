package system

import (
	"fmt"
	"github.com/Trisia/gosysproxy"
	"os"
	"os/signal"
	"proxypin-go/internal/config"
)

func SysProxy() {
	if !config.Conf.System.AutoEnable {
		return
	}

	// 启动时设置系统代理
	addr := fmt.Sprintf("%s:%s", config.Conf.System.Host, config.Conf.System.Port)
	if err := gosysproxy.SetGlobalProxy(addr); err != nil {
		fmt.Errorf("system proxy err: %v", err)
	}
	fmt.Println("system proxy set: ", addr)

	go ExitFunc()
}

func ExitFunc() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("exit: ", s)

	gosysproxy.Off() //关闭代理
	os.Exit(0)
}
