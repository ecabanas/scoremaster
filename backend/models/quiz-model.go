package models

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	Title     string     `json:"title"`
	Questions []Question `json:"questions" gorm:"foreignKey:QuizID" `
}

type Question struct {
	gorm.Model
	QuizID              uint     `json:"quizID"`
	TypeID              uint     `json:"typeID"`
	QuestionText        string   `json:"questionText"`
	QuestionDescription string   `json:"questionDescription"`
	AnswerOptions       []Answer `json:"answerOptions" gorm:"foreignKey:QuestionID" `
}

type Answer struct {
	gorm.Model
	QuestionID uint   `json:"questionID"`
	Text       string `json:"text"`
	Weight     int    `json:"weight"`
}
