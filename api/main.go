package main

import (
	"github.com/gin-gonic/gin"
	"github.com/takahiromitsui/go-tasks/api/routes"
)

func main() {
	router := gin.Default()
	routes.SetupTaskRoutes(router)
	router.Run(":8080")
}