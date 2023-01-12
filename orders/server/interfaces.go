package server

import (
	"context"
	"net/http"
)

type ServerIntf interface {
	Run(ctx context.Context) error
	CreateOrder(http.ResponseWriter, *http.Request)
	UpdateOrder(http.ResponseWriter, *http.Request)
	DeleteOrder(http.ResponseWriter, *http.Request)
}
