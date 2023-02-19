package admin

import (
	"net/http"
)

// Index godoc
//
//	@Tags					admin
//	@Summary			Get User By Email
//	@Description	Get a user from database
//	@Produce			json
//	@Success			200	{object}	web.BaseResponse
//	@Router				/admin/user/{id} [get]
func (route AdminHandlerImpl) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	payload := route.WrapperUtil.SuccessResponseWrap(route.AdminService.GetUserByEmail())
	route.HttpUtil.WriteSuccessJson(w, payload)
}
