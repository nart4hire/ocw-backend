package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type Material struct {
	Id           uuid.UUID     `gorm:"primaryKey" json:"id"`
	CourseId     string        `json:"course_id"`
	CreatorEmail string        `json:"creator_email"`
	Creator      user.User     `gorm:"foreignKey:CreatorEmail;references:Email" json:"creator"`
	Course       course.Course `gorm:"foreignKey:CourseId;references:Id" json:"course"`
	Contents     []Content     `gorm:"foreignKey:MaterialId;references:Id" json:"contents"`
}

func (Material) TableName() string {
	return "material"
}
