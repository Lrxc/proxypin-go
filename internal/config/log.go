package config

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"proxypin-go/internal/constant"
	"proxypin-go/internal/util"
)

func InitLog() {
	util.CreateParentFile(constant.LogPath)

	//保存文件
	file, _ := os.OpenFile(constant.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//同时将日志写入文件和控制台
	writer := io.MultiWriter(file, os.Stdout)

	// 日志级别
	log.SetLevel(log.InfoLevel)
	//日志格式化
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true, //显示颜色
		FullTimestamp: true, //完整时间
	})
	//写入文件
	log.SetOutput(writer)
}
