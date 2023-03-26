package course

import (
	"gitlab.informatika.org/ocw/ocw-backend/repository/cache"
	"gitlab.informatika.org/ocw/ocw-backend/repository/course"
	"gitlab.informatika.org/ocw/ocw-backend/repository/user"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gitlab.informatika.org/ocw/ocw-backend/utils/token"
)

type CourseServiceImpl struct {
	user.UserRepository
	cache.CacheRepository
	course.CourseRepository
	*env.Environment
	token.TokenUtil
	logger.Logger
}
