package reset

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/request"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/confirm"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/validate"
)

type ResetService interface {
	Request(payload request.RequestRequestPayload) error
	Confirm(payload confirm.ConfirmRequestPayload) error
	Validate(payload validate.ValidateRequestPayload) error
}
