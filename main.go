package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sidgim/example-go-web/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func main() {
	router := chi.NewRouter()
	userSrv := user.NewService()
	userEnd := user.MakeEndpoints(userSrv)

	dsn := "host=localhost user=admin password=123456 dbname=example_go_web port=5433 sslmode=disable TimeZone=America/Santiago"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
	}
	_ = db.Debug()

	_ = db.AutoMigrate(&user.User{})

	router.Get("/users/{id}", userEnd.Get)
	router.Get("/users", userEnd.GetAll)
	router.Post("/users", userEnd.Create)
	router.Put("/users", userEnd.Update)
	router.Delete("/users", userEnd.Delete)

	srv := &http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*3, "Timeout!"),
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}
	log.Fatal(srv.ListenAndServe())

}
