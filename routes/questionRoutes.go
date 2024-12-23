package routes

import (
	"just-quizz-server/services"
	"sync"

	"github.com/gin-gonic/gin"
)

func RegisterQuestionGroup(router *gin.Engine, wg *sync.WaitGroup) {
	questionGroup := router.Group("/questions")

	{
		questionGroup.POST("/new", func(ctx *gin.Context) {
			services.CreateQuestion(ctx, wg)
		})

		questionGroup.GET("/random/:theme_id", func(ctx *gin.Context) {
			services.GetRandomQuestionsByTheme(ctx, wg)
		})

		questionGroup.GET("/random", func(ctx *gin.Context) {
			services.GetRandomQuestion(ctx, wg)
		})

		questionGroup.POST("/validation", func(ctx *gin.Context) {
			services.ValidateQuestions(ctx, wg)
		})
	}
}
