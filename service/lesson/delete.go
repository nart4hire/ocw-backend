package lesson

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson/materials"
)

func (l LessonServiceImpl) DeleteLesson(payload lesson.DeleteByUUIDRequestPayload) error {
	// Validate Role
	claim, err := l.TokenUtil.Validate(payload.DeleteLessonToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == user.Student {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	err = l.LessonRepository.DeleteLesson(payload.ID)

	if err != nil {
		// Uncaught error
		return err
	}

	return nil
}

func (l LessonServiceImpl) DeleteLessonMaterial(payload materials.DeleteByUUIDRequestPayload) error {
	// Validate Role
	claim, err := l.TokenUtil.Validate(payload.DeleteLessonMaterialsToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == user.Student {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	err = l.LessonMaterialsRepository.DeleteLessonMaterial(payload.ID)

	if err != nil {
		// Uncaught error
		return err
	}

	return nil
}