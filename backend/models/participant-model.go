package models

import (
	"time"

	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Participant struct {
	gorm.Model
	Name               string              `json:"name" gorm:"not null" validate:"required"`
	Email              string              `json:"email" gorm:"unique" validate:"required,email"`
	ParticipantAnswers []ParticipantAnswer `json:"participant_answers" gorm:"foreignKey:ParticipantID"`
}

type ParticipantAnswer struct {
	gorm.Model
	ParticipantID    uint      `json:"participant_id" gorm:"foreignKey:ParticipantID"`
	QuestionID       uint      `json:"question_id" gorm:"foreignKey:QuestionID"`
	SelectedAnswerID uint      `json:"selected_answer_id"`
	Score            int       `json:"score"`
	AnswerTimestamp  time.Time `json:"answer_timestamp"`
}

func (p *Participant) Validate() error {
	return validate.Struct(p)
}
func (p *Participant) BeforeCreate(tx *gorm.DB) error {
	return p.Validate()
}

func (p *Participant) BeforeUpdate(tx *gorm.DB) error {
	return p.Validate()
}
