package lesson

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/material"
)

type Lesson struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid"`
	Name        string         `json:"name"`
	CourseID    string         `json:"course_id"`
	Course      *course.Course `json:"-" gorm:"foreignKey:CourseID"`
	Order       int            `json:"order"`
	Description string         `json:"description"`
}

type LessonMaterials struct {
	ID         uuid.UUID          `json:"id" gorm:"primaryKey;type:uuid"`
	LessonID   uuid.UUID          `json:"lesson_id"`
	Lesson     *Lesson            `json:"-" gorm:"foreignKey:LessonID"`
	Order      int                `json:"order"`
	MaterialID uuid.UUID          `json:"material_id"`
	Material   *material.Material `json:"-" gorm:"foreignKey:MaterialID"`
	Contents   string             `json:"contents"`
}

func (Lesson) TableName() string {
	return "lesson"
}

func (LessonMaterials) TableName() string {
	return "lesson_materials"
}
