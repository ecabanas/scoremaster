package services

import (
	"encoding/json"
	"net/http"

	"scoremaster/backend/database"
	"scoremaster/backend/models"
)

// GetQuestionsWithAnswers fetches all questions with their answers and categories
func GetQuestionsWithAnswers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var questions []models.Question
	err := database.DBConn.
		Preload("Category").
		Preload("Answers").
		Find(&questions).Error
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

// CreateCategory creates a new category in the database
func CreateCategory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var category models.Category
	if err := json.NewDecoder(request.Body).Decode(&category); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.DBConn.Create(&category).Error; err != nil {
		http.Error(writer, "Failed to create category", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(category)
}

// CreateQuestion creates a new question in the database
func CreateQuestion(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var question models.Question
	if err := json.NewDecoder(request.Body).Decode(&question); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.DBConn.Create(&question).Error; err != nil {
		http.Error(writer, "Failed to create question", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(question)
}

// GetAllCategories fetches all categories from the database
func GetAllCategories(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var categories []models.Category
	if err := database.DBConn.Find(&categories).Error; err != nil {
		http.Error(writer, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(categories)
}

// GetAllQuestions fetches all questions from the database
func GetAllQuestions(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var questions []models.Question
	if err := database.DBConn.Preload("Answers").Preload("Category").Find(&questions).Error; err != nil {
		http.Error(writer, "Failed to fetch questions", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(questions)
}
