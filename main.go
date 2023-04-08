package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Quote struct {
	gorm.Model
	Text   string
	Author string
}

var (
	db *gorm.DB
)

var quotes = []Quote{
	{Text: "Keep calm and code on", Author: "Unknown"},
	{Text: "Talk is cheap, show me the code", Author: "Linus Torvalds"},
	{Text: "Premature optimization is the root of all evil", Author: "Donald Knuth"},
	{Text: "Any fool can write code that a computer can understand. Good programmers write code that humans can understand.", Author: "Martin Fowler"},
	{Text: "Programs must be written for people to read, and only incidentally for machines to execute.", Author: "Harold Abelson"},
}

func main() {
	// Load configuration from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read configuration using viper
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading .env file")
	}

	// Get database configuration values
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")

	// Create DSN string and connect to database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Auto migrate schema
	err = db.AutoMigrate(&Quote{})
	if err != nil {
		log.Fatalf("Error migrating schema: %v", err)
	}

	// Seed the database
	for _, quote := range quotes {
		db.Create(&quote)
	}

	// Start Gin server
	r := gin.Default()
	r.GET("/quotes", getQuotes)
	r.GET("/quotes/:id", getQuoteByID)
	r.POST("/quotes", createQuote)
	r.PUT("/quotes/:id", updateQuote)
	r.DELETE("/quotes/:id", deleteQuote)
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting Gin server: %v", err)
	}
}

func getQuotes(c *gin.Context) {
	var quotes []Quote
	db.Find(&quotes)

	c.JSON(200, quotes)
}

func getQuoteByID(c *gin.Context) {
	var quote Quote
	id := c.Param("id")
	db.First(&quote, id)

	if quote.ID == 0 {
		c.JSON(404, gin.H{"message": "quote not found"})
		return
	}

	c.JSON(200, quote)
}

func createQuote(c *gin.Context) {
	var quote Quote
	if err := c.ShouldBindJSON(&quote); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Create(&quote)

	c.JSON(201, quote)
}

func updateQuote(c *gin.Context) {
	var quote Quote
	id := c.Param("id")

	if err := db.First(&quote, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "quote not found"})
		return
	}

	if err := c.ShouldBindJSON(&quote); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Save(&quote)

	c.JSON(200, quote)
}

func deleteQuote(c *gin.Context) {
	var quote Quote
	id := c.Param("id")

	if err := db.First(&quote, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "quote not found"})
		return
	}

	db.Delete(&quote)

	c.JSON(200, gin.H{"message": "quote deleted"})
}
