//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/wire"
	"github.com/sidgim/example-go-web/internal/course"
	"github.com/sidgim/example-go-web/internal/server"
	"github.com/sidgim/example-go-web/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 1️⃣ Providers básicos

func provideDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		pass = "123456"
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "example_go_web"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
	course.NewService, // si también pide logger, dale provideLogger
	course.NewCourseHandler,
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
)

func InitializeApp() (*server.Server, error) {
	wire.Build(appSet)
	return &server.Server{}, nil
}
