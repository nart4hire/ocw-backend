package admin

import (
	"net/http"
	"path"
)

// Index godoc
//
//	@Tags					admin
//	@Summary			Delete User By Id
//	@Description	Delete a user from database
//	@Produce			json
//	@Accept				json
//	@Success			200	{object}	web.BaseResponse
//	@Router				/admin/user/{email} [delete]
func (route AdminHandlerImpl) DeleteUser(w http.ResponseWriter, r *http.Request){
	email := path.Base(r.URL.Path)

	// get user from database
	err := route.AdminService.DeleteUser(email)

	if err != nil {
		// error handling
		payload := route.WrapperUtil.ErrorResponseWrap("error", err.Error())
		route.HttpUtil.WriteJson(w, http.StatusUnprocessableEntity, payload)
		return
	}

	// return user
	result := route.WrapperUtil.SuccessResponseWrap(email)
	route.HttpUtil.WriteJson(w, http.StatusOK, result)
}