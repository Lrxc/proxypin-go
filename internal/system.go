package internal

import (
	"fmt"
	"github.com/Trisia/gosysproxy"
	"os"
	"os/signal"
)

func SysProxy() {
	// 启动时设置系统代理
	addr := fmt.Sprintf("%s:%s", Conf.System.Host, Conf.System.Port)
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
