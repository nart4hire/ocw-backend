package course

import "github.com/google/uuid"

type Faculty struct {
	id           uuid.UUID `gorm:"primaryKey"`
	name         string
	abbreviation string
}

type Major struct {
	id           uuid.UUID `gorm:"primaryKey;type:uuid"`
	name         string
	fac_id       uuid.UUID `gorm:"type:uuid"`
	abbreviation string
}

type Course struct {
	id           string `gorm:"primaryKey"`
	name         string
	major_id     uuid.UUID `gorm:"type:uuid"`
	description  string
	email        string
	abbreviation string
	lecturer     string
}
