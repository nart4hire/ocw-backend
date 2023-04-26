package quiz

import (
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/service/quiz"
	"gitlab.informatika.org/ocw/ocw-backend/utils/httputil"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type QuizHandlerImpl struct {
	quiz.QuizService
	wrapper.WrapperUtil
	httputil.HttpUtil
	logger.Logger
}
