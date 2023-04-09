package course

import (
	"fmt"
	"net/http"
	"path"

	"github.com/google/uuid"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/course"
)

// Index godoc
//
//	@Summary		Get a major by ID
//	@Description	Get a major by ID
//	@Tags			course
//	@Produce		json
//	@Param			id	path		string				true	"Major ID (UUID)"
//	@Success		200	{object}	web.BaseResponse	"Success"
//	@Failure		400	{object}	web.BaseResponse	"Invalid UUID provided in request path"
//	@Failure		401	{object}	web.BaseResponse	"Unauthorized"
//	@Failure		500	{object}	web.BaseResponse	"Internal Server Error"
//	@Router			/course/major/{id} [get]
func (c CourseHandlerImpl) GetMajor(w http.ResponseWriter, r *http.Request) {
	payload := course.GetByUUIDRequestPayload{}
	id, err := uuid.Parse(path.Base(r.URL.Path))
	
	if err != nil {
		// invalid uuid
		payload := c.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
		c.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	payload.ID = id
	packet, err := c.CourseService.GetMajor(payload)

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