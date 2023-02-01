package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/learningPlatform/internal/models"
	"github.com/learningPlatform/internal/service"
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
	r.HandleFunc("/course/{id}", h.updateCourse())
	r.HandleFunc("/users/create", h.createUser())

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
			w.WriteHeader(http.StatusInternalServerError)
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

func (h Handler) updateCourse() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var crs models.UpdateCourse

		err := json.NewDecoder(r.Body).Decode(&crs)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		err = crs.Validate()
		if err != nil {
			value, ok := err.(models.ErrorFields)

			if !ok {
				w.Write([]byte(err.Error()))
				return
			}

			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(value)
			return
		}

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		crs.ID = id

		course, err := h.service.UpdateCourse(crs)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(course)

	}
}

func (h Handler) createUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var user models.User
		err = json.Unmarshal(bytes, &user)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id, err := h.service.CreateUser(user)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(strconv.Itoa(id)))
		w.WriteHeader(http.StatusOK)

	}

}
