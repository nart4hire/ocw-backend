package lesson

import (
	"gitlab.informatika.org/ocw/ocw-backend/repository/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/utils/token"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

type LessonServiceImpl struct {
	lesson.LessonRepository
	lesson.LessonMaterialsRepository
	*env.Environment
	token.TokenUtil
	logger.Logger
}