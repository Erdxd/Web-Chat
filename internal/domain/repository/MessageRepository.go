package repository

import "Web-Chat/internal/domain/model"

type Message interface {
	Save(model.Message, int) (int64, error)
	CheckMessages(int) ([]model.Message, error)
	DeleteMessage(UserId int, RoomId int, Id int64) error
}
