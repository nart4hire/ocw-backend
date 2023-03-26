package quiz

import "github.com/google/uuid"

type AnswerOption struct {
	Id            uuid.UUID `gorm:"primaryKey"`
	QuizProblemId uuid.UUID `gorm:"primaryKey"`
	Statement     string
	IsAnswer      bool
}

func (AnswerOption) TableName() string {
	return "quiz_choice_answer"
}
