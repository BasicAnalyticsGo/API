package controllers

import (
	"basicanalyticsgo/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Return all sessions in database
func GetAllSessions(c *gin.Context) {
	var sessions []models.Session
	models.DB.Find(&sessions)

	c.JSONP(http.StatusOK, gin.H{"data": sessions})
}

type CreateSessionInput struct {
	IP string `json:"IP" binding:"required"`
}

// Create new session
func CreateSession(c *gin.Context) {

	var input CreateSessionInput

	// If the user input is incorrect
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create session object
	session := models.Session{Timestamp: strconv.FormatInt(time.Now().Unix(), 10), IP: input.IP}
	models.DB.Create(&session)

	c.JSON(http.StatusOK, gin.H{"data": session})
}
