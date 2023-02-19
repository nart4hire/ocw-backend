package admin

import (
	"net/http"
)

// Index godoc
//
//	@Tags					admin
//	@Summary			Delete User By Id
//	@Description	Delete a user from database
//	@Produce			json
//	@Success			200	{object}	web.BaseResponse
//	@Router				/admin/user [delete]
func (route AdminHandlerImpl) DeleteUser(w http.ResponseWriter, r *http.Request){
	payload := route.WrapperUtil.SuccessResponseWrap(route.AdminService.DeleteUser())
	route.HttpUtil.WriteSuccessJson(w, payload)
}