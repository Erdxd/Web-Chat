package repository

import "Web-Chat/internal/domain/model"

type User interface {
	CreateAcc(model.User) error
	Login(Email string) (string, error)
}
