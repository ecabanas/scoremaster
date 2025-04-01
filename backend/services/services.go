package services

import (
	"encoding/json"
	"net/http"
	"scoremaster/backend/models"
)

func GetQuizzes(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var list []models.Quiz
	json.NewEncoder(writer).Encode(list)
}

func GetQuestions(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var list []models.Question
	json.NewEncoder(writer).Encode(list)
}
