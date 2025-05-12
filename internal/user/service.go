package user

import "fmt"

type Service interface {
	Create(req CreateRequest) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Create(req CreateRequest) error {
	fmt.Println("Creating user:", req)
	return nil
}
