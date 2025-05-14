package user

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

		user, err := s.Create(req)
		if err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to create user")
			return
		}
		httphelper.WriteSuccess(w, http.StatusCreated, user, nil)
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
		user, err := s.Get(id)
		if err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to get user")
			return
		}

		if user == nil {
			httphelper.WriteError(w, http.StatusNotFound, "User not found")
			return
		}
		httphelper.WriteSuccess(w, http.StatusOK, user, nil)
	}
}

func makeGetAllEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()

		filters := Filters{
			FirstName: v.Get("first_name"),
			LastName:  v.Get("last_name"),
		}
		limit, _ := strconv.Atoi(v.Get("limit"))
		page, _ := strconv.Atoi(v.Get("offset"))

		count, err := s.Count(filters)
		if err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to count users")
			return
		}
		m, err := meta.New(page, limit, count)

		if err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to create meta")
			return
		}

		user, err := s.GetAll(filters, m.Offset(), m.Limit())
		if err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to get all users")
			return
		}
		httphelper.WriteSuccess(w, http.StatusOK, user, m)
	}
}

func makeUpdateEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1) sacamos y validamos el ID
		idParam := chi.URLParam(r, "id")
		if _, err := uuid.Parse(idParam); err != nil {
			httphelper.WriteError(w, http.StatusBadRequest, "Invalid UUID")
			return
		}

		// 2) decodificamos el body al DTO
		var req UpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httphelper.WriteError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// 3) validamos según las tags `validate:"…"`
		if err := validate.Struct(req); err != nil {
			httphelper.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		// 4) llamamos al service que retorna el usuario actualizado
		updated, err := s.UpdateContact(idParam, req)
		if err != nil {
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				httphelper.WriteError(w, http.StatusNotFound, "User not found")
			default:
				httphelper.WriteError(w, http.StatusInternalServerError, "Failed to update user")
			}
			return
		}

		httphelper.WriteSuccess(w, http.StatusOK, updated, nil)
	}
}

func makeDeleteEndpoint(s Service) http.HandlerFunc {
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

		if err := s.Delete(id); err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to delete user")
			return
		}

		httphelper.WriteSuccess(w, http.StatusNoContent, "User deleted successfully", nil)
	}
}
