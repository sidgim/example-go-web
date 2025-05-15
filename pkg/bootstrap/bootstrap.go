package bootstrap

import (
	"fmt"
	"github.com/sidgim/example-go-web/internal/course"
	"github.com/sidgim/example-go-web/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitLogger() *log.Logger {
	appLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return appLogger
}

func DBConnection() (*gorm.DB, error) {
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
		return nil, err
	}
	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}

	if os.Getenv("DATABASE_MIGRATE") == "true" {
		if err := db.AutoMigrate(&user.User{}); err != nil {
			return nil, err
		}
		if err := db.AutoMigrate(&course.Course{}); err != nil {
			return nil, err
		}
	}

	return db, nil
}
