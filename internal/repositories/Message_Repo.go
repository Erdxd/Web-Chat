package repositories

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"database/sql"
)

type RepoMessage struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) repository.Message {
	return &RepoMessage{db: db}
}
func (R RepoMessage) Save(msg model.Message, IdUser int) error {
	SqlStatement := (`INSERT INTO messages(idmessage,useridmessage,roomidmessage,CAmessage,message) VALUES ($1,$2,$3,$4,$5)`)
	_, err := R.db.Exec(SqlStatement, msg.Id, msg.UserId, msg.RoomId, msg.CreatedAt, msg.Content)
	if err != nil {
		return err
	}

	return nil
}
func (R RepoMessage) CheckMessages(IdUser int) ([]model.Message, error) {
	rows, err := R.db.Query(`SELECT * FROM messages WHERE useridmessage=$1 `, IdUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var messages []model.Message
	for rows.Next() {
		var message model.Message
		err := rows.Scan(message.Id, message.UserId, message.RoomId, message.CreatedAt, message.Content)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, err
}
