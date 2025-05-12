package user

import (
	"log"
)

type Service interface {
	Create(req CreateRequest) error
}

type service struct {
	log  *log.Logger
	repo Repository
}

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		repo: repo,
		log:  log,
	}
}

func (s *service) Create(req CreateRequest) error {
	s.log.Println("Creating user:", req)
	user := User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}
	if err := s.repo.Create(&user); err != nil {
		s.log.Println("Error creating user:", err)
		return err
	}
	return nil
}
