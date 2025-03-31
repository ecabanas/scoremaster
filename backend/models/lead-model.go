package models

import (
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Email           string         `json:"email"`
	QuizID          uint           `json:"quizID"`
	Quiz            Quiz           `json:"quiz" gorm:"foreignKey:QuizID"`
	Responses       []LeadResponse `json:"responses" gorm:"foreignKey:LeadID"`
	FinalScore      int            `json:"finalScore"`
	PersonalizedMsg string         `json:"personalizedMsg"`
}

type LeadResponse struct {
	gorm.Model
	LeadID     uint `json:"leadID"`
	QuestionID uint `json:"questionID"`
	AnswerID   uint `json:"answerID"`
	Weight     int  `json:"weight"`
}
