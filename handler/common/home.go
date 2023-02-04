package common

import (
	"net/http"
)

// Index godoc
//
//	@Summary			Index page
//	@Description	Give server index page response
//	@Produce			json
//	@Success			200	{object}	web.BaseResponse
//	@Router				/ [get]
func (route CommonHandlerImpl) Home(w http.ResponseWriter, r *http.Request) {
	payload := route.WrapperUtil.SuccessResponseWrap(route.CommonService.Home())
	route.HttpUtil.WriteSuccessJson(w, payload)
}
