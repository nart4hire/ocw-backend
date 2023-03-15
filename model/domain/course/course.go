package course

import "github.com/google/uuid"

type Faculty struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	Name         string
	Abbreviation string
}

type Major struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name         string
	Fac_id       uuid.UUID `gorm:"type:uuid"`
	Faculty      Faculty   `gorm:"foreignKey:Fac_id"`
	Abbreviation string
}

type Course struct {
	ID           string `gorm:"primaryKey"`
	Name         string
	Major_id     uuid.UUID `gorm:"type:uuid"`
	Major        Major     `gorm:"foreignKey:Major_id"`
	Description  string
	Email        string
	Abbreviation string
	Lecturer     string
}

func (Faculty) TableName() string {
	return "faculty"
}

func (Major) TableName() string {
	return "major"
}

func (Course) TableName() string {
	return "course"
}
