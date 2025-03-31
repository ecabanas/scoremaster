package test

import (
	"scoremaster/backend/testdata"
	"testing"
)

func TestCreateMockData(t *testing.T) {
	mockData := testdata.CreateMockData()

	if mockData.Title != "Sample Quiz" {
		t.Errorf("expected Title to be 'Sample Quiz', got '%s'", mockData.Title)
	}

	if len(mockData.Questions) != 2 {
		t.Errorf("expected 2 questions, got %d", len(mockData.Questions))
	}

	if mockData.Questions[0].QuestionText != "What is 2 + 2?" {
		t.Errorf("expected first question to be 'What is 2 + 2?', got '%s'", mockData.Questions[0].QuestionText)
	}
}

func TestCreateMockLead(t *testing.T) {
	mockLead := testdata.CreateMockLead(1)

	if mockLead.Email != "test@example.com" {
		t.Errorf("expected Email to be 'test@example.com', got '%s'", mockLead.Email)
	}

	if mockLead.QuizID != 1 {
		t.Errorf("expected QuizID to be 1, got %d", mockLead.QuizID)
	}

	if len(mockLead.Responses) != 2 {
		t.Errorf("expected 2 responses, got %d", len(mockLead.Responses))
	}

	if mockLead.Responses[0].Weight != 5 {
		t.Errorf("expected first response weight to be 5, got %d", mockLead.Responses[0].Weight)
	}
}
