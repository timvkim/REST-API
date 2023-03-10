package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/learningPlatform/internal/models"
)

type CourseService interface {
	GetCourses(ctx context.Context) ([]models.Course, error)
	GetCourseById(ctx context.Context, id int) (models.Course, error)
	CreateCourse(ctx context.Context, course models.Course) (int, error)
	UpdateCourse(ctx context.Context, update models.UpdateCourse) (models.Course, error)
	DeleteCourse(ctx context.Context, id int) error
}

type UserService interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
}

type Handler struct {
	courseService CourseService
	userService   UserService
}

func NewHandler(crs CourseService, usr UserService) *Handler {
	return &Handler{
		courseService: crs,
		userService:   usr,
	}
}

func (h Handler) InitRouters() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/courses", h.getCourses()).Methods(http.MethodGet)
	r.HandleFunc("/course/{id}", h.getCourseById()).Methods(http.MethodGet)
	r.HandleFunc("/course/create", h.createCourse()).Methods(http.MethodPost)
	r.HandleFunc("/course/{id}", h.updateCourse()).Methods(http.MethodPut)
	r.HandleFunc("/course/delete/{id}", h.deleteCourse()).Methods(http.MethodDelete)
	r.HandleFunc("/users/create", h.createUser()).Methods(http.MethodPost)

	return r
}

func (h Handler) getCourses() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		result, err := h.courseService.GetCourses(ctx)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, response{err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}

func (h Handler) getCourseById() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			writeJSON(w, http.StatusBadRequest, response{err.Error()})
			return
		}

		result, err := h.courseService.GetCourseById(ctx, id)
		if err != nil {
			if errors.Is(err, models.ErrDoesNotExist) {
				writeJSON(w, http.StatusNotFound, response{err.Error()})
				return
			}
			writeJSON(w, http.StatusInternalServerError, response{err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}

func (h Handler) createCourse() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		w.Header().Set("Content-Type", "application/json")
		log.Print("trying to create a new course...")

		var course models.Course

		err := json.NewDecoder(r.Body).Decode(&course)
		if err != nil {
			log.Println("provided json is invalid")
			writeJSON(w, http.StatusBadRequest, response{err.Error()})
			return
		}

		id, err := h.courseService.CreateCourse(ctx, course)
		if err != nil {
			log.Println("error while creating a new course")
			writeJSON(w, http.StatusInternalServerError, response{err.Error()})
			return
		}

		log.Println("new course succesfully created")
		writeJSON(w, http.StatusCreated, response{strconv.Itoa(id)})

	}
}

func (h Handler) updateCourse() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		var crs models.UpdateCourse

		err := json.NewDecoder(r.Body).Decode(&crs)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, response{err.Error()})
			return
		}

		err = crs.Validate()
		if err != nil {
			value, ok := err.(models.ErrorFields)

			if !ok {
				writeJSON(w, http.StatusInternalServerError, response{err.Error()})
				return
			}

			writeJSON(w, http.StatusBadRequest, value)
			return
		}

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, response{err.Error()})
			return
		}

		crs.ID = id

		course, err := h.courseService.UpdateCourse(ctx, crs)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, response{err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, course)

	}
}

func (h Handler) deleteCourse() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			log.Println("error while getting id")
			writeJSON(w, http.StatusBadRequest, response{err.Error()})
			return
		}

		log.Printf("deleting course %d\n", id)

		err = h.courseService.DeleteCourse(ctx, id)
		if err != nil {
			if errors.Is(err, models.ErrDoesNotExist) {
				writeJSON(w, http.StatusNotFound, response{err.Error()})
				return
			}
			writeJSON(w, http.StatusInternalServerError, response{err.Error()})
			return
		}

		log.Printf("the course by id %d is succesfully deleted", id)
		writeJSON(w, http.StatusNoContent, nil)
	}

}

func (h Handler) createUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, response{err.Error()})
			return
		}

		id, err := h.userService.CreateUser(ctx, user)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, response{err.Error()})
			return
		}

		writeJSON(w, http.StatusCreated, response{strconv.Itoa(id)})

	}

}
