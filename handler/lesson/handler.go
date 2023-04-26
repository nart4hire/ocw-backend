package lesson

import (
	r "gitlab.informatika.org/ocw/ocw-backend/repository/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/service/lesson"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/httputil"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type LessonHandlerImpl struct {
	r.LessonRepository
	lesson.LessonService
	httputil.HttpUtil
	wrapper.WrapperUtil
	logger.Logger
}