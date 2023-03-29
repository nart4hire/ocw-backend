package material

import (
	"github.com/google/uuid"
)

type Material struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey"`
	CourseId     string    `json:"course_id"`
	CreatorEmail string    `json:"creator_email"`
	Name         string    `json:"name"`
	Contents     []Content `json:"contents"`
}

func (Material) TableName() string {
	return "material"
}
