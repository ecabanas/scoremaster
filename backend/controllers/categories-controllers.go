package controllers

import (
	"encoding/json"
	"net/http"
	"scoremaster/backend/models"
	"scoremaster/backend/services"
)

func GetCategories(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	categories, err := services.GetCategories()
	if err != nil {
		http.Error(writer, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(categories)
}

func CreateCategory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var category models.Category
	if err := json.NewDecoder(request.Body).Decode(&category); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := services.CreateCategory(&category); err != nil {
		http.Error(writer, "Failed to create category", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(category)
}
