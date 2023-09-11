package services

import "github.com/sarulabs/di"

type Service struct {
}

func NewService(ioc di.Container) *Service {
	return &Service{}
}