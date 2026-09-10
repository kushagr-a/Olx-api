package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDb(databaseUrl string) (*sql.DB, error) {

	// connection open
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	// set connection pool
	db.SetMaxOpenConns(25)                 // maximum no. of open connections
	db.SetMaxIdleConns(25)                 // maximum no. of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // maximum no. of idle time

	// fail fast
	// context use for data sending
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // defer used for call function at the time of function exit

	// check connection is alive or not
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db.PingContext: %w", err)
	}

	return db, nil
}
