package course

import (
	"errors"

	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	cadd "gitlab.informatika.org/ocw/ocw-backend/model/web/course/add"
	fadd "gitlab.informatika.org/ocw/ocw-backend/model/web/course/faculty/add"
	madd "gitlab.informatika.org/ocw/ocw-backend/model/web/course/major/add"
	"gorm.io/gorm"
)

func (c CourseServiceImpl) AddCourse(payload cadd.AddCourseRequestPayload) error {
	// Validate Role
	claim, err := c.TokenUtil.Validate(payload.AddCourseToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == user.Student {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	// If payload uses major abbreviation, set id
	if payload.MajAbbr != "" {
		major, err := c.CourseRepository.GetMajorByAbbr(payload.MajAbbr)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return web.NewResponseErrorFromError(err, web.CourseNotExist)
			}
			// Some Uncaught error
			return err
		}

		payload.MajorID = major.ID
	}

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
	// Validate Role
	claim, err := c.TokenUtil.Validate(payload.AddMajorToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role != user.Admin {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	// If payload uses faculty abbreviation, set id
	if payload.FacAbbr != "" {
		faculty, err := c.CourseRepository.GetFacultyByAbbr(payload.FacAbbr)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return web.NewResponseErrorFromError(err, web.CourseNotExist)
			}
			// Some Uncaught error
			return err
		}

		payload.FacultyID = faculty.ID
	}

	err = c.CourseRepository.AddMajor(course.Major{
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
	// Validate Role
	claim, err := c.TokenUtil.Validate(payload.AddFacultyToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role != user.Admin {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}
	
	err = c.CourseRepository.AddFaculty(course.Faculty{
		Name:         payload.Name,
		Abbreviation: payload.Abbreviation,
	})

	if err != nil {
		// Some uncaught error
		return err
	}

	return nil
}
