package protocol

type IBcryptAdapter interface {
	Hash(password string, salt int) (string, error)
	Compare(hashedPassword, password string) error
}
