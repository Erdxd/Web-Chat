package service

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"errors"
	"log"
)

type ServiceMessage struct {
	RepoM repository.Message
}

func NewServiceMessage(Repo repository.Message) *ServiceMessage {
	return &ServiceMessage{RepoM: Repo}
}
func (S *ServiceMessage) Save(msg model.Message, IdUser int) error {
	log.Println(msg)
	if msg.Content == "" {
		return errors.New("Empty field")
	}
	return S.RepoM.Save(msg, IdUser)
}
func (S *ServiceMessage) CheckMessages(IdUser int) ([]model.Message, error) {
	return S.RepoM.CheckMessages(IdUser)
}
