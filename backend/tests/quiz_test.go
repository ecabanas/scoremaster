package tests

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
		&models.Quiz{},
		&models.Question{},
		&models.Answer{},
		&models.Lead{},
		&models.LeadResponse{},
	)
	if err != nil {
		panic("failed to migrate the test database: " + err.Error())
	}

	return db
}

func CreateMockQuiz() models.Quiz {
	return models.Quiz{
		Title: "Sample Quiz",
		Questions: []models.Question{
			{
				QuestionText: "What is 2 + 2?",
				AnswerOptions: []models.Answer{
					{Text: "3", Weight: 0},
					{Text: "4", Weight: 1},
				},
			},
			{
				QuestionText: "What is the capital of France?",
				AnswerOptions: []models.Answer{
					{Text: "Berlin", Weight: 0},
					{Text: "Paris", Weight: 1},
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

	var fetchedQuiz models.Quiz
	if err := db.Preload("Questions.AnswerOptions").First(&fetchedQuiz, quiz.ID).Error; err != nil {
		t.Fatalf("Failed to fetch quiz: %v", err)
	}

	if fetchedQuiz.Title != "Sample Quiz" {
		t.Errorf("Expected quiz title to be 'Sample Quiz', got '%s'", fetchedQuiz.Title)
	}

	if len(fetchedQuiz.Questions) != 2 {
		t.Errorf("Expected 2 questions, got %d", len(fetchedQuiz.Questions))
	}

	if len(fetchedQuiz.Questions[0].AnswerOptions) != 2 {
		t.Errorf("Expected 2 answer options for the first question, got %d", len(fetchedQuiz.Questions[0].AnswerOptions))
	}
}
