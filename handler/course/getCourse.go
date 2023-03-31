package course

import (
	"fmt"
	"net/http"
	"path"

	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course"
)

// Index godoc
//
//	@Summary		Get a course by ID
//	@Description	Retrieve a course by ID
//	@Tags			course
//	@Produce		json
//	@Param			id	path		string				true	"ID of the course to retrieve"
//	@Success		200	{object}	web.BaseResponse	"OK"
//	@Failure		401	{object}	web.BaseResponse	"Unauthorized"
//	@Failure		500	{object}	web.BaseResponse	"Internal Server Error"
//	@Router			/course/{id} [get]
func (c CourseHandlerImpl) GetCourse(w http.ResponseWriter, r *http.Request) {
	payload := course.GetByStringRequestPayload{}
	payload.ID = path.Base(r.URL.Path)
	
	packet, err := c.CourseService.GetCourse(payload)

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

	responsePayload := c.WrapperUtil.SuccessResponseWrap(packet)
	c.HttpUtil.WriteSuccessJson(w, responsePayload)
}