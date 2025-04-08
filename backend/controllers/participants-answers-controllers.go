package controllers

import (
	"encoding/json"
	"net/http"

	"scoremaster/backend/models"
	"scoremaster/backend/services"
)

func CreateParticipantAnswer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var participantAnswer models.ParticipantAnswer
	if err := json.NewDecoder(request.Body).Decode(&participantAnswer); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := services.CreateParticipantAnswer(&participantAnswer); err != nil {
		http.Error(writer, "Failed to create participant answer", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(participantAnswer)
}
