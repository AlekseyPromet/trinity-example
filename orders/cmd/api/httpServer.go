package api

import (
	"github.com/AlekseyPromet/trinity-example/orders/server"
	rt "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func NewHTTPServer(addess string, service server.ServiceHTTP) error {

	router := rt.New()

	router.Handle("GET", "/status", service.Status)

	router.Handle("POST", "/order", service.CreateOrder)

	router.Handle("UPDATE", "/order", service.UpdateOrder)

	router.Handle("DELETE", "/order", service.DeleteOrder)

	return fasthttp.ListenAndServe(addess, router.Handler)
}
