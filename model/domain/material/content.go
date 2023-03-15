package material

import "github.com/google/uuid"

type Content struct {
	Id         uuid.UUID `gorm:"primaryKey"`
	Type       MaterialType
	Link       string
	MaterialId uuid.UUID
}

func (Content) TableName() string {
	return "material_data"
}
