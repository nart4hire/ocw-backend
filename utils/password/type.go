package password

type PasswordUtil interface {
	Hash(password string) (string, error)
	Check(password string, hashedPassword string) error
}
