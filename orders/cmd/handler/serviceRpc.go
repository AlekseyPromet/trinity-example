package handler

import (
	"context"
	"log"

	"github.com/AlekseyPromet/trinity-example/migrations"
	"github.com/AlekseyPromet/trinity-example/models"
	"github.com/AlekseyPromet/trinity-example/orders/cmd/db"
)

func NewServiceRPC(ctx context.Context, cfg *models.ConfigOrder) *ServiceRPC {

	db := db.NewDatabase(cfg.PostgresDB)

	srv := &ServiceRPC{
		CtxService:  ctx,
		DatabaseOrm: db,
	}

	if err := migrations.MigrationUP(db.DB, cfg.PostgresDB, "trinity"); err != nil {
		log.Fatalln(err)
	}

	srv.StatusEndpoint = srv.MakeStatusEndpoint()
	srv.CreateOrderEndpoint = srv.MakeCreateOrderEndpoint()
	srv.UpdateOrderEndpoint = srv.MakeUpdateOrderEndpoint()
	srv.DeleteOrderEndpoint = srv.MakeDeleteOrderEndpoint()

	return srv
}
