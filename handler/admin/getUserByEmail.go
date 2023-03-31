package admin

import (
	"net/http"
	"path"
)

// Index godoc
//
//	@Tags			admin
//	@Summary		Get User By Email
//	@Description	Get a user from database
//	@Produce		json
//	@Success		200	{object}	web.BaseResponse
//	@Router			/admin/user/{email} [get]
func (route AdminHandlerImpl) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	// email := r.URL.Query().Get("email")
	email := path.Base(r.URL.Path)

	// get user from database
	user, err := route.AdminService.GetUserByEmail(email)

	if err != nil {
		// error handling
		payload := route.WrapperUtil.ErrorResponseWrap("error", err.Error())
		route.HttpUtil.WriteJson(w, http.StatusUnprocessableEntity, payload)
		return
	}

	// return user
	result := route.WrapperUtil.SuccessResponseWrap(user)
	route.HttpUtil.WriteJson(w, http.StatusOK, result)
}
