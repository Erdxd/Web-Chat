package auth

type Hash interface {
	Hash(Password string) ([]byte, error)
	Compare(HashedPassword []byte, Password string) (bool, error)
}
