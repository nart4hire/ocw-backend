package quiz

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type Quiz struct {
	Id           uuid.UUID     `gorm:"primaryKey" json:"id"`
	Name         string        `json:"name"`
	CourseId     string        `json:"course_id"`
	CreatorEmail string        `json:"creator_email"`
	Creator      user.User     `gorm:"foreignKey:CreatorEmail;references:Email" json:"creator"`
	Course       course.Course `gorm:"foreignKey:CourseId;references:Id" json:"course"`
	Problems     []QuizProblem `gorm:"foreignKey:QuizId;references:Id" json:"problems"`
}

func (Quiz) TableName() string {
	return "quiz"
}
