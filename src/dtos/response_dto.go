package dtos

type Response struct {
	Status	interface{} `json:"status"`
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}