package dtos

import "go-boilerplate/src/models"

type GetUserByIDResponse struct {
	models.User
}

type GetUserByIDParams struct {
	UserID uint `param:"user_id" validate:"required"`
}

type CreateUserRequest struct {
	Username	string	`json:"username" validate:"required"`
	Password	string	`json:"password" validate:"required"`
}

type UpdateUserParams struct {
	UserID		uint	`param:"user_id" validate:"required"`
	FullName	string	`json:"full_name"`
	PhoneNumber	string	`json:"phone_number"`
}