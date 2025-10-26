// internal/infrastructure/db/postgres.go
package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectPostgres(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}
