package dtos

type CreateTodoRequest struct {
	Title string `json:"title"`
	Content string `json:"content"`
}