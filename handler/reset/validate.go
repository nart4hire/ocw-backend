package reset

import (
	"fmt"
	"net/http"
	"strings"

	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/validate"
)

// Index godoc
//
//	@Tags					reset
//	@Summary			Request Reset Password Token
//	@Description	Send Reset password token to email
//	@Produce			json
//	@Param			  Authorization header string true "Email validation token"
//	@Success			200	{object}	web.BaseResponse "Login Success"
//	@Router				/reset/validate [get]
func (rs ResetHandlerImpl) Validate(w http.ResponseWriter, r *http.Request) {
	payload := validate.ValidateRequestPayload{}
	validateTokenHeader := r.Header.Get("Authorization")

	if validateTokenHeader == "" {
		payload := rs.WrapperUtil.ErrorResponseWrap("token is required", nil)
		rs.HttpUtil.WriteJson(w, http.StatusForbidden, payload)
		return
	}

	token := strings.Split(validateTokenHeader, " ")

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

	payload.ValidateToken = token[1]
	err := rs.ResetService.Validate(payload)

	if err != nil {
		if errData, ok := err.(web.ResponseError); ok {
			payload := rs.WrapperUtil.ErrorResponseWrap(errData.Error(), errData)
			rs.HttpUtil.WriteJson(w, http.StatusUnauthorized, payload)
			return
		}

		rs.Logger.Error(
			fmt.Sprintf("[RESET] some error happened when validating URL: %s", err.Error()),
		)
		payload := rs.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		rs.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	responsePayload := rs.WrapperUtil.SuccessResponseWrap(nil)
	rs.HttpUtil.WriteSuccessJson(w, responsePayload)

}
