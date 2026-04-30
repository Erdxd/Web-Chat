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
func (US *UserService) CreateAcc(Data model.User, repeatpassword string) error {
	if Data.Password != repeatpassword {
		return errors.New("Passwords are different")
	}
	HashedPassword, err := US.Hash.Hash(Data.Password)
	if err != nil {
		return err
	}

	Data.Password = string(HashedPassword)

	return US.User.CreateAcc(Data)
}
func (US *UserService) Login(Email string, PasswordFromUser string) error {
	HashedPassword, err := US.User.Login(Email)
	if err != nil {
		return err
	}

	Compare, err := US.Hash.Compare([]byte(HashedPassword), PasswordFromUser)

	if err != nil {
		return err
	}
	if !Compare {
		return errors.New("Wrong Password or Email")
	}
	return nil
}
func (US *UserService) GetNameById(UserId int) (string, error) {
	return US.User.GetNameUserById(UserId)
}
func (US *UserService) GetUserId(usertag string) (int, error) {
	return US.User.GetUserId(usertag)
}
func (US *UserService) GetDataAboutUserForProfile(UserId int) (model.UserView, error) {

	return US.User.GetDataAboutUserForProfile(UserId)
}
func (US *UserService) RedactUserTag(NewUserTag string, UserId int) error {
	return US.User.RedactUserTag(NewUserTag, UserId)
}
func (US *UserService) RedactPassword(NewPassword, RepeatPassword string, UserId int) error {
	if RepeatPassword != NewPassword {
		return errors.New("Different Passwords")
	}
	NewHashPassword, err := US.Hash.Hash(NewPassword)
	if err != nil {
		return errors.New("Something is wrong")
	}
	return US.User.RedactPassword(string(NewHashPassword), UserId)
}
func (US *UserService) RedactName(NewName string, UserId int) error {
	return US.User.RedactName(NewName, UserId)
}
func (US *UserService) FindUserByUserTag(Usertag string) (model.UserSerchResult, error) {
	return US.User.FindUserByUserTag(Usertag)
}
func (US *UserService) LastSeen(UserId int) error {
	return US.User.LastSeen(UserId)
}
func (US *UserService) IsOnline(UserId int) (bool, error) {
	return US.User.IsOnline(UserId)
}
