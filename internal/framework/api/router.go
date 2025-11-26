package api

import (
	"go-todos/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine  *gin.Engine
	service ports.AppService
}

func NewRouter(svc ports.AppService, debug bool) *Router {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	router := Router{
		Engine:  engine,
		service: svc,
	}

	router.attachRoutes()
	return &router
}

func (r *Router) attachRoutes() {
	v1 := r.Engine.Group("api/v1")
	{
		v1.GET("/todos", r.GetTodos)
		v1.POST("/todos", r.CreateTodo)
		v1.GET("/todos/:id", r.GetTodo)
		v1.PUT("/todos/:id", r.UpdateTodo)
		v1.DELETE("/todos/:id", r.DeleteTodo)
	}
}
