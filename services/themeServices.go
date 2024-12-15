package services

import (
	"just-quizz-server/database"
	"just-quizz-server/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTheme(c *gin.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	// input validation
	var input models.CreateThemeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	theme := models.Themes{Name: input.Name, Icon_url: input.Icon_url}
	database.DB.Create(&theme)

	c.JSON(http.StatusCreated, gin.H{"data": theme})
}

func FindAllThemes(c *gin.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	var themes []models.Themes
	database.DB.Find(&themes)

	c.JSON(http.StatusOK, gin.H{"data": themes})
}

func FindTheme(c *gin.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	var theme models.Themes

	// check if the id is a valid UUID
	themeID := c.Param("id")

	if err := uuid.Validate(themeID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID!"})
		return
	}

	if err := database.DB.Where("ID= ?", themeID).First(&theme).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": theme})
}

func UpdateTheme(c *gin.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	var theme models.Themes

	if err := database.DB.Where("ID= ?", c.Param("id")).First(&theme).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateThemeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&theme).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": theme})
}

func DeleteTheme(c *gin.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	var theme models.Themes

	if err := database.DB.Where("ID= ?", c.Param("id")).First(&theme).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&theme)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
