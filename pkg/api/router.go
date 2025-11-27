package api

func NewRouter(db database.Database, ctx *context.Context) *gin.Engine {
	todoRepository := NewTodoRepository(db, ctx)
	r := gin.Default()
	r.Use(middleware.JSONBindLogger())

	v1 := r.Group("/api/v1")

	todo := v1.Group("/todos")
	{
		todo.GET("/", todoRepository.GetTodos)
		todo.GET("/:id", todoRepository.GetTodoById)
		todo.POST("/", todoRepository.CreateTodo)
		todo.PUT("/:id", todoRepository.UpdateTodo)
		todo.DELETE("/:id", todoRepository.DeleteTodo)
		todo.GET("/health-check", todoRepository.Healthcheck)
	}

	return r
}