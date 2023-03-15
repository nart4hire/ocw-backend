package course

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
)

type UserRepository interface {
	AddCourse(course course.Course) error
	GetCourse(id string) (*course.Course, error)
	GetAllCourse() ([]course.Course, error)
	GetAllCourseByMajor(abbr string) ([]course.Course, error)
	GetAllCourseMyFaculty(abbr string) ([]course.Course, error)
	UpdateCourse(course course.Course) error
	DeleteCourse(id string) error
	IsCourseExist(id string) (bool, error)
}
