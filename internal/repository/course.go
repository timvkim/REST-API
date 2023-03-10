package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/learningPlatform/internal/models"
)

type Course struct {
	conn *pgxpool.Pool
}

func NewCourse(conn *pgxpool.Pool) *Course {
	return &Course{
		conn: conn,
	}
}

func (c Course) CreateCourse(ctx context.Context, course models.Course) (int, error) {

	row := c.conn.QueryRow(ctx, "INSERT INTO courses (title, description, price, author_id) VALUES ($1, $2, $3, $4) RETURNING id",
		course.Title, course.Description, course.Price, course.AuthorID)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (c Course) GetCourses(ctx context.Context) ([]models.Course, error) {
	rows, err := c.conn.Query(ctx, "SELECT id, title, description, price, author_id FROM courses")
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

func (c Course) GetCourseById(ctx context.Context, id int) (models.Course, error) {
	query := `
		SELECT id, title, description, price, author_id
		FROM courses
		WHERE id = $1`

	var course models.Course

	err := c.conn.QueryRow(ctx, query, id).Scan(&course.ID, &course.Title, &course.Description, &course.Price, &course.AuthorID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Course{}, models.ErrDoesNotExist
		}
		return models.Course{}, err
	}

	return course, nil
}

func (c Course) UpdateCourse(ctx context.Context, update models.UpdateCourse) (models.Course, error) {
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

	err := c.conn.QueryRow(ctx, query, args...).Scan(&course.ID, &course.Title, &course.Description, &course.Price, &course.AuthorID)
	if err != nil {
		return models.Course{}, err
	}
	return course, nil
}

func (c Course) DeleteCourse(ctx context.Context, id int) error {
	query := `
		DELETE FROM courses
		WHERE id = $1`

	result, err := c.conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("course with id %d: %w", id, models.ErrDoesNotExist)
	}

	return nil
}
