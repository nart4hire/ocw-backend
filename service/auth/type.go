package auth

type AuthService interface {
	Login(email string, password string)
}
