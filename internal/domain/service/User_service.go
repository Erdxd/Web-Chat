package service

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"Web-Chat/internal/domain/repository/auth"
	"errors"
	"log"
)

type UserService struct {
	User repository.User
	Hash auth.Hash
}

func NewUserService(User repository.User, Hash auth.Hash) *UserService {
	return &UserService{User: User, Hash: Hash}
}
func (US *UserService) CreateAcc(Data model.User, repeatpassword string) error {
	if Data.Password != repeatpassword {
		return errors.New("Passwords are different")
	}
	HashedPassword, err := US.Hash.Hash(Data.Password)
	if err != nil {
		return err
	}

	Data.Password = string(HashedPassword)
	log.Println(Data)
	return US.User.CreateAcc(Data)
}
func (US *UserService) Login(Email string, PasswordFromUser string) error {
	HashedPassword, err := US.User.Login(Email)
	if err != nil {
		return nil
	}
	Compare, err := US.Hash.Compare([]byte(HashedPassword), PasswordFromUser)
	if err != nil {
		return nil
	}
	if !Compare {
		return errors.New("Wrong Password or Email")
	}
	return nil
}
