package quiz

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type Quiz struct {
	Id           uuid.UUID `gorm:"primaryKey"`
	Name         string
	CourseId     string
	CreatorEmail string
	Creator      user.User     `gorm:"foreignKey:CreatorEmail;references:Email"`
	Course       course.Course `gorm:"foreignKey:CourseId;references:Id"`
	Problems     []QuizProblem `gorm:"foreignKey:QuizId;references:Id"`
}

func (Quiz) TableName() string {
	return "quiz"
}
