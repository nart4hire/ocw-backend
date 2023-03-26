package material

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type Material struct {
	Id           uuid.UUID `gorm:"primaryKey"`
	CourseId     string
	CreatorEmail string
	Creator      user.User     `gorm:"foreignKey:CreatorEmail;references:Email"`
	Course       course.Course `gorm:"foreignKey:CourseId;references:Id"`
	Contents     []Content     `gorm:"foreignKey:MaterialId;references:Id"`
}

func (Material) TableName() string {
	return "material"
}
