package repositories

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"database/sql"
	"log"
)

type AdminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) repository.Admin {
	return &AdminRepo{db: db}
}
func (AR *AdminRepo) CheckAllUsers() ([]model.User, error) {
	SqlStatement := (`SELECT userid,name,email,ca FROM users`)
	rows, err := AR.db.Query(SqlStatement)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var Users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		Users = append(Users, user)
	}
	return Users, nil

}
func (AR *AdminRepo) FoundUserByUserId(UserId int) (model.User, error) {
	var user model.User
	SqlStatement := (`SELECT userid,name,email,ca FROM users WHERE userid=$1`)
	err := AR.db.QueryRow(SqlStatement, UserId).Scan(&user.UserId, &user.Name, &user.Email, &user.CreatedAt)
	return user, err

}
func (AR *AdminRepo) DeleteUser(UserId int) error {
	SqlStatemnt := (`DELETE FROM users WHERE userid=$1`)
	_, err := AR.db.Exec(SqlStatemnt, UserId)
	if err != nil {
		return err
	}
	return nil
}
