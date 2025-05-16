package enrollment

import (
	"errors"
	"github.com/sidgim/example-go-web/internal/course"
	"github.com/sidgim/example-go-web/internal/domain"
	"github.com/sidgim/example-go-web/internal/user"
	"log"
)

type (
	Service interface {
		Create(req CreateRequest) (*domain.Enrollment, error)
	}

	service struct {
		userSrv    user.Service
		courseSrv  course.Service
		repository Repository
		log        *log.Logger
	}
)

func NewService(log *log.Logger, repository Repository, userSrv user.Service, courseSrv course.Service) Service {
	return &service{
		repository: repository,
		log:        log,
		userSrv:    userSrv,
		courseSrv:  courseSrv,
	}
}

func (s *service) Create(req CreateRequest) (*domain.Enrollment, error) {
	enrollment := domain.Enrollment{
		UserID:   req.UserID,
		CourseID: req.CourseID,
		Status:   "Create",
	}

	if _, err := s.userSrv.Get(enrollment.UserID); err != nil {
		return nil, errors.New("user not found")
	}

	if _, err := s.courseSrv.GetById(enrollment.CourseID); err != nil {
		return nil, errors.New("course not found")
	}

	if err := s.repository.Create(&enrollment); err != nil {
		s.log.Println("Error creating enrollment:", err)
		return nil, err
	}
	return &enrollment, nil
}
