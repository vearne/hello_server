package server

import (
	"github.com/valyala/fasthttp"
)


func HelloHandler(ctx *fasthttp.RequestCtx){
	ctx.WriteString("hello world!")
	ctx.SetStatusCode(200)
}

