package domain

import "time"

type Enrollment struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique_index"`
	UserID    string    `json:"user_id,omitempty" gorm:"not null"`
	User      *User     `json:"user,omitempty"`
	CourseID  string    `json:"course_id,omitempty" gorm:"not null"`
	Course    *Course   `json:"course,omitempty"`
	Status    string    `json:"status,omitempty" gorm:"not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
