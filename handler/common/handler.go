package common

import (
	"gitlab.informatika.org/ocw/ocw-backend/service/common"
	"gitlab.informatika.org/ocw/ocw-backend/utils/httputil"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type CommonHandlerImpl struct {
	common.CommonService
	httputil.HttpUtil
	wrapper.WrapperUtil
}
