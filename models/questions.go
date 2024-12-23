package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Questions struct {
	ID                uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Question          string         `json:"question" gorm:"type:varchar;not null"`
	Difficulty        int            `json:"difficulty" gorm:"type:int;not null"`
	Proposed_response pq.StringArray `json:"proposed_response" gorm:"type:varchar[];not null"`
	Correct_answer    string         `json:"correct_answer" gorm:"type:varchar;not null"`
	CreatedAt         time.Time      `json:"-"`
	Theme_id          uuid.UUID      `json:"theme_id" gorm:"type:uuid;not null"`
	Themes            Themes         `json:"-" gorm:"foreignKey:Theme_id"`
}

func (question *Questions) BeforeCreate(tx *gorm.DB) (err error) {
	question.ID = uuid.New()
	return
}

type CreateQuestionInput struct {
	Question          string         `json:"question" gorm:"type:varchar;not null"`
	Difficulty        int            `json:"difficulty" gorm:"type:int;not null"`
	Proposed_response pq.StringArray `json:"proposed_response" gorm:"type:varchar[];not null"`
	Correct_answer    string         `json:"correct_answer" gorm:"type:varchar;not null"`
	Theme_id          uuid.UUID      `json:"theme_id" gorm:"type:uuid;not null"`
}

type GetRandomQuestions struct {
	ID                uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Question          string         `json:"question" gorm:"type:varchar;not null"`
	Difficulty        int            `json:"difficulty" gorm:"type:int;not null"`
	Proposed_response pq.StringArray `json:"proposed_response" gorm:"type:varchar[];not null"`
}

type AnswerInput struct {
	Answers []struct {
		Question_id uuid.UUID `json:"question_id"`
		Answer      string    `json:"answer"`
	} `json:"answers"`
}
