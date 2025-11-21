package routes

import (
	"github.com/gin-gonic/gin"

	"task-manager-fullstack/config"
	"task-manager-fullstack/handlers"
	"task-manager-fullstack/middlewares"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.ErrorMiddleware())

	r.GET(config.HEALTH_API, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// public route
	// public := r.Group("/api")
	// {
	// 	public.POST(config.REGISTER_API)
	// 	public.POST(config.LOGIN_API)
	// }

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/tasks", handlers.GetDataTasks)
		// protected.POST("/tasks", handlers.CreateTask)
		// protected.GET("/tasks/:id", handlers.GetTaskByID)
		// protected.PUT("/tasks/:id", handlers.UpdateTask)
		// protected.DELETE("/tasks/:id", handlers.DeleteTask)
	}

	return r
}
