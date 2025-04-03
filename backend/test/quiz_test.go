package test

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"scoremaster/backend/models"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the test database")
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&models.Question{},
		&models.Answer{},
		&models.Participant{},
		&models.ParticipantAnswer{},
		&models.Category{},
	)
	if err != nil {
		panic("failed to migrate the test database: " + err.Error())
	}

	return db
}

func CreateMockQuiz() models.Category {
	return models.Category{
		CategoryName: "Sample Quiz",
		Questions: []models.Question{
			{
				QuestionText: "What is 2 + 2?",
				Answers: []models.Answer{
					{QuestionID: 1, AnswerText: "3", Score: 0},
					{QuestionID: 1, AnswerText: "3", Score: 0},
				},
			},
			{
				QuestionText: "What is the capital of France?",
				Answers: []models.Answer{
					{QuestionID: 2, AnswerText: "Berlin", Score: 0},
					{QuestionID: 2, AnswerText: "Paris", Score: 1},
				},
			},
		},
	}
}

func TestQuizCreation(t *testing.T) {
	db := setupTestDB()

	quiz := CreateMockQuiz()

	if err := db.Create(&quiz).Error; err != nil {
		t.Fatalf("Failed to create quiz: %v", err)
	}

	var fetchedQuiz models.Question
	if err := db.Preload("Questions.AnswerOptions").First(&fetchedQuiz, quiz.ID).Error; err != nil {
		t.Fatalf("Failed to fetch quiz: %v", err)
	}

	if fetchedQuiz.QuestionText != "Sample Quiz" {
		t.Errorf("Expected quiz title to be 'Sample Quiz', got '%s'", fetchedQuiz.QuestionText)
	}

	if len(fetchedQuiz.QuestionText) != 2 {
		t.Errorf("Expected 2 questions, got %d", len(fetchedQuiz.QuestionText))
	}

	if len(fetchedQuiz.Answers[0].AnswerText) != 2 {
		t.Errorf("Expected 2 answer options for the first question, got %d", len(fetchedQuiz.Answers[0].AnswerText))
	}
}
