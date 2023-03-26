package course

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course/get"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course/delete"
	cadd "gitlab.informatika.org/ocw/ocw-backend/model/web/course/add"
	madd "gitlab.informatika.org/ocw/ocw-backend/model/web/course/major/add"
	fadd "gitlab.informatika.org/ocw/ocw-backend/model/web/course/faculty/add"
	cupdate "gitlab.informatika.org/ocw/ocw-backend/model/web/course/update"
	mupdate "gitlab.informatika.org/ocw/ocw-backend/model/web/course/major/update"
	fupdate "gitlab.informatika.org/ocw/ocw-backend/model/web/course/faculty/update"

)

type CourseService interface {
	AddCourse(payload cadd.AddCourseRequestPayload) error
	AddMajor(payload madd.AddMajorRequestPayload) error
	AddFaculty(payload fadd.AddFacultyRequestPayload) error
	GetCourse(payload get.GetByStringRequestPayload) (*course.Course, error)
	GetMajor(payload get.GetByUUIDRequestPayload) (*course.Major, error)
	GetFaculty(payload get.GetByUUIDRequestPayload) (*course.Faculty, error)
	GetAllCourse() ([]course.Course, error)
	GetAllMajor() ([]course.Major, error)
	GetAllFaculty() ([]course.Faculty, error)
	GetAllCourseByMajor(payload get.GetByUUIDRequestPayload) ([]course.Course, error)
	GetAllCourseByFaculty(payload get.GetByUUIDRequestPayload) ([]course.Course, error)
	GetAllMajorByFaculty(payload get.GetByUUIDRequestPayload) ([]course.Major, error)
	UpdateCourse(payload cupdate.UpdateCourseRequestPayload) error
	UpdateMajor(payload mupdate.UpdateMajorRequestPayload) error
	UpdateFaculty(payload fupdate.UpdateFacultyRequestPayload) error
	DeleteCourse(payload delete.DeleteByStringRequestPayload) error
}
