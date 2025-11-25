package main

import (
	"example/todo-go/database"
	"example/todo-go/middleware"
	"example/todo-go/models"
	"example/todo-go/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using default env")
	}
	
	database.Connect()
	database.DB.AutoMigrate(&models.Todo{})

	router := gin.Default()
	router.Use(middleware.JSONBindLogger())
	routers.RegisterTodoRoutes(router)

	router.Run("localhost:9090")
}
