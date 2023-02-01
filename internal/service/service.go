package service

import (
	"github.com/learningPlatform/internal/models"
	"github.com/learningPlatform/internal/repository"
)

type Service struct {
	repo *repository.Repo
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s Service) GetCourses() ([]models.Course, error) {
	courses, err := s.repo.GetCourses()
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (s Service) CreateCourse(course models.Course) (int, error) {
	id, err := s.repo.CreateCourse(course)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s Service) UpdateCourse(update models.UpdateCourse) (models.Course, error) {
	course, err := s.repo.UpdateCourse(update)
	if err != nil {
		return models.Course{}, err
	}

	return course, nil

}

func (s *Service) CreateUser(user models.User) (int, error) {
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
