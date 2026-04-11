package repositories

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"Web-Chat/internal/repositories/entities"
	"database/sql"
	"log"
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
	log.Println(12783457812576812)
	if err != nil {
		return "", err
	}

	return Password, err
}
