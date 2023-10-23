package main

import (
	"basicanalyticsgo/api/controllers"
	"basicanalyticsgo/api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/sessions", controllers.GetAllSessions)
	router.POST("/sessions", controllers.CreateSession)

	router.Run("localhost:8080")
}
