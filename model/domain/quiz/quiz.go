package quiz

import (
	"github.com/google/uuid"
)

type Quiz struct {
	Id           uuid.UUID `gorm:"primaryKey" json:"id"`
	Name         string    `json:"nama"`
	CourseId     string    `json:"course_id"`
	CreatorEmail string    `json:"creator_email"`
	QuizPath     string    `json:"-"`
}

func (Quiz) TableName() string {
	return "quiz"
}
