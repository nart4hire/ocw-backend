package material

import (
	"gitlab.informatika.org/ocw/ocw-backend/service/material"
	"gitlab.informatika.org/ocw/ocw-backend/utils/httputil"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type MaterialHandlerImpl struct {
	material.MaterialService
	material.MaterialContentService
	httputil.HttpUtil
	wrapper.WrapperUtil
}
