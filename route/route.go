package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rikurunico/gin-basic/handler"
)

func SetupRoute(router *gin.Engine) {

	router.GET("/quotes", handler.GetQuotes)
	router.GET("/quotes/:id", handler.GetQuoteByID)
	router.POST("/quotes", handler.CreateQuote)
	router.PUT("/quotes/:id", handler.UpdateQuote)
	router.DELETE("/quotes/:id", handler.DeleteQuote)
}
