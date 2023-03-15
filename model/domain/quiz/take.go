package quiz

import (
	"os/user"
	"time"

	"github.com/google/uuid"
)

type QuizTake struct {
	Id            uuid.UUID `gorm:"primaryKey"`
	QuizId        uuid.UUID
	Email         string
	StartTime     time.Time
	IsFinished    bool
	Score         int
	Quiz          `gorm:"foreignKey:QuizId;references:Id"`
	user.User     `gorm:"foreignKey:Email;references:Email"`
	ChoiceAnswers []TakeChoiceAnswer `gorm:"foreignKey:QuizTakeId;references:Id"`
}

func (QuizTake) TableName() string {
	return "quiz_take"
}
