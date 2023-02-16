package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/register"
)

// Index godoc
//
//		@Tags					auth
//		@Summary			Register New Account
//		@Description	Generate New Account as Member
//		@Produce			json
//		@Accept				json
//		@Param				data body register.RegisterRequestPayload true "Register Payload"
//		@Success			200	{object}	web.BaseResponse
//	  @Failure			400 {object}  web.BaseResponse
//	  @Failure			500 {object}  web.BaseResponse
//		@Router				/auth/register [post]
func (a AuthHandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	payload := register.RegisterRequestPayload{}
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

	err := a.AuthService.Register(payload)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			err = web.NewResponseErrorFromError(err, web.EmailExist)
		}

		respErr, ok := err.(web.ResponseError)
		if ok {
			payload := a.WrapperUtil.ErrorResponseWrap("email was registered by other account", respErr)
			a.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		} else {
			a.Logger.Error(
				fmt.Sprintf("[AUTH] some error happened when do register: %s", err.Error()),
			)
			payload := a.WrapperUtil.ErrorResponseWrap("internal server error", nil)
			a.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		}
		return
	}

	responsePayload := a.WrapperUtil.SuccessResponseWrap(nil)
	a.HttpUtil.WriteSuccessJson(w, responsePayload)
}
