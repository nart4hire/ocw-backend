package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type Material struct {
	Id           uuid.UUID      `json:"id" gorm:"primaryKey"`
	CourseId     string         `json:"course_id"`
	CreatorEmail string         `json:"creator_email"`
	Creator      *user.User     `json:"creator" gorm:"foreignKey:Email;references:Creator"`
	Course       *course.Course `json:"course" gorm:"foreignKey:Id;references:CourseId"`
	Contents     []Content      `json:"contents" gorm:"foreignKey:Id;references:MaterialId"`
}

func (Material) TableName() string {
	return "material"
}
