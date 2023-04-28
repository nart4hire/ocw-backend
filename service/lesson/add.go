package lesson

import (
	domLesson "gitlab.informatika.org/ocw/ocw-backend/model/domain/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson/materials"
)

func (l LessonServiceImpl) AddLesson(payload lesson.AddLessonRequestPayload) error {
	// Validate Role
	claim, err := l.TokenUtil.Validate(payload.AddLessonToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == user.Student {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	err = l.LessonRepository.AddLesson(domLesson.Lesson{
		Name:        payload.Name,
		CourseID:    payload.CourseID,
		Order:       payload.Order,
		Description: payload.Description,
	})

	if err != nil {
		// Some uncaught error
		return err
	}

	return nil
}

func (l LessonServiceImpl) AddLessonMaterial(payload materials.AddLessonMaterialsRequestPayload) error {
	// Validate Role
	claim, err := l.TokenUtil.Validate(payload.AddLessonMaterialsToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == user.Student {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	err = l.LessonMaterialsRepository.AddLessonMaterial(domLesson.LessonMaterials{
		LessonID:   payload.LessonID,
		Order:      payload.Order,
		MaterialID: payload.MaterialID,
		Contents:   payload.Contents,
	})

	if err != nil {
		// Some uncaught error
		return err
	}

	return nil
}
