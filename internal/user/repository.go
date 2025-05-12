package user

import (
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	Create(user *User) error
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepository(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		db:  db,
		log: log}
}

func (r *repo) Create(user *User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	r.log.Println("User created:", user)
	return nil
}
