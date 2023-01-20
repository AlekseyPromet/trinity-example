package handler

import (
	"context"

	"github.com/AlekseyPromet/trinity-example/models"
	"github.com/AlekseyPromet/trinity-example/orders/cmd/db"
)

func NewServiceRPC(ctx context.Context, cfg *models.ConfigOrder) *ServiceRPC {

	srv := &ServiceRPC{
		CtxService:  ctx,
		DatabaseOrm: db.NewDatabase("orders"),
	}

	srv.StatusEndpoint = srv.MakeStatusEndpoint()
	srv.CreateOrderEndpoint = srv.MakeCreateOrderEndpoint()
	srv.UpdateOrderEndpoint = srv.MakeUpdateOrderEndpoint()
	srv.DeleteOrderEndpoint = srv.MakeDeleteOrderEndpoint()

	return srv
}
