package api

import "github.com/gin-gonic/gin"

type Controller interface {
	GetTodos(ctx *gin.Context)
	CreateTodo(ctx *gin.Context)
	GetTodo(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
}
