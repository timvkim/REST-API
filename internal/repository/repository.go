package repository

import (
	"github.com/jackc/pgx"
	"github.com/learningPlatform/internal/models"
)

type Repo struct {
	conn *pgx.Conn
}

func NewRepo(conn *pgx.Conn) *Repo {
	return &Repo{
		conn: conn,
	}
}

func (r Repo) CreateCourse(course models.Course) (int, error) {
	row := r.conn.QueryRow("INSERT INTO courses (title, description, price, author_id) VALUES ($1, $2, $3, $4) RETURNING id",
		course.Title, course.Description, course.Price, course.AuthorID)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r Repo) GetCourses() ([]models.Course, error) {
	rows, err := r.conn.Query("SELECT id, title, description, price, author_id FROM courses")
	if err != nil {
		return nil, err
	}

	var rowSlice []models.Course
	for rows.Next() {
		var r models.Course
		err := rows.Scan(&r.ID, &r.Title, &r.Description, &r.Price, &r.AuthorID)
		if err != nil {
			return nil, err
		}
		rowSlice = append(rowSlice, r)
	}
	return rowSlice, nil
}

func (r Repo) UpdateCourse(update models.UpdateCourse) (models.Course, error) {
	query := `
		UPDATE courses
		SET title = $1, description = $2, price = $3
		WHERE id = $4
		RETURNING id, title, description, price, author_id`

	args := []interface{}{
		update.Title,
		update.Description,
		update.Price,
		update.ID,
	}

	course := models.Course{}

	err := r.conn.QueryRow(query, args...).Scan(&course.ID, &course.Title, &course.Description, &course.Price, &course.AuthorID)
	if err != nil {
		return models.Course{}, err
	}
	return course, nil
}

func (r Repo) CreateUser(user models.User) (int, error) {
	row := r.conn.QueryRow("INSERT INTO users (name, login, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Login, user.Password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
