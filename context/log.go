package context

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/pkg/errors"
	"path"
	"time"
	//"os"
	"os"
	"bufio"
)

func setNull(){
	// 1. load iplib
	src, _ := os.Open(os.DevNull)
	writer := bufio.NewWriter(src)
	log.SetOutput(writer)
}

func InitLogger() {
	baseLogPath := path.Join(GlobalConfig.LogConf.Logdir,
		GlobalConfig.LogConf.Filename)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	//log.SetFormatter(&log.TextFormatter{})
	switch level := GlobalConfig.LogConf.LogLevel; level {
	/*
	如果日志级别不是debug就不要打印日志到控制台了
 	*/
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		setNull()
		log.SetLevel(log.InfoLevel)
	case "warn":
		setNull()
		log.SetLevel(log.WarnLevel)
	case "error":
		setNull()
		log.SetLevel(log.ErrorLevel)
	default:
		setNull()
		log.SetLevel(log.InfoLevel)
	}


	if err != nil {
		log.Errorf("config local file system logger error. %v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	})
	log.AddHook(lfHook)
}
