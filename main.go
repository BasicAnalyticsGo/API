package main

import (
	"basicanalyticsgo/api/controllers"
	"basicanalyticsgo/api/models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/sessions", controllers.GetAllSessions)
	router.POST("/sessions", controllers.CreateSession)

	err := router.Run(":8080")

	if err != nil {
		log.Fatal(err)
		return
	}

}
