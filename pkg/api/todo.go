package api

import ()

type TodoRepository interface {
	Healthcheck(c *gin.Context)
	CreateTodo(c *gin.Context)
	GetTodos(c *gin.Context)
	GetTodoById(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type todoRepository struct {
	DB database.Database
	Ctx *context.Context
}

func NewTodoRepository(db database.Database, ctx *context.Context) *todoRepository {
	return &todoRepository{
		DB: db,
		Ctx: ctx,
	}
}

func (r *todoRepository) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func (r *todoRepository) CreateTodo(c *gin.Context) {
	var req dto.CreateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Item:      req.Item,
		Completed: req.Completed,
	}

	r.DB.Create(&todo)
	c.JSON(200, dto.ConvertToTodoResponse(todo))
}

func (r *todoRepository) GetTodos(c *gin.Context) {
	var todos []models.Todo
	r.DB.Find(&todos)
	// var response []dto.TodoResponse /// response == nil -> true
	response := make([]dto.TodoResponse, 0)
	for _, todo := range todos {
		response = append(response, dto.ConvertToTodoResponse(todo))
	}

	c.JSON(http.StatusOK, response)
}

func (r *todoRepository) GetTodoById(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := r.DB.First(&todo, id).Error(); err != nil {
		c.JSON(404, gin.H{"error": "Todo khong ton tai"})
		return
	}

	c.JSON(200, dto.TodoResponse(todo))
}

func (r *todoRepository) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := r.DB.First(&todo, id).Error(); err != nil {
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

	r.DB.Save(&todo)
	c.JSON(200, dto.ConvertToTodoResponse(todo))
}

func (r *todoRepository) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := r.DB.First(&todo, id).Error(); err != nil {
		c.JSON(404, gin.H{"error": "Todo khong ton tai"})
		return
	}

	r.DB.Delete(&todo)
	c.JSON(200, gin.H{"message": "Deleted"})
}
