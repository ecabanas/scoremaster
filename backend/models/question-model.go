package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string     `json:"category_name"`
	CategoryDesc string     `json:"category_desc"`
	Questions    []Question `json:"questions" gorm:"foreignKey:CategoryID"`
}

type Question struct {
	gorm.Model
	CategoryID   uint     `json:"category_id" gorm:"foreignKey:CategoryID"`
	Category     Category `json:"category" gorm:"foreignKey:CategoryID"`
	QuestionText string   `json:"question_text"`
	QuestionDesc string   `json:"question_desc"`
	Order        int      `json:"order"`
	Answers      []Answer `json:"answers" gorm:"foreignKey:QuestionID"`
}

type Answer struct {
	gorm.Model
	QuestionID uint   `json:"question_id" gorm:"foreignKey:QuestionID"`
	AnswerText string `json:"answer_text"`
	Score      int    `json:"score"`
	ImageURL   string `json:"image_url"`
}
