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
	Creator      *user.User     `json:"creator" gorm:"foreignKey:CreatorEmail;references:Email"`
	Course       *course.Course `json:"course" gorm:"foreignKey:CourseId;references:Id"`
	Contents     []Content      `json:"contents" gorm:"foreignKey:MaterialId;references:Id"`
}

func (Material) TableName() string {
	return "material"
}
