package admin

import (
	"net/http"
)

// Index godoc
//
//	@Tags					admin
//	@Summary			Add User
//	@Description	Add a user to database
//	@Produce			json
//	@Success			200	{object}	web.BaseResponse
//	@Router				/admin/user [post]
func (route AdminHandlerImpl) AddUser(w http.ResponseWriter, r *http.Request){
	payload := route.WrapperUtil.SuccessResponseWrap(route.AdminService.AddUser())
	route.HttpUtil.WriteSuccessJson(w, payload)
}