package dtos

import "go-boilerplate/src/models"

type GetTodoByIDResponse struct {
	models.Todo
}

type CreateTodoRequest struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type TodoIDParams struct {
	ID uint `param:"id" validate:"required"`
}