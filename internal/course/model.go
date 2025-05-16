package course

import (
	"net/http"
	"time"
)

type (
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
