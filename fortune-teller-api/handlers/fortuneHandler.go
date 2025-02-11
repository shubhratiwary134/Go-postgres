package handlers

import (
	"encoding/json"
	"fortune-teller-api/models"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadFortunesFromJSON(db *gorm.DB) {
	var count int64
	db.Model(&models.Fortune{}).Count(&count)

	if count > 0 {
		log.Println("Database already has fortunes, skipping JSON load.")
		return
	}
	file, err := os.ReadFile("data/fortunes.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}
	var fortunes []models.Fortune
	if err := json.Unmarshal(file, &fortunes); err != nil {
		log.Fatal("Error parsing JSON:", err)
	}
	db.Create(&fortunes)
	log.Println("Seeded fortunes from JSON successfully.")
}
func GetRandomFortune(c, *gin.Context) {
	var fortunes []models.Fortune
	result := database.DB.Find(&fortunes)
	if result.Error != nil || len(fortunes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No fortunes found"})
		return
	}
	randomIndex := rand.Intn(len(fortunes))
	c.JSON(http.StatusOK, gin.H{"fortune": fortunes[randomIndex].Fortune})
}
