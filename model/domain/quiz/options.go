package quiz

import "github.com/google/uuid"

type AnswerOption struct {
	Id            uuid.UUID `gorm:"primaryKey" json:"id"`
	QuizProblemId uuid.UUID `gorm:"primaryKey" json:"problem_id"`
	Statement     string    `json:"statement"`
	IsAnswer      bool      `json:"isAnswer"`
}

func (AnswerOption) TableName() string {
	return "quiz_choice_answer"
}
