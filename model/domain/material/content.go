package material

import "github.com/google/uuid"

type Content struct {
	Id         uuid.UUID    `gorm:"primaryKey" json:"id"`
	Type       MaterialType `json:"type"`
	Link       string       `json:"link"`
	MaterialId uuid.UUID    `json:"material_id"`
	Material   `gorm:"foreignKey:MaterialId;references:Id" json:"material"`
}

func (Content) TableName() string {
	return "material_data"
}
