package api

import (
	"go-todos/internal/core/domain"
	"go-todos/internal/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPController struct {
	todoService ports.TodoService
}

func NewHTTPController(service ports.TodoService) *HTTPController {
	return &HTTPController{
		todoService: service,
	}
}

func (h *HTTPController) GetTodos(c *gin.Context) {
	todos, err := h.todoService.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Could not fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *HTTPController) CreateTodo(c *gin.Context) {
	var data domain.TodoPayload

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	todo, err := h.todoService.CreateTodo(data.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *HTTPController) GetTodo(c *gin.Context) {
	todo_id := c.Param("id")
	todo, err := h.todoService.GetTodo(todo_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *HTTPController) UpdateTodo(c *gin.Context) {
	var data domain.TodoPayload
	todo_id := c.Param("id")

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	todo, err := h.todoService.UpdateTodo(todo_id, data.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *HTTPController) DeleteTodo(c *gin.Context) {
	todo_id := c.Param("id")
	if err := h.todoService.DeleteTodo(todo_id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
