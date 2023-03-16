package quiz

import "github.com/google/uuid"

type TakeChoiceAnswer struct {
	QuizTakeId    uuid.UUID `gorm:"primaryKey"`
	AnswerChoice  uuid.UUID
	QuizProblemId uuid.UUID `gorm:"primaryKey"`
	AnswerOption  `gorm:"foreignKey:AnswerChoice,QuizProblemId;references:Id,QuizProblemId"`
}

func (TakeChoiceAnswer) TableName() string {
	return "quiz_take_choice_answer"
}
