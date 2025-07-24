package system

import (
	"fmt"
	"github.com/Trisia/gosysproxy"
	"os"
	"os/signal"
	"proxypin-go/internal/config"
	"sync"
)

var Once sync.Once

func SysProxyOn() error {
	// 启动时设置系统代理
	addr := fmt.Sprintf("%s:%d", config.Conf.System.Host, config.Conf.System.Port)
	err := gosysproxy.SetGlobalProxy(addr)
	if err != nil {
		fmt.Errorf("system proxy err: %v", err)
		return err
	}
	fmt.Println("system proxy on: ", addr)

	go ExitFunc()
	return nil
}

func SysProxyOff() error {
	err := gosysproxy.Off()
	fmt.Println("system proxy off: ", err == nil)
	return err
}

func ExitFunc() {
	Once.Do(func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		s := <-c
		fmt.Println("exit: ", s)

		SysProxyOff()
		os.Exit(0)
	})
}
