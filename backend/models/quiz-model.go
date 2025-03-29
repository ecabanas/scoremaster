package models

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	ID        uint       `json:"primaryKey"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

type Question struct {
	gorm.Model
	ID                  uint         `json:"primaryKey"`
	Type                questionType `json:"type"`
	QuestionText        string       `json:"questionText"`
	QuestionDescription string       `json:"questionDescription"`
	AnswerOptions       []Answer     `json:"answerOptions"`
}

type questionType struct {
	Name       string   `json:"name"`       // e.g., "single_choice"
	ValidRules []string `json:"validRules"` // e.g., ["required", "max_length:100"]
}

type Answer struct {
	gorm.Model
	ID     uint   `json:"primaryKey"`
	Text   string `json:"text"`
	Weight int    `json:"weight"`
}
