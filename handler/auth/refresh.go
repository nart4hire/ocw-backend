package auth

import (
	"fmt"
	"net/http"
	"strings"

	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/refresh"
)

// Index godoc
//
//		@Tags					auth
//		@Summary			Refresh Token
//		@Description	Generate new access token
//		@Produce			json
//		@Accept				json
//		@Param				Authorization header string true "Refresh token"
//		@Success			200	{object}	web.BaseResponse{data=refresh.RefreshResponsePayload}
//	  @Failure			400 {object}  web.BaseResponse
//	  @Failure			401 {object}  web.BaseResponse
//	  @Failure			500 {object}  web.BaseResponse
//		@Router				/auth/refresh [post]
func (a AuthHandlerImpl) Refresh(w http.ResponseWriter, r *http.Request) {
	payload := refresh.RefreshRequestPayload{}

	refreshTokenHeader := r.Header.Get("Authorization")

	if refreshTokenHeader == "" {
		payload := a.WrapperUtil.ErrorResponseWrap("token is required", nil)
		a.HttpUtil.WriteJson(w, http.StatusForbidden, payload)
		return
	}

	token := strings.Split(refreshTokenHeader, " ")

	if len(token) != 2 {
		payload := a.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		a.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	if token[0] != "Bearer" {
		payload := a.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		a.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	payload.RefreshToken = token[1]
	response, err := a.AuthService.Refresh(payload)

	if err != nil {
		if errData, ok := err.(web.ResponseError); ok {
			payload := a.WrapperUtil.ErrorResponseWrap(errData.Error(), errData)
			a.HttpUtil.WriteJson(w, http.StatusUnauthorized, payload)
			return
		}

		a.Logger.Error(
			fmt.Sprintf("[AUTH] some error happened when do refresh: %s", err.Error()),
		)
		payload := a.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		a.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	responsePayload := a.WrapperUtil.SuccessResponseWrap(response)
	a.HttpUtil.WriteSuccessJson(w, responsePayload)
}
