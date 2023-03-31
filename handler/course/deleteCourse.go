package course

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course"
)

// Index godoc
//
//	@Summary		Delete course by ID
//	@Description	Delete a course with the specified ID
//	@Tags			course
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string				true	"Course ID"
//	@Param			Authorization	header		string				true	"DeleteCourseToken"
//	@Success		200				{object}	web.BaseResponse	"Success"
//	@Failure		400				{object}	web.BaseResponse	"Bad Request"
//	@Failure		401				{object}	web.BaseResponse	"Unauthorized"
//	@Failure		403				{object}	web.BaseResponse	"Forbidden"
//	@Failure		422				{object}	web.BaseResponse	"Unprocessable Entity"
//	@Failure		500				{object}	web.BaseResponse	"Internal Server Error"
//	@Router			/course/{id} [delete]
func (c CourseHandlerImpl) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	payload := course.DeleteByStringRequestPayload{}
	validateTokenHeader := r.Header.Get("Authorization")

	if validateTokenHeader == "" {
		payload := c.WrapperUtil.ErrorResponseWrap("token is required", nil)
		c.HttpUtil.WriteJson(w, http.StatusForbidden, payload)
		return
	}

	token := strings.Split(validateTokenHeader, " ")

	if len(token) != 2 {
		payload := c.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		c.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	if token[0] != "Bearer" {
		payload := c.WrapperUtil.ErrorResponseWrap("invalid token", nil)
		c.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	payload.DeleteCourseToken = token[1]
	err := c.CourseService.DeleteCourse(payload)

	if err != nil {
		if errData, ok := err.(web.ResponseError); ok {
			payload := c.WrapperUtil.ErrorResponseWrap(errData.Error(), errData)
			c.HttpUtil.WriteJson(w, http.StatusUnauthorized, payload)
			return
		}

		c.Logger.Error(
			fmt.Sprintf("[RESET] some error happened when validating URL: %s", err.Error()),
		)
		payload := c.WrapperUtil.ErrorResponseWrap("internal server error", nil)
		c.HttpUtil.WriteJson(w, http.StatusInternalServerError, payload)
		return
	}

	responsePayload := c.WrapperUtil.SuccessResponseWrap(nil)
	c.HttpUtil.WriteSuccessJson(w, responsePayload)
	
}