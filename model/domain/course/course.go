package course

import "github.com/google/uuid"

// TODO: Abbreviations should be unique constrainted as identifiers
type Faculty struct {
	ID           uuid.UUID `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
}

type Major struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	Name         string    `json:"name"`
	Fac_id       uuid.UUID `gorm:"type:uuid" json:"faculty_id"`
	Faculty      Faculty   `gorm:"foreignKey:Fac_id" json:"faculty"`
	Abbreviation string    `json:"abbreviation"`
}

type Course struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	Major_id     uuid.UUID `gorm:"type:uuid" json:"major_id"`
	Major        Major     `gorm:"foreignKey:Major_id" json:"major"`
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
