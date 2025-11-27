package main

import (
	"example/todo-go/database"
	"example/todo-go/middleware"
	"example/todo-go/models"
	"example/todo-go/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using default env")
	}

	db := database.NewDatabase()
	dbWrapper := &database.GormDatabase{DB: db}
	ctx := context.Backgroud()

	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := api.NewRouter(dbWrapper, &ctx)

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("============= Running default port 9090 =============")
		port = "9090" // default
	}
	r.Run(":" + port)
}
