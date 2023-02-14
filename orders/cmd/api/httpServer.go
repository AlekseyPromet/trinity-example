package api

import (
	"github.com/AlekseyPromet/trinity-example/orders/server"
	rt "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func NewHTTPServer(addess string, service server.ServiceHTTP) error {

	router := rt.New()

	groupV1 := router.Group("/v1")

	groupV1.Handle("GET", "/status", service.Status)

	groupV1.Handle("POST", "/order", service.CreateOrder)

	groupV1.Handle("UPDATE", "/order", service.UpdateOrder)

	groupV1.Handle("DELETE", "/order/{id}", service.DeleteOrder)

	return fasthttp.ListenAndServe(addess, router.Handler)
}
