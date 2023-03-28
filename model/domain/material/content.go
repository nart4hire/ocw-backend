package material

import "github.com/google/uuid"

type Content struct {
	Id         uuid.UUID    `json:"id" gorm:"primaryKey"`
	Type       MaterialType `json:"type"`
	Link       string       `json:"link"`
	MaterialId uuid.UUID    `json:"material_id"`
}

func (Content) TableName() string {
	return "material_data"
}
