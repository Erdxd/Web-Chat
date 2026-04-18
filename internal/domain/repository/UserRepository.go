package repository

import "Web-Chat/internal/domain/model"

type User interface {
	CreateAcc(model.User) error
	Login(Email string) (string, error)
	GetUserDataForJWT(email string) (int, bool, error)
<<<<<<< HEAD
	GetNameUserById(UserId int) (string, error)
=======
	GetUserId(usertag string) (int, error)
>>>>>>> 2633c2f777616eef751127b29db2bd214789ed21
}
