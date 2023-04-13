package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rikurunico/gin-basic/database"
	"github.com/rikurunico/gin-basic/model"
)

func GetQuotes(c *gin.Context) {
	var quotes []model.Quote
	if err := database.DB.Find(&quotes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quotes)
}

func GetQuoteByID(c *gin.Context) {
	var quote model.Quote
	id := c.Param("id")
	if err := database.DB.First(&quote, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "quote not found"})
		return
	}
	c.JSON(http.StatusOK, quote)
}

func CreateQuote(c *gin.Context) {
	var quote model.Quote
	if err := c.ShouldBindJSON(&quote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&quote).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, quote)
}

func UpdateQuote(c *gin.Context) {
	var quote model.Quote
	id := c.Param("id")
	if err := database.DB.First(&quote, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "quote not found"})
		return
	}
	if err := c.ShouldBindJSON(&quote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&quote).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quote)
}

func DeleteQuote(c *gin.Context) {
	var quote model.Quote
	id := c.Param("id")
	if err := database.DB.First(&quote, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "quote not found"})
		return
	}
	if err := database.DB.Delete(&quote).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "quote deleted"})
}
