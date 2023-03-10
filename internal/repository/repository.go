package repository

import "github.com/jackc/pgx/v4/pgxpool"

type Repo struct {
	Course Course
	User   User
}

func NewRepo(conn *pgxpool.Pool) *Repo {
	return &Repo{
		Course: *NewCourse(conn),
		User:   *NewUser(conn),
	}
}
