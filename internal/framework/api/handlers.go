package api

import (
	"go-todos/internal/core/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mwinyimoha/commons/pkg/errors"
)

func (r *Router) GetTodos(c *gin.Context) {
	todos, err := r.service.GetTodos()
	if err != nil {
		if cerr, ok := err.(*errors.Error); ok {
			code, detail := cerr.HTTPStatus()
			c.JSON(code, detail)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (r *Router) CreateTodo(c *gin.Context) {
	var data domain.TodoPayload
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	todo, err := r.service.CreateTodo(data.Description)
	if err != nil {
		if cerr, ok := err.(*errors.Error); ok {
			code, detail := cerr.HTTPStatus()
			c.JSON(code, detail)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (r *Router) GetTodo(c *gin.Context) {
	todoId := c.Param("id")
	todo, err := r.service.GetTodo(todoId)
	if err != nil {
		if cerr, ok := err.(*errors.Error); ok {
			code, detail := cerr.HTTPStatus()
			c.JSON(code, detail)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (r *Router) UpdateTodo(c *gin.Context) {
	todoId := c.Param("id")

	var data domain.TodoPayload
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	todo, err := r.service.UpdateTodo(todoId, data.Description)
	if err != nil {
		if cerr, ok := err.(*errors.Error); ok {
			code, detail := cerr.HTTPStatus()
			c.JSON(code, detail)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (r *Router) DeleteTodo(c *gin.Context) {
	todo_id := c.Param("id")
	if err := r.service.DeleteTodo(todo_id); err != nil {
		if cerr, ok := err.(*errors.Error); ok {
			code, detail := cerr.HTTPStatus()
			c.JSON(code, detail)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
