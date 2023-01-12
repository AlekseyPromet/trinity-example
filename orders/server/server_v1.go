package server

import (
	"context"
	"fmt"
	"net/http"

	env "github.com/caarlos0/env/v6"

	"github.com/AlekseyPromet/trinity-example/models"
)

type ServerV1 struct {
	port string
}

var _ ServerV1 = (*ServerIntf)(nil)

func NewServer() ServerIntf {
	cfg := &models.ConfigGoods{
		ConfigPostgres: &models.ConfigPostgres{},
		ConfigKafka:    &models.ConfigKafka{},
	}

	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
		return nil
	}

	srv := &ServerV1{
		port: cfg.HttpBind,
	}

	return srv
}

func (srv *ServerV1) Run(ctx context.Context) error {
	return nil
}

func (srv *ServerV1) CreateOrder(w http.ResponseWriter, r *http.Request) {

}

func (srv *ServerV1) UpdateOrder(http.ResponseWriter, *http.Request) {

}

func (srv *ServerV1) DeleteOrder(http.ResponseWriter, *http.Request) {

}
