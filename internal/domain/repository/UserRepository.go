package repository

import (
	"Web-Chat/internal/domain/model"
	"time"
)

type User interface {
	CreateAcc(model.User) error
	Login(Email string) (string, error)
	GetUserDataForJWT(email string) (int, bool, error)
	GetNameUserById(UserId int) (string, error)
	GetUserId(usertag string) (int, error)
	GetDataAboutUserForProfile(UserId int) (model.UserView, error)
	RedactUserTag(NewUserTag string, UserId int) error
	RedactPassword(NewPassword string, UserId int) error
	RedactName(NewName string, UserId int) error
	FindUserByUserTag(UserTag string) (model.UserSerchResult, error)
	LastSeen(UserId int) error
	HowOnline(UserId int) (time.Time, error)
}
