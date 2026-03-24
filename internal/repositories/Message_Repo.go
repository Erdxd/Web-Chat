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
func (R RepoMessage) Save(model.Message) error {
	var err error
	return err
}
func (R RepoMessage) Last(count int) ([]model.Message, error) {
	var err error
	return []model.Message{}, err
}
