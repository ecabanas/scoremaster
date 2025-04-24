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

// CheckEmailExists checks if an email is already registered
func CheckEmailExists(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")

    var requestBody struct {
        Email string `json:"email"`
    }
    
    if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
        http.Error(writer, "Invalid request payload", http.StatusBadRequest)
        return
    }

    exists, err := services.CheckEmailExists(requestBody.Email)
    if err != nil {
        http.Error(writer, "Error checking email", http.StatusInternalServerError)
        return
    }
    
    writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode(map[string]bool{"exists": exists})
}
