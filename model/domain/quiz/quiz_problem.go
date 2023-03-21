package quiz

import "github.com/google/uuid"

type QuizProblem struct {
	Id        uuid.UUID      `gorm:"primaryKey" json:"id"`
	Statement string         `json:"statement"`
	Type      ProblemType    `json:"type"`
	QuizId    uuid.UUID      `json:"quiz_id"`
	Options   []AnswerOption `gorm:"foreignKey:QuizProblemId;references:Id" json:"options"`
}

func (QuizProblem) TableName() string {
	return "quiz_problem"
}
