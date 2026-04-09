package hasher

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
}

func NewHasher() *Hasher {
	return &Hasher{}
}
func (H *Hasher) Hash(password string) ([]byte, error) {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return HashedPassword, nil
}
func (H *Hasher) Compare(HashedPassword []byte, Password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(HashedPassword, []byte(Password))
	if err != nil {
		return false, err
	}
	return err == nil, nil
}
