package service

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"Web-Chat/internal/domain/repository/auth"
	"errors"
	"fmt"
	"time"
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
func (US *UserService) HowOnline(UserId int) (string, error) {
	Time, err := US.User.HowOnline(UserId)
	if err != nil {
		return "", err
	}
	if Time.IsZero() {
		return "🔴 OFFLINE", nil
	}

	diff := time.Since(Time)
	if diff > 60*time.Minute && diff < 24*time.Hour {
		Hours := int(diff.Hours())
		return fmt.Sprintf("WAS ONLINE %d HOURS AGO", Hours), nil
	} else if diff > 24*time.Hour {
		return fmt.Sprintf("LAST SEEN %s", Time.Format("02.01.2006")), nil
	} else if diff < 60*time.Minute && diff > 1*time.Minute {
		Minute := int(diff.Minutes())
		return fmt.Sprintf("WAS ONLINE %d MINUTES AGO", Minute), nil
	} else if diff < time.Minute {
		return "🟢 ONLINE", nil
	} else {
		return fmt.Sprintf("LAST SEEN %s", Time.Format("02.01.2006")), nil
	}

}
