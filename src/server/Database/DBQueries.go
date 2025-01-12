package database

import (
	"context"
	"fmt"
	s "timeTrackerApp/src/server/Structures"
	"timeTrackerApp/src/utils"

	"github.com/jackc/pgx/v5"
)

// var conn *pgx.Conn

func connect() (*pgx.Conn, error) {
	conn, err := DatabaseConnection()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GetUserByName(username string) (*s.User, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(context.Background(),
		"SELECT user_id, username, email, password_hash, created_at, updated_at, roles.name FROM users INNER JOIN roles ON roles.role_id=users.role_id WHERE username=$1 ORDER BY user_id ASC", username)
	if err != nil {
		return nil, err
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[s.User])
	if err != nil {
		return nil, err
	} else if len(users) == 0 {
		return nil, fmt.Errorf("Данный пользователь не найден")
	}
	return &users[0], nil
}

func CreateUser(user *s.User) error {
	conn, err := connect()
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), "INSERT INTO users (username, password_hash, email, role_id) VALUES ($1, $2, $3, $4)", user.Name, utils.Sha512Hashing(user.Password), user.Email, 2)
	if err != nil {
		return err
	}
	return nil
}
