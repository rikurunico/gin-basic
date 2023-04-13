package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rikurunico/gin-basic/model"
	"github.com/rikurunico/gin-basic/database"

)

func GetQuotes(c *gin.Context) {
	var quotes []model.Quote
	database.DB.Find(&quotes)

	c.JSON(200, quotes)
}

func GetQuoteByID(c *gin.Context) {
	var quote model.Quote
	id := c.Param("id")
	database.DB.First(&quote, id)

	if quote.ID == 0 {
		c.JSON(404, gin.H{"message": "quote not found"})
		return
	}

	c.JSON(200, quote)
}

func CreateQuote(c *gin.Context) {
	var quote model.Quote
	if err := c.ShouldBindJSON(&quote); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&quote)

	c.JSON(201, quote)
}

func UpdateQuote(c *gin.Context) {
	var quote model.Quote
	id := c.Param("id")

	if err := database.DB.First(&quote, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "quote not found"})
		return
	}

	if err := c.ShouldBindJSON(&quote); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&quote)

	c.JSON(200, quote)
}

func DeleteQuote(c *gin.Context) {
	var quote model.Quote
	id := c.Param("id")

	if err := database.DB.First(&quote, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "quote not found"})
		return
	}

	database.DB.Delete(&quote)

	c.JSON(200, gin.H{"message": "quote deleted"})
}
