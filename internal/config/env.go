package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"proxypin-go/internal/util"
)

type Config struct {
	System System
	Proxy  Proxy
	Rule   []Rule
}

type System struct {
	MinExit     bool
	Https       bool
	GlobalProxy bool
}

type Proxy struct {
	Host       string
	Port       int
	AutoEnable bool `mapstructure:"auto_enable" yaml:"auto_enable"`
}

type Rule struct {
	Enable bool
	Name   string
	Source string
	Target string
}

var Conf = new(Config)

const confname = "conf.yml"

// 初始化环境参数
func InitConfig() {
	exit := util.FileExist(confname)
	if !exit {
		json := &Config{
			Proxy: Proxy{Host: "127.0.0.1", Port: 10086, AutoEnable: true},
			Rule:  []Rule{{Enable: true, Name: "baidu", Source: "https://www.baidu.com/", Target: "http://www.bing.com/"}},
		}
		//写入默认配置文件
		WriteConf(json)
	}

	viper.AddConfigPath("../../") //配置文件路径
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
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}

	fmt.Printf("config %s ok\n", msg)
}

func WriteJson(msg string) error {
	err := json.Unmarshal([]byte(msg), &Conf.Rule)
	if err != nil {
		return err
	}
	return WriteConf(Conf)
}

func WriteConf(conf *Config) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	join, _ := filepath.Abs(confname)
	return os.WriteFile(join, data, 0644)
}
