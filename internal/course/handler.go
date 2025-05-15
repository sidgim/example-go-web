package course

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sidgim/example-go-web/internal/shared/httphelper"
	"github.com/sidgim/example-go-web/pkg/meta"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var validate = validator.New()

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httphelper.WriteError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := validate.Struct(req); err != nil {
			httphelper.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		course, err := s.Create(req)
		if err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to create course")
			return
		}

		httphelper.WriteSuccess(w, http.StatusCreated, course, nil)
	}
}

func makeGetEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if _, err := uuid.Parse(id); err != nil {
			httphelper.WriteError(w, http.StatusBadRequest, "Invalid UUID")
			return
		}

		if id == "" {
			httphelper.WriteError(w, http.StatusBadRequest, "UUID is required")
			return
		}

		course, err := s.GetById(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				httphelper.WriteError(w, http.StatusNotFound, "Course not found")
				return
			}
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to get course")
			return
		}

		httphelper.WriteSuccess(w, http.StatusOK, course, nil)
	}
}
func makeGetAllEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		filters := Filters{
			Name: v.Get("name"),
		}
		limit, _ := strconv.Atoi(v.Get("limit"))
		offset, _ := strconv.Atoi(v.Get("offset"))

		total, err := s.Count(filters)
		if err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to count courses")
			return
		}

		m, err := meta.New(offset, limit, total)
		if err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to create meta")
			return
		}

		courses, err := s.GetAll(filters, m.Offset(), m.Limit())
		if err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to get all courses")
			return
		}

		httphelper.WriteSuccess(w, http.StatusOK, courses, m)
	}

}
func makeUpdateEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		var req UpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httphelper.WriteError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := validate.Struct(req); err != nil {
			httphelper.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		course, err := s.Update(id, req)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				httphelper.WriteError(w, http.StatusNotFound, "Course not found")
				return
			}
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to update course")
			return
		}

		httphelper.WriteSuccess(w, http.StatusOK, course, nil)
	}
}

func makeDeleteEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if err := s.Delete(id); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				httphelper.WriteError(w, http.StatusNotFound, "Course not found")
				return
			}
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to delete course")
			return
		}

		httphelper.WriteSuccess(w, http.StatusNoContent, nil, nil)
	}
}
