package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	env "github.com/caarlos0/env/v6"
	"github.com/go-kit/log"
	"github.com/valyala/fasthttp"

	"github.com/AlekseyPromet/trinity-example/models"
	"github.com/AlekseyPromet/trinity-example/orders/cmd/handler"
)

type serverV1 struct {
	port   string
	logger log.Logger
	srv    *handler.ServiceRPC
	ctx    context.Context
}

var _ ServiceHTTP = (*serverV1)(nil)

func NewService(ctx context.Context) ServiceHTTP {
	cfg := &models.ConfigOrder{
		ConfigPostgres: &models.ConfigPostgres{},
		ConfigKafka:    &models.ConfigKafka{},
	}

	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
		return nil
	}

	serviceRpc := handler.NewServiceRPC(ctx, cfg)

	srv := &serverV1{
		port:   cfg.HttpBind,
		logger: log.NewJSONLogger(os.Stdout),
		srv:    serviceRpc,
	}

	return srv
}

func (srv *serverV1) GetHttpAddress() string {
	return ":" + srv.port
}

func midleware(request *fasthttp.RequestCtx) {
	request.Response.Header.Add("Content-Type", "application/json; charset=utf-8")
}

func (sv1 *serverV1) Status(request *fasthttp.RequestCtx) {

	midleware(request)

	response, err := sv1.srv.StatusEndpoint(sv1.ctx, nil)
	if err != nil {
		request.Response.Header.Add("Error", err.Error())
		request.SetStatusCode(http.StatusBadRequest)
		return
	}

	if resp, ok := response.(models.StatusResponse); ok {

		if body, err := json.Marshal(resp); err == nil {
			request.Write(body)
			return
		}

	}

	request.Response.Header.Add("Error", err.Error())
	request.SetStatusCode(http.StatusBadRequest)
}

func (sv1 *serverV1) CreateOrder(request *fasthttp.RequestCtx) {

	midleware(request)

	body := request.Request.Body()

	order := new(models.OrderReq)
	err := json.Unmarshal(body, order)
	if err != nil {
		request.Response.Header.Add("Error", err.Error())
		request.SetStatusCode(http.StatusBadRequest)
		return
	}

	response, err := sv1.srv.CreateOrderEndpoint(sv1.ctx, order)
	if err != nil {
		request.Response.Header.Add("Error", err.Error())
		request.SetStatusCode(http.StatusBadRequest)
		return
	}

	if resp, ok := response.(models.OrderRes); ok {

		if body, err := json.Marshal(resp); err == nil {
			request.SetStatusCode(http.StatusOK)
			request.Write(body)
			return
		}

	}

	request.Response.Header.Add("Error", err.Error())
	request.SetStatusCode(http.StatusBadRequest)
}

func (sv1 *serverV1) UpdateOrder(request *fasthttp.RequestCtx) {

	midleware(request)

	body := request.Request.Body()

	order := new(models.OrderRes)
	err := json.Unmarshal(body, order)
	if err != nil {
		request.Response.Header.Add("Error", err.Error())
		request.SetStatusCode(http.StatusBadRequest)
		return
	}

	response, err := sv1.srv.UpdateOrderEndpoint(sv1.ctx, body)
	if err != nil {
		request.Response.Header.Add("Error", err.Error())
		request.SetStatusCode(http.StatusBadRequest)
		return
	}

	if resp, ok := response.(models.OrderRes); ok {

		if body, err := json.Marshal(resp); err == nil {
			request.SetStatusCode(http.StatusOK)
			request.Write(body)
			return
		}

	}

	request.Response.Header.Add("Error", err.Error())
	request.SetStatusCode(http.StatusBadRequest)
}

func (sv1 *serverV1) DeleteOrder(request *fasthttp.RequestCtx) {

	midleware(request)

	id := request.UserValue("id")

	response, err := sv1.srv.DeleteOrderEndpoint(sv1.ctx, id)
	if err != nil {
		request.Response.Header.Add("Error", err.Error())
		request.SetStatusCode(http.StatusBadRequest)
		return
	}

	if _, ok := response.(string); ok {
		request.SetStatusCode(http.StatusOK)
		return
	}

	request.Response.Header.Add("Error", err.Error())
	request.SetStatusCode(http.StatusBadRequest)
}
