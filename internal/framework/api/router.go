package api

import "github.com/gin-gonic/gin"

type Router struct {
	controller Controller
	engine     *gin.Engine
}

func NewRouter(controller Controller) *Router {
	return &Router{
		controller: controller,
		engine:     gin.Default(),
	}
}

func (r *Router) AddRoutes() {
	v1 := r.engine.Group("api/v1")
	{
		v1.GET("/todos", r.controller.GetTodos)
		v1.POST("/todos", r.controller.CreateTodo)
		v1.GET("/todos/:id", r.controller.GetTodo)
		v1.PUT("/todos/:id", r.controller.UpdateTodo)
		v1.DELETE("/todos/:id", r.controller.DeleteTodo)
	}
}

func (r *Router) Serve() error {
	if err := r.engine.Run(); err != nil {
		return err
	}

	return nil
}
