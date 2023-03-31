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
//	@Summary		Get faculty by ID
//	@Description	Retrieve faculty data by UUID
//	@Tags			course
//	@Produce		json
//	@Param			id	path		string				true	"Faculty ID (UUID)"
//	@Success		200	{object}	web.BaseResponse	"Success"
//	@Failure		400	{object}	web.BaseResponse	"Invalid UUID provided in request path"
//	@Failure		401	{object}	web.BaseResponse	"Unauthorized"
//	@Failure		500	{object}	web.BaseResponse	"Internal Server Error"
//	@Router			/course/faculty/{id} [get]
func (c CourseHandlerImpl) GetFaculty(w http.ResponseWriter, r *http.Request) {
	payload := course.GetByUUIDRequestPayload{}
	id, err := uuid.Parse(path.Base(r.URL.Path))
	
	if err != nil {
		// invalid uuid
		payload := c.WrapperUtil.ErrorResponseWrap(err.Error(), nil)
		c.HttpUtil.WriteJson(w, http.StatusBadRequest, payload)
		return
	}

	payload.ID = id
	packet, err := c.CourseService.GetFaculty(payload)

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