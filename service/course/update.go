package course

import (
	"errors"

	"gitlab.informatika.org/ocw/ocw-backend/model/domain/course"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	fupdate "gitlab.informatika.org/ocw/ocw-backend/model/web/course/faculty/update"
	mupdate "gitlab.informatika.org/ocw/ocw-backend/model/web/course/major/update"
	cupdate "gitlab.informatika.org/ocw/ocw-backend/model/web/course/update"
	"gorm.io/gorm"
)

// TODO: Authorization Checks

func (c CourseServiceImpl) UpdateCourse(payload cupdate.UpdateCourseRequestPayload) error {
	// Validate Role
	claim, err := c.TokenUtil.Validate(payload.UpdateCourseToken, token.Access)

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

	err = c.CourseRepository.UpdateCourse(course.Course{
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
	
	// Validate Role
	claim, err := c.TokenUtil.Validate(payload.UpdateMajorToken, token.Access)

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
	
	err = c.CourseRepository.UpdateMajor(course.Major{
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
	// Validate Role
	claim, err := c.TokenUtil.Validate(payload.UpdateFacultyToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role != user.Admin {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}
	
	err = c.CourseRepository.UpdateFaculty(course.Faculty{
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
