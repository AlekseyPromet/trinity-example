package db

import (
	"fmt"

	"github.com/AlekseyPromet/trinity-example/models"
	"github.com/caarlos0/env"
	"github.com/jmoiron/sqlx"
)

func NewDatabase(dbname string) (*sqlx.DB, error) {

	cfg := models.ConfigPostgres{}

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	// postgresql://user:password@]netloc:port/dbname[?param1=value1&...]
	return sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%s	dbname=%s sslmode=%s",
			cfg.PostgresUser, cfg.PostgresUser, cfg.HostDB, cfg.PosrtDB, dbname, "false"),
	)
}
