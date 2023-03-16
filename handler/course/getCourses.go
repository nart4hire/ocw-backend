package course

import (
	"fmt"
	"net/http"

	"gitlab.informatika.org/ocw/ocw-backend/model/web"
)

func (c CourseHandlerImpl) GetCourses(w http.ResponseWriter, r *http.Request) {
	packet, err := c.CourseService.GetAllCourse()

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