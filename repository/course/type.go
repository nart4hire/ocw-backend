package course

import (
	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
)

type CourseRepository interface {
	AddCourse(course course.Course) error
	AddMajor(major course.Major) error
	AddFaculty(faculty course.Faculty) error
	GetCourse(id string) (*course.Course, error)
	GetMajor(id uuid.UUID) (*course.Major, error)
	GetFaculty(id uuid.UUID) (*course.Faculty, error)
	GetAllCourse() ([]course.Course, error)
	GetAllMajor() ([]course.Major, error)
	GetAllFaculty() ([]course.Faculty, error)
	GetAllCourseByMajor(id uuid.UUID) ([]course.Course, error)
	GetAllCourseByFaculty(id uuid.UUID) ([]course.Course, error)
	GetAllMajorByFaculty(id uuid.UUID) ([]course.Major, error)
	UpdateCourse(course course.Course) error
	UpdateMajor(major course.Major) error
	UpdateFaculty(faculty course.Faculty) error
	DeleteCourse(id string) error
	IsCourseExist(id string) (bool, error)
	IsMajorExist(id uuid.UUID) (bool, error)
	IsFacultyExist(id uuid.UUID) (bool, error)
	IsUserCourseContributor(id string, email string) (bool, error)

	// Internal Method Only

	GetMajorByAbbr(abbr string) (*course.Major, error)
	GetFacultyByAbbr(abbr string) (*course.Faculty, error)
}
