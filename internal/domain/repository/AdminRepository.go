package repository

import "Web-Chat/internal/domain/model"

type Admin interface {
	CheckAllUsers() ([]model.User, error)
	FoundUserByUserId(UserId int) (model.User, error)
	DeleteUser(UserId int) error
}
