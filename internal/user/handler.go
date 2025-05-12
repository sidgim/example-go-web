package user

import (
	"encoding/json"
	"github.com/sidgim/example-go-web/internal/shared/httphelper"
	"net/http"
)

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

		if req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Phone == "" {
			httphelper.WriteError(w, http.StatusBadRequest, "All fields are required")
			return
		}

		err := s.Create(req)
		if err != nil {
			httphelper.WriteError(w, http.StatusInternalServerError, "Failed to create user")
			return
		}

		json.NewEncoder(w).Encode(req)
	}
}

func makeDeleteEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeGetEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeGetAllEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeUpdateEndpoint(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
