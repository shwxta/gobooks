package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:secret@localhost:5432/book")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Conn = conn
}

func Disconnect() {
	Conn.Close(context.Background())
}
