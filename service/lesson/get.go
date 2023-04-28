package lesson

import (
	"errors"

	domLesson "gitlab.informatika.org/ocw/ocw-backend/model/domain/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/lesson/materials"
	"gorm.io/gorm"
)

func (l LessonServiceImpl) GetLesson(payload lesson.GetByUUIDRequestPayload) (*domLesson.Lesson, error) {
	packet, err := l.LessonRepository.GetLesson(payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.LessonNotExist)
		}
		// Some Uncaught error
		return nil, err
	}
	
	return packet, nil
}

func (l LessonServiceImpl) GetLessons(payload lesson.GetByStringRequestPayload) ([]domLesson.Lesson, error) {
	packet, err := l.LessonRepository.GetLessons(payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.LessonNotExist)
		}
		// Some Uncaught error
		return nil, err
	}
	
	return packet, nil
}

func (l LessonServiceImpl) GetLessonMaterial(payload materials.GetByUUIDRequestPayload) (*domLesson.LessonMaterials, error) {
	packet, err := l.LessonMaterialsRepository.GetLessonMaterial(payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.LessonMaterialNotExist)
		}
		// Some Uncaught error
		return nil, err
	}
	
	return packet, nil
}

func (l LessonServiceImpl) GetLessonMaterials(payload materials.GetByUUIDRequestPayload) ([]domLesson.LessonMaterials, error) {
	packet, err := l.LessonMaterialsRepository.GetLessonMaterials(payload.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, web.NewResponseErrorFromError(err, web.LessonMaterialNotExist)
		}
		// Some Uncaught error
		return nil, err
	}
	
	return packet, nil
}