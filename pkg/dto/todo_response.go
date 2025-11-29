package dto

import "example/todo-go/pkg/models"

type TodoResponse struct {
	ID        int    `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func ConvertToTodoResponse(todo models.Todo) TodoResponse {
	return TodoResponse{
		ID:        todo.ID,
		Item:      todo.Item,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}
