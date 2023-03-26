package course

import (
	r "gitlab.informatika.org/ocw/ocw-backend/repository/course"
	"gitlab.informatika.org/ocw/ocw-backend/service/course"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/httputil"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type CourseHandlerImpl struct {
	r.CourseRepository
	course.CourseService
	httputil.HttpUtil
	wrapper.WrapperUtil
	logger.Logger
}
