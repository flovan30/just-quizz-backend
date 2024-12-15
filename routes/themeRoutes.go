package routes

import (
	"just-quizz-server/services"
	"sync"

	"github.com/gin-gonic/gin"
)

func RegisterThemeGroup(router *gin.Engine, wg *sync.WaitGroup) {
	themeGroup := router.Group("/themes")

	{
		themeGroup.POST("/new", func(ctx *gin.Context) {
			services.CreateTheme(ctx, wg)
		})

		themeGroup.GET("", func(ctx *gin.Context) {
			services.FindAllThemes(ctx, wg)
		})

		themeGroup.GET("/:id", func(ctx *gin.Context) {
			services.FindTheme(ctx, wg)
		})

		themeGroup.PATCH("/:id", func(ctx *gin.Context) {
			services.UpdateTheme(ctx, wg)
		})

		themeGroup.DELETE("/:id", func(ctx *gin.Context) {
			services.DeleteTheme(ctx, wg)
		})
	}
}
