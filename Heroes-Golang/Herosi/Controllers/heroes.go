//controllers/heroes.go

package controllers

import (
	"net/http"

	"herosi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//GetHeroes
func GetHeroes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var heroes []models.Hero
	db.Find(&heroes)

	c.JSON(http.StatusOK, heroes)
}

//AddHero
func AddHero(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.HeroInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hero := models.Hero{Name: input.Name}
	db.Create(&hero)

	c.JSON(http.StatusOK, hero)
}

// GetHero
func GetHero(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("id")
	// Get model if exist
	var hero models.Hero
	if err := db.Where("id = ?", id).First(&hero).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hero with id: " + id + " not found!"})
		return
	}

	c.JSON(http.StatusOK, hero)
}

//PutHero
func PutHero(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("id")
	// Get model if exist
	var hero models.Hero
	if err := db.Where("id = ?", id).First(&hero).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hero with id: " + id + " not found!"})
		return
	}

	// Validate input
	var input models.HeroInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&hero).Updates(input)

	c.JSON(http.StatusOK, hero)
}

// DeleteHero
func DeleteHero(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("id")
	// Get model if exist
	var hero models.Hero
	if err := db.Where("id = ?", id).First(&hero).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hero with id: " + id + " not found!"})
		return
	}

	db.Delete(&hero)

	c.JSON(http.StatusOK, true)
}
