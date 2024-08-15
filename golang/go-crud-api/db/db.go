package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	var err error
	DB, err = pgx.Connect(context.Background(), "postgres://postgres:1122@localhost:5432/cruddb")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}
