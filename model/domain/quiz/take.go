package quiz

import (
	"time"

	"github.com/google/uuid"
)

type QuizTake struct {
	Id         uuid.UUID `gorm:"primaryKey" json:"id"`
	QuizId     uuid.UUID `json:"quiz_id"`
	Email      string    `json:"email"`
	StartTime  time.Time `json:"start"`
	IsFinished bool      `json:"finished"`
	Score      int       `json:"score"`
}

func (QuizTake) TableName() string {
	return "quiz_take"
}
