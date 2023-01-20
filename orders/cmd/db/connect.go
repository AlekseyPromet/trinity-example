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

	// postgresql://user:password@]netloc:port/dbname[?param1=value1&...]
	connect, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%s	dbname=%s sslmode=%s",
			cfg.PostgresUser, cfg.PostgresUser, cfg.HostDB, cfg.PosrtDB, dbname, "false"),
	)

	if err != nil {
		log.Fatalln(err)
	}

	return connect
}
