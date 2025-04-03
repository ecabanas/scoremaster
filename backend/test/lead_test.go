package test

import (
	"testing"
	"time"

	"scoremaster/backend/models"
)

func TestLeadCreation(t *testing.T) {
	db := setupTestDB()

	// Create a sample quiz
	quiz := models.Category{CategoryName: "Sample Quiz"}
	if err := db.Create(&quiz).Error; err != nil {
		t.Fatalf("Failed to create quiz: %v", err)
	}

	// Create a lead associated with the quiz
	lead := models.Participant{
		Email:          "test@example.com",
		SubmissionTime: time.Now(),
	}

	if err := db.Create(&lead).Error; err != nil {
		t.Fatalf("Failed to create lead: %v", err)
	}

	// Fetch the lead and validate
	var fetchedLead models.Participant
	if err := db.First(&fetchedLead, lead.ID).Error; err != nil {
		t.Fatalf("Failed to fetch lead: %v", err)
	}

	if fetchedLead.Email != "test@example.com" {
		t.Errorf("Expected email to be 'test@example.com', got '%s'", fetchedLead.Email)
	}
}
