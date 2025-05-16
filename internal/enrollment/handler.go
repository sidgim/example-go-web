package enrollment

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/sidgim/example-go-web/internal/shared/httphelper"
	"net/http"
)

type Handler struct {
	svc Service
}

var validate = validator.New()

func NewEnrollmentHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Mount(r chi.Router) {
	r.Post("/", h.Create)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.WriteError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}
	if err := validate.Struct(req); err != nil {
		httphelper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	enrollment, err := h.svc.Create(req)
	if err != nil {
		httphelper.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	httphelper.WriteSuccess(w, http.StatusCreated, enrollment, nil)
}
