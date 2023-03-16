package course

import (
	"errors"

	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	fupdate "gitlab.informatika.org/ocw/ocw-backend/model/web/course/faculty/update"
	mupdate "gitlab.informatika.org/ocw/ocw-backend/model/web/course/major/update"
	cupdate "gitlab.informatika.org/ocw/ocw-backend/model/web/course/update"
	"gorm.io/gorm"
)

func (c CourseServiceImpl) UpdateCourse(payload cupdate.UpdateCourseRequestPayload) error {
	err := c.CourseRepository.UpdateCourse(course.Course{
		ID: payload.ID,
		Name: payload.Name,
		Major_id: payload.MajorID,
		Description: payload.Description,
		Email: payload.Email,
		Abbreviation: payload.Abbreviation,
		Lecturer: payload.Lecturer,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.NewResponseErrorFromError(err, web.CourseNotExist)
		}
		// Uncaught error
		return err
	}

	return nil
}

func (c CourseServiceImpl) UpdateMajor(payload mupdate.UpdateMajorRequestPayload) error {
	err := c.CourseRepository.UpdateMajor(course.Major{
		ID: payload.ID,
		Name: payload.Name,
		Fac_id: payload.FacultyID,
		Abbreviation: payload.Abbreviation,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.NewResponseErrorFromError(err, web.MajorNotExist)
		}
		// Uncaught error
		return err
	}

	return nil
}

func (c CourseServiceImpl) UpdateFaculty(payload fupdate.UpdateFacultyRequestPayload) error {
	err := c.CourseRepository.UpdateFaculty(course.Faculty{
		ID: payload.ID,
		Name: payload.Name,
		Abbreviation: payload.Abbreviation,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.NewResponseErrorFromError(err, web.FacultyNotExist)
		}
		// Uncaught error
		return err
	}

	return nil
}
