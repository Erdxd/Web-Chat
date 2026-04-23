package repositories

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"Web-Chat/internal/repositories/entities"
	"database/sql"
	"log"
	"time"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repository.User {
	return &UserRepo{db: db}
}
func (U *UserRepo) CreateAcc(Data model.User) error {
	DataUserToDB := entities.User{
		UserId:    Data.UserId,
		Name:      Data.Name,
		Password:  Data.Password,
		Email:     Data.Email,
		CreatedAt: Data.CreatedAt,
	}
	SqlStatement := (`INSERT INTO "users" (name,password,email,ca) VALUES ($1,$2,$3,$4)`)
	_, err := U.db.Exec(SqlStatement, DataUserToDB.Name, DataUserToDB.Password, DataUserToDB.Email, DataUserToDB.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (U *UserRepo) Login(Email string) (string, error) {
	var Password string
	SqlStatement := (`SELECT password FROM users WHERE email=$1`)
	err := U.db.QueryRow(SqlStatement, Email).Scan(&Password)
	log.Println(Password)
	if err != nil {
		return "", err
	}

	return Password, err
}
func (U *UserRepo) GetUserDataForJWT(Email string) (int, bool, error) {
	var userid int
	var admin bool
	SqlStatement := (`SELECT userid,admin FROM users WHERE email=$1`)
	err := U.db.QueryRow(SqlStatement, Email).Scan(&userid, &admin)

	if err != nil {
		return 0, false, err
	}
	return userid, admin, nil
}
func (U *UserRepo) GetNameUserById(UserId int) (string, error) {
	var name string
	SqlStatement := (`SELECT name FROM users WHERE userid=$1`)
	err := U.db.QueryRow(SqlStatement, UserId).Scan(&name)

	if err != nil {
		return "", err
	}
	return name, nil
}
func (U *UserRepo) GetUserId(usertag string) (int, error) {
	var userid int
	SqlStatemnt := (`SELECT userid FROM users WHERE usertag=$1`)
	err := U.db.QueryRow(SqlStatemnt, usertag).Scan(&userid)
	if err != nil {
		return 0, err
	}
	return userid, nil
}
func (U *UserRepo) GetDataAboutUserForProfile(UserId int) (model.UserView, error) {
	var email string
	var name string
	var ca time.Time
	var usertag string
	SqlStatemnt := (`SELECT email,name,ca,usertag FROM users WHERE userid=$1`)
	err := U.db.QueryRow(SqlStatemnt, UserId).Scan(&email, &name, &ca, &usertag)
	if err != nil {
		return model.UserView{}, err
	}
	return model.UserView{Email: email, Name: name, Ca: ca, Usertag: usertag}, nil

}
func (U *UserRepo) RedactUserTag(NewUserTag string, UserId int) error {
	SqlStatement := (`UPDATE users SET usertag =$1 WHERE userid=$2`)
	_, err := U.db.Exec(SqlStatement, NewUserTag, UserId)
	if err != nil {
		return err
	}
	return nil
}
func (U *UserRepo) RedactPassword(NewPassword string, UserId int) error {
	SqlStatement := (`UPDATE users SET password =$1 WHERE userid=$2`)
	_, err := U.db.Exec(SqlStatement, NewPassword, UserId)
	if err != nil {
		return err
	}
	return nil
}
func (U *UserRepo) RedactName(NewName string, UserId int) error {
	SqlStatement := (`UPDATE users SET name =$1 WHERE userid=$2`)
	_, err := U.db.Exec(SqlStatement, NewName, UserId)
	if err != nil {
		return err
	}
	return nil
}
func (U *UserRepo) FindUserByUserTag(UserTag string) (model.UserSerchResult, error) {
	var UserSearchResult model.UserSerchResult
	SqlStatemnt := (`SELECT name,userid FROM users WHERE usertag = $1`)
	err := U.db.QueryRow(SqlStatemnt, UserTag).Scan(&UserSearchResult.Name, &UserSearchResult.UserId)
	if err != nil {
		return model.UserSerchResult{}, err
	}
	return UserSearchResult, nil
}
