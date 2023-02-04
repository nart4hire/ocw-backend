package wrapper

import "gitlab.informatika.org/ocw/ocw-backend/model/web"

type WrapperUtilImpl struct{}

func (WrapperUtilImpl) SuccessResponseWrap(data interface{}) *web.BaseResponse {
	return &web.BaseResponse{
		Status:  web.Success,
		Message: "success",
		Data:    data,
	}
}

func (WrapperUtilImpl) ErrorResponseWrap(message string, payload interface{}) *web.BaseResponse {
	return &web.BaseResponse{
		Status:  web.Failed,
		Message: message,
		Data:    payload,
	}
}
