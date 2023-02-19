package admin

import (
	"gitlab.informatika.org/ocw/ocw-backend/service/admin"
	"gitlab.informatika.org/ocw/ocw-backend/utils/httputil"
	"gitlab.informatika.org/ocw/ocw-backend/utils/wrapper"
)

type AdminHandlerImpl struct {
	admin.AdminService
	httputil.HttpUtil
	wrapper.WrapperUtil
}
