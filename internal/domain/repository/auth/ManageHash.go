package auth

type Hash interface {
	Hash(Password string) (string, error)
	Compare(HashedPassword []byte, Password string) (bool, error)
}
