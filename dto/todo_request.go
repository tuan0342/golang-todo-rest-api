package dto

type CreateTodoRequest struct {
	Item      string `json:"item" binding:"required"`
	Completed bool   `json:"completed"`
}
