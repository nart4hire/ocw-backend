package admin

import (
	"net/http"
)

// Index godoc
//
//	@Tags			admin
//	@Summary		Get All User
//	@Description	Get all users from database
//	@Produce		json
//	@Success		200	{object}	web.BaseResponse
//	@Router			/admin/user [get]
func (route AdminHandlerImpl) GetAllUser(w http.ResponseWriter, r *http.Request){
	// get all user from service
	users, err := route.AdminService.GetAllUser()
	if err != nil {
		payload := route.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
		route.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}
	
	// wrap the response
	payload := route.WrapperUtil.SuccessResponseWrap(users)
	route.HttpUtil.WriteSuccessJson(w, payload)
}