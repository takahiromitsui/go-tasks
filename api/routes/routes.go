package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/takahiromitsui/go-tasks/api/handlers"
)

func SetupTaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/api/tasks")
	taskGroup.GET("/", handlers.GetTasks)
}