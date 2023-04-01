package auth

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/login"
)

// Index godoc
//
//	@Tags			auth
//	@Summary		Login
//	@Description	Login and generate new pair of token
//	@Produce		json
//	@Accept			json
//	@Param			data	body		login.LoginRequestPayload							true	"Login payload"
//	@Success		200		{object}	web.BaseResponse{data=login.LoginResponsePayload}	"Login Success"
//	@Failure		400		{object}	web.BaseResponse{data=[]string}						"Bad Input"
//	@Failure		403		{object}	web.BaseResponse									"Login Credential Error"
//	@Failure		415		{object}	web.BaseResponse									"Not a json request"
//	@Failure		422		{object}	web.BaseResponse									"Invalid JSON input"
//	@Failure		500		{object}	web.BaseResponse									"Unknown Internal Error"
//	@Router			/auth/login [post]
func (a AuthHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	payload := login.LoginRequestPayload{}
	validate := validator.New()

	if r.Header.Get("Content-Type") != "application/json" {
		payload := a.WrapperUtil.ErrorResponseWrap("this service only receive json input", nil)
		a.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	if err := a.HttpUtil.ParseJson(r, &payload); err != nil {
		payload := a.WrapperUtil.ErrorResponseWrap("invalid json input", err.Error())
		a.HttpUtil.WriteJson(w, http.StatusUnprocessableEntity, payload)
		return
	}

	if err := validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			payload := a.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
			a.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			return
		}

		errPayload := web.NewResponseErrorFromValidator(err.(validator.ValidationErrors))
		payload := a.WrapperUtil.ErrorResponseWrap(errPayload.Error(), errPayload)
		a.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	response, err := a.AuthService.Login(payload)

	if err != nil {
		respErr, ok := err.(web.ResponseError)
		if ok {
			payload := a.WrapperUtil.ErrorResponseWrap(respErr.Error(), respErr)

			if respErr.Code != web.InvalidLogin {
				a.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			} else {
				a.HttpUtil.WriteJson(w, http.StatusForbidden, payload)
			}
		} else {
			a.Logger.Error(
				fmt.Sprintf("[AUTH] some error happened when do login: %s", err.Error()),
			)
			payload := a.WrapperUtil.ErrorResponseWrap("internal server error", nil)
			a.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		}
		return
	}

	responsePayload := a.WrapperUtil.SuccessResponseWrap(response)
	a.HttpUtil.WriteSuccessJson(w, responsePayload)
}
