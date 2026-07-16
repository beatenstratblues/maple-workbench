package store

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connection(connectionString string, maxActCnn int, maxIdlCnn int, maxCnnTime int) (*sql.DB, error, string) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err, "FAILED"
	}
	db.SetMaxOpenConns(maxActCnn)
	db.SetMaxIdleConns(maxIdlCnn)
	db.SetConnMaxLifetime(time.Duration(maxCnnTime) * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err, "FAILED"
	}

	return db, nil, "SUCCESS"
}
