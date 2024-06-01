package router

import (
	"github.com/gabriel1734/goopportunities/handler/opening"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", opening.ShowOpeningHandler)

		v1.POST("/opening", opening.CreateOpeningHandler)

		v1.DELETE("/opening", opening.DeleteOpeningHandler)

		v1.PUT("/opening", opening.UpdateOpeningHandler)

		v1.GET("/openings", opening.ListOpeningHandler)
	}
}
