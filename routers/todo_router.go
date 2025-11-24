package routers

import (
	"example/todo-go/database"
	"example/todo-go/dto"
	"example/todo-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(r *gin.Engine) {
	r.POST("/todos", createTodo)
	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getTodoById)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)
}

func createTodo(c *gin.Context) {
	var req dto.CreateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Item:      req.Item,
		Completed: req.Completed,
	}

	database.DB.Create(&todo)
	c.JSON(200, dto.ConvertToTodoResponse(todo))
}

func getTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	// var response []dto.TodoResponse /// response == nil -> true
	response := make([]dto.TodoResponse, 0)
	for _, todo := range todos {
		response = append(response, dto.ConvertToTodoResponse(todo))
	}

	c.JSON(http.StatusOK, response)
}

func getTodoById(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo khong ton tai"})
		return
	}

	c.JSON(200, dto.TodoResponse(todo))
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo khong ton tai"})
		return
	}

	var req dto.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todo.Item = req.Item
	todo.Completed = req.Completed

	database.DB.Save(&todo)
	c.JSON(200, dto.ConvertToTodoResponse(todo))
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo khong ton tai"})
		return
	}

	database.DB.Delete(&todo)
	c.JSON(200, gin.H{"message": "Deleted"})
}
