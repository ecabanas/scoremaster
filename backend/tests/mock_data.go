package tests

import (
	"scoremaster/backend/models"
)

func CreateMockData() models.Quiz {
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

func CreateMockLead(quizID uint) models.Lead {
	return models.Lead{
		Email:      "test@example.com",
		QuizID:     quizID,
		FinalScore: 15,
		Responses: []models.LeadResponse{
			{QuestionID: 1, AnswerID: 2, Weight: 5},
			{QuestionID: 2, AnswerID: 4, Weight: 10},
		},
	}
}
