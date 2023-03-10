package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/learningPlatform/internal/models"
)

type User struct {
	conn *pgxpool.Pool
}

func NewUser(conn *pgxpool.Pool) *User {
	return &User{
		conn: conn,
	}
}

func (u User) CreateUser(ctx context.Context, user models.User) (int, error) {
	row := u.conn.QueryRow(ctx, "INSERT INTO users (name, login, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Login, user.Password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
