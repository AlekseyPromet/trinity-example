package server

import "github.com/valyala/fasthttp"

type ServiceHTTP interface {
	GetHttpAddress() string
	Status(ctx *fasthttp.RequestCtx)
	CreateOrder(ctx *fasthttp.RequestCtx)
	UpdateOrder(ctx *fasthttp.RequestCtx)
	DeleteOrder(ctx *fasthttp.RequestCtx)
}
