package api

import (
	"go-todos/internal/core/ports"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	Engine  *gin.Engine
	service ports.AppService
}

func NewRouter(svc ports.AppService, logger *zap.Logger, debug bool) *Router {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(logger, true))

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
