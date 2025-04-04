package services

import (
	"encoding/json"
	"net/http"

	"scoremaster/backend/database"
	"scoremaster/backend/models"
)

// func GetQuizzes(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	var list []models.Question
// 	json.NewEncoder(writer).Encode(list)
// }

// func GetQuestions(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	var list []models.Question
// 	json.NewEncoder(writer).Encode(list)
// }

// GetQuestionsWithAnswers fetches all questions with their answers and categories
func GetQuestionsWithAnswers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var questions []models.Question
	err := database.DBConn.Preload("Answers").Preload("Category").Find(&questions).Error
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(questions)
}

// GetCategories fetches all categories with their associated questions
func GetCategories(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var categories []models.Category
	err := database.DBConn.Preload("Questions.Answers").Find(&categories).Error
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(categories)
}
