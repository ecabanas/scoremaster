package models

import (
	"time"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	Email          string    `json:"email" gorm:"unique"`
	SubmissionTime time.Time `json:"submission_time"`
	Metadata       string    `json:"metadata" gorm:"type:json"`
}

type ParticipantAnswer struct {
	gorm.Model
	ParticipantID    uint      `json:"participant_id" gorm:"foreignKey:ParticipantID"`
	QuestionID       uint      `json:"question_id" gorm:"foreignKey:QuestionID"`
	SelectedAnswerID uint      `json:"selected_answer_id"`
	Score            int       `json:"score"`
	AnswerTimestamp  time.Time `json:"answer_timestamp"`
}
