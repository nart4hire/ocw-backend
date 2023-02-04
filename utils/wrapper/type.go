package wrapper

import "gitlab.informatika.org/ocw/ocw-backend/model/web"

type WrapperUtil interface {
	SuccessResponseWrap(data interface{}) *web.BaseResponse
	ErrorResponseWrap(message string, payload interface{}) *web.BaseResponse
}
