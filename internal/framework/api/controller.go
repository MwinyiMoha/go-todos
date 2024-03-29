package api

import (
	"go-todos/internal/core/domain"
	"go-todos/internal/core/ports"
	"go-todos/internal/utils/exceptions"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPController struct {
	TodoService ports.TodoService
}

func NewHTTPController(service ports.TodoService) *HTTPController {
	return &HTTPController{
		TodoService: service,
	}
}

func (h *HTTPController) GetTodos(c *gin.Context) {
	todos, err := h.TodoService.GetTodos()
	if err != nil {
		if verr, ok := err.(*exceptions.Exception); ok {
			c.JSON(verr.StatusCode, gin.H{"detail": verr.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *HTTPController) CreateTodo(c *gin.Context) {
	var data domain.TodoPayload

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	todo, err := h.TodoService.CreateTodo(data.Description)
	if err != nil {
		if verr, ok := err.(*exceptions.Exception); ok {
			c.JSON(verr.StatusCode, gin.H{"detail": verr.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *HTTPController) GetTodo(c *gin.Context) {
	todo_id := c.Param("id")
	todo, err := h.TodoService.GetTodo(todo_id)
	if err != nil {
		if verr, ok := err.(*exceptions.Exception); ok {
			c.JSON(verr.StatusCode, gin.H{"detail": verr.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *HTTPController) UpdateTodo(c *gin.Context) {
	var data domain.TodoPayload
	todo_id := c.Param("id")

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	todo, err := h.TodoService.UpdateTodo(todo_id, data.Description)
	if err != nil {
		if verr, ok := err.(*exceptions.Exception); ok {
			c.JSON(verr.StatusCode, gin.H{"detail": verr.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *HTTPController) DeleteTodo(c *gin.Context) {
	todo_id := c.Param("id")
	if err := h.TodoService.DeleteTodo(todo_id); err != nil {
		if verr, ok := err.(*exceptions.Exception); ok {
			c.JSON(verr.StatusCode, gin.H{"detail": verr.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
