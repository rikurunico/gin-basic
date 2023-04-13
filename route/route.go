package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rikurunico/gin-basic/handler"
)

func SetupRoute(router *gin.Engine) {

	quotes := router.Group("/quotes")
	{
		quotes.GET("", handler.GetQuotes)
		quotes.GET("/:id", handler.GetQuoteByID)
		quotes.POST("", handler.CreateQuote)
		quotes.PUT("/:id", handler.UpdateQuote)
		quotes.DELETE("/:id", handler.DeleteQuote)
	}

}
