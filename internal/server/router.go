package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/sidgim/example-go-web/internal/course"
	"github.com/sidgim/example-go-web/internal/enrollment"
	"github.com/sidgim/example-go-web/internal/user"
	"net/http"
)

func NewRouter(
	uH *user.Handler,
	cH *course.Handler,
	eH *enrollment.Handler,
) http.Handler {
	r := chi.NewRouter()

	r.Route("/users", uH.Mount)
	r.Route("/courses", cH.Mount)
	r.Route("/enrollments", eH.Mount)
	return r
}
