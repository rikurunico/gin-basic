package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Quote struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

var quotes = []Quote{
	{ID: "1", Text: "Keep calm and code on", Author: "Unknown"},
	{ID: "2", Text: "Talk is cheap, show me the code", Author: "Linus Torvalds"},
	{ID: "3", Text: "Any fool can write code that a computer can understand. Good programmers write code that humans can understand.", Author: "Martin Fowler"},
	{ID: "4", Text: "Programs must be written for people to read, and only incidentally for machines to execute.", Author: "Harold Abelson and Gerald Jay Sussman"},
	{ID: "5", Text: "The only way to do great work is to love what you do.", Author: "Steve Jobs"},
	{ID: "6", Text: "The best way to predict the future is to invent it.", Author: "Alan Kay"},
	{ID: "7", Text: "Quality is not an act, it is a habit.", Author: "Aristotle"},
	{ID: "8", Text: "It's not that I'm so smart, it's just that I stay with problems longer.", Author: "Albert Einstein"},
	{ID: "9", Text: "In order to be irreplaceable, one must always be different.", Author: "Coco Chanel"},
}

func getQuotes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": quotes})
}

func getQuoteByID(c *gin.Context) {
	id := c.Param("id")
	for _, q := range quotes {
		if q.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": q})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "quote not found"})
}

func createQuote(c *gin.Context) {
	var input Quote
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newQuote := Quote{
		ID:     strconv.Itoa(len(quotes) + 1),
		Text:   input.Text,
		Author: input.Author,
	}
	quotes = append(quotes, newQuote)
	c.JSON(http.StatusCreated, gin.H{"data": newQuote})
}

func updateQuote(c *gin.Context) {
	id := c.Param("id")
	var input Quote
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, q := range quotes {
		if q.ID == id {
			quotes[i] = Quote{
				ID:     q.ID,
				Text:   input.Text,
				Author: input.Author,
			}
			c.JSON(http.StatusOK, gin.H{"data": quotes[i]})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "quote not found"})
}

func deleteQuote(c *gin.Context) {
	id := c.Param("id")
	for i, q := range quotes {
		if q.ID == id {
			quotes = append(quotes[:i], quotes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"data": "quote deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "quote not found"})
}

func main() {
	r := gin.Default()

	r.GET("/quotes", getQuotes)
	r.GET("/quotes/:id", getQuoteByID)
	r.POST("/quotes", createQuote)
	r.PUT("/quotes/:id", updateQuote)
	r.DELETE("/quotes/:id", deleteQuote)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
