package controllers

import (
	"encoding/json"
	"net/http"
	"scoremaster/backend/models"
	"scoremaster/backend/services"
)

func GetQuestionsWithAnswers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	questions, err := services.GetQuestionsWithAnswers()
	if err != nil {
		http.Error(writer, "Failed to fetch questions", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(questions)
}

func CreateQuestion(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var question models.Question
	if err := json.NewDecoder(request.Body).Decode(&question); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := services.CreateQuestion(&question); err != nil {
		http.Error(writer, "Failed to create question", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(question)
}
