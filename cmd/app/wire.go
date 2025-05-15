package main

import (
	"github.com/google/wire"
	"github.com/sidgim/example-go-web/internal/course"
	"github.com/sidgim/example-go-web/internal/server"
	"github.com/sidgim/example-go-web/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ----------- DB Provider -----------
func provideDB() (*gorm.DB, error) {
	dsn := "host=localhost user=... dbname=... password=..."
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// ----------- Course Module -----------
var courseSet = wire.NewSet(
	provideDB,            // inyecta *gorm.DB
	course.NewRepository, // func NewRepository(*gorm.DB) *Repo
	course.NewService,    // func NewService(*Repo) *Service
)

// ----------- User Module -----------
var userSet = wire.NewSet(
	provideDB,
	user.NewRepository, // func NewRepository(*gorm.DB) *Repo
	user.NewService,    // func NewService(*Repo) *Service
)

// ----------- Server & Router -----------
var serverSet = wire.NewSet(
	server.NewRouter, // func NewRouter(*user.Service, *course.Service) http.Handler
	server.NewServer, // func NewServer(http.Handler) *http.Server
)

// ----------- Application Set -----------
var appSet = wire.NewSet(
	courseSet,
	userSet,
	serverSet,
)

// InitializeApp construye *http.Server con todos los módulos
func InitializeApp() (*server.Server, error) {
	wire.Build(appSet)
	return &server.Server{}, nil
}
