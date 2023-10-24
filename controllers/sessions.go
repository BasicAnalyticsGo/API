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

// Create new session
func CreateSession(c *gin.Context) {

	var ipAddress = c.ClientIP()

	// Create session object
	session := models.Session{Timestamp: strconv.FormatInt(time.Now().Unix(), 10), IP: ipAddress}
	models.DB.Create(&session)

	// Faire une requête à la base pour savoir si l'adresse IP a déjà été utilisé
	// Si oui : C'est une nouvelle visite, on enregistre mais pas de requête à une API externe pour récupérer la localisation, on prend celle en base
	// Si non : On fait une requête à une API externe pour récupérer le pays et ou région et on stocke

	c.JSON(http.StatusOK, gin.H{"data": session})
}
