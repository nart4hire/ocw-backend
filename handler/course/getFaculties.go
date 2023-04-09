package course

import (
	"fmt"
	"net/http"

	"gitlab.informatika.org/ocw/ocw-backend/model/web"
)
// Index godoc
//
//	@Summary		Get all faculties
//	@Description	Retrieves a list of all faculties
//	@Tags			course
//	@Produce		json
//	@Success		200	{object}	web.BaseResponse	"OK"
//	@Failure		401	{object}	web.BaseResponse	"Unauthorized"
//	@Failure		500	{object}	web.BaseResponse	"Internal Server Error"
//	@Router			/course/faculty [get]
func (c CourseHandlerImpl) GetFaculties(w http.ResponseWriter, r *http.Request) {
	packet, err := c.CourseService.GetAllFaculty()

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