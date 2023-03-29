package material

import "github.com/google/uuid"

type Content struct {
	Id         uuid.UUID    `gorm:"primaryKey" json:"id"`
	Type       MaterialType `json:"type"`
	Link       string       `json:"link"`
	MaterialID uuid.UUID    `json:"material_id"`
}

func (Content) TableName() string {
	return "material_data"
}
