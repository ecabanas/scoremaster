package controllers

import (
	"encoding/json"
	"net/http"

	"scoremaster/backend/models"
	"scoremaster/backend/services"
)

func CreateParticipant(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var participant models.Participant
	if err := json.NewDecoder(request.Body).Decode(&participant); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := services.CreateParticipant(&participant); err != nil {
		http.Error(writer, "Failed to create participant", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(participant)
}
