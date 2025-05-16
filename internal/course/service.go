package course

import (
	"github.com/sidgim/example-go-web/internal/domain"
	"log"
)

type (
	Service interface {
		Create(req CreateRequest) (*domain.Course, error)
		GetById(id string) (*domain.Course, error)
		GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
		Update(id string, updateRequest UpdateRequest) (*domain.Course, error)
		Delete(id string) error
		Count(filters Filters) (int, error)
	}

	service struct {
		repository Repository
		log        *log.Logger
	}

	Filters struct {
		Name string
	}
)

func NewService(log *log.Logger, repository Repository) Service {
	return &service{
		repository: repository,
		log:        log,
	}
}

func (s *service) Create(req CreateRequest) (*domain.Course, error) {
	course := domain.Course{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}
	if err := s.repository.Create(&course); err != nil {
		s.log.Println("Error creating course:", err)
		return nil, err
	}
	return &course, nil
}
func (s *service) GetById(id string) (*domain.Course, error) {
	course, err := s.repository.GetById(id)
	if err != nil {
		s.log.Println("Error getting course:", err)
		return nil, err
	}
	return course, nil
}

func (s *service) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	courses, err := s.repository.GetAll(filters, offset, limit)
	if err != nil {
		s.log.Println("Error getting all courses:", err)
		return nil, err
	}
	return courses, nil
}
func (s *service) Update(id string, updateRequest UpdateRequest) (*domain.Course, error) {
	course, err := s.repository.GetById(id)
	if err != nil {
		s.log.Println("Error getting course:", err)
		return nil, err
	}
	if course == nil {
		s.log.Println("domain not found")
		return nil, nil
	}

	course.EndDate = updateRequest.EndDate
	course.StartDate = updateRequest.StartDate
	course.Name = updateRequest.Name

	if err := s.repository.Update(id, updateRequest); err != nil {
		s.log.Println("Error updating course:", err)
		return nil, err
	}
	return course, nil
}
func (s *service) Delete(id string) error {
	course, err := s.repository.GetById(id)
	if err != nil {
		s.log.Println("Error getting course:", err)
		return err
	}
	if course == nil {
		s.log.Println("domain.Course not found")
		return nil
	}

	if err := s.repository.Delete(id); err != nil {
		s.log.Println("Error deleting course:", err)
		return err
	}
	return nil
}

func (s *service) Count(filters Filters) (int, error) {
	return s.repository.Count(filters)
}
