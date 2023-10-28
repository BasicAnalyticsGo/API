package controllers

import (
	"basicanalyticsgo/api/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Return all sessions in database
func GetAllSessions(c *gin.Context) {

	fmt.Println("Get all sessions")

	var sessions []models.Session
	models.DB.Find(&sessions)

	c.JSONP(http.StatusOK, gin.H{"data": sessions})
}

type IPDetails struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func getCountry(ip string) IPDetails {
	ipAddress := ip
	url := fmt.Sprintf("http://ip-api.com/json/%s", ipAddress)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP error :", err)
		//return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ReadAll error :", err)
		//return
	}

	var ipDetails IPDetails
	err = json.Unmarshal(body, &ipDetails)
	if err != nil {
		fmt.Println("JSON error :", err)
		//return
	}

	// Debug
	fmt.Println("Country :", ipDetails.Country)
	fmt.Println("City :", ipDetails.City)

	return ipDetails
}

// Create new session
func CreateSession(c *gin.Context) {

	fmt.Println("Create session")

	// Ajouter un tableau qui stocke les 5 dernières IP avec timestamp
	// Si nouvelle requête avec même IP < x secondes alors on refuse

	var ipAddress = c.ClientIP()

	details := getCountry(ipAddress)

	ipDetails := IPDetails{City: details.City, Country: details.Country}

	// Create session object
	session := models.Session{Timestamp: strconv.FormatInt(time.Now().Unix(), 10), IP: ipAddress, Country: ipDetails.Country, City: ipDetails.City}
	models.DB.Create(&session)

	// Faire une requête à la base pour savoir si l'adresse IP a déjà été utilisé
	// Si oui : C'est une nouvelle visite, on enregistre mais pas de requête à une API externe pour récupérer la localisation, on prend celle en base
	// Si non : On fait une requête à une API externe pour récupérer le pays et ou région et on stocke

	c.JSON(http.StatusOK, gin.H{"data": session})
}
