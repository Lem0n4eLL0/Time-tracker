package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func main() {

	var err error

	conn, err = pgx.Connect(context.Background(), "postgres://postgres:zerr0@[::1]:8000/USERS?sslmode=disable")
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer conn.Close(context.Background())
	test()
}

func test() {
	// rows, _ := conn.Query(context.Background(), "SELECT * FROM users WHERE username='vladik'") // WHERE username='vladik'
	// fmt.Println(rows)
	// users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	// if err != nil {
	// 	fmt.Printf("CollectRows error: %v", err)
	// 	return
	// }
	// for _, p := range users {
	// 	p.displayUser()
	// }

}
