package course

import "github.com/google/uuid"

type Major struct {
	Id   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string
}

func (Major) TableName() string {
	return "major"
}
