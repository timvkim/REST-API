package service

import (
	"context"

	"github.com/learningPlatform/internal/models"
)

type CourseStorage interface {
	GetCourses(ctx context.Context) ([]models.Course, error)
	GetCourseById(ctx context.Context, id int) (models.Course, error)
	CreateCourse(ctx context.Context, course models.Course) (int, error)
	UpdateCourse(ctx context.Context, update models.UpdateCourse) (models.Course, error)
	DeleteCourse(ctx context.Context, id int) error
}

type UserStorage interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
}

type Service struct {
	CourseStorage CourseStorage
	UserStorage   UserStorage
}

func NewService(crs CourseStorage, usr UserStorage) *Service {
	return &Service{
		CourseStorage: crs,
		UserStorage:   usr,
	}
}
