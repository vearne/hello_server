package server

import (
	"github.com/valyala/fasthttp"
	"p2p_dispatch/db"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"runtime"
)

func BasicRecover(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				// 参数错误
				state := db.State{Code: 1001, Message: "unknow error"}
				fields := log.Fields{}
				fields["uri"] = string(ctx.RequestURI())
				fields["error"] = err
				fields["stack"] = string(stack())

				log.WithFields(fields).Error("RequestHandler error")
				res := make(map[string]db.State)
				res["state"] = state
				out, _ := json.Marshal(res)
				ctx.Write(out)
				ctx.SetStatusCode(501)
			}
		}()

		// 实际执行
		h(ctx)
	})
}

func stack() []byte {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	return buf[:n]
}