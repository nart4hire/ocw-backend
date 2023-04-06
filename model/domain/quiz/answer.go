package quiz

import "github.com/google/uuid"

type Answer struct {
	QuestionId uuid.UUID
	OptionId   uuid.UUID
}
