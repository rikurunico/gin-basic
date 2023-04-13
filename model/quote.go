package model

import (
	"gorm.io/gorm"

)

type Quote struct {
	gorm.Model
	Text   string
	Author string
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Quote{})
}

func SeedDatabase(db *gorm.DB) error {
	quotes := []Quote{
		{Text: "Keep calm and code on", Author: "Unknown"},
		{Text: "Talk is cheap, show me the code", Author: "Linus Torvalds"},
		{Text: "Premature optimization is the root of all evil", Author: "Donald Knuth"},
		{Text: "Any fool can write code that a computer can understand. Good programmers write code that humans can understand.", Author: "Martin Fowler"},
		{Text: "Programs must be written for people to read, and only incidentally for machines to execute.", Author: "Harold Abelson"},
	}

	for _, quote := range quotes {
		if err := db.Create(&quote).Error; err != nil {
			return err
		}
	}
	return nil
}
