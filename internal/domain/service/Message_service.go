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
func (S *ServiceMessage) Save(msg model.Message, IdUser int) (int64, error) {
	log.Println(msg)
	if msg.Content == "" {
		log.Println(msg)
		return 0, errors.New("Empty field")

	}

	return S.RepoM.Save(msg, IdUser)
}
func (S *ServiceMessage) CheckMessage(RoomId int) ([]model.Message, error) {
	log.Println(RoomId)
	return S.RepoM.CheckMessages(RoomId)
}
func (S *ServiceMessage) DeleteMessage(UserId int, RoomdId int, Id int64) error {
	return S.RepoM.DeleteMessage(UserId, RoomdId, Id)

}
