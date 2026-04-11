package repositories

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"Web-Chat/internal/repositories/entities"
	"database/sql"
	"log"
)

type RepoMessage struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) repository.Message {
	return &RepoMessage{db: db}
}
func (R RepoMessage) Save(msg model.Message, IdUser int) error {
	log.Println(msg)
	modelDB := entities.Message{
		Id:        msg.Id,
		UserId:    msg.UserId,
		RoomId:    msg.RoomId,
		CreatedAt: msg.CreatedAt,
		Content:   msg.Content,
	}
	SqlStatement := (`INSERT INTO messages(useridmessage,roomidmessage,CAmessage,message) VALUES ($1,$2,$3,$4)`)
	_, err := R.db.Exec(SqlStatement, IdUser, modelDB.RoomId, modelDB.CreatedAt, modelDB.Content)
	if err != nil {
		log.Println(err)
	}

	return nil
}
func (R RepoMessage) CheckMessages(RoomId int) ([]model.Message, error) {
	log.Println(RoomId)
	var messages []model.Message
	rows, err := R.db.Query(`SELECT * FROM messages WHERE roomidmessage=$1 `, RoomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var messageDB entities.Message
		err := rows.Scan(&messageDB.Id, &messageDB.UserId, &messageDB.RoomId, &messageDB.CreatedAt, &messageDB.Content)
		if err != nil {
			return nil, err
		}

		messages = append(messages, model.Message{
			Id:        messageDB.Id,
			UserId:    messageDB.UserId,
			RoomId:    messageDB.RoomId,
			CreatedAt: messageDB.CreatedAt,
			Content:   messageDB.Content,
		})
	}
	log.Println(messages)
	return messages, err
}
