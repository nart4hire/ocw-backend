package auth

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/login"
)

// Index godoc
//
//		@Tags					auth
//		@Summary			Login
//		@Description	Login and generate new pair of token
//		@Produce			json
//		@Accept				json
//		@Param				data body login.LoginRequestPayload true "Login payload"
//		@Success			200	{object}	web.BaseResponse{data=login.LoginResponsePayload}
//	  @Failure			403 {object}  web.BaseResponse
//		@Router				/auth/login [post]
func (a AuthHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	payload := login.LoginRequestPayload{}
	validate := validator.New()

	if err := a.HttpUtil.ParseJson(r, &payload); err != nil {
		payload := a.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
		a.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	if err := validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			payload := a.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
			a.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			return
		}

		errList := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errList = append(errList, err.Error())
		}

		payload := a.WrapperUtil.ErrorResponseWrap("input validation error", errList)
		a.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	response, err := a.AuthService.Login(payload)

	if err != nil {
		payload := a.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
		a.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	responsePayload := a.WrapperUtil.SuccessResponseWrap(response)
	a.HttpUtil.WriteSuccessJson(w, responsePayload)
}
