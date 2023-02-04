package common

import "net/http"

func (route CommonHandlerImpl) NotFound(w http.ResponseWriter, r *http.Request) {
	payload := route.WrapperUtil.ErrorResponseWrap(route.CommonService.NotFound(), nil)
	route.HttpUtil.WriteJson(w, http.StatusNotFound, payload)
}
