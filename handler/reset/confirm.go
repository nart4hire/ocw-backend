package reset

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/confirm"
)

// Index godoc
//
//	@Tags			reset
//	@Summary		Confirm Reset Password
//	@Description	Do confirmation to reset password
//	@Produce		json
//	@Param			Authorization	header		string							true	"Email validation token"
//	@Param			data			body		confirm.ConfirmRequestPayload	true	"payload"
//	@Success		200				{object}	web.BaseResponse				"Login Success"
//	@Router			/reset/confirm [put]
func (rs ResetHandlerImpl) Confirm(w http.ResponseWriter, r *http.Request) {
	payload := confirm.ConfirmRequestPayload{}
	validate := validator.New()

	// Validate payload
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

	// Confirm Valid Website Token
	confirmTokenHeader := r.Header.Get("Authorization")

	if confirmTokenHeader == "" {
		payload := rs.WrapperUtil.ErrorResponseWrap("token is required", nil)
		rs.HttpUtil.WriteJson(w, http.StatusForbidden, payload)
		return
	}

	token := strings.Split(confirmTokenHeader, " ")

	if len(token) != 2 {
		payload := rs.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		rs.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	if token[0] != "Bearer" {
		payload := rs.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		rs.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	// Service to handle
	payload.ConfirmToken = token[1]
	err := rs.ResetService.Confirm(payload)

	if err != nil {
		if strings.Contains(err.Error(), "expired/not exist") {
			err = web.NewResponseErrorFromError(err, web.EmailNotExist)
		}

		respErr, ok := err.(web.ResponseError)
		if ok {
			payload := rs.WrapperUtil.ErrorResponseWrap("email was not found", respErr)
			rs.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		} else {
			rs.Logger.Error(
				fmt.Sprintf("[RESET] some error happened when requesting password reset: %s", err.Error()),
			)
			payload := rs.WrapperUtil.ErrorResponseWrap("internal server error", nil)
			rs.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		}
		return
	}

	responsePayload := rs.WrapperUtil.SuccessResponseWrap(nil)
	rs.HttpUtil.WriteSuccessJson(w, responsePayload)
}
