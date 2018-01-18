package server

import (
	"hello_server/context"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"os"
	"net"
	"sync"
	"time"
)


var Listener net.Listener

func NewServer(wgp * sync.WaitGroup) {
	wgp.Add(1)
	defer wgp.Done()

	router := fasthttprouter.New()
	// 1. 获取IP信息
	// /v1/ipinfo?ip=xxx
	router.GET("/hello", BasicRecover(HelloHandler))

	addr := fmt.Sprintf("%v:%v", context.GlobalConfig.Bind.Host,
		context.GlobalConfig.Bind.Port)

	log.WithFields(log.Fields{
		"addr": addr,
	}).Info("start server")

	ln, err := net.Listen("tcp4", addr)
	if err != nil {
		log.WithField("error", err).Error("start server fail")
		os.Exit(-5)
	}

	// 最大等待3秒
	Listener = NewGracefulListener(ln, time.Second * 3)
	err = fasthttp.Serve(Listener, router.Handler)
	if err != nil {
		log.WithField("error", err).Error("start server fail")
		os.Exit(-5)
	}

	log.Info("server stop listen")

}




