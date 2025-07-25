package system

import (
	"fmt"
	"github.com/Trisia/gosysproxy"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"proxypin-go/internal/config"
	"sync"
)

var Once sync.Once

func SysProxyOn() error {
	// 启动时设置系统代理
	addr := fmt.Sprintf("%s:%d", config.Conf.Proxy.Host, config.Conf.Proxy.Port)
	err := gosysproxy.SetGlobalProxy(addr)
	if err != nil {
		log.Errorf("system proxy err: %v", err)
		return err
	}
	log.Warn("system proxy on: ", addr)

	go ExitFunc()
	return nil
}

func SysProxyOff() error {
	err := gosysproxy.Off()
	log.Warn("system proxy off: ", err == nil)
	return err
}

func ExitFunc() {
	Once.Do(func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		s := <-c
		log.Warn("exit: ", s)

		SysProxyOff()
		os.Exit(0)
	})
}
