package course

import "github.com/google/uuid"

// TODO: Abbreviations should be unique constrainted as identifiers
type Faculty struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
}

type Major struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name         string    `json:"name"`
	Fac_id       uuid.UUID `json:"fac_id" gorm:"type:uuid"`
	Faculty      *Faculty  `json:"-" gorm:"foreignKey:Fac_id"`
	Abbreviation string    `json:"abbreviation"`
}

type Course struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Major_id     uuid.UUID `json:"major_id" gorm:"type:uuid"`
	Major        *Major    `json:"-" gorm:"foreignKey:Major_id"`
	Description  string    `json:"description"`
	Email        string    `json:"email"`
	Abbreviation string    `json:"abbreviation"`
	Lecturer     string    `json:"lecturer"`
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
