package auth

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/login"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/refresh"
)

type AuthService interface {
	Login(payload login.LoginRequestPayload) (*login.LoginResponsePayload, error)
	Refresh(payload refresh.RefreshRequestPayload) (*refresh.RefreshResponsePayload, error)
}
