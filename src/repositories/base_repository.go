package repositories

import "github.com/sarulabs/di"

type Repository struct {
}

func NewRepository(ioc di.Container) *Repository {
	return &Repository{}
}