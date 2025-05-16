//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"github.com/sidgim/example-go-web/internal/domain"
	"log"
	"os"

	"github.com/google/wire"
	"github.com/sidgim/example-go-web/internal/course"
	"github.com/sidgim/example-go-web/internal/enrollment"
	"github.com/sidgim/example-go-web/internal/server"
	"github.com/sidgim/example-go-web/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 1️⃣ Providers básicos

func provideDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
		return nil, err
	}
	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}

	if os.Getenv("DATABASE_MIGRATE") == "true" {
		fmt.Println("DEBUG MODE ACTIVATED")

		if err := db.AutoMigrate(&domain.User{}); err != nil {
			return nil, err
		}
		if err := db.AutoMigrate(&domain.Course{}); err != nil {
			return nil, err
		}
		if err := db.AutoMigrate(&domain.Enrollment{}); err != nil {
			return nil, err
		}
	}
	return db, nil
}

// Aquí le dices a Wire cómo crear el logger
func provideLogger() *log.Logger {
	return log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lshortfile)
}

var baseSet = wire.NewSet(
	provideDB,
	provideLogger,
)

// 2️⃣ Módulos

var userSet = wire.NewSet(
	user.NewRepository,  // func NewRepository(*gorm.DB) *UserRepo
	user.NewService,     // func NewService(*UserRepo, *log.Logger) *Service
	user.NewUserHandler, // func NewUserHandler(*Service) *UserHandler
)

var courseSet = wire.NewSet(
	course.NewRepository,
	course.NewService,
	course.NewCourseHandler,
)

var enrollmentSet = wire.NewSet(
	enrollment.NewRepository,
	enrollment.NewService,
	enrollment.NewEnrollmentHandler,
)

// 3️⃣ Server & Router

var serverSet = wire.NewSet(
	server.NewRouter, // func NewRouter(*UserHandler, *CourseHandler) http.Handler
	server.NewServer, // func NewServer(http.Handler) *Server
)

var appSet = wire.NewSet(
	baseSet,
	userSet,
	courseSet,
	serverSet,
	enrollmentSet,
)

func InitializeApp() (*server.Server, error) {
	wire.Build(appSet)
	return &server.Server{}, nil
}
