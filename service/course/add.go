package course

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	cadd "gitlab.informatika.org/ocw/ocw-backend/model/web/course/add"
	fadd "gitlab.informatika.org/ocw/ocw-backend/model/web/course/faculty/add"
	madd "gitlab.informatika.org/ocw/ocw-backend/model/web/course/major/add"
)

func (c CourseServiceImpl) AddCourse(payload cadd.AddCourseRequestPayload) error {
	exist, err := c.CourseRepository.IsCourseExist(payload.ID)

	if err != nil {
		// Some uncaught error
		return err
	}

	if exist {
		return web.NewResponseError("Course ID Already Exists", web.IDExists)
	}

	err = c.CourseRepository.AddCourse(course.Course{
		ID:           payload.ID,
		Name:         payload.Name,
		Major_id:     payload.MajorID,
		Description:  payload.Description,
		Email:        payload.Email,
		Abbreviation: payload.Abbreviation,
	})

	if err != nil {
		// Some uncaught error
		return err
	}

	return nil
}

func (c CourseServiceImpl) AddMajor(payload madd.AddMajorRequestPayload) error {
	err := c.CourseRepository.AddMajor(course.Major{
		Name:         payload.Name,
		Fac_id:       payload.FacultyID,
		Abbreviation: payload.Abbreviation,
	})

	if err != nil {
		// Some uncaught error
		return err
	}

	return nil
}

func (c CourseServiceImpl) AddFaculty(payload fadd.AddFacultyRequestPayload) error {
	err := c.CourseRepository.AddFaculty(course.Faculty{
		Name:         payload.Name,
		Abbreviation: payload.Abbreviation,
	})

	if err != nil {
		// Some uncaught error
		return err
	}

	return nil
}
