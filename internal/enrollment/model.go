package enrollment

type (
	CreateRequest struct {
		UserID   string `json:"user_id" validate:"required"`
		CourseID string `json:"course_id" validate:"required"`
	}
)
