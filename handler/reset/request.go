package reset

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/request"
)

// Index godoc
//
//	@Tags					reset
//	@Summary			Request Reset Password Token
//	@Description	Send Reset password token to email
//	@Produce			json
//	@Param				data body request.RequestRequestPayload true "payload"
//	@Success			200	{object}	web.BaseResponse "Login Success"
//	@Router				/reset/request [post]
func (rs ResetHandlerImpl) Request(w http.ResponseWriter, r *http.Request) {
	payload := request.RequestRequestPayload{}
	validate := validator.New()

	if r.Header.Get("Content-Type") != "application/json" {
		payload := rs.WrapperUtil.ErrorResponseWrap("this service only receive json input", nil)
		rs.HttpUtil.WriteJson(w, http.StatusUnsupportedMediaType, payload)
		return
	}

	if err := rs.HttpUtil.ParseJson(r, &payload); err != nil {
		payload := rs.WrapperUtil.ErrorResponseWrap("invalid json input", err.Error())
		rs.HttpUtil.WriteJson(w, http.StatusUnprocessableEntity, payload)
		return
	}

	if err := validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			payload := rs.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
			rs.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
			return
		}

		errPayload := web.NewResponseErrorFromValidator(err.(validator.ValidationErrors))
		payload := rs.WrapperUtil.ErrorResponseWrap(errPayload.Error(), errPayload)
		rs.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	err := rs.ResetService.Request(payload)

	if err != nil {
		if strings.Contains(err.Error(), "unknown key") {
			err = web.NewResponseErrorFromError(err, web.EmailNotExist)
		}

		respErr, ok := err.(web.ResponseError)
		if ok {
			payload := rs.WrapperUtil.ErrorResponseWrap("email was not found", respErr)
			rs.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		} else {
			rs.Logger.Error(
				fmt.Sprintf("[RESET] some error happened when requesting reset email: %s", err.Error()),
			)
			payload := rs.WrapperUtil.ErrorResponseWrap("internal server error", nil)
			rs.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		}
		return
	}

	responsePayload := rs.WrapperUtil.SuccessResponseWrap(nil)
	rs.HttpUtil.WriteSuccessJson(w, responsePayload)
}
