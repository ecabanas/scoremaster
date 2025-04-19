package controllers

import (
	"encoding/json"
	"net/http"
	"scoremaster/backend/models"
	"scoremaster/backend/services"
	"strconv"

	"github.com/gorilla/mux"
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

func GetQuestionByID(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")

    // Extract the ID from the URL parameters
    params := mux.Vars(request)
    id, err := strconv.ParseUint(params["id"], 10, 32)
    if err != nil {
        http.Error(writer, "Invalid question ID", http.StatusBadRequest)
        return
    }

    question, err := services.GetQuestionByID(uint(id))
    if err != nil {
        http.Error(writer, "Question not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(writer).Encode(question)
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
