package main

import (
	"fmt"
	"flag"
	"os"
	"runtime"
	"hello_server/constants"
	"hello_server/context"
	"hello_server/server"
	log "github.com/sirupsen/logrus"
	"github.com/erikdubbelboer/gspt"
	"os/signal"
	"syscall"
	"sync"
)

// 设置环境变量
func setupEnv(){
	// 设置最大的CPU数量
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 设置进程名称
	gspt.SetProcTitle(constants.Name)
}


func GracefulExit(){
	log.Info("register signal handler")
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	switch <-ch {
	case syscall.SIGTERM:
		Exit()
		break
	}
}

func Exit(){
	//singal.Notify(quit, syscall.SIGTERM)
	log.Info("get signal SIGTERM, prepare exit!!!")
	// 停止webserver
	server.Listener.Close()
	log.Info("get signal SIGTERM, success exit!!!")

}


func main() {
	setupEnv()

	version := flag.Bool("v", false, "show version")
	cfg := flag.String("c", "./etc/config.yaml", "configuration file")

	flag.Parse()

	if *version {
		fmt.Println(constants.Version)
		os.Exit(0)
	}


	fmt.Println("configure file:", *cfg)
	// 1. 加载配置
	// 配置信息保存在 context.t 中
	context.ParseConfig(*cfg)

	// 2. 日志
	context.InitLogger()
	log.Info("init logger success")

	// 3. 初始化webserver并启动
	log.Info("start server ... ...")
	wg := sync.WaitGroup{}
	go server.NewServer(&wg)

	// 4. 注册信号处理逻辑
	GracefulExit()

	wg.Wait()

}
