package reset

import (
	"fmt"
	"net/http"

	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/request"
)

func (rs ResetHandlerImpl) Test(w http.ResponseWriter, r *http.Request) {
	payload := request.RequestRequestPayload{Email: "test@test.com",}

	err := rs.ResetService.Request(payload)

	if err != nil {
		fmt.Print("Oh no :)")
	}

	responsePayload := rs.WrapperUtil.SuccessResponseWrap(nil)
	rs.HttpUtil.WriteSuccessJson(w, responsePayload)
}