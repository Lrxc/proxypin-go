package internal

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	System struct {
		Host string
		Port string
	}
	Rule []struct {
		Enable bool
		Name   string
		Source string
		Target string
	}
}

var Conf = new(Config)

// 初始化环境参数
func InitConfig() {
	viper.AddConfigPath("../") //配置文件路径
	viper.SetConfigFile("conf.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	ReadConf("init")

	// 监听配置文件的变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		ReadConf("reload")
	})
	viper.WatchConfig()
}

func ReadConf(msg string) {
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}

	fmt.Printf("config %s ok\n", msg)
}
