package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sidgim/example-go-web/internal/user"
	"github.com/sidgim/example-go-web/pkg/bootstrap"
	"log"
	"net/http"
	"time"
)

func main() {
	router := chi.NewRouter()
	_ = godotenv.Load()
	l := bootstrap.InitLogger()
	db, err := bootstrap.DBConnection()
	if err != nil {
		l.Fatal(err)
	}

	userRepo := user.NewRepository(l, db)
	userSrv := user.NewService(l, userRepo)
	userEnd := user.MakeEndpoints(userSrv)

	router.Get("/users/{id}", userEnd.Get)
	router.Get("/users", userEnd.GetAll)
	router.Post("/users", userEnd.Create)
	router.Put("/users/{id}", userEnd.Update)
	router.Delete("/users/{id}", userEnd.Delete)

	srv := &http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*3, "Timeout!"),
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}
	log.Fatal(srv.ListenAndServe())

}
