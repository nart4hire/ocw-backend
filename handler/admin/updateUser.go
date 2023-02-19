package admin

import (
	"net/http"
)

// Index godoc
//
//	@Tags					admin
//	@Summary			Update User By Id
//	@Description	Update a user from database
//	@Produce			json
//	@Success			200	{object}	web.BaseResponse
//	@Router				/admin/user [patch]
func (route AdminHandlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request){
	payload := route.WrapperUtil.SuccessResponseWrap(route.AdminService.UpdateUser())
	route.HttpUtil.WriteSuccessJson(w, payload)
}