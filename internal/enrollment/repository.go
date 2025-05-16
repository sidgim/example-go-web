package enrollment

import (
	"github.com/sidgim/example-go-web/internal/domain"
	"gorm.io/gorm"
	"log"
)

type (
	Repository interface {
		Create(enrollment *domain.Enrollment) error
	}

	repo struct {
		log *log.Logger
		db  *gorm.DB
	}
)

func NewRepository(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (r *repo) Create(enrollment *domain.Enrollment) error {
	if err := r.db.Create(enrollment).Error; err != nil {
		return err
	}
	r.log.Println("Enrollment created:", enrollment)
	return nil
}
