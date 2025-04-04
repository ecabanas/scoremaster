package services

import (
	"scoremaster/backend/database"
	"scoremaster/backend/models"
)

// Fetch all categories with their associated questions and answers
func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := database.DBConn.Preload("Questions.Answers").Find(&categories).Error
	return categories, err
}

// Create a new category
func CreateCategory(category *models.Category) error {
	return database.DBConn.Create(category).Error
}
