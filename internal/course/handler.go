package course

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sidgim/example-go-web/internal/shared/httphelper"
	"github.com/sidgim/example-go-web/pkg/meta"
	"gorm.io/gorm"
)

var validate = validator.New()

type CourseHandler struct {
	svc Service
}

func NewCourseHandler(svc Service) *CourseHandler {
	return &CourseHandler{svc: svc}
}

// Mount engancha los endpoints de course al router
func (h *CourseHandler) Mount(r chi.Router) {
	r.Get("/", h.GetAll)
	r.Post("/", h.Create)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.GetByID)
		r.Put("/", h.Update)
		r.Delete("/", h.Delete)
	})
}

func (h *CourseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.WriteError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}
	if err := validate.Struct(req); err != nil {
		httphelper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	crs, err := h.svc.Create(req)
	if err != nil {
		httphelper.WriteError(w, http.StatusInternalServerError, "failed to create course")
		return
	}
	httphelper.WriteSuccess(w, http.StatusCreated, crs, nil)
}

func (h *CourseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if _, err := uuid.Parse(id); err != nil {
		httphelper.WriteError(w, http.StatusBadRequest, "invalid UUID")
		return
	}

	crs, err := h.svc.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httphelper.WriteError(w, http.StatusNotFound, "course not found")
			return
		}
		httphelper.WriteError(w, http.StatusInternalServerError, "failed to fetch course")
		return
	}
	httphelper.WriteSuccess(w, http.StatusOK, crs, nil)
}

func (h *CourseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))
	filters := Filters{
		Name: q.Get("name"),
	}

	total, err := h.svc.Count(filters)
	if err != nil {
		httphelper.WriteError(w, http.StatusInternalServerError, "count failed")
		return
	}
	m, err := meta.New(offset, limit, total)
	if err != nil {
		httphelper.WriteError(w, http.StatusInternalServerError, "meta generation failed")
		return
	}

	list, err := h.svc.GetAll(filters, m.Offset(), m.Limit())
	if err != nil {
		httphelper.WriteError(w, http.StatusInternalServerError, "failed to fetch courses")
		return
	}
	httphelper.WriteSuccess(w, http.StatusOK, list, m)
}

func (h *CourseHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if _, err := uuid.Parse(id); err != nil {
		httphelper.WriteError(w, http.StatusBadRequest, "invalid UUID")
		return
	}

	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.WriteError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}
	if err := validate.Struct(req); err != nil {
		httphelper.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	crs, err := h.svc.Update(id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httphelper.WriteError(w, http.StatusNotFound, "course not found")
			return
		}
		httphelper.WriteError(w, http.StatusInternalServerError, "failed to update course")
		return
	}
	httphelper.WriteSuccess(w, http.StatusOK, crs, nil)
}

func (h *CourseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if _, err := uuid.Parse(id); err != nil {
		httphelper.WriteError(w, http.StatusBadRequest, "invalid UUID")
		return
	}

	if err := h.svc.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httphelper.WriteError(w, http.StatusNotFound, "course not found")
			return
		}
		httphelper.WriteError(w, http.StatusInternalServerError, "failed to delete course")
		return
	}

	httphelper.WriteSuccess(w, http.StatusNoContent, nil, nil)
}
