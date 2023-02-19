package reset

import (
	"gitlab.informatika.org/ocw/ocw-backend/service/reset"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/httputil"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type ResetHandlerImpl struct {
	reset.ResetService
	httputil.HttpUtil
	wrapper.WrapperUtil
	logger.Logger
}
