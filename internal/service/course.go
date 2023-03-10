package service

import (
	"context"

	"github.com/learningPlatform/internal/models"
)

type Course struct {
	courseStorage CourseStorage
}

func NewCourse(courseStorage CourseStorage) *Course {
	return &Course{
		courseStorage: courseStorage,
	}
}

func (c Course) GetCourses(ctx context.Context) ([]models.Course, error) {
	return c.courseStorage.GetCourses(ctx)
}

func (c Course) GetCourseById(ctx context.Context, id int) (models.Course, error) {
	return c.courseStorage.GetCourseById(ctx, id)
}

func (c Course) CreateCourse(ctx context.Context, course models.Course) (int, error) {
	return c.courseStorage.CreateCourse(ctx, course)
}

func (c Course) UpdateCourse(ctx context.Context, update models.UpdateCourse) (models.Course, error) {
	return c.courseStorage.UpdateCourse(ctx, update)
}

func (c Course) DeleteCourse(ctx context.Context, id int) error {
	return c.courseStorage.DeleteCourse(ctx, id)
}
