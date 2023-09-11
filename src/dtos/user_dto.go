package dtos

import "go-boilerplate/src/models"

type GetUserByIDResponse struct {
	models.User
}

type GetUserByIDParams struct {
	UserID uint `param:"user_id" validate:"required"`
}