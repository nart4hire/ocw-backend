package lesson

import (
	"errors"

	domLesson "gitlab.informatika.org/ocw/ocw-backend/model/domain/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/token"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson/materials"
	"gorm.io/gorm"
)

func (l LessonServiceImpl) UpdateLesson(payload lesson.UpdateLessonRequestPayload) error {
	// Validate Role
	claim, err := l.TokenUtil.Validate(payload.UpdateLessonToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == user.Student {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	err = l.LessonRepository.UpdateLesson(domLesson.Lesson{
		ID:          payload.ID,
		Name:        payload.Name,
		CourseID:    payload.CourseID,
		Order:       payload.Order,
		Description: payload.Description,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.NewResponseErrorFromError(err, web.LessonNotExist)
		}
		// Uncaught error
		return err
	}

	return nil
}

func (l LessonServiceImpl) UpdateLessonMaterial(payload materials.UpdateLessonMaterialsRequestPayload) error {
	// Validate Role
	claim, err := l.TokenUtil.Validate(payload.UpdateLessonMaterialsToken, token.Access)

	// Invalid Token
	if err != nil {
		return web.NewResponseErrorFromError(err, web.TokenError)
	}

	// Unauthorized Role
	if claim.Role == user.Student {
		return web.NewResponseErrorFromError(err, web.UnauthorizedAccess)
	}

	err = l.LessonMaterialsRepository.UpdateLessonMaterial(domLesson.LessonMaterials{
		ID:         payload.ID,
		LessonID:   payload.LessonID,
		Order:      payload.Order,
		MaterialID: payload.MaterialID,
		Contents:   payload.Contents,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return web.NewResponseErrorFromError(err, web.LessonMaterialNotExist)
		}
		// Uncaught error
		return err
	}

	return nil
}
