package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// initialize routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(
			ctx *gin.Context,
		) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
}
