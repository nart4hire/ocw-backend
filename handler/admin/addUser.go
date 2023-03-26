package admin

import (
	"net/http"
	req "gitlab.informatika.org/ocw/ocw-backend/model/web/admin/addUser"
)

// Index godoc
//
//	@Tags					admin
//	@Summary			Add User
//	@Description	Add a user to database
//	@Produce			json
//	@Accept				json
//	@Param				data body req.AdminAddUserPayload true "Admin Add User Payload"
//	@Success			200	{object}	web.BaseResponse
//	@Router				/admin/user [post]
func (route AdminHandlerImpl) AddUser(w http.ResponseWriter, r *http.Request){
	payload := req.AdminAddUserPayload{}

	if r.Header.Get("Content-Type") != "application/json" {
		payload := route.WrapperUtil.ErrorResponseWrap("this service only receive json input", nil)
		route.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	if err := route.HttpUtil.ParseJson(r, &payload); err != nil {
		payload := route.WrapperUtil.ErrorResponseWrap("invalid json input", err.Error())
		route.HttpUtil.WriteJson(w, http.StatusUnprocessableEntity, payload)
		return
	}

	err := route.AdminService.AddUser(payload)

	if err != nil {
		payload := route.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
		route.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}
	
	route.HttpUtil.WriteSuccessJson(w, payload)
}