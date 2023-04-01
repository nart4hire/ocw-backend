package course

import (
	domCourse "gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course/major"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course/faculty"

)

type CourseService interface {
	AddCourse(payload course.AddCourseRequestPayload) error
	AddMajor(payload major.AddMajorRequestPayload) error
	AddFaculty(payload faculty.AddFacultyRequestPayload) error
	GetCourse(payload course.GetByStringRequestPayload) (*domCourse.Course, error)
	GetMajor(payload course.GetByUUIDRequestPayload) (*domCourse.Major, error)
	GetFaculty(payload course.GetByUUIDRequestPayload) (*domCourse.Faculty, error)
	GetAllCourse() ([]domCourse.Course, error)
	GetAllMajor() ([]domCourse.Major, error)
	GetAllFaculty() ([]domCourse.Faculty, error)
	GetAllCourseByMajor(payload course.GetByUUIDRequestPayload) ([]domCourse.Course, error)
	GetAllCourseByFaculty(payload course.GetByUUIDRequestPayload) ([]domCourse.Course, error)
	GetAllMajorByFaculty(payload course.GetByUUIDRequestPayload) ([]domCourse.Major, error)
	UpdateCourse(payload course.UpdateCourseRequestPayload) error
	UpdateMajor(payload major.UpdateMajorRequestPayload) error
	UpdateFaculty(payload faculty.UpdateFacultyRequestPayload) error
	DeleteCourse(payload course.DeleteByStringRequestPayload) error
}
