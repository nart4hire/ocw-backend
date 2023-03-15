package course

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
)

type Course struct {
	Id           string `gorm:"primaryKey"`
	Name         string
	MajorId      uuid.UUID
	Description  string
	Major        Major       `gorm:"foreignKey:MajorId;references:Id"`
	Contributors []user.User `gorm:"many2many:course_contributor;foreignKey:Id;joinForeignKey:CourseId;references:Email;joinReferences:Email"`
}

func (Course) TableName() string {
	return "course"
}
