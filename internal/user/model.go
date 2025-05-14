package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type (
	Endpoints struct {
		Create http.HandlerFunc
		Get    http.HandlerFunc
		GetAll http.HandlerFunc
		Update http.HandlerFunc
		Delete http.HandlerFunc
	}

	CreateRequest struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
		Phone     string `json:"phone" validate:"required,e164"`
	}

	User struct {
		ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4();not null;unique_index"`
		FirstName string         `json:"first_name" gorm:"not null";type:varchar(50);not null"`
		LastName  string         `json:"last_name" gorm:"not null";type:varchar(50);not null"`
		Email     string         `json:"email" gorm:"not null";type:varchar(50);not null"`
		Phone     string         `json:"phone" gorm:"not null";type:varchar(30);not null"`
		CreatedAt time.Time      `json:"-"`
		UpdatedAt time.Time      `json:"-"`
		DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	}

	UpdateRequest struct {
		Email string `json:"email" validate:"required,email"`
		Phone string `json:"phone" validate:"required,e164"`
	}
)
