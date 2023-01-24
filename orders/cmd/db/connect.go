package db

import (
	"fmt"
	"log"

	"github.com/AlekseyPromet/trinity-example/models"
	"github.com/caarlos0/env"
	"github.com/jmoiron/sqlx"
)

func NewDatabase(dbname string) *sqlx.DB {

	cfg := models.ConfigPostgres{}

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalln(err)
	}

	// postgresql://user:password@netloc:port/dbname[?param1=value1&...]
	connect, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.PostgresUser, cfg.PosttgresPassword, cfg.HostDB, cfg.PortDB, dbname, "disable"),
	)

	if err != nil {
		log.Fatalln(err)
	}

	return connect
}
