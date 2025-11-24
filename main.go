package main

import (
	"example/todo-go/database"
	"example/todo-go/middleware"
	"example/todo-go/models"
	"example/todo-go/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Todo{})

	router := gin.Default()
	router.Use(middleware.JSONBindLogger())
	routers.RegisterTodoRoutes(router)

	router.Run("localhost:9090")
}
