package services

import (
	"scoremaster/backend/database"
	"scoremaster/backend/models"
)

// Fetch all questions with their answers and categories
func GetQuestionsWithAnswers() ([]models.Question, error) {
	var questions []models.Question
	err := database.DBConn.
		Preload("Category").
		Preload("Answers").
		Find(&questions).Error
	return questions, err
}

func GetQuestionByID(id uint) (models.Question, error) {
	var question models.Question
	err := database.DBConn.
		Preload("Category").
		Preload("Answers").
		First(&question, id).Error
	return question, err
}

// Create a new question
func CreateQuestion(question *models.Question) error {
	return database.DBConn.Create(question).Error
}
