package services

import (
	"errors"
	"io"
	"just-quizz-server/database"
	"just-quizz-server/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateQuestion(c *gin.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	// input validation
	var input models.CreateQuestionInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {

		if errors.Is(err, io.EOF) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty or invalid"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		return
	}

	var theme models.Themes
	if err := database.DB.First(&theme, input.Theme_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Theme ID"})
		return
	}

	question := models.Questions{Question: input.Question, Difficulty: input.Difficulty, Proposed_response: input.Proposed_response, Correct_answer: input.Correct_answer, Theme_id: input.Theme_id}
	database.DB.Create(&question)

	c.JSON(http.StatusCreated, gin.H{"data": question})
}

func GetRandomQuestionsByTheme(c *gin.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	themeID := c.Param("theme_id")

	parsedThemeID, err := uuid.Parse(themeID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Theme ID"})
		return
	}

	var questions []models.GetRandomQuestions

	if err := database.DB.
		Model(models.Questions{}).
		Where("Theme_id = ?", parsedThemeID).
		Order("RANDOM()").
		Limit(10).
		Find(&questions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "No questions found for this theme"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": questions})
}

func GetRandomQuestion(c *gin.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	var questions []models.GetRandomQuestions

	if err := database.DB.
		Model(models.Questions{}).
		Order("RANDOM()").
		Limit(10).Find(&questions).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No questions found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": questions})
}

func ValidateQuestions(c *gin.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	var input models.AnswerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var QuestionIDs []string
	for _, answer := range input.Answers {
		QuestionIDs = append(QuestionIDs, answer.Question_id.String())
	}

	var questions []models.Questions
	if err := database.DB.Where("id IN ?", QuestionIDs).Find(&questions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Questions not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	questionMap := make(map[string]models.Questions)
	for _, question := range questions {
		questionMap[question.ID.String()] = question
	}

	var totalScore int
	for _, userAnswer := range input.Answers {
		if question, exists := questionMap[userAnswer.Question_id.String()]; exists {
			// Comparer la réponse de l'utilisateur avec la réponse correcte
			if userAnswer.Answer == question.Correct_answer {
				points := 10 + ((question.Difficulty - 1) * 5)
				totalScore += points
			}
		}
	}

	// TO DO -> add totalScore to current user global score
	// code ...

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"score": totalScore}})

}
