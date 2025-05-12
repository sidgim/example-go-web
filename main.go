package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sidgim/example-go-web/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	router := chi.NewRouter()
	l := log.New(os.Stdout, "api: ", log.LstdFlags)
	_ = godotenv.Load()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
	}
	_ = db.Debug()

	_ = db.AutoMigrate(&user.User{})

	userRepo := user.NewRepository(l, db)
	userSrv := user.NewService(l, userRepo)
	userEnd := user.MakeEndpoints(userSrv)

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
