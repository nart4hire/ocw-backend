package auth

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/login"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/refresh"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/register"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/auth/verification"
)

type AuthService interface {
	Login(payload login.LoginRequestPayload) (*login.LoginResponsePayload, error)
	Refresh(payload refresh.RefreshRequestPayload) (*refresh.RefreshResponsePayload, error)
	Register(payload register.RegisterRequestPayload) error
	SendVerifyEmail(payload verification.VerificationSendRequestPayload) error
}
