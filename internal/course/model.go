package course

import (
	"gorm.io/gorm"
	"net/http"
	"time"
)

type (
	Course struct {
		ID        string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique_index"`
		Name      string         `json:"name" gorm:"not null"`
		StartDate time.Time      `json:"start_date" gorm:"not null"`
		EndDate   time.Time      `json:"end_date" gorm:"not null"`
		CreatedAt time.Time      `json:"-"`
		UpdatedAt time.Time      `json:"-"`
		DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	}

	CreateRequest struct {
		Name      string    `json:"name" validate:"required"`
		StartDate time.Time `json:"start_date" validate:"required"`
		EndDate   time.Time `json:"end_date" validate:"required"`
	}

	Endpoints struct {
		Create http.HandlerFunc
		Get    http.HandlerFunc
		GetAll http.HandlerFunc
		Update http.HandlerFunc
		Delete http.HandlerFunc
	}

	UpdateRequest struct {
		Name      string    `json:"name" validate:"required"`
		StartDate time.Time `json:"start_date" `
		EndDate   time.Time `json:"end_date" `
	}
)
