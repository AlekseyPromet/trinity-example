package migrations

import (
	"database/sql"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MigrationUP(db *sql.DB, dbName, schemaName string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName: dbName,
		SchemaName:   schemaName,
	})
	if err != nil {
		return err
	}
	migr, err := migrate.NewWithDatabaseInstance(
		"file:///sql///"+dbName,
		dbName,
		driver,
	)
	if err != nil {
		return err
	}

	return migr.Up()
}
