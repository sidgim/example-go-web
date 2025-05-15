package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/sidgim/example-go-web/internal/course"
	"github.com/sidgim/example-go-web/internal/user"
	"net/http"
)

func NewRouter(
	uH *user.UserHandler,
	cH *course.CourseHandler,
) http.Handler {
	r := chi.NewRouter()
	// middlewares…
	r.Route("/users", uH.Mount)
	r.Route("/courses", cH.Mount)
	return r
}
