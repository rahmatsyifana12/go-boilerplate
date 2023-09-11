package services

import "github.com/sarulabs/di"

type Service struct {
	User	UserService
}

func NewService(ioc di.Container) *Service {
	return &Service{
		User: NewUserService(ioc),
	}
}