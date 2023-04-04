package course

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course"
)

func (c CourseServiceImpl) DeleteCourse(payload course.DeleteByStringRequestPayload) error {
	// Validate Role
	claim, err := c.TokenUtil.Validate(payload.DeleteCourseToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == user.Student {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	err = c.CourseRepository.DeleteCourse(payload.ID)

	if err != nil {
		// Uncaught error
		return err
	}

	return nil
}