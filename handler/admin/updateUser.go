package admin

import (
	"net/http"
	"path"
	req "gitlab.informatika.org/ocw/ocw-backend/model/web/admin/updateUser"
)

// Index godoc
//
//	@Tags					admin
//	@Summary			Update User By Id
//	@Description	Update a user from database
//	@Produce			json
//	@Accept				json
//	@Param				data body req.AdminUpdateUserPayload true "Admin Update User Payload"
//	@Success			200	{object}	web.BaseResponse
//	@Router				/admin/user/{email} [patch]
func (route AdminHandlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request){
	email := path.Base(r.URL.Path)
	// TODO: how to change email

	payload := req.AdminUpdateUserPayload{}

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

	err := route.AdminService.UpdateUser(email, payload)

	if err != nil {
		payload := route.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
		route.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}
	
	route.HttpUtil.WriteSuccessJson(w, payload)
}