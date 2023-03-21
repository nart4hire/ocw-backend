package quiz

import (
	"os/user"
	"time"

	"github.com/google/uuid"
)

type QuizTake struct {
	Id            uuid.UUID `gorm:"primaryKey" json:"id"`
	QuizId        uuid.UUID `json:"quiz_id"`
	Email         string    `json:"email"`
	StartTime     time.Time `json:"start"`
	IsFinished    bool      `json:"finished"`
	Score         int       `json:"score"`
	Quiz          `gorm:"foreignKey:QuizId;references:Id" json:"quiz"`
	user.User     `gorm:"foreignKey:Email;references:Email" json:"user"`
	ChoiceAnswers []TakeChoiceAnswer `gorm:"foreignKey:QuizTakeId;references:Id" json:"-"`
}

func (QuizTake) TableName() string {
	return "quiz_take"
}
