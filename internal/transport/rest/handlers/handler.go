package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/learningPlatform/internal/models"
	"github.com/learningPlatform/internal/service"
	"io"
	"net/http"
	"strconv"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) InitRouters() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/courses", h.getCourses())
	r.HandleFunc("/courses/create", h.createCourse())

	return r
}

func (h Handler) getCourses() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.service.GetCourses()
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}

		bytes, err := json.Marshal(result)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
		w.WriteHeader(http.StatusOK)
	}
}

func (h Handler) createCourse() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		var course models.Course
		err = json.Unmarshal(bytes, &course)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id, err := h.service.CreateCourse(course)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(strconv.Itoa(id)))
		w.WriteHeader(http.StatusOK)

	}
}
