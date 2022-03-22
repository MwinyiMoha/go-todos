package api

import "github.com/gin-gonic/gin"

type Router struct {
	Controller Controller
	Engine     *gin.Engine
}

func NewRouter(controller Controller) *Router {
	return &Router{
		Controller: controller,
		Engine:     gin.Default(),
	}
}

func (r *Router) AddRoutes() {
	v1 := r.Engine.Group("api/v1")
	{
		v1.GET("/todos", r.Controller.GetTodos)
		v1.POST("/todos", r.Controller.CreateTodo)
		v1.GET("/todos/:id", r.Controller.GetTodo)
		v1.PUT("/todos/:id", r.Controller.UpdateTodo)
		v1.DELETE("/todos/:id", r.Controller.DeleteTodo)
	}
}

func (r *Router) Serve() error {
	if err := r.Engine.Run(); err != nil {
		return err
	}

	return nil
}
