package domain

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique_index"`
	Name      string         `json:"name" gorm:"not null"`
	StartDate time.Time      `json:"start_date" gorm:"not null"`
	EndDate   time.Time      `json:"end_date" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
