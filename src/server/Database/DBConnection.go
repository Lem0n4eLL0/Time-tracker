package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var DATABASE_ADDR string = "postgres://postgres:zerr0@[::1]:8000/USERS?sslmode=disable"

func DatabaseConnection() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), DATABASE_ADDR)
}
