package quiz

import "github.com/google/uuid"

type QuizProblem struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Statement string
	Type      ProblemType
	QuizId    uuid.UUID
	Options   []AnswerOption `gorm:"foreignKey:QuizProblemId;references:Id"`
}

func (QuizProblem) TableName() string {
	return "quiz_problem"
}
