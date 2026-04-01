package service

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"Web-Chat/internal/domain/repository/auth"
	"errors"
)

type UserService struct {
	User repository.User
	Hash auth.Hash
}

func NewUserService(User repository.User, Hash auth.Hash) *UserService {
	return &UserService{User: User, Hash: Hash}
}
func (US *UserService) CreateAcc(Data model.User) error {
	HashedPassword, err := US.Hash.Hash(Data.Password)
	if err != nil {
		return err
	}
	Data.Password = HashedPassword
	return US.User.CreateAcc(Data)
}
func (US *UserService) Login(Email string, HashedPasssword string, PasswordFromUser string) error {
	Compare, err := US.Hash.Compare([]byte(HashedPasssword), PasswordFromUser)
	if err != nil {
		return nil
	}
	if !Compare {
		return errors.New("Wrong Password or Email")
	}
	return nil
}
